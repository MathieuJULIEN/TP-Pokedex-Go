package main

import (
	"TP-Pokedex-Go/api"
	"TP-Pokedex-Go/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var _temp = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/{$}", homeHendler)
	router.HandleFunc("/pokedex", PokedexHendler)
	router.HandleFunc("/pokemon/{id}", PokedexByHendler)
	router.HandleFunc("/", notFound)

	fileServer := http.FileServer(http.Dir("./assets"))
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("Erreur (Serveur) : Impossible de lancer le serveur : %v", err)
		return
	}
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	buffer := bytes.Buffer{}

	errRender := _temp.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		log.Printf("Erreur (Template) : Impossible de rendre le template %s : %v", name, errRender)
		http.Error(w, "Erreur chargement de la page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	buffer.WriteTo(w)
}

func homeHendler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	renderTemplate(w, "index", nil)
}

func PokedexHendler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	pokemon, err := api.GetAllPokemon()
	if err != nil {
		http.Error(w, "Erreur", http.StatusNotFound)
		return
	}
	search := strings.TrimSpace(r.URL.Query().Get("search"))
	renderTemplate(w, "pokedex", SearchPokemon(pokemon, search))
}

func SearchPokemon(list []models.Pokemon, search string) []models.Pokemon {
	search = strings.ToLower(search)
	recherche := []models.Pokemon{}
	for _, tempo := range list {
		if tempo.PokedexId == 0 {
			continue
		}
		nom := strings.ToLower(tempo.Name.Fr)
		if strings.Contains(nom, search) == true {
			recherche = append(recherche, tempo)
		}
	}
	return recherche
}

func PokedexByHendler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	id, errId := strconv.Atoi(r.PathValue("id"))
	if errId != nil {
		http.Error(w, "Erreur données invalides...", http.StatusBadRequest)
		return
	}

	pokemon, err := api.GetPokemon(id)
	if err != nil {
		http.Error(w, "Pokemon inexistant", http.StatusNotFound)
		return
	}
	renderTemplate(w, "pokemon", pokemon)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	renderTemplate(w, "error", nil)
}
