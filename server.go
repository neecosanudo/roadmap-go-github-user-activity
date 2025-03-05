package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	ERR_FETCH_USER_ACTIVITY string = "error fetching user activity for %s: %w"

	PUSH_EVENT_MESSAGE string = "Pushed %d commits to %s\n"
)

type GithubEvent struct {
	Type string `json:"type"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"repo"`
}

type summaryHelper struct {
	counter          int
	currentRepo      string
	currentEventType string
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

func formatActivitySummary(username string, events []GithubEvent) string {
	/* El usuario no tiene eventos */
	if len(events) == 0 {
		return fmt.Sprintf("User @%s has no activity", username)
	}

	var text strings.Builder

	/*
		El ejercicio pide este tipo de output:
		- Pushed 3 commits to kamranahmedse/developer-roadmap

		Recorro el slice de eventos para contar cuantas veces consecutivas
		ocurre el mismo tipo de evento en el mismo repositorio.

		Bien ocurre un cambio, utiliza la información en summaryHelper
		para crear el mensaje correspondiente a ese tipo de evento.

		La información en summaryHelper se reinicia con los nuevos cambios
		y repite el ciclo.

		De esta forma tengo un contador de eventos en X repositorio
		por cada evento consecutivo.

	*/
	sh := summaryHelper{
		counter:          0,
		currentRepo:      events[0].Repo.Name,
		currentEventType: events[0].Type,
	}

	title := fmt.Sprintf("Activity on GitHub from user @%s:\n", username)
	text.WriteString(title)

	for _, event := range events {
		if event.Repo.Name == sh.currentRepo && event.Type == sh.currentEventType {
			sh.counter++
			continue
		}

		// Escribo con la información actual en el summaryHelper con el strings.Builder
		newMessage := findCorrectMessage(sh)
		text.WriteString(newMessage)

		// Reinicio de eventos
		sh.counter = 1
		sh.currentRepo = event.Repo.Name
		sh.currentEventType = event.Type

	}

	return text.String()
}

func findCorrectMessage(sh summaryHelper) string {
	text := ""

	// Todo: pendiente de completar con todos los tipos de eventos y manejar el caso de un evento no registrado
	switch sh.currentEventType {
	case "PushEvent":
		text = fmt.Sprintf(PUSH_EVENT_MESSAGE, sh.counter, sh.currentRepo)
	}

	return text
}
