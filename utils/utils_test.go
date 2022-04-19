package utils

import (
	"context"
	"net/http"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_GetClientDB_OK(t *testing.T) {
	clientdb := sqlx.DB{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "clientdb", &clientdb)
	context := GetClientDB(&ctx)
	assert.NotNil(t, context)
}

func Test_Messages_OK(t *testing.T) {
	m1 := NewNotFound("message")
	m2 := NewInternalError("message")
	m3 := NewBadRequest("message")

	assert.Equal(t, m1.Code, http.StatusNotFound)
	assert.Equal(t, m2.Code, http.StatusInternalServerError)
	assert.Equal(t, m3.Code, http.StatusBadRequest)
}
