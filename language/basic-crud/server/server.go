package server

import (
	"basic-crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	created, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		// http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		w.Write([]byte("Failed to execute statement"))
		return
	}
	
	idCreated, err := created.LastInsertId()
	if err != nil {
		// http.Error(w, "Failed to get last insert ID", http.StatusInternalServer
		w.Write([]byte("Failed to get last insert ID"))
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created with ID: %d", idCreated)))
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}
}


func GetUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, err := strconv.ParseUint(param["id"], 10, 32)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = ?", ID)
	
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}
	
	defer row.Close()
	
	var user user

	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, err := strconv.ParseUint(param["id"], 10, 32)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var user user
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := db.Connect()

	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusBadRequest)
		return
	}

	defer db.Close()

	hasUser, err := db.Query("SELECT id FROM users WHERE id = ?", ID)

	if err != nil {
		http.Error(w, "Erro query check user", http.StatusInternalServerError)
		return
	}

	if !hasUser.Next() {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}	

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		http.Error(w, "Erro query user", http.StatusInternalServerError)
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}