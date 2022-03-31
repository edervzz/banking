package domain

import (
	"time"
)

const (
	PAYM_CREATED  = 1
	PAYM_REVERSED = 2
)

type PaymItem struct {
	DocumentId int       `db:"documentId"`
	AccountId  int       `db:"account_id"`
	TAmount    float64   `db:"tamount"`
	TransType  string    `db:"transType"`
	Status     int       `db:"status"`
	Concept    string    `db:"concept"`
	DatePost   time.Time `db:"datePost"`
	DateValue  time.Time `db:"dateValue"`
}

type PaymItemRepository interface {
	Create(p PaymItem) (documentId int, e error)
	Reverse(documentId int) error
	GetById(document int) (p *PaymItem, e error)
}
