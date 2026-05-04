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
