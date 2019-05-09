package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type parsed struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	CreatedBy     string   `json:"created_by"`
	ReleasedAt    string   `json:"released_at"`
	Repositories  int64    `json:"repositories"`
	RelatedTopics []string `json:"related_topics"`
}

func TestSubmitParseHTML(t *testing.T) {
	err := parseHTML("golang.html")
	assert.Nil(t, err)
	assert.FileExists(t, "golang.json")
	raw, err := ioutil.ReadFile("./golang.json")
	assert.Nil(t, err)
	var r, expected parsed
	err = json.Unmarshal(raw, &r)
	expected = parsed{
		Name:          "Go",
		Description:   "Go is a programming language built to resemble a simplified version of the C programming language. It compiles at the machine level. Go was created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson.",
		CreatedBy:     "Robert Griesemer, Rob Pike, Ken Thompson",
		ReleasedAt:    "November 10, 2009",
		Repositories:  21022,
		RelatedTopics: []string{"docker", "bot", "godoc", "database", "language", "python", "marbles", "heroku"},
	}
	assert.Equal(t, expected, r)
}
