package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"series-tracker/repository"
	"strconv"
	"strings"
)

func RatingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 4 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Formato de ruta inválido. Use /series/{id}/ratings",
			})
			return
		}
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "El id debe ser un número",
			})
			return
		}
		switch r.Method {
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			var body struct {
				Puntaje int `json:"puntaje"`
			}
			err = json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Error al decodificar el rating",
				})
				return
			}
			if body.Puntaje < 1 || body.Puntaje > 5 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "El rating debe estar entre 0 y 5",
				})
				return
			}

			err = repository.CreateRating(db, id, body.Puntaje)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Error al crear el rating",
				})
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "Rating creado exitosamente"})
		case http.MethodGet:
			avgRating, err := repository.GetAvgRating(db, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Error al obtener el rating promedio",
				})
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Rating promedio obtenido exitosamente",
				"data":    avgRating,
			})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Método no permitido",
			})
			return
		}
	}
}
