package scraper

import "github.com/AmazonScraper/ScraperAPI/models"

type Scraper interface {
	Scrape(URL string) (*models.ProductDocument, error)
}
