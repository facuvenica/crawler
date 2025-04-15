package main

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetURLs(t *testing.T) {

	// Absolute and Relative URLs
	path := "https://blog.boot.dev"
	body := `<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>`
	num := strings.Count(body, "<a href")
	result, err := getURLsFromHTML(body, path)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, num, len(result))
	assert.ElementsMatch(t, []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"}, result)
	for _, elem := range result {
		log.Println(elem)
	}

	// Missing url
	path = ""
	body = `<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>`
	_, err = getURLsFromHTML(body, path)
	require.Error(t, err)
	log.Println(err)
	log.Println(body)
}
