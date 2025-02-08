package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kurdilesmana/backend-product-service/deps"
	"github.com/kurdilesmana/backend-product-service/server"
)

func ExecuteHTTP(dependency deps.Dependency) {
	handler := server.SetupHandler(dependency)
	httpServer := server.Http(handler, dependency.Logger, dependency.Cfg.AppConfig)

	// Start HTTP server
	go func() {
		if err := httpServer.Start(fmt.Sprintf(":%d", dependency.Cfg.AppConfig.Port)); err != nil {
			log.Fatalf("Failed to listen on port %d, err: %v", dependency.Cfg.AppConfig.Port, err)
		}
	}()

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Error during HTTP server shutdown: %v", err)
	}
	log.Println("HTTP server gracefully stopped.")
}
