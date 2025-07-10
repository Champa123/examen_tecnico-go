package model

import "encoding/json"

type Summary struct {
	TotalBalance        float64
	MonthlyTransactions []*MonthlyTransaction
	AvarageDebit        float64
	AvarageCredit       float64
}

func (t Summary) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(b)
}
