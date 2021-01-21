package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/delivery/rest"
	movie "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/movie/repository"
	movie_mysql "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/movie/repository/mysql"
	omdb "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/omdb/repository"
	omdb_http "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/omdb/repository/http"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg := loadConfig()

	log.Println("Connecting database...")
	db := initDB(cfg)
	defer db.Close()

	log.Println("Initializing repository...")
	var omdbRepo omdb.Repository = omdb_http.New(&omdb_http.Config{BaseURL: cfg.Omdb.URL, Key: cfg.Omdb.Key})
	var movieRepo movie.Repository = movie_mysql.New(db)

	log.Println("Initializing service...")
	service := service.New(omdbRepo, movieRepo)

	log.Println("Initializing delivery...")
	router := httprouter.New()
	rest.New(service).Register(router)

	log.Println("Initializing grpc...")
	go initGrpc(cfg, service)
	log.Printf("Running grpc on port %s", cfg.GrpcPort)

	log.Printf("Running server on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(cfg.ServerPort, router))
}
