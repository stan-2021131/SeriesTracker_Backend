package repository

import (
	"database/sql"
	"series-tracker/model"
)

func GetSeries(database *sql.DB, limit int, offset int, q string) ([]model.Serie, error) {
	var series []model.Serie
	var rows *sql.Rows
	var err error
	if q != "" {
		rows, err = database.Query(`SELECT * FROM series WHERE titulo ILIKE '%' || $3 || '%' LIMIT $1 OFFSET $2`, limit, offset, q) // Realiza la consulta a la base de datos
	} else {
		rows, err = database.Query("SELECT * FROM series LIMIT $1 OFFSET $2", limit, offset) // Realiza la consulta a la base de datos
	}
	if err != nil {
		return series, err
	}
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
	if err == sql.ErrNoRows {
		return s, sql.ErrNoRows
	}
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
	res, err := database.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func UpdateSerie(database *sql.DB, serie *model.Serie, id int) error {
	query := `UPDATE series SET titulo = $1, sinopsis = $2, episodios = $3, pais_origen = $4, genero_principal = $5, portada_url = $6 WHERE id = $7`
	res, err := database.Exec(
		query,
		serie.Titulo,
		serie.Sinopsis,
		serie.Episodios,
		serie.PaisOrigen,
		serie.GeneroPrincipal,
		serie.PortadaURL,
		id,
	)
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}
