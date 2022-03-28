package domain

type Customer struct {
	CustomerId int    `db:"id"`
	Fullname   string `db:"name"`
	City       string `db:"city"`
	Zipcode    string `db:"zipcode"`
}

type CustomerRepository interface {
	Create(*Customer) (int, error)
	Get(int) (*Customer, error)
}
