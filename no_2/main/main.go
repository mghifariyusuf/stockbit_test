package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
	grpc_handler "github.com/mghifariyusuf/stockbit_test.git/no_2/delivery/grpc/handler"
	grpc_schema "github.com/mghifariyusuf/stockbit_test.git/no_2/delivery/grpc/schema"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/delivery/rest"
	movie "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/movie/repository"
	movie_http "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/movie/repository/http"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
	"google.golang.org/grpc"
)

// input configuration here
var (
	movieURL   = "http://www.movieapi.com"
	movieKey   = "faf7e5bb"
	serverPort = ":8080"
	grpcPort   = "9000"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Initializing repository...")
	var movieRepo movie.Repository = movie_http.New(&movie_http.Config{BaseURL: movieURL, Key: movieKey})

	log.Println("Initializing service...")
	service := service.New(movieRepo)

	log.Println("Initializing delivery...")
	router := httprouter.New()
	rest.New(service).Register(router)

	log.Println("Initializing grpc...")
	go initGrpc(service)
	log.Printf("Running grpc on port %s", grpcPort)

	log.Printf("Running server on port %s", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}

func initGrpc(service service.Service) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen grpc: %s", err)
	}

	grpcServer := grpc.NewServer()
	grpc_schema.RegisterSearchServiceServer(
		grpcServer,
		grpc_handler.New(service),
	)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc: %s", err)
	}
}
