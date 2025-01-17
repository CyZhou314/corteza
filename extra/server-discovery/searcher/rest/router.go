package rest

import (
	"github.com/cyzhou314/corteza/extra/server-discovery/pkg/auth"
	"github.com/cyzhou314/corteza/extra/server-discovery/searcher/rest/handlers"
	"github.com/go-chi/chi/v5"
)

func MountRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(auth.HttpTokenValidator("discovery"))
			handlers.NewSearch(Search()).MountRoutes(r)
		})
	}
}
