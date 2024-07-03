package main

import (
	"net/http"

	"github.com/joshua468/basic-webserver/api"
)

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	api.Handler(w, r)
}
