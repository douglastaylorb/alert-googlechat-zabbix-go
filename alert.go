package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func enviaAlertaGoogleChat(url string, message string) error {
	requestBody, err := json.Marshal(map[string]string{
		"text": message,
	})
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	url := os.Args[1]
	message := os.Args[2]

	err := enviaAlertaGoogleChat(url, message)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
