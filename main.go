package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	EnvTeamsWebhook = "TEAMS_WEBHOOK"
	EnvTeamsTitle   = "TEAMS_TITLE"
	EnvTeamsMessage = "TEAMS_MESSAGE"
	EnvTeamsColor   = "TEAMS_COLOR"
)

type Webhook struct {
	Text       string `json:"text,omitempty"`
	Title      string `json:"title,omitempty"`
	ThemeColor string `json:"themeColor,omitempty"`
}

func main() {
	endpoint := os.Getenv(EnvTeamsWebhook)
	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "URL is required")
		os.Exit(1)
	}
	text := os.Getenv(EnvTeamsMessage)
	if text == "" {
		fmt.Fprintln(os.Stderr, "Message is required")
		os.Exit(1)
	}

	// Reference fields https://docs.microsoft.com/en-us/outlook/actionable-messages/card-reference

	msg := Webhook{
		Title:      os.Getenv(EnvTeamsTitle),
		Text:       text,
		ThemeColor: os.Getenv(EnvTeamsColor),
	}

	if err := send(endpoint, msg); err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %s\n", err)
		os.Exit(2)
	}
}

func send(endpoint string, msg Webhook) error {
	enc, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(enc)
	res, err := http.Post(endpoint, "application/json", b)
	if err != nil {
		return err
	}

	if res.StatusCode >= 299 {
		return fmt.Errorf("Error on message: %s\n", res.Status)
	}
	fmt.Println(res.Status)
	return nil
}
