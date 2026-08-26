package user

import (
	"ecommerce-backend/config"
	"encoding/json"
	"log"
	"net/http"
)

func UserRegistration(w http.ResponseWriter, r *http.Request) {
	var NewUser CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		log.Println("Failed to decode json, err:", err)
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	tx, err := config.Server.Db.Begin()
	if err != nil {
		log.Println("Failed to start transaction, err:", err)
	}

	result, err := RegistorUser(tx, NewUser)
	if err != nil {

	}

}
