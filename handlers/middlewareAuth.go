package handlers

import (
	"fmt"
	"net/http"

	"github.com/sidsharma96/go-aggregator/internal/auth"
	"github.com/sidsharma96/go-aggregator/internal/database"
	"github.com/sidsharma96/go-aggregator/util"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			util.RespondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}

