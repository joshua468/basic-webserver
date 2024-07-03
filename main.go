package main

import (
	"net/http"

	"github.com/joshua468/basic-webserver/api"
)

func main() {
	http.HandleFunc("/api/hello", api.Handler)
	http.ListenAndServe(":3000", nil)
}
