package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"examen-tecnico-stori/internal/service"

	"examen-tecnico-stori/internal/email"
)

func GetTransactions(c *gin.Context) {

	var req SummaryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	var transactions, err = service.ReadTransactions(req.PathToFile)

	if err != nil {
		panic(err)
	}

	summary := service.ProcessTransactions(transactions)

	email.SendEmail(summary, req.UserMail)

	c.IndentedJSON(http.StatusOK, summary)
}
