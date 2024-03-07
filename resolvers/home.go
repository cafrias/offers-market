package resolvers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/cafrias/offers-market/db"
	"github.com/cafrias/offers-market/models"
	"github.com/cafrias/offers-market/ui"
)

const resultsPerPage = 15

type HomeResolver struct {
	Db db.Db
}

func (h HomeResolver) Resolver(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	pStr := qs.Get("p")
	pNum, err := strconv.Atoi(pStr)
	if err != nil {
		pNum = 1
	}

	var offers []models.Offer
	var totalPages uint
	qStr := qs.Get("q")

	if len(qStr) > 0 {
		offers, err = db.SearchAvailableOffers(h.Db, qStr, uint(pNum), resultsPerPage)
	} else {
		offers, totalPages, err = db.GetAvailableOffers(h.Db, uint(pNum), resultsPerPage)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	page := ui.Home(offers, uint(pNum), totalPages, qStr)

	page.Render(context.Background(), w)
}
