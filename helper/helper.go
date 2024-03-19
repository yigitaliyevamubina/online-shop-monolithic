package helper

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"online_shop/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	database := client.Database("onlineshop")
	return database
}

func HandleError(err error, w http.ResponseWriter, statusCode int, methodName string) {
	log.Println(err, methodName)

	response := models.ErrorModel{
		StatusCode: statusCode,
		Message:    err.Error(),
	}

	message, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
