package main

import (
	"database/sql"
	"net/http"
)

type user struct {
	ID    int
	Name  string
	Email string
}

func main() {
	http.ListenAndServe("8082", nil)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
