package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/game/0", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
}
