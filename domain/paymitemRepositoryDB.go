package domain

import (
	"banking/logger"
	"banking/utils"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type PaymItemRepositoryDB struct {
	client *sqlx.DB
}

func (db PaymItemRepositoryDB) Create(p PaymItem) (int, error) {
	n := time.Now().Local()
	dt := n.Format("2006-01-02 15:04:05")

	balance, err := db.GetCurrentBalance(p.AccountId)
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	balance = balance + p.TAmount
	if balance < 0 {
		err = errors.New("impossible to overdraw")
		logger.Warn(err.Error())
		return 0, err
	}

	tx, err := db.client.Begin() // insert paymitem
	result, err := tx.Exec(`INSERT INTO paymitems 
		(account_id, tamount, transType, status, concept, datePost, dateValue)
		VALUES(?, ?, ?, ?, ?, ?, ?)`,
		p.AccountId,
		p.TAmount,
		p.TransType,
		1,
		p.Concept,
		dt,
		dt)
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		logger.Warn(err.Error())
		return 0, err
	}

	documentId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		logger.Warn(err.Error())
		return 0, err
	}

	err = db.UpdateAccountBalance(balance, p.AccountId, tx)
	if err != nil {
		tx.Rollback()
		logger.Warn(err.Error())
		return 0, err
	}

	tx.Commit()

	return int(documentId), nil
}

func (db PaymItemRepositoryDB) Reverse(documentId int) error {
	paymitem, err := db.GetPaymitem(documentId)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	balance, err := db.GetCurrentBalance(paymitem.AccountId)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	tx, err := db.client.Begin()
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	result, err := tx.Exec(`UPDATE paymitems
		SET status=6
		WHERE documentId=?`,
		documentId)
	if err != nil {
		tx.Rollback()
		logger.Warn(err.Error())
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		logger.Warn(err.Error())
		return err
	}

	balance = balance - paymitem.TAmount
	err = db.UpdateAccountBalance(balance, paymitem.AccountId, tx)
	if err != nil {
		tx.Rollback()
		logger.Warn(err.Error())
		return err
	}

	return nil
}

func (db PaymItemRepositoryDB) GetById(documentId int) (*PaymItem, error) {
	p, err := db.GetPaymitem(documentId)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func NewPaymItemRepositoryDB(ctx *context.Context) *PaymItemRepositoryDB {
	return &PaymItemRepositoryDB{
		client: utils.GetClientDB(ctx),
	}
}

func (db PaymItemRepositoryDB) GetCurrentBalance(accountId int) (float64, error) {
	var balance float64

	sqlRow := db.client.QueryRow(`SELECT balance
		FROM account
		WHERE account_id = ?`,
		accountId)

	err := sqlRow.Scan(&balance)
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}
	return balance, nil
}

func (db PaymItemRepositoryDB) GetPaymitem(documentId int) (*PaymItem, error) {
	var p PaymItem
	err := db.client.Select(&p, `SELECT 
		documentId, account_id, tamount, transType, status, concept, datePost, dateValue
		FROM paymitems
		WHERE documentId = ?`,
		documentId)
	if err != nil {
		logger.Warn(err.Error())
		return nil, err
	}
	return &p, nil
}

func (db PaymItemRepositoryDB) UpdateAccountBalance(balance float64, accountId int, tx *sql.Tx) error {
	result, err := tx.Exec(`UPDATE account
		SET balance=?
		WHERE account_id=?`,
		balance,
		accountId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return err
	}
	return nil
}
