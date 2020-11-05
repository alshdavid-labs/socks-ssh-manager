package handlers

import (
	"io/ioutil"
	"net/http"
)

func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data, _ := ioutil.ReadFile("/Users/alshdavid/Development/alshdavid/socks-ssh-manager/daemon/src/httpd/handlers/static/index.html")
		w.Write(data)
	}
}

var indexPage = /* html */ `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Sox5 SSH Manager</title>
</head>
<body>
  test
</body>
</html>
`
