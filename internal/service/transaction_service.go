package service

import (
	"errors"
	"examen-tecnico-stori/internal/model"
	"examen-tecnico-stori/internal/reader"
	"strconv"
	"time"
)

func ReadTransactions(path string) ([]*model.Transaction, error) {
	records, err := reader.ReadCSV(path)

	if err != nil {
		panic(err)
	}

	transactions := []*model.Transaction{}
	for i, record := range records {
		if i == 0 {
			if record[0] != "Id" && record[1] != "Date" && record[2] != "Transaction" {
				err := errors.New("invalid csv format")
				return nil, err
			}
		} else {

			id, err := strconv.Atoi(record[0])

			if err != nil {
				return nil, err
			}

			date, err := time.Parse("1/2", record[1])

			if err != nil {
				return nil, err
			}

			amount, err := strconv.ParseFloat(record[2], 64)

			if err != nil {
				return nil, err
			}

			transaction := &model.Transaction{
				Id:     id,
				Date:   date,
				Amount: amount,
			}

			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}

func ProcessTransactions(transactions []*model.Transaction) *model.Summary {
	transactionMap := make(map[time.Month]int)
	avarageCredit := 0.0
	avarageDebit := 0.0
	totalBalance := 0.0
	totalCredits := 0.0
	totalDebits := 0.0

	for _, transaction := range transactions {
		transactionMap[transaction.Date.Month()] += 1

		if transaction.Amount > 0 {
			avarageCredit += transaction.Amount
			totalCredits++
		} else {
			avarageDebit += transaction.Amount
			totalDebits++
		}

		totalBalance += transaction.Amount
	}

	monthlyTransactions := []*model.MonthlyTransaction{}

	for k, v := range transactionMap {

		monthlyTransaction := &model.MonthlyTransaction{
			Month: k,
			Total: v,
		}
		monthlyTransactions = append(monthlyTransactions, monthlyTransaction)

	}

	summary := &model.Summary{
		TotalBalance:        totalBalance,
		MonthlyTransactions: monthlyTransactions,
		AvarageDebit:        avarageDebit / totalDebits,
		AvarageCredit:       avarageCredit / totalCredits,
	}

	return summary

}
