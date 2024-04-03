package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sidsharma96/go-aggregator/internal/database"
	"github.com/sidsharma96/go-aggregator/model"
	"github.com/sidsharma96/go-aggregator/util"
)

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}
	util.Respond(w, http.StatusCreated, model.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context()) 
	if err != nil {
		util.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}
	util.Respond(w, http.StatusOK, model.DatabaseFeedsToFeeds(feeds))
}