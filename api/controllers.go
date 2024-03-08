package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cafrias/offers-market/db"
	"github.com/go-chi/render"
)

const resultsPerPage = 15

type Controllers struct {
	Db db.Db
}

type OffersResponse struct {
	Offers []db.OfferResult `json:"offers"`
	Pages  int32            `json:"pages"`
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
		offers, err = db.SearchAvailableOffers(c.Db, qStr, uint(pNum), resultsPerPage)
	} else {
		offers, totalPages, err = db.GetAvailableOffers(c.Db, uint(pNum), resultsPerPage)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	res := OffersResponse{
		Offers: offers,
		Pages:  int32(totalPages),
	}

	render.Render(w, r, &res)
}
