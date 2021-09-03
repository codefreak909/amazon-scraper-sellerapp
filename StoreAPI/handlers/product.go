package handlers

import (
	"log"

	"github.com/AmazonScraper/StoreAPI/models"
	"github.com/AmazonScraper/StoreAPI/stores"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	store stores.Store
}

func GetProductHandler(s stores.Store) Creater {
	return &productHandler{store: s}
}

func (ph *productHandler) Create(c *gin.Context) {
	var productDocument models.ProductDocument

	err := c.BindJSON(&productDocument)
	if err != nil {
		log.Println("Error in json Unmarshaling : ", err.Error())
		c.JSON(500, gin.H{"error": "Internal Server Error"})
	}

	err = ph.store.Create(c.Request.Context(), productDocument)
	if err != nil {
		log.Println("Error in retrieving Document : ", err.Error())
		c.JSON(500, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(201, gin.H{})
}
