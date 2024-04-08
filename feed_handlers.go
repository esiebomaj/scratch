package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/esiebomaj/scratch/internal/database"
	"github.com/google/uuid"
)

type Params struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (apiConfig *ApiConfig) CreateFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	params := Params{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not create feed %v", err))
		return
	}

	_, er := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if er != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not follow new feed %v", err))
		return
	}

	HandleSuccessJson(w, 201, feed)
}

func (apiConfig *ApiConfig) GetAllFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiConfig.DB.GetAllFeeds(r.Context())
	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not retrieve feeds %v", err))
		return
	}
	HandleSuccessJson(w, 200, feeds)
}
