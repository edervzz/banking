package domain

import (
	"banking/logger"
	"banking/utils"
	"context"

	"github.com/jmoiron/sqlx"
)

type MigrationRepositoryDB struct {
	client *sqlx.DB
}

func (db MigrationRepositoryDB) Prepare() bool {
	_, err := db.client.Exec(`CREATE DATABASE banking`)
	if err != nil {
		logger.Warn(err.Error())
		return false
	}

	_, err = db.client.Exec(`CREATE TABLE banking.user (
		user_id INT auto_increment NOT NULL,
		password varchar(100) NOT NULL,
		email varchar(100) NOT NULL,
		role varchar(10) NOT NULL,
		CONSTRAINT user_PK PRIMARY KEY (user_id),
		CONSTRAINT user_UN UNIQUE KEY (email)
		)`)
	if err != nil {
		logger.Warn(err.Error())
	}

	_, err = db.client.Exec(`CREATE TABLE banking.customer (
		customer_id int NOT NULL AUTO_INCREMENT,
		name varchar(100) NOT NULL,
		city varchar(100) NOT NULL,
		zipcode varchar(10) NOT NULL,
		PRIMARY KEY (customer_id))`)
	if err != nil {
		logger.Warn(err.Error())
	}

	_, err = db.client.Exec(`CREATE TABLE banking.account (
		account_id int NOT NULL AUTO_INCREMENT,
		customer_id int NOT NULL,
		name varchar(100) NOT NULL,
		city varchar(100) NOT NULL,
		zipcode varchar(10) NOT NULL,
		PRIMARY KEY (account_id),
		CONSTRAINT fk_account_customer FOREIGN KEY (customer_id) REFERENCES customer(customer_id))`)
	if err != nil {
		logger.Warn(err.Error())
	}

	_, err = db.client.Exec(`CREATE TABLE banking.paymitems (
		payment_id int NOT NULL AUTO_INCREMENT,
		account_id int NOT NULL,
		t_amount float NOT NULL,
		trans_type varchar(4) NOT NULL,
		status int NOT NULL,
		concept varchar(50) NOT NULL,
		date_post date NOT NULL,
		date_value date NOT NULL,
		created_at datetime NOT NULL,
		PRIMARY KEY (account_id),
		UNIQUE KEY un_payment_account (payment_id,account_id),
		CONSTRAINT fk_payment_account FOREIGN KEY (account_id) REFERENCES account(account_id))`)
	if err != nil {
		logger.Warn(err.Error())
	}

	return true
}

func NewMigrationRepositoryDB(ctx *context.Context) *MigrationRepositoryDB {
	return &MigrationRepositoryDB{
		client: utils.GetClientDB(ctx),
	}
}
