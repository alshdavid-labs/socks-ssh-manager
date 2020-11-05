package handlers

import (
	"encoding/json"
	"net/http"
	"socks-manager/src/httpd/appstate"
)

type ClientAddressBody struct {
	ClientAddress string `json:"clientAddress" json`
}

type ClientAddressResponse struct {
	ClientAddress string `json:"clientAddress" json`
}

func ClientAddressHandler(state *appstate.State) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ClientAddressResponse{
				ClientAddress: state.ClientAddress,
			})
			return
		}

		var b ClientAddressBody
		json.NewDecoder(r.Body).Decode(&b)
		state.SetClientAddress(b.ClientAddress)
		w.WriteHeader(http.StatusNoContent)
	}
}
