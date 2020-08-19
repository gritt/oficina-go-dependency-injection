package rest

import (
	"errors"
	"log"
	"net/http"

	"github.com/gritt/oficina-go-dependency-injection/core"
)

type (
	TransactionCreator interface {
		Create(t core.Transaction) error
	}

	API struct {
		TransactionCreator
	}
)

func NewAPI(c TransactionCreator) *API {
	return &API{c}
}

func (api *API) Serve() {
	http.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		// TODO: read request payload
		trs := core.Transaction{
			Name:   "test",
			Amount: 10,
		}

		if err := api.Create(trs); err != nil {
			log.Println(err)

			dbErr := &core.DatabaseError{}
			if errors.As(err, &dbErr) {
				w.WriteHeader(http.StatusInternalServerError)
			}

			tsErr := &core.TransactionError{}
			if errors.As(err, &tsErr) {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	})

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
