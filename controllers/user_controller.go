package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"s3-microservice/services"
)

type UserController struct {
	userService services.UserService
}

func UserControllerInstance() *UserController {
	return &UserController{}
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
