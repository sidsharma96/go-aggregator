package handlers

import (
	"net/http"

	"github.com/sidsharma96/go-aggregator/util"
)

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	util.Respond(w, 200, struct{}{})
}