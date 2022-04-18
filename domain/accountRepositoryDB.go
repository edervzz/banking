package domain

import (
	"banking/logger"
	"banking/utils"
	"context"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (db AccountRepositoryDB) Create(a *Account) (int, error) {
	result, err := db.client.Exec(`INSERT INTO banking.account
		(customer_id, opening_date, account_type, balance, status)
		VALUES(?, ?, ?, ?, ?)`,
		a.CustomerId,
		a.OpeningDate,
		a.AccountType,
		a.Balance,
		a.Status)
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	accountId, err := result.LastInsertId()
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	return int(accountId), nil
}

func (db AccountRepositoryDB) GetBalance(accountId int) (float64, error) {
	var account Account
	sqlRow := db.client.QueryRow(`SELECT balance
		FROM banking.account
		WHERE account_id = ?`,
		accountId)

	err := sqlRow.Scan(&account.Balance)
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	return account.Balance, nil
}

func (db AccountRepositoryDB) Lock(accountId int) error {
	result, err := db.client.Exec(`UPDATE account
		SET status = ?
		WHERE account_id=?`,
		ACCT_LOCK, // lock
		accountId)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	return nil
}

func (db AccountRepositoryDB) Unlock(accountId int) error {
	result, err := db.client.Exec(`UPDATE banking.account
		SET status = ?
		WHERE account_id=?`,
		3, // active
		accountId)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	return nil
}

func NewAccountRepositoryDB(ctx *context.Context) AccountRepositoryDB {
	return AccountRepositoryDB{
		client: utils.GetClientDB(ctx),
	}
}
