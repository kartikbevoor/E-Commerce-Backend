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

	//tx, err := config.DB.Beginx() //config.Server.Db.Begin()
	tx, err := config.DB.Beginx()
	if err != nil {
		log.Println("Failed to start transaction, err:", err)
	}

	result, err := RegisterUser(tx, NewUser)
	if err != nil {
		log.Println("Failed to registor user")
		http.Error(w, "Failed to registor user, err:", http.StatusInternalServerError)
		return
	}

	var CreateUserResponse CreateUserResponse

	CreateUserResponse.ID, err = result.LastInsertId()
	if err != nil {

	}
}
