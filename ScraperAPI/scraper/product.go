package scraper

import (
	"log"
	"strconv"
	"strings"

	"github.com/AmazonScraper/ScraperAPI/models"
	"github.com/gocolly/colly"
)

type productScraper struct {
	collector *colly.Collector
}

func NewProductScaper(c *colly.Collector) Scraper {
	return &productScraper{collector: c}
}

func (ps *productScraper) Scrape(URL string) (*models.ProductDocument, error) {
	c := ps.collector

	productDocument := models.ProductDocument{URL: URL}

	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		name := e.DOM.Find(`span[id=productTitle]`).Text()
		productDocument.Product.Name = strings.Trim(name, "\n")
		log.Println(productDocument.Product.Name)

		reviewsText := e.DOM.Find(`span[id=acrCustomerReviewText]`).Text()
		reviewsText = strings.Trim(reviewsText, "\n")
		review := strings.Split(reviewsText, " ")
		productDocument.Product.TotalReviews, _ = strconv.Atoi(strings.Replace(review[0], ",", "", 1))
		log.Println(productDocument.Product.TotalReviews)

		price := e.DOM.Find(`span[id=priceblock_ourprice]`).Text()
		productDocument.Product.Price = strings.Trim(price, "\n")
		log.Println(productDocument.Product.Price)

		images := e.DOM.Find(`div[id=altImages]`)
		productDocument.Product.ImageURL, _ = images.Find(`img`).First().Attr("src")
		log.Println(productDocument.Product.ImageURL)

		description := e.DOM.Find(`div[id=productDescription]`)
		productDocument.Product.Description = description.Find(`p`).First().Text()
		log.Println(productDocument.Product.Description)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visting", r.URL)
	})

	c.Visit(URL)

	return &productDocument, nil
}
