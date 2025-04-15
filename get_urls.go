package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	if rawBaseURL == "" {
		return nil, fmt.Errorf("missing base url")
	}

	urls := make([]string, 0)
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link, err := url.Parse(attr.Val)
					if err != nil {
						return nil, err
					}

					if link.Host == "" {
						link.Host = rawBaseURL
					}
					urls = append(urls, link.String())
				}
			}
		}
	}

	return urls, nil
}
