package main

import (
	"context"
	"github.com/alexlucaci-go/ports-service/cmd/ports-service/handlers"
	"github.com/alexlucaci-go/ports-service/domain/ports"
	"github.com/alexlucaci-go/ports-service/domain/ports/store/inmemorydb"
	"github.com/alexlucaci-go/ports-service/loader"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run() error {
	// Load configuration using any libraries, will just hardcode for this example
	type Config struct {
		Web struct {
			Host            string
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
			ShutdownTimeout time.Duration
		}
		Loader struct {
			Timeout time.Duration
		}
	}

	cfg := Config{
		Web: struct {
			Host            string
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
			ShutdownTimeout time.Duration
		}{
			Host:            "0.0.0.0:8000",
			ReadTimeout:     5 * time.Second,
			WriteTimeout:    10 * time.Second,
			ShutdownTimeout: 5 * time.Second,
		},
		Loader: struct{ Timeout time.Duration }{
			Timeout: 5 * time.Minute,
		},
	}
	// load the initial state of the ports from the json file

	db := inmemorydb.NewInMemoryDB()
	portdomain := ports.NewDomain(db)
	jsonLoader := loader.NewJson(portdomain)

	// configure loader to timeout after a certain time
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Loader.Timeout)
	defer cancel()

	err := jsonLoader.LoadFromFile(ctx, "ports.json")
	if err != nil {
		return errors.Wrap(err, "loading ports from json")
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr:         cfg.Web.Host,
		Handler:      handlers.API(shutdown, db),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Asking listener to shutdown and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
