package schemas

import "time"

type CreateTransactionParams struct {
	Name string
	Cost float64
	Time time.Time
}

type PartialUpdateTransactionParams struct {
	UpdateName bool
	Name       string
	UpdateCost bool
	Cost       float32
	UpdateTime bool
	Time       time.Time
	ID         int64
}
