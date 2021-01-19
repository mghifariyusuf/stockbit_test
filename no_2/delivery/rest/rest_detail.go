package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
)

func (rest *Rest) getDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), ctxTimeout)
	defer cancel()

	e, err := rest.service.GetDetail(ctx, &service.GetDetailRequest{
		ID: ps.ByName("id"),
	})
	if err != nil {
		log.Println(err)
		errorHandler(w, err)
		return
	}

	responseHandler(w, e)
}
