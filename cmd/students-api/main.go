package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shahsidar-k-s/students-api/internal/config"
	"github.com/shahsidar-k-s/students-api/internal/http/handler/students"
)

func main() {
	// Load the application configuration from environment variables or a YAML file.
	configs := config.MustLoad()

	// Initialize a new HTTP router (ServeMux) to register your API endpoints.
	router := http.NewServeMux()

	// Register HTTP handlers (routes) for student-related APIs.
	// Each handler is responsible for handling a specific type of HTTP request.
	router.HandleFunc("GET /getAllStudents", students.GetAllStudents())
	router.HandleFunc("POST /addNewStudent", students.AddNewStudent())
	router.HandleFunc("PUT /updateStudent", students.UpdateStudent())

	// Create and configure the HTTP server using the address from the config file.
	server := http.Server{
		Addr:    configs.Addr,
		Handler: router,
	}

	// Log that the server is starting.
	slog.Info("server started @", slog.String("address", configs.Addr))

	// Create a channel to listen for OS signals (interrupts or termination requests).
	done := make(chan os.Signal, 1)

	// Notify the channel when an interrupt or termination signal is received.
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine so it doesn't block the shutdown logic.
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Error while starting server, please restart again..!")
		}
	}()

	// Block the main goroutine until a termination signal is received.
	<-done

	// Once a signal is received, log shutdown initiation.
	slog.Info("Shutting down the server==>")

	// Create a context with timeout to allow the server to shutdown gracefully.
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server.
	if errors := server.Shutdown(ctx); errors != nil {
		slog.Error("Failed to shutDown the server==>", slog.String("error", errors.Error()))
	}

	// Log successful shutdown.
	slog.Info("Server shut down successFully")
}
