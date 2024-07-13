package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go/qualityWater/src/handlers"
	"github.com/go/qualityWater/src/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading .env file")
	}

	PORT := os.Getenv("PORT")
	DATABASE_URL := os.Getenv("MONGODB_URI")
	DB_NAME := os.Getenv("DB_NAME")
	COLLECTION_NAME := os.Getenv("COLLECTION_NAME")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:         PORT,
		DatabaseUrl:  DATABASE_URL,
		DatabaseName: DB_NAME,
		Collection:   COLLECTION_NAME,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	log.Println("handlers")
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/iotdevice/{id}", handlers.GetIotDeviceByHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/iotdevice", handlers.InsertIotDeviceByHandler(s)).Methods(http.MethodPost)
}
