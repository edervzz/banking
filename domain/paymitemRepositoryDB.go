package domain

import (
	"banking/logger"
	"time"

	"github.com/jmoiron/sqlx"
)

// type PaymItemRepository interface {
// 	Create(PaymItem) (int, error)
// 	Reverse(int) error
// 	GetById(int) (*PaymItem, error)
// }

type PaymItemRepositoryDB struct {
	client *sqlx.DB
}

func (db PaymItemRepositoryDB) Create(p PaymItem) (int, error) {
	n := time.Now().Local()
	dt := n.Format("2006-01-02 15:04:05")

	sqlResult, err := db.client.Exec(`INSERT INTO paymitems
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
		logger.Warn(err.Error())
		return 0, err
	}

	documentId, err := sqlResult.LastInsertId()
	if err != nil {
		logger.Warn(err.Error())
		return 0, err
	}

	return int(documentId), nil
}

func (db PaymItemRepositoryDB) Reverse(documentId int) error {
	sqlResult, err := db.client.Exec(`UPDATE paymitems
		SET status=6
		WHERE documentId=?`,
		documentId)
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	_, err = sqlResult.RowsAffected()
	if err != nil {
		logger.Warn(err.Error())
		return err
	}
	return nil
}

func (db PaymItemRepositoryDB) GetById(documentId int) (*PaymItem, error) {
	p := &PaymItem{}
	err := db.client.Select(p, `SELECT 
		documentId, account_id, tamount, transType, status, concept, datePost, dateValue
		FROM paymitems
		WHERE documentId = ?`,
		documentId)
	if err != nil {
		logger.Warn(err.Error())
		return nil, err
	}

	return p, nil
}
