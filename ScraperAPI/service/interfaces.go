package service

import (
	"github.com/AmazonScraper/ScraperAPI/models"
)

type Caller interface {
	Call(document *models.ProductDocument) error
}
