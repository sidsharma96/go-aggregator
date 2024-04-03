package handlers

import (
	"net/http"

	"github.com/sidsharma96/go-aggregator/util"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}