package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/esiebomaj/scratch/internal/database"
	"github.com/google/uuid"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	HandleSuccessJson(w, 200, struct{}{})
}

type userData struct {
	Name string `json:"name"`
}

func (apiConfig *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	data := userData{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      data.Name,
	})
	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not create user %v", err))
		return
	}

	HandleSuccessJson(w, 201, user)
}

func (apiConfig *ApiConfig) GetAllUser(w http.ResponseWriter, r *http.Request) {

	users, err := apiConfig.DB.GetAllUser(r.Context())
	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not retrieve users %v", err))
		return
	}
	HandleSuccessJson(w, 200, users)
}

func (apiConfig *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	HandleSuccessJson(w, 200, user)
}
