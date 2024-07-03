package handler

import (
	"net/http"

	"github.com/joshua468/basic-webserver/api"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    api.Handler(w, r)
}
