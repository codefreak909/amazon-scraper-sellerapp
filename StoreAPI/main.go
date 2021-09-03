package main

import (
	"context"
	"log"
	"os"

	"github.com/AmazonScraper/StoreAPI/handlers"
	"github.com/AmazonScraper/StoreAPI/stores"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	collections := getMongoCollections()

	productStore := stores.GetProductStore(collections)
	productHandler := handlers.GetProductHandler(productStore)

	r.GET("/ping", handlers.PingHandler)
	r.POST("/product", productHandler.Create)

	r.Run(":8000")
}

func getMongoCollections() *mongo.Collection {
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://"+mongoHost+":"+mongoPort))
	if err != nil {
		log.Fatal("Database connection error: ", err.Error())
	}

	return client.Database("test").Collection("products")
}
