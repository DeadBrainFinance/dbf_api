package models

type Installment struct {
	ID            int64
	Name          string
	TotalCost     float64
	InterestRate  float64
	PeriodNum     int32
	PaidCost      float64
	CurrentPeriod int32
	PeriodCost    float64
	AccountID     int32
}
