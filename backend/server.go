package main

import (
	"log"
	"net/http"
	"decentralized-chat-app/backend/handlers"
	"github.com/ipfs/go-ipfs-api" // Import IPFS Shell library
)

var sh *shell.Shell // Initialize IPFS Shell globally

func main() {
	// Connect to local IPFS node
	sh = shell.NewShell("localhost:5001")

	// Register handlers
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/create-room", handlers.CreateRoomHandler)
	http.HandleFunc("/send-message", func(w http.ResponseWriter, r *http.Request) {
		handlers.SendMessageHandler(w, r, sh) // Pass IPFS Shell to handler
	})
	http.HandleFunc("/create-location-room", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateLocationRoomHandler(w, r, sh) // Pass IPFS Shell to handler
	})
	http.HandleFunc("/map", handlers.MapHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./client/static/"))))

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
