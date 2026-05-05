package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"series-tracker/model"
	"series-tracker/repository"
	"strconv"
	"strings"
)

// SeriesHandler maneja las peticiones a la ruta "/series"
func SeriesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
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
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			var newSerie model.Serie
			err := json.NewDecoder(r.Body).Decode(&newSerie)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Formato de serie inválido",
				})
				return
			}
			if newSerie.Titulo == "" || newSerie.Sinopsis == "" || newSerie.PaisOrigen == "" || newSerie.GeneroPrincipal == "" {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Los campos titulo, sinopsis, pais_origen y genero_principal son obligatorios",
				})
				return
			}

			if newSerie.Episodios <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "El campo episodios debe ser mayor o igual a 0",
				})
				return
			}

			err = repository.CreateSerie(db, &newSerie)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Error al crear la serie",
				})
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Serie creada exitosamente",
				"data":    newSerie,
			})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Método no permitido",
			})
		}
	}
}

// GetSerieById maneja las peticiones a la ruta "/series/{id}"
func SeriesById(db *sql.DB) http.HandlerFunc {
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

		switch r.Method {
		case http.MethodGet:
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
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			var newSerie model.Serie
			err = json.NewDecoder(r.Body).Decode(&newSerie)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Formato de serie inválido",
				})
				return
			}
			if newSerie.Titulo == "" || newSerie.Sinopsis == "" || newSerie.PaisOrigen == "" || newSerie.GeneroPrincipal == "" {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Los campos titulo, sinopsis, pais_origen y genero_principal son obligatorios",
				})
				return
			}

			if newSerie.Episodios <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "El campo episodios debe ser mayor o igual a 0",
				})
				return
			}

			err = repository.UpdateSerie(db, &newSerie, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Error al actualizar la serie",
				})
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Serie actualizada exitosamente",
				"data":    newSerie,
			})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Método no permitido",
			})
		}
	}
}
