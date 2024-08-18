package config

import (
	"myExpenses/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func Router(db gorm.DB) *chi.Mux {
	trhdl := handler.NewTransactionRecordHandler(db)

	r := chi.NewRouter()
	r.Use(JSONMiddleware)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	}))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/transaction_records", func(r chi.Router) {
			r.Get("/", trhdl.ListTransactionRecord)
			r.Post("/", trhdl.CreateTransactionRecord)
			r.Put("/{id}", trhdl.UpdateTransactionRecord)
		})

	})

	return r
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
