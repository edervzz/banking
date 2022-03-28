package domain

import (
	"time"
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
	Create(PaymItem) (int, error)
	Reverse(int) error
	GetById(int) (*PaymItem, error)
}
