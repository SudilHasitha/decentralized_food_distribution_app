package handlers

import (
	"backend/utils"
	"log"
	http "net/http"

	shell "github.com/ipfs/go-ipfs-api" // Import IPFS Shell library
)

// SendMessageHandler handles sending messages to a specific chat room
func SendMessageHandler(w http.ResponseWriter, r *http.Request, sh *shell.Shell) {
	if r.Method == http.MethodPost {
		roomID := r.FormValue("room_id")  // Get the room ID from the form data
		message := r.FormValue("message") // Get the message content from the form data

		// Add the message to the chat room in IPFS
		err := utils.AddMessageToChatRoom(roomID, message, sh)
		if err != nil {
			http.Error(w, "Failed to send message", http.StatusInternalServerError)
			log.Println("Error sending message:", err)
			return
		}

		// Return the new message as HTML to be added to the messages list dynamically
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<p>` + message + `</p>`))
	}
}
