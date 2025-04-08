package repositories

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os" 
	"errors"
)

type DB interface {
	Connect() error
	Disconnect() error
	GetClient() *mongo.Client
}

type MongoDB struct {
	MongoClient *mongo.Client
}

func NewMongoDB() *MongoDB {
	instancia := &MongoDB{}
	instancia.Connect()

	return instancia
}

func (mongoDB *MongoDB) GetClient() *mongo.Client {
	return mongoDB.MongoClient
}

func (mongoDB *MongoDB) Connect() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	connectionURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		err = errors.New("mongo connection failed")
		return err 
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		err = errors.New("ping failed")
		return err
	}

	mongoDB.MongoClient = client
	return nil

}

func (mongoDB *MongoDB) Disconnect() error {
	err := mongoDB.MongoClient.Disconnect(context.TODO())
	if err != nil {
		err = errors.New("disconnection failed")
		return err
	}
	return nil
}

