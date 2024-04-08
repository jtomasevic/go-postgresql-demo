package imdb

import (
	"context"

	"github.com/pkg/errors"

	datasource "github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_source"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/repos"
)

var (
	ErrServerErrorOpenConnection   = "error in server while openening connection to db"
	ErrServerErrorCloseConnection  = "error in server while closing connection to db"
	ErrServerErrorStartTransaction = "error in server while starting transaction"
	ErrServerCommitTransaction     = "error in server while commiting transaction"
	ErrServerRollbackTransaction   = "error in server while rollback transaction"
)

var txStarted bool

func InitServices(ctx context.Context, withTx bool) (ImdbAPI, error) {
	dataSource := datasource.NewDataSource()
	var dataStore *repos.ImdbStores
	// **** with Transaction
	if withTx {
		_, err := dataSource.OpenConnection(ctx)
		tx, err := dataSource.StartTransaction(ctx)
		if err != nil {
			return ImdbAPI{}, errors.Wrap(err, ErrServerErrorStartTransaction)
		}
		txStarted = true
		dataStore = repos.NewImdbStore(tx)
	} else {
		// **** without Transaction
		_, err := dataSource.OpenConnection(ctx)
		if err != nil {
			return ImdbAPI{}, errors.Wrap(err, ErrServerErrorOpenConnection)
		}
		txStarted = false
		dataStore = repos.NewImdbStore(dataSource.GetConnection())
	}

	return ImdbAPI{
		DatSource: dataSource,
		ActorAPI:  NewActorAPI(dataStore.ActorStore),
		MovieAPI:  NewMovieAPI(dataStore.MovieStore),
	}, nil
}
