package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
)

func (rest *Rest) getDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// define context
	ctx, cancel := context.WithTimeout(r.Context(), ctxTimeout)
	defer cancel()

	// service get detail
	e, err := rest.service.GetDetail(ctx, &service.GetDetailRequest{
		ID: ps.ByName("id"),
	})
	if err != nil {
		log.Println(err)
		errorHandler(w, err, http.StatusInternalServerError)
		return
	}

	responseHandler(w, e)
}
