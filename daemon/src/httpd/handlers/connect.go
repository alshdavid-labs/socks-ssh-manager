package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socks-manager/src/httpd/appstate"
	"socks-manager/src/platform/socks5client"
	"socks-manager/src/platform/socks5server"
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
				Connected: state.IsConnected(),
			})
			return
		}

		var b ConnectionBody
		json.NewDecoder(r.Body).Decode(&b)

		if b.Action == "CONNECT" {
			state.Connect()

			dialSocksProxy := socks5client.DialCtx(fmt.Sprintf("socks5://127.0.0.1:%v", state.GetConnectionPort()))

			conf := &socks5server.Config{
				Dial:  dialSocksProxy,
				State: state,
			}

			server, _ := socks5server.New(conf)

			if err := server.ListenAndServe("tcp", "127.0.0.1:1337"); err != nil {
				panic(err)
			}
		}

		if b.Action == "DISCONNECT" {
			state.Disconnect()
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
