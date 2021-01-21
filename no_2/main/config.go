package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	godotenv "github.com/joho/godotenv"
	envconfig "github.com/kelseyhightower/envconfig"
	grpc_handler "github.com/mghifariyusuf/stockbit_test.git/no_2/delivery/grpc/handler"
	grpc_schema "github.com/mghifariyusuf/stockbit_test.git/no_2/delivery/grpc/schema"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
	"google.golang.org/grpc"
)

// wrap configuration needed
type config struct {
	Omdb       omdbConfig `envconfig:"OMDB"`
	ServerPort string     `envconfig:"SERVER_PORT"`
	GrpcPort   string     `envconfig:"GRPC_PORT"`
	Database   dbConfig   `envconfig:"DB"`
}

type dbConfig struct {
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
	Name     string `envconfig:"NAME"`
}

type omdbConfig struct {
	URL string `envconfig:"URL"`
	Key string `envconfig:"KEY"`
}

func loadConfig() config {
	var (
		cfg     config
		appName = "MOVIE"
	)

	log.Println("Loading configuration...")

	// load configuration from env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed load configuration: %s", err)
	}

	// parse configuration
	err = envconfig.Process(appName, &cfg)
	if err != nil {
		log.Fatalf("Failed load configuration: %s", err)
	}

	return cfg
}

func initDB(cfg config) *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)
	}

	return db
}

func initGrpc(cfg config, service service.Service) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen grpc: %s", err)
	}

	grpcServer := grpc.NewServer()
	grpc_schema.RegisterSearchServiceServer(
		grpcServer,
		grpc_handler.New(service),
	)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve grpc: %s", err)
	}
}
