package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestQuote(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/quote", quote())
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/v1/quote")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestQuoteByActor(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/quote/{actor}", quoteByActor())
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/v1/quote/John+Cleese")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
