package models

type Tracker struct {
	Expenses []*Expense
}

func NewTracker() *Tracker {
	return &Tracker{}
}
