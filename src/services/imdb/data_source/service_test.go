package datasource

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPSQLDataSource_OpenConnection(t *testing.T) {
	ds := NewDataSource()
	ctx := context.Background()
	Current_connection_string = Local_connection_string
	conn, err := ds.OpenConnection(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	err = ds.CloseConnection(ctx)
	assert.NoError(t, err)
}

func TestPSQLDataSource_StartTransaction(t *testing.T) {
	ds := NewDataSource()
	ctx := context.Background()

	tx, err := ds.StartTransaction(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	err = ds.RollbackTransaction(ctx)
	assert.NoError(t, err)
}

func TestPSQLDataSource_CommitTransaction(t *testing.T) {
	ds := NewDataSource()
	ctx := context.Background()

	tx, err := ds.StartTransaction(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	err = ds.CommitTransaction(ctx)
	assert.NoError(t, err)
}
