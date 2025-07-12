package main

import (
	"examen-tecnico-stori/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/transactions/:path", controller.GetTransactions)

	router.Run("0.0.0.0:8080")
}
