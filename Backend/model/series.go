package model

import "time"

type Serie struct {
	ID              int       `json:"id"`
	Titulo          string    `json:"titulo"`
	Sinopsis        string    `json:"sinopsis"`
	Episodios       int       `json:"episodios"`
	PaisOrigen      string    `json:"pais_origen"`
	GeneroPrincipal string    `json:"genero_principal"`
	PortadaURL      string    `json:"portada_url"`
	FechaCreacion   time.Time `json:"fecha_creacion"`
}

type Rating struct {
	ID            int       `json:"id"`
	SerieID       int       `json:"serie_id"`
	Puntaje       int       `json:"puntaje"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}
