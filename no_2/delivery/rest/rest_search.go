package rest

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
)

func (rest *Rest) search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), ctxTimeout)
	defer cancel()

	var (
		params     = r.URL.Query()
		err        error
		searchWord = ""
		page       = 1
	)

	if params.Get("searchword") != "" {
		searchWord = params.Get("searchword")
	}

	if params.Get("pagination") != "" {
		page, err = strconv.Atoi(params.Get("pagination"))
		if err != nil {
			log.Println(err)
			errorHandler(w, err)
			return
		}
	}

	e, err := rest.service.Search(ctx, &service.SearchRequest{
		SearchWord: searchWord,
		Page:       page,
	})
	if err != nil {
		log.Println(err)
		errorHandler(w, err)
		return
	}

	responseHandler(w, e)
}
