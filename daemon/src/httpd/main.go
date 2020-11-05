package main

import (
	"net/http"
	"socks-manager/src/cmd/httpd/appstate"
	"socks-manager/src/cmd/httpd/environment"
	"socks-manager/src/cmd/httpd/handlers"
)

var env = environment.NewEnvironment()
var state = appstate.NewState(env.ConfigPath)

func main() {
	http.HandleFunc("/", handlers.IndexHandler())
	http.HandleFunc("/proxy-list", handlers.ProxyListHandler(state))
	http.HandleFunc("/proxy-bypass-list", handlers.BypassListHandler(state))
	http.HandleFunc("/proxy-strategy", handlers.ProxyStrategyHandler(state))
	http.HandleFunc("/client-address", handlers.ClientAddressHandler(state))
	http.HandleFunc("/connection", handlers.ConnectionHandler(state))

	http.ListenAndServe(":3000", nil)
}
