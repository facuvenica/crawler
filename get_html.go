package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if res.StatusCode > 400 {
		return "", fmt.Errorf("error %d", res.StatusCode)
	}

	header := res.Header.Get("Content-Type")

	if !strings.Contains(header, "text/html") {
		return "", fmt.Errorf("wrong header: %s", header)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
