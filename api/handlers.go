package api

import (
	"net/http"
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("0.0.1"))
}

func PongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
