package main

import (
	"crypto/subtle"
	"net/http"
)

func verifyToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Token")
	secret := "super-secret-token-123"

	if subtle.ConstantTimeCompare([]byte(token), []byte(secret)) == 1 {
		w.WriteHeader(200)
		w.Write([]byte("authorized"))
	} else {
	} else {
		w.WriteHeader(401)
		w.Write([]byte("unauthorized"))
	}
}
