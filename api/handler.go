package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

// DB_PASSWORD should be set via the DB_PASSWORD environment variable.
// Retrieve it at runtime with os.Getenv("DB_PASSWORD") — never hardcode credentials.

// GetUser fetches a user by ID from the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	userID := r.URL.Query().Get("id")

	// SQL injection: user input directly interpolated into query
	query := fmt.Sprintf("SELECT name, email FROM users WHERE id = '%s'", userID)
	row := db.QueryRow(query)

	var name, email string
	row.Scan(&name, &email)

	// Missing error handling on Scan
	fmt.Fprintf(w, "User: %s (%s)", name, email)
}

// DeleteUser removes a user without authorization check
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	userID := r.URL.Query().Get("id")

	// SQL injection + no auth check + no error handling
	query := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", userID)
	db.Exec(query)

	fmt.Fprintf(w, "Deleted user %s", userID)
}
