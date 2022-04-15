package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/game", func(c *gin.Context) {
		c.JSON(http.StatusOK, CreateNewGame("DEFAULT", 1, 4))
	})

	router.GET("/game/reveal", func(c *gin.Context) {
		sectorId := c.Query("sectorId")
		sectorIdAsInt, _ := strconv.Atoi(sectorId)
		c.JSON(http.StatusOK, RevealSector("DEFAULT", sectorIdAsInt))
	})

	router.GET("/game/mark", func(c *gin.Context) {
		sectorId := c.Query("sectorId")
		sectorIdAsInt, _ := strconv.Atoi(sectorId)
		c.JSON(http.StatusOK, MarkSector("DEFAULT", sectorIdAsInt))
	})
}
