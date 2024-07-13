package db

import (
	"context"
	"log"

	"github.com/go/qualityWater/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoCollectionRepo struct {
	MongoCollection *mongo.Collection
}

type MongoClientRepo struct {
	MongoClient *mongo.Client
}

func ConnectMongoDB(url string) (*MongoClientRepo, error) {
	var err error
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("connection error", err)
	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed")
	}

	log.Print("mongo connected")
	return &MongoClientRepo{MongoClient: mongoClient}, nil
}

func NewMongoDBRepository(databaseName string, collectionName string, client *MongoClientRepo) (*MongoCollectionRepo, error) {
	coll := client.MongoClient.Database(databaseName).Collection(collectionName)
	return &MongoCollectionRepo{MongoCollection: coll}, nil
}

// metodo
func (repo *MongoCollectionRepo) InsertIotDevice(ctx context.Context, iotdevice *models.IotDevice) (interface{}, error) {
	result, err := repo.MongoCollection.InsertOne(ctx, iotdevice)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *MongoCollectionRepo) GetIotDeviceById(ctx context.Context, id string) (*models.IotDevice, error) {
	var iotdevice models.IotDevice

	err := repo.MongoCollection.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&iotdevice)

	if err != nil {
		return nil, err
	}

	return &iotdevice, nil
}
