package auth

import (
	"errors"
	"net/http"
	"strings"
)

//Extracts API key from the headers of an HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authentication provided")
	}
	splitHeader := strings.Split(authHeader, " ")
	if len(splitHeader) != 2 || splitHeader[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitHeader[1], nil
}