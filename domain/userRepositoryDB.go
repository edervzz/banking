package domain

import (
	"banking/utils"
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryDB struct {
	client *sqlx.DB
}

func (db UserRepositoryDB) Create(u *User) error {
	sqlResult, err := db.client.Exec(`INSERT INTO user
		(userId, email, password)
		VALUES(?, ?, ?)`,
		u.Username, u.HashedPassword, u.Email)
	if err != nil {
		return err
	}
	_, err = sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// constructor
func NewUserRepositoryDB(ctx *context.Context) *UserRepositoryDB {
	return &UserRepositoryDB{
		client: utils.GetClientDB(ctx),
	}
}
