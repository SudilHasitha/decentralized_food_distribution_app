package main

import (
	"backend/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell // Initialize IPFS Shell globally

func main() {
	// Connect to local IPFS node
	sh = shell.NewShell("localhost:5001")

	// Register handlers
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/create-room", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateLocationRoomHandler(w, r, sh)
	})
	http.HandleFunc("/send-message", func(w http.ResponseWriter, r *http.Request) {
		handlers.SendMessageHandler(w, r, sh)
	})
	http.HandleFunc("/create-location-room", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateLocationRoomHandler(w, r, sh) // Pass IPFS Shell to handler
	})
	http.HandleFunc("/map", handlers.MapHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./client/static/"))))

	srv := &http.Server{
		Addr: ":8080",
	}

	go func() {
		log.Println("Starting server on :8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
