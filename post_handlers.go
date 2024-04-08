package main

import (
	"fmt"
	"net/http"

	"github.com/esiebomaj/rss-aggregator/internal/database"
)

func (apiConfig *ApiConfig) GetRecentPosts(w http.ResponseWriter, r *http.Request, user database.User) {

	posts, err := apiConfig.DB.GetRecentPostForUser(r.Context(), database.GetRecentPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not retrieve posts %v", err))
		return
	}
	HandleSuccessJson(w, 200, posts)
}
