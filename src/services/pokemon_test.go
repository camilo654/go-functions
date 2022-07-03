package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemonName(t *testing.T) {
	idPokemon := 25
	expectNamePokemon := "pikachu"

	resultNamePokemon := GetPokemonName(idPokemon)

	assert.Equal(t, expectNamePokemon, resultNamePokemon)
}
