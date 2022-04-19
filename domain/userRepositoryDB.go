package domain

import (
	"banking/utils"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryDB struct {
	client *sqlx.DB
}

func (db UserRepositoryDB) Create(u *User) error {
	sqlResult, err := db.client.Exec(`INSERT INTO user
		(userId, email, password,role)
		VALUES(?,?,?,?)`,
		u.Username, u.Email, u.HashedPassword, u.Role)
	if err != nil {
		return err
	}
	_, err = sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (db UserRepositoryDB) Find(u *User) error {
	fmt.Println(u.Username, u.HashedPassword)
	sqlRow := db.client.QueryRow(`SELECT email, role 
		FROM user
		WHERE userId = ?
		AND password = ?`,
		u.Username,
		u.HashedPassword)

	err := sqlRow.Scan(&u.Email, &u.Role)
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
