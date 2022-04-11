package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

var router *gin.Engine

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router = gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	initializeRoutes()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"https://space-mines.github.io"}
	corsConfig.AddAllowMethods("GET")

	// Register the middleware
	router.Use(cors.New(corsConfig))
	router.Run(":" + port)
}
