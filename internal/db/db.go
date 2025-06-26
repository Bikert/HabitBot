package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
	"os"
)

func NewDB(lc fx.Lifecycle) (*sqlx.DB, error) {
	filepath := os.Getenv("DB_FILE_PATH")

	fmt.Println("Get file path = ", filepath)
	database, err := sqlx.Open("sqlite3", filepath)
	if err != nil {
		fmt.Println("not connected to db", err)
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("🚀 Connecting to SQLite...")

			if err := database.Ping(); err != nil {
				return fmt.Errorf("ping db: %w", err)
			}

			fmt.Println("✅ Running migrations...")

			driver, err := sqlite.WithInstance(database.DB, &sqlite.Config{}) //sqlite.WithInstance принимает *sql.DB и возвращает драйвер для миграций.
			if err != nil {
				fmt.Println("could not create driver", err)
				return err
			}
			m, err := migrate.NewWithDatabaseInstance(
				"file://internal/db/migrations", // путь относительно рабочей директории
				"sqlite3",
				driver,
			)
			err = m.Up()
			if err != nil && !errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("migration failed", err)
				return err
			}

			fmt.Println("migrations applied")
			return nil
		},
		OnStop: func(context.Context) error {
			return database.Close()
		},
	})

	return database, nil
}
