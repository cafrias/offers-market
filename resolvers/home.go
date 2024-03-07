package resolvers

import (
	"context"
	"log"
	"net/http"

	"github.com/cafrias/offers-market/db"
	"github.com/cafrias/offers-market/ui"
)

type HomeResolver struct {
	Db db.Db
}

func (h HomeResolver) Resolver(w http.ResponseWriter, r *http.Request) {
	offers, err := db.GetAvailableOffers(h.Db, 1, 15)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	page := ui.Home(offers)

	page.Render(context.Background(), w)
}
