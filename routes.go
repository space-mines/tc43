package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/tc43/services"
	"net/http"
)

func initializeRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/game", func(c *gin.Context) {
		c.JSON(http.StatusOK, services.CreateNewGame("DEFAULT", 1, 3))
	})

	router.GET("/game/reveal", func(c *gin.Context) {
		sectorId := c.Query("sectorId")
		c.JSON(http.StatusOK, services.RevealSector("DEFAULT", sectorId))
	})

	router.GET("/game/mark", func(c *gin.Context) {
		sectorId := c.Query("sectorId")
		c.JSON(http.StatusOK, services.RevealSector("DEFAULT", sectorId))
	})
}
