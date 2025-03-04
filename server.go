package main

import (
	"fmt"
	"net/http"
)

const (
	ERR_FETCH_USER_ACTIVITY string = "error fetching user activity for %s: %w"
)

func getUserActivity(username string) (*http.Response, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))

	if err != nil {
		return nil, fmt.Errorf(ERR_FETCH_USER_ACTIVITY, username, err)
	}

	return response, nil
}
