package schemas


type CreateBalanceSheetParams struct {
	Month      int32
	Year       int32
	Allocation float64
	Paid       float64
	Remaining  float64
	CategoryID int32
}

type PartialUpdateBalanceSheetParams struct {
	UpdateMonth      bool
	Month            int32
	UpdateYear       bool
	Year             int32
	UpdateAllocation bool
	Allocation       float64
	UpdatePaid       bool
	Paid             float64
	UpdateRemaining  bool
	Remaining        float64
	UpdateCategories bool
	CategoryID       int64
	ID               int64
}
