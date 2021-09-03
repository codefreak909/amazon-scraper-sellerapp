package stores

import (
	"context"
	"log"

	"github.com/AmazonScraper/StoreAPI/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type productStore struct {
	collection *mongo.Collection
}

func GetProductStore(c *mongo.Collection) Store {
	return &productStore{collection: c}
}

func (ps *productStore) Create(ctx context.Context, product models.ProductDocument) error {
	result, err := ps.collection.InsertOne(ctx, product)
	if err != nil {
		log.Println("Error inserting document: ", err.Error())
		return err
	}

	log.Printf("Document %d created successfully.", result.InsertedID)

	return nil
}
