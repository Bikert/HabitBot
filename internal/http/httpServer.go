package http

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type Server struct {
	router *gin.Engine
	server *http.Server
}

func NewHttpServer(router *gin.Engine) *Server {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	return &Server{
		router: router,
		server: server,
	}
}

func RunHttpServer(lc fx.Lifecycle, s *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Starting HTTP server on :8080")
				if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("Listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping HTTP server on :8080")
			return s.server.Shutdown(ctx)
		},
	})
}
