package handlers

import (
	"github.com/gin-gonic/gin"
)

type Scraper interface {
	Scrape(c *gin.Context)
}
