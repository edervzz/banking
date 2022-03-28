package domain

import (
	"banking/logger"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

// from AccountRepository interface
func (db AccountRepositoryDB) CreateAccount(a *Account) (int, error) {
	result, err := db.client.Exec(`INSERT INTO account
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
	err := db.client.Select(&account, `SELECT balance
		FROM account
		WHERE account_id = ?`,
		accountId)

	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	return account.Balance, nil
}

func (db AccountRepositoryDB) LockAccount(accountId int) error {
	result, err := db.client.Exec(`UPDATE account
		SET status = ?
		WHERE account_id=?`,
		5, // lock
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
