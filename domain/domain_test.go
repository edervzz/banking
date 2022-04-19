package domain

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// **************************************************
// Account
// **************************************************
func Test_NewAccountRepositoryDB_OK(t *testing.T) {
	clientdb := sqlx.DB{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "clientdb", &clientdb)
	result := NewCustomerRepositoryDB(&ctx)
	assert.NotNil(t, result)
}
