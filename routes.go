package main

import (
	"github.com/gin-gonic/gin"
	"github.com/twcrone/space-mines/tc43/internal"
	"net/http"
)

func initializeRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/game", func(c *gin.Context) {
		c.JSON(http.StatusOK, internal.CreateNewGame("DEFAULT", 1, 4))
	})

	router.GET("/game/reveal", func(c *gin.Context) {
		sectorId := c.Query("sectorId")
		c.JSON(http.StatusOK, internal.RevealSector("DEFAULT", sectorId))
	})

	router.GET("/game/mark", func(c *gin.Context) {
		sectorId := c.Query("sectorId")
		c.JSON(http.StatusOK, internal.MarkSector("DEFAULT", sectorId))
	})
}
