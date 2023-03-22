package main

import (
	"net/http"

	"github.com/MarioGN/desafio-multithreading/internal/infra/webserver/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	handler := handlers.NewCepHandler()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/consulta-cep/{cep}", handler.GetCep)

	http.ListenAndServe(":8000", router)
}
