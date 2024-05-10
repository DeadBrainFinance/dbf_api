package models

type BalanceSheet struct {
	ID         int64
	Month      int32
	Year       int32
	Allocation float64
	Paid       float64
	Remaining  float64
	CategoryID int32
}
