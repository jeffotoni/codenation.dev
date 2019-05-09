package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test10MaioresEstadosDoBrasil(t *testing.T) {
	estados, err := os10maioresEstadosDoBrasil()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(estados))
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Acre"},
		{"Minas Gerais"},
		{"Distrito Federal"},
		{"Minas Gerais"},
		{"São Paulo"},
		{"Amazonas"},
		{"Rio grande do sul"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_os10maioresEstadosDoBrasil(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Acre", []string{"minas", "sao paulo"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := os10maioresEstadosDoBrasil()
			if (err != nil) != tt.wantErr {
				t.Errorf("os10maioresEstadosDoBrasil() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("os10maioresEstadosDoBrasil() = %v, want %v", got, tt.want)
			}
		})
	}
}
