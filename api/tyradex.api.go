package api

import (
	"TP-Pokedex-Go/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

const baseURL = "https://tyradex.app/api/v1/pokemon"

func GetPokemon(id int) (models.Pokemon, error) {
	endpoint := fmt.Sprintf("%s/%d", baseURL, id)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return models.Pokemon{}, err
	}
	request.Header.Set("User-Agent", "DataDex/1.0")

	reponse, err := client.Do(request)
	if err != nil {
		return models.Pokemon{}, err
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != http.StatusOK {
		return models.Pokemon{}, fmt.Errorf("réponse inattendue : %s", reponse.Status)
	}
	var pokemon models.Pokemon
	err = json.NewDecoder(reponse.Body).Decode(&pokemon)
	return pokemon, err
}

func GetAllPokemon() ([]models.Pokemon, error) {
	endpoint := fmt.Sprintf("%s", baseURL)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return []models.Pokemon{}, err
	}

	request.Header.Set("User-Agent", "DataDex/1.0")

	reponse, err := client.Do(request)
	if err != nil {
		return []models.Pokemon{}, err
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != http.StatusOK {
		return []models.Pokemon{}, fmt.Errorf("réponse inattendue : %s", reponse.Status)
	}
	var pokemon []models.Pokemon
	err = json.NewDecoder(reponse.Body).Decode(&pokemon)
	return pokemon, err
}
