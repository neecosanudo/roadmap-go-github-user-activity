package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ERR_FETCH_USER_ACTIVITY string = "error fetching user activity for %s: %w"
)

type GithubEvent struct {
	Type string `json:"type"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"repo"`
}

func getUserActivity(username string) (*http.Response, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))

	if err != nil {
		return nil, fmt.Errorf(ERR_FETCH_USER_ACTIVITY, username, err)
	}

	return response, nil
}

/* Solo trabajo con la respuesta de la petición porque esta capa no decide que hacer con los errores, solo los expone */
func getUserEvents(response *http.Response) []GithubEvent {
	var events []GithubEvent
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	json.Unmarshal([]byte(body), &events)

	return events
}
