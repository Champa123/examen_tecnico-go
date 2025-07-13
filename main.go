package main

import (
	"examen-tecnico-stori/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/transactions/summary/email", controller.GetTransactions)

	router.Run("0.0.0.0:8080")
}
