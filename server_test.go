package main

import (
	"net/http"
	"testing"
)

const (
	USERNAME       string = "neecosanudo"
	NOT_A_USERNAME string = "not_a_real_username_on_github"
)

func TestGetUserActivity(t *testing.T) {

	/* Almaceno en variables dos peticiones distintas para no realizar peticiones a la API en cada test */
	responseOk := getUserActivity(USERNAME)
	responseNotFound := getUserActivity(NOT_A_USERNAME)

	t.Run("assert correct status", func(t *testing.T) {
		assertStatus(t, responseOk.StatusCode, http.StatusOK)
		assertStatus(t, responseNotFound.StatusCode, http.StatusNotFound)
	})
}

/* Abstracciones para hacer el código más legible */
func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot status code %d\nwant status code %d", got, want)
	}
}
