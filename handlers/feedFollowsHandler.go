package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/sidsharma96/go-aggregator/internal/database"
	"github.com/sidsharma96/go-aggregator/model"
	"github.com/sidsharma96/go-aggregator/util"
)

func (apiCfg *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Couldn't get feed follow")
		return
	}

	util.Respond(w, http.StatusOK, model.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) HandlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	util.Respond(w, http.StatusCreated, model.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdParam := chi.URLParam(r, "feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIdParam)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid feed follow id %v", feedFollowIdParam))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't delete feed follow with ID %v", feedFollowIdParam))
		return
	}
	
	util.Respond(w, http.StatusOK, struct{}{})
}
