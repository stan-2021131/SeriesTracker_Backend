package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"series-tracker/db"
	"series-tracker/handlers"
	"series-tracker/services"

	_ "github.com/lib/pq"
)

func main() {
	// Conectar con retry
	var err error
	var database *sql.DB
	for i := 0; i < 10; i++ {
		database, err = db.ConnectDB()
		if err == nil {
			err = database.Ping()
			if err == nil {
				break
			}
		}
		log.Println("Esperando a la base de datos...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("No se pudo conectar a la DB:", err)
	}

	log.Println("Conectado a PostgreSQL")

	// Endpoint de prueba
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API up and running"))
	})

	http.HandleFunc("/series", handlers.SeriesHandler(database))
	http.HandleFunc("/series/", handlers.SeriesById(database))

	// Servir imágenes (para más adelante)
	http.Handle("/uploads/",
		http.StripPrefix("/uploads/",
			http.FileServer(http.Dir("./uploads")),
		),
	)

	// Middleware CORS
	handler := services.EnableCORS(http.DefaultServeMux)
	log.Println("Servidor corriendo en puerto 3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
