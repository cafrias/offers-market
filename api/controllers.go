package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cafrias/offers-market/db"
	"github.com/cafrias/offers-market/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const resultsPerPage = 15

type Controllers struct {
	Db db.Db
}

type OffersResponse struct {
	Items []db.OfferResult `json:"items"`
	Total int32            `json:"total"`
}

func (o *OffersResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *Controllers) ListOffers(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	pStr := qs.Get("p")
	pNum, err := strconv.Atoi(pStr)
	if err != nil {
		pNum = 1
	}

	var offers []db.OfferResult
	var totalPages uint
	qStr := qs.Get("q")

	if len(qStr) > 0 {
		offers, totalPages, err = db.SearchAvailableOffers(c.Db, qStr, uint(pNum), resultsPerPage)
	} else {
		offers, totalPages, err = db.GetAvailableOffers(c.Db, uint(pNum), resultsPerPage)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	res := OffersResponse{
		Items: offers,
		Total: int32(totalPages),
	}

	render.Render(w, r, &res)
}

type GetStoreResponse struct {
	models.Store
}

func (g *GetStoreResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *Controllers) GetStore(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if len(idStr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	store, err := db.GetStore(c.Db, uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	res := GetStoreResponse{
		Store: *store,
	}

	render.Render(w, r, &res)
}
