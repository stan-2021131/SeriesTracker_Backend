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

func RatingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 4 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Message: "Formato inválido. Use /series/{id}/ratings",
			})
			return
		}

		id, err := strconv.Atoi(parts[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Message: "El id debe ser un número",
			})
			return
		}

		switch r.Method {

		// ================= POST =================
		case http.MethodPost:

			var body model.RatingRequest

			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error al decodificar el rating",
				})
				return
			}

			if body.Puntaje < 1 || body.Puntaje > 5 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "El rating debe estar entre 1 y 5",
				})
				return
			}

			err = repository.CreateRating(db, id, body.Puntaje)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error al crear el rating",
				})
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(model.Response{
				Message: "Rating creado exitosamente",
			})

		// ================= GET =================
		case http.MethodGet:

			avgRating, err := repository.GetAvgRating(db, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error al obtener el rating promedio",
				})
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(model.Response{
				Message: "Rating promedio obtenido exitosamente",
				Data:    avgRating,
			})

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Message: "Método no permitido",
			})
		}
	}
}

// @Summary Crear rating
// @Description Agrega un rating (1 a 5) a una serie
// @Tags ratings
// @Accept json
// @Produce json
// @Param id path int true "ID de la serie"
// @Param rating body model.RatingRequest true "Puntaje"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /series/{id}/ratings [post]
func _docCreateRating() {}

// @Summary Obtener rating promedio
// @Description Retorna el promedio de ratings de una serie
// @Tags ratings
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} model.Response
// @Failure 500 {object} model.ErrorResponse
// @Router /series/{id}/ratings [get]
func _docGetAvgRating() {}
