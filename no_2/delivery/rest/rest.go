package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
)

// Rest ...
type Rest struct {
	service service.Service
}

const ctxTimeout = 30 * time.Second

// New ...
func New(
	service service.Service,
) *Rest {
	return &Rest{
		service: service,
	}
}

// Register ...
func (rest *Rest) Register(router *httprouter.Router) {
	router.GET("/search", rest.search)
	router.GET("/detail/:id", rest.getDetail)
}

func responseHandler(w http.ResponseWriter, object interface{}) {
	status := http.StatusOK

	jsonResp, err := json.Marshal(object)
	if err != nil {
		errorHandler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(jsonResp))
}

func errorHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
