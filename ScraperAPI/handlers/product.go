package handlers

import (
	"log"

	"github.com/AmazonScraper/ScraperAPI/scraper"
	"github.com/AmazonScraper/ScraperAPI/service"
	"github.com/gin-gonic/gin"
)

type productScraper struct {
	sc     scraper.Scraper
	caller service.Caller
}

func NewScrapeHandler(sc scraper.Scraper, caller service.Caller) Scraper {
	return &productScraper{sc: sc, caller: caller}
}

func (ps *productScraper) Scrape(c *gin.Context) {
	params := struct {
		URL string
	}{}

	err := c.BindJSON(&params)
	if err != nil {
		log.Println("Error in parsing body: ", err.Error())
		c.JSON(500, "Internal Server Error")
	}

	document, err := ps.sc.Scrape(params.URL)
	if err != nil {
		log.Println("Error in scraping: ", err.Error())
		c.JSON(500, "Internal Server Error")
	}

	err = ps.caller.Call(document)
	if err != nil {
		log.Println("Error in service call: ", err.Error())
		c.JSON(500, "Internal Server Error")
	}

	c.JSON(201, gin.H{})
}
