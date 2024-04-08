package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/esiebomaj/scratch/internal/database"
)

type AuthHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiConfig *ApiConfig) AuthMiddleWare(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		if apiKey == "" {
			HandleErrorJson(w, 401, "Authorization header required")
			return
		}

		splits := strings.Split(apiKey, " ")
		if len(splits) != 2 || splits[0] != "Bearer" {
			HandleErrorJson(w, 401, "Authorization header malformed")
			return
		}

		key := splits[1]

		user, err := apiConfig.DB.GetUser(r.Context(), key)
		if err != nil {
			HandleErrorJson(w, 400, fmt.Sprintf("Could not retrieve user %v", err))
			return
		}
		handler(w, r, user)
	}
}
