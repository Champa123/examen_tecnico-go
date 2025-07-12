package service

import (
	"examen-tecnico-stori/internal/model"
	"testing"
	"time"
)

func TestReadTransactions(t *testing.T) {
	transactions, err := ReadTransactions("../../transactions.csv")

	if err != nil {
		t.Errorf("err should be nil, is %s", err)
	}

	expected := generateExpectedTransactions()

	for i, transaction := range transactions {
		if transaction.Id != expected[i].Id {
			t.Errorf("Id = %d; want %d", transaction.Id, expected[i].Id)
		}
		if transaction.Date != expected[i].Date {
			t.Errorf("Date = %v; want %v", transaction.Date, expected[i].Date)
		}
		if int(transaction.Amount) != int(expected[i].Amount) {
			t.Errorf("Amount = %v; want %v", transaction.Amount, expected[i].Amount)
		}

	}
}

func TestReadTransactionsBrokenCsv(t *testing.T) {
	_, err := ReadTransactions(getBasePath() + "badCSVFormat.csv")

	if err == nil {
		t.Errorf("err should not be nil, is %s", err)
	}

	expected := generateExpectedErrorOnBadCSVFormat()

	if err.Error() != expected {
		t.Errorf("error is different from expected, got: %s, wanted: %s", err.Error(), expected)
	}

}

func TestProcessTransactions(t *testing.T) {

	summary := ProcessTransactions(generateExpectedTransactions())

	expected := generateExpectedSummary()

	if summary.AvarageCredit != expected.AvarageCredit {
		t.Errorf("AverageCredit = %f; want %f", summary.AvarageCredit, expected.AvarageCredit)
	}

	if summary.AvarageDebit != expected.AvarageDebit {
		t.Errorf("AvarageDebit = %f; want %f", summary.AvarageDebit, expected.AvarageDebit)
	}

	if summary.TotalBalance != expected.TotalBalance {
		t.Errorf("TotalBalance = %f; want %f", summary.TotalBalance, expected.TotalBalance)
	}

	for i, monthlyTransaction := range summary.MonthlyTransactions {

		if monthlyTransaction.Month != expected.MonthlyTransactions[i].Month {
			t.Errorf("Month = %d; want %d in MonthlyTransaction[%d]", monthlyTransaction.Month, expected.MonthlyTransactions[i].Month, i)
		}

		if monthlyTransaction.Total != expected.MonthlyTransactions[i].Total {
			t.Errorf("Total = %d; want %d in MonthlyTransaction[%d]", monthlyTransaction.Total, expected.MonthlyTransactions[i].Total, i)
		}

	}

}

func generateExpectedTransactions() []*model.Transaction {
	return []*model.Transaction{
		{
			Id:     0,
			Date:   time.Date(0000, time.July, 15, 0, 0, 0, 0, time.UTC),
			Amount: 60.5,
		},
		{
			Id:     1,
			Date:   time.Date(0000, time.July, 28, 0, 0, 0, 0, time.UTC),
			Amount: -10.3,
		},
		{
			Id:     2,
			Date:   time.Date(0000, time.August, 2, 0, 0, 0, 0, time.UTC),
			Amount: -20.46,
		},
		{
			Id:     3,
			Date:   time.Date(0000, time.August, 13, 0, 0, 0, 0, time.UTC),
			Amount: 10,
		},
	}
}

func generateExpectedSummary() *model.Summary {

	return &model.Summary{
		TotalBalance: 39.74,
		MonthlyTransactions: []*model.MonthlyTransaction{
			{
				Month: time.July,
				Total: 2,
			},
			{
				Month: time.August,
				Total: 2,
			},
		},
		AvarageDebit:  -15.38,
		AvarageCredit: 35.25,
	}

}

func generateExpectedErrorOnBadCSVFormat() string {
	return "invalid csv format"
}

func getBasePath() string {
	return "../../resources/test/"
}
