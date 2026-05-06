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

func SeriesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:

			query := r.URL.Query()

			page, err := strconv.Atoi(query.Get("page"))
			limit, err2 := strconv.Atoi(query.Get("limit"))

			if err != nil || err2 != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "page y limit deben ser números",
				})
				return
			}

			q := query.Get("q")
			sort := query.Get("sort")
			order := query.Get("order")

			offset := (page - 1) * limit

			series, err := repository.GetSeries(db, limit, offset, q, sort, order)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error al obtener series",
				})
				return
			}

			if len(series) == 0 {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "No se encontraron series",
				})
				return
			}

			json.NewEncoder(w).Encode(model.Response{
				Message: "OK",
				Data:    series,
			})
		case http.MethodPost:

			var newSerie model.Serie

			path, err := services.SaveImage(r, "imagen")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: err.Error(),
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
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Episodios debe ser número",
				})
				return
			}

			if newSerie.Titulo == "" || newSerie.Sinopsis == "" ||
				newSerie.PaisOrigen == "" || newSerie.GeneroPrincipal == "" {

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Campos obligatorios faltantes",
				})
				return
			}

			if newSerie.Episodios <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Episodios inválidos",
				})
				return
			}

			err = repository.CreateSerie(db, &newSerie)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error al crear serie",
				})
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(model.Response{
				Message: "Serie creada",
				Data:    newSerie,
			})

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Message: "Método no permitido",
			})
		}
	}
}

func SeriesById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Message: "Ruta inválida",
			})
			return
		}

		id, err := strconv.Atoi(parts[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Message: "ID inválido",
			})
			return
		}

		switch r.Method {

		case http.MethodGet:

			serie, err := repository.GetSerie(db, id)
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "No encontrada",
				})
				return
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error interno",
				})
				return
			}

			json.NewEncoder(w).Encode(model.Response{
				Message: "OK",
				Data:    serie,
			})

		case http.MethodPut:

			data, err := repository.GetSerie(db, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error obteniendo serie",
				})
				return
			}

			var updated model.Serie

			path, err := services.SaveImage(r, "imagen")
			if err == nil && path != "/uploads/default.jpg" {
				updated.PortadaURL = path
				if data.PortadaURL != "/uploads/default.jpg" {
					_ = os.Remove("." + data.PortadaURL)
				}
			} else {
				updated.PortadaURL = data.PortadaURL
			}

			updated.Titulo = r.FormValue("titulo")
			updated.Sinopsis = r.FormValue("sinopsis")
			updated.PaisOrigen = r.FormValue("pais_origen")
			updated.GeneroPrincipal = r.FormValue("genero_principal")

			updated.Episodios, err = strconv.Atoi(r.FormValue("episodios"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Episodios inválido",
				})
				return
			}

			err = repository.UpdateSerie(db, &updated, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error actualizando",
				})
				return
			}

			json.NewEncoder(w).Encode(model.Response{
				Message: "Actualizado",
				Data:    updated,
			})

		case http.MethodDelete:

			err = repository.DeleteSerie(db, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Message: "Error eliminando",
				})
				return
			}

			json.NewEncoder(w).Encode(model.Response{
				Message: "Eliminado",
			})

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Message: "Método no permitido",
			})
		}
	}
}

// @Summary Listar series
// @Description Obtiene lista paginada de series
// @Tags series
// @Produce json
// @Param page query int true "Número de página"
// @Param limit query int true "Cantidad por página"
// @Param q query string false "Búsqueda"
// @Param sort query string false "Campo de orden"
// @Param order query string false "asc o desc"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /series [get]
func _docGetSeries() {}

// @Summary Crear serie
// @Description Crea una nueva serie
// @Tags series
// @Accept multipart/form-data
// @Produce json
// @Param titulo formData string true "Título"
// @Param sinopsis formData string true "Sinopsis"
// @Param pais_origen formData string true "País de origen"
// @Param genero_principal formData string true "Género principal"
// @Param episodios formData int true "Número de episodios"
// @Param imagen formData file false "Imagen de portada"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /series [post]
func _docCreateSeries() {}

// @Summary Obtener serie por ID
// @Tags series
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /series/{id} [get]
func _docGetSeriesById() {}

// @Summary Actualizar serie
// @Tags series
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "ID de la serie"
// @Param titulo formData string true "Título"
// @Param sinopsis formData string true "Sinopsis"
// @Param pais_origen formData string true "País"
// @Param genero_principal formData string true "Género"
// @Param episodios formData int true "Episodios"
// @Param imagen formData file false "Imagen"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /series/{id} [put]
func _docUpdateSeries() {}

// @Summary Eliminar serie
// @Tags series
// @Param id path int true "ID de la serie"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /series/{id} [delete]
func _docDeleteSeries() {}
