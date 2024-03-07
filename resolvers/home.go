package resolvers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/cafrias/offers-market/db"
	"github.com/cafrias/offers-market/ui"
)

const resultsPerPage = 15

type HomeResolver struct {
	Db db.Db
}

func (h HomeResolver) Resolver(w http.ResponseWriter, r *http.Request) {
	pStr := r.URL.Query().Get("p")
	pNum, err := strconv.Atoi(pStr)
	if err != nil {
		pNum = 1
	}

	offers, totalPages, err := db.GetAvailableOffers(h.Db, uint(pNum), resultsPerPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	page := ui.Home(offers, uint(pNum), totalPages)

	page.Render(context.Background(), w)
}
