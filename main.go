package main

import (
	"net/http"

	"github.com/joshua468/basic-webServer/api"
)

func main() {
	http.HandleFunc("/api/hello", api.Handler)
	http.ListenAndServe(":3000", nil)
}
