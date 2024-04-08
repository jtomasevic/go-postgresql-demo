package datasource

import (
	"context"

	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/store"
)

// put this into configuration.
var (
	Local_connection_string   = "user=imdb_user password=imdb_user dbname=imdb sslmode=disable port=5432"
	Docker_connection_string  = "user=imdb_user password=imdb_user dbname=imdb sslmode=disable port=5432 host=127.0.0.1"
	Current_connection_string = Docker_connection_string
)

func connectionString() string {
	return Current_connection_string
}

type DataSource interface {
	OpenConnection(ctx context.Context) (store.DBTX, error)
	CloseConnection(ctx context.Context) error
	StartTransaction(ctx context.Context) (store.DBTX, error)
	CommitTransaction(ctx context.Context) error
	RollbackTransaction(ctx context.Context) error
	GetConnection() store.DBTX
}
