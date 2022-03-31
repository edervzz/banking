package domain

type Customer struct {
	CustomerId int    `db:"id"`
	Fullname   string `db:"name"`
	City       string `db:"city"`
	Zipcode    string `db:"zipcode"`
}

type CustomerRepository interface {
	Create(customer *Customer) (customerId int, e error)
	Get(customerId int) (customer *Customer, e error)
}
