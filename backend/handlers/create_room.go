package handlers

import (
	"backend/utils"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	shell "github.com/ipfs/go-ipfs-api" // Import IPFS Shell library
)

// CreateLocationRoomHandler handles creating chat rooms based on map location clicks
func CreateLocationRoomHandler(w http.ResponseWriter, r *http.Request, sh *shell.Shell) {
	if r.Method == http.MethodPost {
		var data struct {
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
			Address string  `json:"address"`
		}

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		// Create a unique room ID based on latitude and longitude
		roomID := utils.GenerateRoomID(data.Lat, data.Lng)

		// Create a chat room in IPFS
		err = utils.CreateChatRoom(roomID, sh)
		if err != nil {
			http.Error(w, "Failed to create location-based chat room", http.StatusInternalServerError)
			log.Println("Error creating location-based chat room:", err)
			return
		}

		// Return the new room as HTML
		w.Header().Set("Content-Type", "text/html")
		tmpl := template.Must(template.New("room-entry").Parse(`<div><a href="/chat/{{.RoomID}}">{{.Address}}</a></div>`))
		tmpl.Execute(w, struct {
			RoomID  string
			Address string
		}{RoomID: roomID, Address: data.Address})
	}
}
