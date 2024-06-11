package webserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// file go para abstrair o servidor web :)

const TIMEOUT = 30 * time.Second

type ServerOption func(server *http.Server)

// starting server with gracefully shutdown
func Start(port string, handler http.Handler, options ...ServerOption) {

	server := &http.Server{
		ReadTimeout:  TIMEOUT,
		WriteTimeout: TIMEOUT,
		Addr:         ":" + port,
		Handler:      handler,
	}

	for _, o := range options {
		o(server)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server http")
		}
	}()
		
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Shutting Down server http")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("error shutting down server http")
	}
	fmt.Println("Server stopped")
}

// WithReadTimeout configure http.Server parameter ReadTimeout
func WithReadTimeout(t time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.ReadTimeout = t
	}
}

// WithWriteTimeout configure http.Server parameter WriteTimeout
func WithWriteTimeout(t time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.WriteTimeout = t
	}
}
