package router

import (
	"net/http"
	"sayban/internal/config"
	"sayban/internal/handlers"
	"sayban/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(cfg *config.Config, userService service.UserServiceInterFace) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	// Key-Handler
	r.Route("/keys", func(r chi.Router) {
		// r.Get("/bundle", handlers.NewBundleHandler().ServeHTTP)
		r.Post("/register", handlers.NewKeyRegisterHadnler().ServeHTTP)
	})

	// Data-Base
	r.Route("/db/v1", func(r chi.Router) {
		r.Post("/lock", handlers.NewLockHandler(userService).ServeHTTP)
		r.Post("/unlock", handlers.NewUnlockHandler(cfg, userService).ServeHTTP)
		r.Post("/register", handlers.NewRegisterHandler(cfg, userService).ServeHTTP)
	})

	// account
	r.Route("/account", func(r chi.Router) {
		r.Post("/register", handlers.NewAccountRegisterHandler().ServeHTTP)
		r.Post("/delete", handlers.NewAccountDeleteHandler().ServeHTTP)
		r.Post("/get", handlers.NewAccountListHandler().ServeHTTP)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		// r.Get("/auth/v1/me", handlers.MeHandler)
	})

	return r
}
