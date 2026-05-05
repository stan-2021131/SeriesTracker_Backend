package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"series-tracker/model"
	"series-tracker/repository"
	"series-tracker/services"
	"strconv"
	"strings"
)

// SeriesHandler maneja las peticiones a la ruta "/series"
func SeriesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			query := r.URL.Query()
			//Extrae los valores de page y limit de la query
			pageStr := query.Get("page")
			limitStr := query.Get("limit")
			q := query.Get("q")
			sort := query.Get("sort")
			order := query.Get("order")

			//Convierte los valores de page y limit a enteros
			page, err := strconv.Atoi(pageStr)
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Los campos page y limit deben ser números",
				})
				return
			}

			//Calcula el offset
			offset := (page - 1) * limit

			//Obtiene las series de la base de datos
			series, err := repository.GetSeries(db, limit, offset, q, sort, order)
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

			path, err := services.SaveImage(r, "imagen")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": err.Error(),
				})
				return
			}
			newSerie.PortadaURL = path
			newSerie.Titulo = r.FormValue("titulo")
			newSerie.Sinopsis = r.FormValue("sinopsis")
			newSerie.PaisOrigen = r.FormValue("pais_origen")
			newSerie.GeneroPrincipal = r.FormValue("genero_principal")
			newSerie.Episodios, err = strconv.Atoi(r.FormValue("episodios"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "El campo episodios debe ser un número",
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

			data, err := repository.GetSerie(db, id)
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"message": `No se encontro la serie`,
				})
				return
			}
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Error interno al obtener la serie",
				})
				return
			}

			path, err := services.SaveImage(r, "imagen")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": err.Error(),
				})
				return
			}
			if path == "/uploads/default.jpg" {
				// no se subió imagen → mantener la actual
				newSerie.PortadaURL = data.PortadaURL
			} else {
				// sí se subió imagen → reemplazar
				newSerie.PortadaURL = path
				if data.PortadaURL != "/uploads/default.jpg" {
					_ = os.Remove("." + data.PortadaURL)
				}
			}

			newSerie.Titulo = r.FormValue("titulo")
			newSerie.Sinopsis = r.FormValue("sinopsis")
			newSerie.PaisOrigen = r.FormValue("pais_origen")
			newSerie.GeneroPrincipal = r.FormValue("genero_principal")
			newSerie.Episodios, err = strconv.Atoi(r.FormValue("episodios"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "El campo episodios debe ser un número",
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
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Serie actualizada exitosamente",
				"data":    newSerie,
			})
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			err = repository.DeleteSerie(db, id)
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
					"message": "Error al eliminar la serie",
				})
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Serie eliminada exitosamente",
			})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Método no permitido",
			})
		}
	}
}
