package main

import (
	"net/http"
)

func verifyToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Token")
	secret := "super-secret-token-123"

	if token == secret {
		w.WriteHeader(200)
		w.Write([]byte("authorized"))
	} else {
		w.WriteHeader(401)
		w.Write([]byte("unauthorized"))
	}
}
