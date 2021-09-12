package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abba5/sample-user/models"
)

func StartServerWithGracefulShutdown(server models.Server) {
	// Initialize server.
	httpServer := http.Server{
		Handler: server.Router,
		Addr:    fmt.Sprintf("%s:%s", server.Config.ServerHost, server.Config.ServerPort),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Server is starting at %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Server is shutting down...")
	os.Exit(0)
}
