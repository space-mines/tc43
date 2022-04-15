package main

import (
	"github.com/gin-contrib/cors"
	"github.com/twcrone/space-mines/tc43/internal"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	internal.InitializeRoutes(router)

	err := router.Run(":" + port)
	if err != nil {
		println(err.Error())
		return
	}
}
