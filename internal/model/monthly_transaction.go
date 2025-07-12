package model

import (
	"encoding/json"
	"time"
)

type MonthlyTransaction struct {
	Month time.Month
	Total int
}

func (t MonthlyTransaction) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(b)
}
