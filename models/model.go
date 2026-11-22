package models

type Pokemon struct {
	PokedexId int     `json:"pokedex_id"`
	Category  string  `json:"category"`
	Name      Name    `json:"name"`
	Sprites   Sprites `json:"sprites"`
	Types     []Types `json:"types"`
	Stats     Stats   `json:"stats"`
}

type Name struct {
	Fr string `json:"fr"`
	En string `json:"en"`
}

type Sprites struct {
	Regular string `json:"regular"`
}

type Types struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Stats struct {
	Hp     int `json:"hp"`
	Atk    int `json:"atk"`
	Def    int `json:"def"`
	SpeAtk int `json:"spe_atk"`
	SpeDef int `json:"spe_def"`
	Vit    int `json:"vit"`
}
