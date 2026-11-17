package models

import "errors"

var (
	ErrNotFound    = errors.New("ressource introuvable")
	ErrBadRequest  = errors.New("Requete innaproprié")
	ErrTimeout     = errors.New("Temps de recherche max atteint // Erreur Timeout")
	ErrUnavailable = errors.New("Ressource Indisponible")
)
