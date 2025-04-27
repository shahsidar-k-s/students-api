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
)

func main() {
	//loading the config files
	configs := config.MustLoad()
	// setup router
	router := http.NewServeMux()

	//create APIs
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students API world"))
	})
	// setup server
	server := http.Server{
		Addr:    configs.Addr,
		Handler: router,
	}
	slog.Info("server started @", slog.String("address", configs.Addr))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Error while starting error please restart again..!")
		}
	}()
	/// blocked untill the server recive the termination or ctr+c signals from the os once its recive signa.Notify
	// will send message to channle and then it will be continued
	<-done
	// grace shutdow
	slog.Info("Shutting down the server==>")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()
	if errors := server.Shutdown(ctx); errors != nil {
		slog.Error("Failed to shutDown the server==>", slog.String("error", errors.Error()))
	}
	slog.Info("Server shut down successFully")
}
