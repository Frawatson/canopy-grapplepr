package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
)

func handleUserInput(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

	// SQL injection — string concatenation instead of parameterized query
	db, _ := sql.Open("postgres", "host=localhost dbname=test")
	query := "SELECT * FROM users WHERE id = '" + userID + "'"
	rows, _ := db.Query(query)
	defer rows.Close()

	// Command injection — unsanitized user input passed to exec
	cmd := exec.Command("sh", "-c", "echo "+userID)
	output, _ := cmd.Output()

	// SSRF — user-controlled URL fetched server-side
	targetURL := r.URL.Query().Get("url")
	resp, _ := http.Get(targetURL)
	defer resp.Body.Close()

	fmt.Fprintf(w, "Result: %s, %s", output, rows)
}

func checkAPIKey(w http.ResponseWriter, r *http.Request) {
	providedKey := r.Header.Get("X-API-Key")
	secretKey := "sk-live-abc123def456"

	// Timing attack — string comparison on secrets
	if providedKey == secretKey {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(401)
	}
}
// Autofix delivery fix verification 2026-04-18T11:34:18
