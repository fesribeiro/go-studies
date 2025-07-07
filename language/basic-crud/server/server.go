package server

import (
	"basic-crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Here you would typically parse the request body to get user data,
	// validate it, and then insert it into the database.
	// For now, we'll just send a placeholder response.

	body, err := io.ReadAll(r.Body)

	if err != nil {
		// http.Error(w, "Failed to read request body", http.StatusBadRequest)
		w.Write([]byte("Failed to read request body"))
		return
	}

	var user user

	if err := json.Unmarshal(body, &user); err != nil {
		// http.Error(w, "Invalid request body", http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	db, err := db.Connect()

	if err != nil {
		// http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		w.Write([]byte("Failed to connect to database"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		// http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		w.Write([]byte("Failed to prepare statement"))
		return
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email)
	if err != nil {
		// http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		w.Write([]byte("Failed to execute statement"))
		return
	}

	
	fmt.Printf("User created: %+v\n", user)
	
}
