package schemas

type CreateAccountParams struct {
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

type PartialUpdateAccountParams struct {
	UpdateName         bool
	Name               string
	UpdateAccBalance   bool
	AccBalance         int32
	UpdateAccNum       bool
	AccNum             string
	UpdateCardNum      bool
	CardNum            string
	UpdatePin          bool
	Pin                string
	UpdateSecurityCode bool
	SecurityCode       string
	UpdateCreditLimit  bool
	CreditLimit        float64
	UpdateMethod       bool
	MethodID           int64
	ID                 int64
}
