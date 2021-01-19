package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/delivery/rest"
	omdb "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/omdb/repository"
	omdb_http "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/omdb/repository/http"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
)

// input configuration here
var (
	omdbURL    = "http://www.omdbapi.com"
	omdbKey    = "faf7e5bb"
	serverPort = ":8080"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Initializing repository...")
	var omdbRepo omdb.Repository = omdb_http.New(&omdb_http.Config{BaseURL: omdbURL, Key: omdbKey})

	log.Println("Initializing service...")
	service := service.New(omdbRepo)

	log.Println("Initializing delivery...")
	router := httprouter.New()
	rest.New(service).Register(router)

	log.Printf("Start running server on port %s", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}
