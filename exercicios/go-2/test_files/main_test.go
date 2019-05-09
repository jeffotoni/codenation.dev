package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ1(t *testing.T) {
	r, err := q1()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, r)
}

func TestQ2(t *testing.T) {
	r, err := q2()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, r)
}

func TestQ3(t *testing.T) {
	r, err := q3()
	assert.Nil(t, err)
	assert.Equal(t, 20, len(r))
}

func TestQ4(t *testing.T) {
	r, err := q4()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(r))
}

func TestQ5(t *testing.T) {
	r, err := q5()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(r))
}

func TestQ6(t *testing.T) {
	r, err := q6()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(r))
}
