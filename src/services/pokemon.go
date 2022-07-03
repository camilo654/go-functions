package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/camilo654/go-functions/src/models"
)

func GetPokemonName(id int) string {

	pokemon := getPokemonFromService(id)

	return pokemon.Name
}

func getPokemonFromService(id int) models.Pokemon {
	response, err := http.Get(fmt.Sprint("https://pokeapi.co/api/v2/pokemon-form/", id))
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responsePokemonService models.ResponsePokemonService
	json.Unmarshal(responseBody, &responsePokemonService)

	return responsePokemonService.Pokemon
}
