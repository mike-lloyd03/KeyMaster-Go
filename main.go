package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	initializeRoutes()

	router.Run("localhost:8080")
}
