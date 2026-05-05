package repository

import (
	"database/sql"
)

func CreateRating(db *sql.DB, id int, rating int) error {
	query := "INSERT INTO ratings (serie_id, puntaje) VALUES ($1, $2)"
	_, err := db.Exec(query, id, rating)
	if err != nil {
		return err
	}
	return nil
}

func GetAvgRating(db *sql.DB, serieID int) (float64, error) {
	var avg sql.NullFloat64
	err := db.QueryRow("SELECT AVG(puntaje) FROM ratings WHERE serie_id = $1", serieID).Scan(&avg)
	if err != nil {
		return 0, err
	}
	if !avg.Valid {
		return 0, nil
	}
	return avg.Float64, nil
}
