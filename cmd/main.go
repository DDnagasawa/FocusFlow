package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/DDnagasawa/FocusFlow/internal/repository"
	"github.com/gorilla/mux"
	"github.com/yDDnagasawa/FocusFlow/internal/service"
)

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/adhd_app_db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Create repositories
	userRepo := repository.NewMySQLUserRepository(db)

	// Create services
	userService := service.NewUserService(userRepo)

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/users", getUsersHandler(userService)).Methods("GET")
	r.HandleFunc("/users", createUserHandler(userService)).Methods("POST")
	r.HandleFunc("/users/{id}", getUserHandler(userService)).Methods("GET")
	r.HandleFunc("/users/{id}", updateUserHandler(userService)).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUserHandler(userService)).Methods("DELETE")

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getUsersHandler(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

func createUserHandler(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user service.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdUser, err := userService.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdUser)
	}
}

func getUserHandler(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		user, err := userService.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

func updateUserHandler(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var user service.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedUser, err := userService.UpdateUser(id, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedUser)
	}
}

func deleteUserHandler(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		err := userService.DeleteUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
