package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cafrias/offers-market/db"
	"github.com/cafrias/offers-market/ui"
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

	// templates := initTemplates()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "text/html", "text/css", "application/javascript", "application/json"))
	r.Use(render.SetContentType(render.ContentTypeHTML))

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	FileServer(r, "/static", filesDir)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		offers, err := db.GetAvailableOffers(session, 1, 15)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		page := ui.Home(offers)

		page.Render(context.Background(), w)
	})

	// offers, err := db.GetAvailableOffers(session, 1, 15)
	// if err != nil {
	// 	log.Fatalf("db.GetAvailableOffers(): %q\n", err)
	// }

	// for _, offer := range offers {
	// 	log.Printf("Offer: %+v\n", offer)
	// }

	// offers, err := db.SearchAvailableOffers(session, "car", 1, 15)
	// if err != nil {
	// 	log.Fatalf("db.SearchAvailableOffers(): %q\n", err)
	// }

	// for _, offer := range offers {
	// 	log.Printf("Offer: %+v\n", offer)
	// }

	// store := models.Store{
	// 	Name:    "Store 2",
	// 	Address: "Address 1",
	// }
	// err = db.CreateStore(session, &store)
	// if err != nil {
	// 	log.Fatalf("db.CreateStore(): %q\n", err)
	// }

	// brand := models.Brand{
	// 	Name: "Brand 1",
	// }

	// err = db.CreateBrand(session, &brand)
	// if err != nil {
	// 	log.Fatalf("db.CreateBrand(): %q\n", err)
	// }

	// offer := models.Offer{
	// 	Name:           "Offer 1",
	// 	BrandId:        2,
	// 	StoreId:        1,
	// 	Price:          100,
	// 	Quantity:       10,
	// 	Available:      10,
	// 	ExpirationDate: time.Now().Add(time.Hour * 24),
	// 	Picture:        "https://www.google.com",
	// }

	// err = db.CreateOffer(session, &offer)
	// if err != nil {
	// 	log.Fatalf("db.CreateOffer(): %q\n", err)
	// }

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
