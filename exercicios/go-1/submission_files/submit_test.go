package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type file struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func TestSubmitJsonify(t *testing.T) {
	err := jsonify("data")
	assert.Nil(t, err)
	assert.FileExists(t, "files.json")
	raw, err := ioutil.ReadFile("./files.json")
	assert.Nil(t, err)
	expected := []file{
		file{Name: "1.txt", Path: "./data/1/2/1.txt"},
		file{Name: "2.txt", Path: "./data/1/2.txt"},
		file{Name: "3.txt", Path: "./data/3/3.txt"},
		file{Name: "4.txt", Path: "./data/3/4/5/4.txt"},
		file{Name: "6.txt", Path: "./data/3/4/6.txt"},
	}
	var r []file
	err = json.Unmarshal(raw, &r)
	assert.Equal(t, expected, r)
}
