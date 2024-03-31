package main

import "net/http"

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, 200, struct{}{})
}