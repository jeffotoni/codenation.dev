package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGithubStars(t *testing.T) {
	err := githubStars("go")
	assert.Nil(t, err)
	assert.FileExists(t, "stars.json")
}
