package main

import (
	"TP-Pokedex-Go/api"
	"fmt"
)

func main() {
	fmt.Println(api.GetPokemon(2))
	fmt.Println(api.GetAllPokemon())
}
