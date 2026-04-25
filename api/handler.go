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

	// Use a parameterized query to prevent SQL injection (mirrors DeleteUser pattern)
	row := db.QueryRow("SELECT name, email FROM users WHERE id = $1", userID)

	var name, email string
	row.Scan(&name, &email)

	// Missing error handling on Scan
	fmt.Fprintf(w, "User: %s (%s)", name, email)
}

// DeleteUser removes a user without authorization check
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Enforce POST-only to prevent CSRF and accidental deletion via browser links
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Authorization: require a valid Bearer token
	expectedToken := os.Getenv("AUTH_TOKEN")
	if expectedToken == "" {
		// Log the detailed reason server-side; never expose config details to callers
		log.Println("ERROR: AUTH_TOKEN environment variable is not set")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	authHeader := r.Header.Get("Authorization")
	if authHeader != "Bearer "+expectedToken {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validate the id parameter
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Bad Request: missing id parameter", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Use a parameterized query to prevent SQL injection
	result, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Not Found: no user with that id", http.StatusNotFound)
		return
	}

	// Set explicit Content-Type to prevent browser sniffing the body as text/html,
	// which would allow reflected XSS via a crafted id parameter.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Deleted user %s", userID)
}
