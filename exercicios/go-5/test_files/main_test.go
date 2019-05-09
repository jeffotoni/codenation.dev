package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHTML(t *testing.T) {
	err := parseHTML("golang.html")
	assert.Nil(t, err)
	assert.FileExists(t, "golang.json")
}
