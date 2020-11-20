package model

import "time"

type Transaction struct {
	ID              int
	AccountID       int
	OperationTypeID int
	Amount          float32
	EventDate       time.Time
}
