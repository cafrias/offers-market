package main

import (
	"log"
	"net/http"

	"github.com/cafrias/offers-market/api"
	"github.com/cafrias/offers-market/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	session, err := db.Connect()
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer session.Close()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "application/json"))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	apiControllers := api.Controllers{
		Db: session,
	}
	r.Route("/api", func(r chi.Router) {
		r.Get("/offers", apiControllers.ListOffers)
		r.Get("/store/{id}", apiControllers.GetStore)
	})

	http.ListenAndServe(":1234", r)
}
