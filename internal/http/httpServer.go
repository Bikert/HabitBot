package http

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func NewHttpServer(lc fx.Lifecycle, router *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("http server listening on :8080")
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Starting HTTP server on :8080")
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("Listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping HTTP server on :8080")
			return server.Shutdown(ctx)
		},
	})
	return server
}
