package main

import (
	"fmt"
	"net/http"
)

func getUserActivity(username string) *http.Response {
	response, _ := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))

	return response
}
