package main

import (
	"keymaster_go/models"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	models.DBConnect()
	initializeRoutes()

	router.Run("localhost:8080")
}
