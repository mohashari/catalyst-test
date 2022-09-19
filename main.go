package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mohashari/catalyst-test/handler"
	"github.com/mohashari/catalyst-test/repository/postgres"
	"github.com/mohashari/catalyst-test/service"
)

func init() {
	godotenv.Load()
}

func main() {

	db, err := postgres.NewPostgres(&postgres.ConnParam{
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		DBName:      os.Getenv("DB_NAME"),
		User:        os.Getenv("DB_USER"),
		Pass:        os.Getenv("DB_PASSWORD"),
		Options:     os.Getenv("DB_OPTIONS"),
		MaxOpenConn: 5,
		MaxIdleConn: 5,
		MaxLifetime: time.Duration(5 * time.Minute),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	svc := service.NewService(db)

	router := http.DefaultServeMux

	router = handler.Router(context.Background(), router, svc)

	port := os.Getenv("SERVICE_PORT")

	log.Printf("=========> runnning server %s 🚀🚀🚀🚀🚀🚀", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("running err: %v", err)
	}

}
