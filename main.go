package main

type Pokemon struct {
	PokedexId int    `json:"pokedex_id"`
	Category  string `json:"category"`
	Name      Name   `json:"name"`
	Sprite    Sprite `json:"sprite"`
	Types     Types  `json:"types"`
	Stats     Stats  `json:"stats"`
}

type Name struct {
	Fr string `json:"fr"`
	En string `json:"en"`
}

type Sprite struct {
	Regular string `json:"regular"`
}

type Types struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Stats struct {
	Hp     string `json:"hp"`
	Atk    string `json:"atk"`
	Def    string `json:"def"`
	SpeAtk string `json:"spe_atk"`
	SpeDef string `json:"spe_def"`
	Vit    string `json:"vit"`
}

func main() {

}
