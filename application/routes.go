package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/arthurspedine/orders-api/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHanlder := &handler.Order{}

	router.Post("/", orderHanlder.Create)
	router.Get("/", orderHanlder.List)
	router.Get("/{id}", orderHanlder.GetById)
	router.Put("/{id}", orderHanlder.UpdateById)
	router.Delete("/{id}", orderHanlder.DeleteById)
}