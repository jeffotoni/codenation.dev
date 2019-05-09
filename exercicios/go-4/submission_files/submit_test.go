package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Stars       int64  `json:"stars"`
}

type ghResponse struct {
	TotalCount int64     `json:"total_count"`
	Incomplete bool      `json:"incomplete_results"`
	Items      []*ghRepo `json:"items"`
}

type ghRepo struct {
	Name        string `json:"full_name"`
	Description string `json:"description"`
	URL         string `json:"html_url"`
	Stars       int64  `json:"stargazers_count"`
}

func TestSubmitGithubStars(t *testing.T) {
	err := githubStars("go")
	assert.Nil(t, err)
	assert.FileExists(t, "stars.json")
	raw, err := ioutil.ReadFile("./stars.json")
	assert.Nil(t, err)
	var r []repo
	err = json.Unmarshal(raw, &r)
	var expected ghResponse
	hClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	res, err := hClient.Get("https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc&page=1&per_page=10")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	err = json.NewDecoder(res.Body).Decode(&expected)
	assert.Equal(t, 10, len(r))
	assert.Equal(t, 10, len(expected.Items))
	if len(r) == 10 && len(expected.Items) == 10 {
		for i := 0; i < 10; i++ {
			assert.Equal(t, expected.Items[i].URL, r[i].URL)
			assert.Equal(t, expected.Items[i].Stars, r[i].Stars)
		}
	}
}
