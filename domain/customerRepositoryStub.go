package domain

import (
	"context"
)

type CustomerRepositoryStub struct{}

func (stub CustomerRepositoryStub) Create(customer *Customer) (int, error) {
	return 1, nil
}

func (stub CustomerRepositoryStub) Get(customerId int) (*Customer, error) {
	c := Customer{
		CustomerId: 1,
		Fullname:   "Eder Velázquez",
		City:       "Atizapan",
		Zipcode:    "123321",
	}
	return &c, nil
}

// contructor
func NewCustomerRepositoryStub(ctx *context.Context) *CustomerRepositoryStub {
	return &CustomerRepositoryStub{}
}
