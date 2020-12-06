package db

import (
	"context"
	"database/sql"
)

// DB provide the contract that needs to be adhered to by any database used in this service.
type DB interface {
	// Connect should establish the connection with DB and return a connection pool
	Connect(ctx context.Context, dbURL string, maxConn, maxIdleConn int) (*sql.DB, error)
}

func New(environment string) DB {
	switch environment {
	case "Local":
		return &Postgres{}
	default:
		return &Postgres{}
	}
}
