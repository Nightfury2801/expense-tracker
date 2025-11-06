package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"expense-tracker/internal/config"
	"expense-tracker/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(cfg)
	defer db.Close()

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("🚀 Server running on :8080")
	http.ListenAndServe(":8080", r)
}
