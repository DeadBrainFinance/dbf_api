package schemas

type PartialUpdateInstallmentParams struct {
	UpdateName          bool
	Name                string
	UpdateTotalCost     bool
	TotalCost           float64
	UpdateInterestRate  bool
	InterestRate        float64
	UpdatePeriodNum     bool
	PeriodNum           int32
	UpdatePaidCost      bool
	PaidCost            float64
	UpdateCurrentPeriod bool
	CurrentPeriod       int32
	UpdatePeriodCost    bool
	PeriodCost          float64
	UpdateAccount       bool
	AccountID           int64
	ID                  int64
}

type CreateInstallmentParams struct {
	Name          string
	TotalCost     float64
	InterestRate  float64
	PeriodNum     int32
	PaidCost      float64
	CurrentPeriod int32
	PeriodCost    float64
	AccountID     int32
}
