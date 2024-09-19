package models

import (
	"time"
)

type Debt struct {
	ID              int64
	Name            string
	Lender          string
	Borrower        string
	InterestRate    float64
	BorrowedAmount  float64
	PaidAmount      float64
	RemainingAmount float64
	LendDate        time.Time // change this to datetime
}
