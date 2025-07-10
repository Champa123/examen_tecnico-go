package cmd

import (
	"examen-tecnico-stori/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/transactions", controller.GetTransactions)

	router.Run("localhost:8080")
}
