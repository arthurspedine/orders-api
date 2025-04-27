package application

import (
	"net/http"

	"github.com/arthurspedine/orders-api/handler"
	"github.com/arthurspedine/orders-api/repository/order"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", a.loadOrderRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	orderHanlder := &handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}

	router.Post("/", orderHanlder.Create)
	router.Get("/", orderHanlder.List)
	router.Get("/{id}", orderHanlder.GetById)
	router.Put("/{id}", orderHanlder.UpdateById)
	router.Delete("/{id}", orderHanlder.DeleteById)
}