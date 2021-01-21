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
	router.GET("/", rest.search)
	router.GET("/:id", rest.getDetail)
}

// handler to return json response
func responseHandler(w http.ResponseWriter, object interface{}) {
	status := http.StatusOK

	jsonResp, err := json.Marshal(object)
	if err != nil {
		errorHandler(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(jsonResp))
}

// handler to return error
func errorHandler(w http.ResponseWriter, err error, status int) {
	m := map[string]string{
		"error": err.Error(),
	}
	jsonResp, err := json.Marshal(m)
	if err != nil {
		errorHandler(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(jsonResp))
}
