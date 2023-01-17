package main

import (
	"net/http"

	"github.com/georgeikani/Bookings/pkg/config"
	"github.com/georgeikani/Bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// mux.Get("/signup", http.HandlerFunc(handlers.Repo.Signup))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)


	 mux.Get("/home", handlers.Repo.Home)
	 mux.Get("/about", handlers.Repo.About)
	 mux.Get("/signup", handlers.Repo.Signup)


	return mux
}




