package models

type Account struct {
	ID           int64
	Name         string
	// AccBalance   sql.NullFloat64
	AccBalance   float64
	AccNum       string
	CardNum      string
	Pin          string
	SecurityCode string
	// CreditLimit  sql.NullFloat64
	CreditLimit  float64
	MethodID     int32
}
