package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bagusays/efishery-technical-test/internal/app"
	"github.com/bagusays/efishery-technical-test/internal/config"
	"github.com/bagusays/efishery-technical-test/internal/delivery"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	RunE: runHTTPServer,
}

// Execute :nodoc:
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// runHTTPServer :nodoc:
func runHTTPServer(_ *cobra.Command, _ []string) error {
	cfg := config.New()
	container := app.NewContainer(cfg)

	// - Setup Handler
	server := delivery.New(container)

	errChan := make(chan error, 2)
	defer close(errChan)

	go func() {
		//- start service
		fmt.Println("HTTP server is running on port", cfg.Port)
		if err := server.Start(cfg.AppAddress()); err != nil {
			errChan <- err
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sigCh
		errChan <- fmt.Errorf("getting signal to shutdown the app")
	}()

	fmt.Println("error: ", <-errChan)
	fmt.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		return err
	}

	return nil
}
