package router

import (
	"net/http"

	"github.com/1-Utkarsh/temp/api/tasks"
	"github.com/go-chi/chi"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		// API routes
		r.Route("/v1", func(r chi.Router) {
			// v1 routes
			r.Mount("/tasks", tasks.Routes())
		})
	})
	return r
}
