package schemas

import (
	"time"
)

type PartialUpdateDebtParams struct {
	ID                   int64
	UpdateName           bool
	Name                 string
	UpdateLender         bool
	Lender               string
	UpdateBorrower       bool
	Borrower             string
	UpdateInterestRate   bool
	InterestRate         float64
	UpdateBorrowedAmount bool
	BorrowedAmount       float64
	UpdatePaidAmount     bool
	PaidAmount           float64
	UpdateLendDate       bool
	LendDate             time.Time
}

type CreateDebtParams struct {
	Name           string
	Lender         string
	Borrower       string
	InterestRate   float64
	BorrowedAmount float64
	PaidAmount     float64
	LendDate       time.Time
}
