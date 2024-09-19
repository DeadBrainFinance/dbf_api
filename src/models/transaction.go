package models

import (
	"time"
)

type Transaction struct {
	ID         int64
	Name       string
	Cost       float64
	Time       time.Time
	CategoryID int32
}
