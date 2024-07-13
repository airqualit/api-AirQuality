package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	db "github.com/go/qualityWater/src/DB"
	"github.com/go/qualityWater/src/repository"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port         string
	DatabaseUrl  string
	Collection   string
	DatabaseName string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("DB_Url is required")
	}

	if config.DatabaseName == "" {
		return nil, errors.New("DatabaseName is required")
	}

	if config.Collection == "" {
		return nil, errors.New("Collection is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	handler := cors.Default().Handler(b.router)

	mongoClient, err := db.ConnectMongoDB(b.config.DatabaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	coll, err := db.NewMongoDBRepository(b.config.DatabaseName, b.config.Collection, mongoClient)

	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(coll)

	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("Listen And Server", err)
	}
}
