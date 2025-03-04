package main

import (
	"net/http"
	"testing"
)

const (
	NOT_A_USERNAME string = "not_a_real_username_on_github"
)

func TestGetUserActivity(t *testing.T) {
	t.Run("assert status Not Found", func(t *testing.T) {
		response := getUserActivity(NOT_A_USERNAME)

		assertStatus(t, response.StatusCode, http.StatusNotFound)
	})
}

/* Abstracciones para hacer el código más legible */
func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot status code %d\nwant status code %d", got, want)
	}
}
