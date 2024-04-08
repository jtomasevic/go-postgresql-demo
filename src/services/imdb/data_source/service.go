package datasource

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/store"
)

func NewDataSource() *PSQLDataSource {
	return &PSQLDataSource{}
}

type PSQLDataSource struct {
	conn *pgx.Conn
	tx   pgx.Tx
}

func (ds *PSQLDataSource) OpenConnection(ctx context.Context) (store.DBTX, error) {
	conn, err := pgx.Connect(ctx, connectionString())
	if err != nil {
		return nil, err
	}
	ds.conn = conn // Assign the connection to the conn field of PSQLDataSource
	return ds.conn, nil
}

func (ds *PSQLDataSource) CloseConnection(ctx context.Context) error {
	err := ds.conn.Close(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ds *PSQLDataSource) StartTransaction(ctx context.Context) (store.DBTX, error) {
	tx, err := ds.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	ds.tx = tx
	return tx, nil
}

func (ds *PSQLDataSource) RollbackTransaction(ctx context.Context) error {
	err := ds.tx.Rollback(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ds *PSQLDataSource) CommitTransaction(ctx context.Context) error {
	err := ds.tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ds *PSQLDataSource) GetConnection() store.DBTX {
	return ds.conn
}
