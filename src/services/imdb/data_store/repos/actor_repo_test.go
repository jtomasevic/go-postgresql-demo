package repos

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"

	datasource "github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_source"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/store"
)

func Test_AllActors(t *testing.T) {
	ctx := context.Background()

	dataSource := datasource.NewDataSource()
	conn, err := dataSource.OpenConnection(ctx)
	require.NoError(t, err)

	imbdStores := NewImdbStore(conn)
	actorStore := imbdStores.ActorStore

	actors, err := actorStore.AllActors(context.Background())
	require.NoError(t, err)
	fmt.Println(actors)
}
func ToDbUUID(v uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: v,
		Valid: true,
	}
}

func FromDbUUID(v pgtype.UUID) uuid.UUID {
	return v.Bytes
}
func Int32ToInt4(i int32) pgtype.Int4 {
	return pgtype.Int4{Int32: i, Valid: true}
}
func Test_NewActor(t *testing.T) {
	ctx := context.Background()

	dataSource := datasource.NewDataSource()
	conn, err := dataSource.OpenConnection(ctx)
	require.NoError(t, err)

	imbdStores := NewImdbStore(conn)
	actorStore := imbdStores.ActorStore

	// q := store.New(conn)
	id := uuid.New()
	var year int = 1972
	err = actorStore.CreateActor(context.Background(),
		store.CreateActorParams{
			ID:        id,
			Name:      "Gvinet Paltrou",
			Birthyear: &year,
		},
	)
	println(err)
	actor, err := actorStore.GetActor(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, "Gvinet Paltrou", actor.Name)
	require.Equal(t, year, *actor.Birthyear)
	require.Equal(t, id, actor.ID)
	fmt.Println(actor)
	require.NoError(t, err)

	err = dataSource.CloseConnection(ctx)
	require.NoError(t, err)
}
