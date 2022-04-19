package domain

import (
	"banking/logger"
	"banking/utils"
	"context"
)

type CustomerRepositoryDB struct {
	client SqlDB
}

func (db *CustomerRepositoryDB) Create(customer *Customer) (int, error) {
	result, err := db.client.Exec(`INSERT INTO customer
		(name, city, zipcode)
		VALUES(?, ?, ?)`,
		customer.Fullname,
		customer.City,
		customer.Zipcode)
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	customerId, err := result.LastInsertId()
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	return int(customerId), nil
}

func (db *CustomerRepositoryDB) Get(customerID int) (*Customer, error) {
	var c Customer
	sqlRow := db.client.QueryRow(`SELECT id, name, city, zipcode
		FROM customer
		WHERE id = ?`,
		customerID,
	)
	err := sqlRow.Scan(&c.CustomerId, &c.Fullname, &c.City, &c.Zipcode)
	if err != nil {
		logger.Warn(err.Error())
		return nil, err
	}

	return &c, nil
}

// contructor
func NewCustomerRepositoryDB(ctx *context.Context) *CustomerRepositoryDB {
	return &CustomerRepositoryDB{
		client: utils.GetClientDB(ctx),
	}
}
