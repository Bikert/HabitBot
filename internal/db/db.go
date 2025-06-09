package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitDB(filepath string) (*sql.DB, error) {
	fmt.Println("Connecting to db ...")
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		fmt.Println("not connected to db", err)
		return nil, err
	}
	fmt.Println("connected to db")

	driver, err := sqlite.WithInstance(database, &sqlite.Config{})
	if err != nil {
		fmt.Println("could not create driver", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations", // путь относительно рабочей директории
		"sqlite3",
		driver,
	)
	if err != nil {
		fmt.Println("could not create migrate instance", err)
		return nil, err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println("migration failed", err)
		return nil, err
	}

	fmt.Println("migrations applied")

	return database, nil
}
