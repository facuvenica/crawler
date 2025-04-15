package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeURL(t *testing.T) {

	// Remove Scheme
	path := "https://blog.boot.dev/path"
	result := normalizeURL(path)
	require.NotNil(t, result)
	assert.Equal(t, "blog.boot.dev/path", result)

	// Remove trailing slash
	path = "http://another.url/path/"
	result = normalizeURL(path)
	require.NotNil(t, result)
	assert.Equal(t, "another.url/path", result)

}
