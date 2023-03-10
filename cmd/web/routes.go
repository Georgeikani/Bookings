package main

import (
	"net/http"

	"github.com/georgeikani/Bookings/internal/config"
	"github.com/georgeikani/Bookings/internal/handlers"
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
	 
	 mux.Get("/general", handlers.Repo.General)
	 mux.Get("/major", handlers.Repo.Major)

	 mux.Get("/search-availability", handlers.Repo.Search )
	 mux.Post("/reservation", handlers.Repo.PostReservation)
	 mux.Post("/reservation-json", handlers.Repo.ReservationJSON)


	 mux.Get("/make-reservation", handlers.Repo.Reserve)
	 mux.Get("/contact", handlers.Repo.Contact)

	 
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}




