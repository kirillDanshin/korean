package db

import (
	"github.com/pkg/errors"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/powerman/structlog"
	"golang.org/x/net/context"
)

type (
	db struct {
		conn       *sqlx.DB
		setProduct chan<- Product
		delProduct chan<- int
	}

	Configuration struct {
		Type     string
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	// Сtx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

// New - return new db.
func New(conn *sqlx.DB, chSet chan<- Product, chDel chan<- int) (*db, error) {
	db := &db{conn: conn, setProduct: chSet, delProduct: chDel}
	if err := db.migrate(); err != nil {
		return nil, errors.Wrapf(err, "failed to migrate")
	}

	return db, nil
}

var tables = []string{schemeTableBrands, schemeTableProducts}

func (db *db) migrate() error {
	_, err := db.conn.Exec(strings.Join(tables, ";"))

	return errors.Wrapf(err, "failed migrate tables")
}
