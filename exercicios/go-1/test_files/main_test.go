package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonify(t *testing.T) {
	err := jsonify("./")
	assert.Nil(t, err)
	assert.FileExists(t, "files.json")
}
