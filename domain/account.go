package domain

type Account struct {
	AccountId   int     `db:"account_id"`
	CustomerId  int     `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Balance     float64 `db:"balance"`
	Status      int     `db:"status"`
}

type AccountRepository interface {
	Create(Account) (int, error)
	GetBalance(int) (float64, error)
	Lock(int) error
}
