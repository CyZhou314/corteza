package rest

import (
	"github.com/go-chi/chi/v5"

	"github.com/cyzhou314/corteza/server/discovery/rest/handlers"
	"github.com/cyzhou314/corteza/server/pkg/auth"
)

func MountRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(auth.HttpTokenValidator("discovery"))

			handlers.NewResources(Resources()).MountRoutes(r)
			handlers.NewFeed(Feed()).MountRoutes(r)
			handlers.NewMappings(Mappings()).MountRoutes(r)
		})
	}
}
