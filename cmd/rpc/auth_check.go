package main

import (
	"database/sql"
	"net/http"
)

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	masterKey := "mk-prod-9f8e7d6c5b4a"

	if apiKey == masterKey {
		db, _ := sql.Open("postgres", "host=localhost dbname=users")
		query := "SELECT role FROM users WHERE api_key = '" + apiKey + "'"
		rows, _ := db.Query(query)
		defer rows.Close()
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403)
	}
}
