package main

import (
	"os"

	"github.com/AmazonScraper/ScraperAPI/handlers"
	"github.com/AmazonScraper/ScraperAPI/scraper"
	"github.com/AmazonScraper/ScraperAPI/service"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func main() {
	r := gin.Default()

	c := colly.NewCollector()

	productStoreAPIURL := os.Getenv("PRODUCT_STORE_API_URL")

	productScraper := scraper.NewProductScaper(c)
	serviceCaller := service.NewCaller(productStoreAPIURL)
	scrapeHandler := handlers.NewScrapeHandler(productScraper, serviceCaller)

	r.GET("/ping", handlers.PingHandler)
	r.POST("/scrape", scrapeHandler.Scrape)

	r.Run()
}
