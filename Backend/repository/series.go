package repository

import (
	"database/sql"
	"series-tracker/model"
)

func GetSeries(database *sql.DB) ([]model.Serie, error) {
	var series []model.Serie
	rows, err := database.Query("SELECT * FROM series") // Realiza la consulta a la base de datos
	if err != nil {
		return series, err
	} // Si hay error, retorna el slice vacío

	defer rows.Close() // Cierra la conexión cuando la función termine
	for rows.Next() {
		var s model.Serie
		err = rows.Scan(&s.ID, &s.Titulo, &s.Sinopsis, &s.Episodios, &s.PaisOrigen, &s.GeneroPrincipal, &s.PortadaURL, &s.FechaCreacion)
		if err != nil {
			return series, err
		}
		series = append(series, s)
	}
	return series, nil
}

func GetSerie(database *sql.DB, id int) (model.Serie, error) {
	var s model.Serie
	row := database.QueryRow("SELECT * FROM series WHERE id = $1", id)
	err := row.Scan(&s.ID, &s.Titulo, &s.Sinopsis, &s.Episodios, &s.PaisOrigen, &s.GeneroPrincipal, &s.PortadaURL, &s.FechaCreacion)
	if err != nil {
		return s, err
	}
	return s, nil
}

func CreateSerie(database *sql.DB, serie *model.Serie) error {
	query := `INSERT INTO series (titulo, sinopsis, episodios, pais_origen, genero_principal, portada_url) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := database.Exec(
		query,
		serie.Titulo,
		serie.Sinopsis,
		serie.Episodios,
		serie.PaisOrigen,
		serie.GeneroPrincipal,
		serie.PortadaURL,
	)
	return err
}

func DeleteSerie(database *sql.DB, id int) error {
	query := "DELETE FROM series WHERE id = $1"
	_, err := database.Exec(query, id)
	return err
}
