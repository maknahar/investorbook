package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	// sql driver for database/sql.
	_ "github.com/lib/pq"
	// file system migrations.
	_ "github.com/mattes/migrate/source/file"
)

//nolint:gochecknoglobals
var (
	db *sql.DB
)

type Postgres struct {
}

func (p *Postgres) Connect(ctx context.Context, dbURL string, maxConn, maxIdleConn int) (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("%w; Unable to open db connection", err)
	}

	if maxConn == 0 {
		maxConn = 25
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)

	for i := 0; ; i++ {
		err = db.PingContext(ctx)
		if err != nil {
			if i < 10 {
				logrus.WithError(err).Warn("Unable to ping database. Retrying after 1 second")
				time.Sleep(time.Second)

				continue
			}

			return nil, fmt.Errorf("%w; Unable to ping database", err)
		}
		break
	}

	return db, nil
}
