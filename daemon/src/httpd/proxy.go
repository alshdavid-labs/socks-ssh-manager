package main

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"socks-manager/src/cmd/cli/handlers"
// 	"socks-manager/src/platform/socks5client"
// 	"socks-manager/src/platform/socks5server"
// )

// var proxyList map[string]bool = make(map[string]bool)

// func main() {
// 	http.HandleFunc("/", handlers.IndexHandler())
// 	http.HandleFunc("/proxy-list", func(w http.ResponseWriter, r *http.Request) {
// 		var b map[string]string
// 		json.NewDecoder(r.Body).Decode(&b)
// 		domain := b["domain"]

// 		if r.Method == "PUT" {
// 			proxyList[domain] = true
// 			fmt.Println("Set " + domain)
// 		}

// 		if r.Method == "DELETE" {
// 			proxyList[domain] = false
// 			fmt.Println("Unset " + domain)
// 			fmt.Println(proxyList)
// 		}

// 		w.WriteHeader(http.StatusNoContent)
// 		w.Write([]byte{})
// 	})
// 	go http.ListenAndServe(":3000", nil)

// 	dialSocksProxy := socks5client.DialCtx("socks5://127.0.0.1:1339")

// 	conf := &socks5server.Config{
// 		Dial:      dialSocksProxy,
// 		ProxyList: proxyList,
// 	}

// 	server, err := socks5server.New(conf)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := server.ListenAndServe("tcp", "127.0.0.1:1337"); err != nil {
// 		panic(err)
// 	}
// }
