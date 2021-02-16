package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type server struct {
	srv *http.Server
}

func newServer() *server {
	return &server{
		srv: &http.Server{
			Addr: ":8080",
		},
	}
}

func (s *server) setRoutes() {
	r := gin.Default()
	r.GET("/health", requestTimeMiddleware(healthCheck))

	s.srv.Handler = r
}

func (s *server) run() error {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		return fmt.Errorf("error while running server: %w", err)
	}

	return nil
}

func requestTimeMiddleware(h gin.HandlerFunc) gin.HandlerFunc {
	v := "REQUEST TIME"

	return func(ctx *gin.Context) {
		now := time.Now()

		h(ctx)

		fmt.Printf("%s: %dμs\n", v, time.Since(now).Microseconds())
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "hello world"})
}

func main() {
	srv := newServer()
	srv.setRoutes()

	if err := srv.run(); err != nil {
		log.Fatal(err)
	}
}
