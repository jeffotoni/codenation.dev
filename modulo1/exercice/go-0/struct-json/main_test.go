package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test10MaioresEstadosDoBrasil2(t *testing.T) {
	estados, err := os10maioresEstadosDoBrasil()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(estados))
}
