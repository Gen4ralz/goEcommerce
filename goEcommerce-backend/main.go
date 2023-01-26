package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gen4ralz/go-Ecommerce/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	// app := controllers.
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	})
	// router.Use(middleware.Authentication())
	log.Fatal(router.Run(":" + port))
}