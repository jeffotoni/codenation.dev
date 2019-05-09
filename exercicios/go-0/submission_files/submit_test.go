package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubmit10MaioresEstadosDoBrasil(t *testing.T) {
	estados, err := os10maioresEstadosDoBrasil()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(estados))
	expected := []string{"Amazonas", "Pará", "Mato Grosso", "Minas Gerais", "Bahia", "Mato Grosso do Sul", "Goiás", "Maranhão", "Rio Grande do Sul", "Tocantins"}
	assert.Equal(t, expected, estados)
}
