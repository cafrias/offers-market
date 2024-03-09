package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cafrias/offers-market/api"
	"github.com/cafrias/offers-market/db"
	"github.com/cafrias/offers-market/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

var port = 1234

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

	log.Printf("Server listening on %v:%v\n", utils.GetLocalIP(), port)

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), r)
}
