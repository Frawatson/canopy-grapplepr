package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
)

func unsafeQuery(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	db, _ := sql.Open("postgres", "host=localhost dbname=app")
	query := "SELECT * FROM products WHERE name = '" + name + "'"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	fmt.Fprintf(w, "Results: %v", rows)
}
