package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"series-tracker/repository"
	"strconv"
	"strings"
)

// GetSeries maneja las peticiones a la ruta "/series"
func GetSeries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		series, err := repository.GetSeries(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Error al obtener las series",
			})
			return
		}
		if len(series) == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "No se encontraron series",
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(series)
	}
}

// GetSerieById maneja las peticiones a la ruta "/series/{id}"
func GetSerieById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path                // Obtiene el path de la petición
		parts := strings.Split(path, "/") // Divide el path en partes
		if len(parts) != 3 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Formato de ruta inválido. Use /series/{id}",
			})
			return
		}
		id, err := strconv.Atoi(parts[2]) // Convierte el id a entero
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "El id debe ser un número",
			})
			return
		}
		serie, err := repository.GetSerie(db, id)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "No se encontro la serie",
			})
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Error al obtener la serie",
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(serie)
	}
}
