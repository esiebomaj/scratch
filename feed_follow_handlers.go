package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/esiebomaj/scratch/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiConfig *ApiConfig) FollowFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := Params{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	follow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not follow this feed %v", err))
		return
	}

	HandleSuccessJson(w, 201, follow)
}

func (apiConfig *ApiConfig) GetUserFollowedFeeds(w http.ResponseWriter, r *http.Request, user database.User) {

	follows, err := apiConfig.DB.GetAllUserFeedFollows(r.Context(), user.ID)
	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not retrieve follows %v", err))
		return
	}
	HandleSuccessJson(w, 200, follows)
}

func (apiConfig *ApiConfig) UnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follow_id_s := chi.URLParam(r, "feedFollowID")
	feed_follow_id, err := uuid.Parse(feed_follow_id_s)

	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("invalid feed id %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollowsByID(r.Context(), database.DeleteFeedFollowsByIDParams{
		ID:     feed_follow_id,
		UserID: user.ID,
	})

	if err != nil {
		HandleErrorJson(w, 400, fmt.Sprintf("Could not unfollow feed %v", err))
		return
	}

	HandleSuccessJson(w, 200, struct{}{})
}
