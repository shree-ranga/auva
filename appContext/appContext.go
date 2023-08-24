package appContext

import (
	"context"

	"github.com/edgedb/edgedb-go"
)

type appContext struct {
	edgedbClient *edgedb.Client
}

var appCtx *appContext

func Init() {
	ctx := context.Background()
	appCtx = &appContext{}
	appCtx.edgedbClient = newDbConnection(ctx)
}

func newDbConnection(ctx context.Context) *edgedb.Client {
	opts := edgedb.Options{
		Concurrency: 4,
	}
	client, err := edgedb.CreateClient(ctx, opts)
	if err != nil {
		panic("could not connect to db")
	}
	return client
}

func GetDB() *edgedb.Client {
	return appCtx.edgedbClient
}
