package handlers

import (
	"encoding/json"
	"net/http"
	"socks-manager/src/cmd/httpd/appstate"
)

type ConnectionBody struct {
	Action string `json:"action" json`
}

type ConnectionResponse struct {
	Connected bool `json:"status" json`
}

func ConnectionHandler(state *appstate.State) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ConnectionResponse{
				Connected: state.Connected,
			})
			return
		}

		var b ConnectionBody
		json.NewDecoder(r.Body).Decode(&b)

		if b.Action == "CONNECT" {
			state.Connected = true
		}

		if b.Action == "DISCONNECT" {
			state.Connected = false
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
