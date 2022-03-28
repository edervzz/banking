package utils

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func GetClientDB(ctx *context.Context) *sqlx.DB {
	return (*ctx).Value("clientdb").(*sqlx.DB)
}
