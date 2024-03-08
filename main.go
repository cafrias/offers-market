package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

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
	})

	http.ListenAndServe(":1234", r)
}

func initTemplates() (templates map[string]*template.Template) {
	templates = make(map[string]*template.Template)

	home, err := template.ParseFiles("pages/index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles(): %q\n", err)
	}

	templates["home"] = home

	return templates
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
