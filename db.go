package collecta

import (
	"context"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/lib/pq" // need that for open the ent client
	"github.com/minskylab/collecta/ent"
	"github.com/pkg/errors"
)

// DB represents a DB structure to wrap all related to database operations
type DB struct {
	Ent *ent.Client
}

// NewDB tries to open the db connection and return a DB instance
func NewDB(ctx context.Context) (*DB, error) {
	client, err := openDBConnection(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at open db connection")
	}
	return &DB{Ent: client}, nil
}

func openDBConnection(ctx context.Context) (*ent.Client, error) {
	drv, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=collecta password=collecta sslmode=disable")
	if err != nil {
		return nil, errors.Wrap(err, "error at open sql connection")
	}

	db := drv.DB()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(ctx); err != nil {
		return nil, errors.Wrap(err, "failed creating schema resources")
	}

	return client, nil
}