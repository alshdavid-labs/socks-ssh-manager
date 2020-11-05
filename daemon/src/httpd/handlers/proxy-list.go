package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socks-manager/src/httpd/appstate"
	"sort"
)

type ProxyListBody struct {
	Domain string `json:"domain" json`
}

func ProxyListHandler(state *appstate.State) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			filtered := []string{}
			for domain, value := range state.ProxyList {
				if value == true {
					filtered = append(filtered, domain)
				}
			}
			sort.Strings(filtered)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(filtered)
			return
		}

		var b ProxyListBody
		json.NewDecoder(r.Body).Decode(&b)

		if b.Domain == "" {
			w.WriteHeader(http.StatusNoContent)
		}

		if r.Method == "PUT" {
			state.PutProxyList(b.Domain)
			fmt.Println("Set " + b.Domain)
		}

		if r.Method == "DELETE" {
			state.DeleteProxyList(b.Domain)
			fmt.Println("Unset " + b.Domain)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
