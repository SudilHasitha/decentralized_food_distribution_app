package utils

import (
	"fmt"
	"log"
	"strings"

	shell "github.com/ipfs/go-ipfs-api" // Import IPFS Shell library
)

// GenerateRoomID generates a unique room ID based on latitude and longitude
func GenerateRoomID(lat, lng float64) string {
	return fmt.Sprintf("%f_%f", lat, lng)
}

// CreateChatRoom initializes a new location-based chat room using IPFS
func CreateChatRoom(roomID string, sh *shell.Shell) error {
	// Example: Add initial room data to IPFS
	_, err := sh.Add(strings.NewReader(fmt.Sprintf("Chat Room: %s", roomID)))
	if err != nil {
		log.Printf("Error creating IPFS entry for room %s: %v", roomID, err)
		return err
	}
	log.Printf("Created location-based chat room: %s", roomID)
	return nil
}

// AddMessageToChatRoom adds a new message to the specified chat room in IPFS
func AddMessageToChatRoom(roomID, message string, sh *shell.Shell) error {
	// Example: Add message to IPFS
	_, err := sh.Add(strings.NewReader(message))
	if err != nil {
		log.Printf("Error adding message to room %s: %v", roomID, err)
		return err
	}
	log.Printf("Added message to room %s: %s", roomID, message)
	return nil
}
