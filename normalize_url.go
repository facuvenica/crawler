package main

import (
	"log"
	"net/url"
	"strings"
)

func normalizeURL(path string) string {

	url, err := url.Parse(path)
	if err != nil {
		log.Fatal("error parsing url")
	}
	fixed := url.Host + url.Path
	return strings.TrimSuffix(fixed, "/")
}
