package handlers

import "github.com/gin-gonic/gin"

type Creater interface {
	Create(c *gin.Context)
}
