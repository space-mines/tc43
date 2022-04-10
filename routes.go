package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/tc43/models"
	"net/http"
)

func initializeRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/game/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.FindGameById(c.Param("id")))
	})
}
