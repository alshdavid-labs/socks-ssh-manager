package handlers

import (
	"encoding/json"
	"net/http"
	"socks-manager/src/cmd/httpd/appstate"
)

type ProxyStrategyBody struct {
	ProxyStrategy string `json:"proxyStrategy" json`
}

type ProxyStrategyResponse struct {
	ProxyStrategy string `json:"proxyStrategy" json`
}

func ProxyStrategyHandler(state *appstate.State) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ProxyStrategyResponse{
				ProxyStrategy: state.ProxyStrategy,
			})
			return
		}

		var b ProxyStrategyBody
		json.NewDecoder(r.Body).Decode(&b)
		state.SetProxyStrategy(b.ProxyStrategy)
		w.WriteHeader(http.StatusNoContent)
	}
}
