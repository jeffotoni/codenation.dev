package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type Quote struct {
	Actor string `json:"actor"`
	Quote string `json:"quote"`
}

func TestSubmitQuote(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/quote", quote())
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/v1/quote")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var result Quote
	var expected string
	json.NewDecoder(resp.Body).Decode(&result)
	assert.NotEqual(t, "", result.Quote)
	db, err := sql.Open("sqlite3", "./database.sqlite")
	assert.Nil(t, err)
	rows, err := db.Query("SELECT detail FROM scripts where detail=?", result.Quote)
	assert.Nil(t, err)
	rows.Next()
	rows.Scan(&expected)
	assert.Equal(t, expected, result.Quote)
}

func TestSubmitQuoteByActor(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/quote/{actor}", quoteByActor())
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/v1/quote/John+Cleese")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var result Quote
	var expected string
	json.NewDecoder(resp.Body).Decode(&result)
	assert.NotEqual(t, "", result.Quote)
	db, err := sql.Open("sqlite3", "./database.sqlite")
	assert.Nil(t, err)
	rows, err := db.Query("SELECT detail FROM scripts where detail=? and actor='John Cleese'", result.Quote)
	assert.Nil(t, err)
	defer rows.Close()
	rows.Next()
	rows.Scan(&expected)
	assert.Equal(t, expected, result.Quote)
}
