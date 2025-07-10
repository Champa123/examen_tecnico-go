package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"examen-tecnico-stori/internal/service"

	"examen-tecnico-stori/internal/email"
)

func GetTransactions(c *gin.Context) {
	var transactions, err = service.ReadTransactions("../../transactions.csv")

	if err != nil {
		panic(err)
	}

	summary := service.ProcessTransactions(transactions)

	email.SendEmail(summary)

	c.IndentedJSON(http.StatusOK, summary)
}
