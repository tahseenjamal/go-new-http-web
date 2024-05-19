package controllers

import (
	"encoding/json"
	"net/http"

	"go-new-http-web/common"
)

type User struct{}

func (user *User) GetAll(w http.ResponseWriter, r *http.Request) {
	// Get all user
	response := common.Response{
		Status:  http.StatusOK,
		Message: map[string]interface{}{"message": "Get all user"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response.Message)
}

func (user *User) Get(w http.ResponseWriter, r *http.Request) {
	// Get a user
	response := common.Response{
		Status:  http.StatusOK,
		Message: map[string]interface{}{"message": "Get a user"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response.Message)
}

func (user *User) Post(w http.ResponseWriter, r *http.Request) {
	// Create user
	response := common.Response{
		Status:  http.StatusOK,
		Message: map[string]interface{}{"message": "Create a user"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response.Message)
}

func (user *User) Put(w http.ResponseWriter, r *http.Request) {
	// Update user
	response := common.Response{
		Status:  http.StatusOK,
		Message: map[string]interface{}{"message": "Update a user"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response.Message)
}

func (user *User) Delete(w http.ResponseWriter, r *http.Request) {
	// Delete user
	response := common.Response{
		Status:  http.StatusOK,
		Message: map[string]interface{}{"message": "Delete a user"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response.Message)
}
