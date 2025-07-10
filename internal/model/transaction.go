package model

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	Id     int
	Date   time.Time
	Amount float64
}

func (t Transaction) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(b)
}
