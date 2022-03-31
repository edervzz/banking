package domain

import "context"

type CustomerRepositoryDBMock struct {
}

func (mock CustomerRepositoryDBMock) Create(customer *Customer) (int, error) {
	return 1, nil
}

func (mock CustomerRepositoryDBMock) Get(customerID int) (*Customer, error) {
	c := Customer{
		CustomerId: 1,
		Fullname:   "Eder Velázquez",
		City:       "EDOMEX",
		Zipcode:    "12321",
	}
	return &c, nil
}

func NewCustomerRepositoryDBMock(ctx *context.Context) *CustomerRepositoryDBMock {
	return &CustomerRepositoryDBMock{}
}
