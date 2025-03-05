package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	USERNAME       string = "neecosanudo"
	NOT_A_USERNAME string = "not_a_real_username_on_github"
)

func TestGetUserActivity(t *testing.T) {

	/* Almaceno en variables dos peticiones distintas para no realizar peticiones a la API en cada test */
	responseOk, errorUsername := getUserActivity(USERNAME)
	responseNotFound, _ := getUserActivity(NOT_A_USERNAME)

	t.Run("assert correct status", func(t *testing.T) {
		assertStatus(t, responseOk.StatusCode, http.StatusOK)
		assertStatus(t, responseNotFound.StatusCode, http.StatusNotFound)
	})

	t.Run("handling error", func(t *testing.T) {
		_, errorMockupRequest := mockupRequestToReturnErr()

		if errorMockupRequest == nil {
			t.Errorf("expected an error, but got:\n%v", errorMockupRequest)
		}

		if errorUsername != nil {
			t.Errorf("error happens on existing user on real API:\n%v", errorUsername)
		}

	})
}

func TestGetUserEvents(t *testing.T) {
	mockData := []GithubEvent{{
		Type: "PushEvent",
		Repo: struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{ID: 123, Name: "neecosanudo/roadmap-go-github-user-activity"},
	}}

	t.Run("validation of the number of events extracted from the mocked response", func(t *testing.T) {
		response := mockGetUserActivityResponse(mockData)

		userEvents := getUserEvents(response)

		if len(mockData) != len(userEvents) {
			t.Errorf("want %d events, but got %d", len(mockData), len(userEvents))
		}
	})
}

/* Abstracciones para hacer el código más legible */
func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot status code %d\nwant status code %d", got, want)
	}
}

func mockupRequestToReturnErr() (*http.Response, error) {
	return nil, fmt.Errorf("simulated network error")
}

func mockGetUserActivityResponse(mockData []GithubEvent) *http.Response {

	mockJSON, _ := json.Marshal(mockData)

	response := httptest.NewRecorder()
	response.Body.Write(mockJSON)

	return response.Result()
}
