package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubmitQ1(t *testing.T) {
	r, err := q1()
	assert.Nil(t, err)
	assert.Equal(t, 164, r)
}

func TestSubmitQ2(t *testing.T) {
	r, err := q2()
	assert.Nil(t, err)
	assert.Equal(t, 647, r)
}

func TestSubmitQ3(t *testing.T) {
	r, err := q3()
	assert.Nil(t, err)
	assert.Equal(t, 20, len(r))
	expected := []string{
		"C. Ronaldo dos Santos Aveiro",
		"Lionel Messi",
		"Neymar da Silva Santos Jr.",
		"Luis Suárez",
		"Manuel Neuer",
		"Robert Lewandowski",
		"David De Gea Quintana",
		"Eden Hazard",
		"Toni Kroos",
		"Gonzalo Higuaín",
		"Sergio Ramos García",
		"Kevin De Bruyne",
		"Thibaut Courtois",
		"Alexis Sánchez",
		"Luka Modrić",
		"Gareth Bale",
		"Sergio Agüero",
		"Giorgio Chiellini",
		"Gianluigi Buffon",
		"Paulo Dybala",
	}
	assert.Equal(t, expected, r)
}

func TestSubmitQ4(t *testing.T) {
	r, err := q4()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(r))
	expected := []string{
		"C. Ronaldo dos Santos Aveiro",
		"Lionel Messi",
		"Luis Suárez",
		"Gareth Bale",
		"Robert Lewandowski",
		"Toni Kroos",
		"Luka Modrić",
		"Sergio Agüero",
		"Sergio Ramos García",
		"Eden Hazard",
	}
	assert.Equal(t, expected, r)
}

func TestSubmitQ5(t *testing.T) {
	r, err := q5()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(r))
	expected := []string{
		"Barry Richardson",
		"Óscar Pérez",
		"Essam El Hadary",
		"Danny Coyne",
		"Jimmy Walker",
		"Joaquim Manuel Sampaio Silva",
		"Kjetil Wæhler",
		"Chris Day",
		"Marco Storari",
		"Benjamin Nivet",
	}
	assert.Equal(t, expected, r)
}

func TestSubmitQ6(t *testing.T) {
	r, err := q6()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(r))
	expected := map[int]int{
		32: 506,
		30: 807,
		25: 1515,
		31: 666,
		28: 1053,
		26: 1199,
		27: 1153,
		29: 1127,
		39: 18,
		23: 1395,
		24: 1321,
		35: 188,
		33: 610,
		34: 271,
		36: 137,
		21: 1275,
		22: 1324,
		18: 682,
		20: 1252,
		19: 1088,
		37: 69,
		38: 38,
		40: 4,
		17: 270,
		44: 2,
		41: 3,
		16: 18,
		43: 2,
		47: 1,
	}
	assert.Equal(t, expected, r)
}
