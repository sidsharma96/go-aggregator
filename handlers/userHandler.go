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

func (apiCfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	util.Respond(w, http.StatusCreated, model.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	userIdsAndKeys, err := apiCfg.DB.GetAllUserIdsAndKeys(r.Context())
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	util.Respond(w, http.StatusOK, model.DatabaseUsersToIds(userIdsAndKeys))
}

func (apiCfg *ApiConfig) HandleAuthorizedUser(w http.ResponseWriter, r *http.Request, user database.User) {
	util.Respond(w, http.StatusOK, model.DatabaseUserToUser(user))
}