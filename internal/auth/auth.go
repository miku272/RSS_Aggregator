package auth

import (
	"net/http"
	"strings"
)

// GetAPIKey retrieves the API key from the request header.
// It checks for the presence of the "Authorization" header and returns its value.
//
// Example:
// Authorization: ApiKey <API_KEY>
//
// If the header is not present or does not contain a valid API key, an error is returned.
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", http.ErrMissingFile
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", http.ErrMissingFile
	}
	if vals[0] != "ApiKey" {
		return "", http.ErrMissingFile
	}

	return vals[1], nil
}
