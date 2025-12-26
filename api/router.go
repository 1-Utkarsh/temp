package router

import (
	"context"
	"net/http"

	"github.com/1-Utkarsh/temp/api/tasks"
	db "github.com/1-Utkarsh/temp/store"
	"github.com/1-Utkarsh/temp/util"
	"github.com/go-chi/chi"
)

// InitRoutes initializes the API routes and returns the router
func InitRoutes() http.Handler {
	r := chi.NewRouter()

	// Pass db connection via middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), util.DbKey, db.GetDB())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	// API routes
	r.Route("/", func(r chi.Router) {
		// tasks endpoints
		r.Mount("/tasks", tasks.Routes())
	})

	return r
}
