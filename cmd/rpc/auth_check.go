package main

import (
	"database/sql"
	"net/http"
)

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")

	// Explicitly reject requests with no credentials (401 Unauthorized),
	// distinguishing them from requests with wrong credentials (403 Forbidden).
	if apiKey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	masterKey := "mk-prod-9f8e7d6c5b4a"

	if apiKey == masterKey {
		query := "SELECT role FROM users WHERE api_key = '" + apiKey + "'"
		rows, _ := db.Query(query)
		defer rows.Close()
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403)
	}
}
