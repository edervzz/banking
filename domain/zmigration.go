package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type MigrationRepository interface {
	Prepare() bool
}

// ---------------------------------
type MigrationRepositoryDB struct {
	client *sqlx.DB
}

func (m MigrationRepositoryDB) Prepare() bool {
	var customers []Customer = []Customer{}
	findAllSql := "select id, name, city, zipcode from customer"
	err := m.client.Select(&customers, findAllSql)
	if err != nil {
		return false
	}
	return true
}

func NewMigrationRepositoryDB(ctx context.Context) MigrationRepositoryDB {
	clientdb := ctx.Value("clientdb").(*sqlx.DB)
	return MigrationRepositoryDB{
		client: clientdb,
	}
}
