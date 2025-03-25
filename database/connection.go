package database

import (
	"context"

	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var Client *mongo.Client

func Collection() *mongo.Collection{
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("DATABASE")))
	if err != nil {
		log.Fatal(err)
	}




	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	return client.Database("GOLANG_DB").Collection("usuarios")
}


