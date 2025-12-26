package router

import (
	"net/http"

	"github.com/1-Utkarsh/temp/api/tasks"
	"github.com/go-chi/chi"
)

// InitRoutes initializes the API routes and returns the router
func InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		// API routes
		r.Route("/v1", func(r chi.Router) {
			// v1 routes
			// tasks endpoints
			r.Mount("/tasks", tasks.Routes())
		})
	})
	return r
}
