package db

import (
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nolint:golint
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"golang.org/x/net/context"
)

type (
	// Storage - main store products and brands.
	Storage interface {
		BrandCreate(ctx Ctx, brand NewBrand) ([]Brand, error)
		BrandDelete(ctx Ctx, brandID int) (err error)
		GetBrands(ctx Ctx) ([]Brand, error)

		ProductCreate(ctx Ctx, product NewProduct) (*Product, error)
		ProductDelete(ctx Ctx, productID int) (err error)
		ProductByID(ctx Ctx, id int) (*Product, error)
		ListProduct(ctx Ctx, params Params) ([]Product, error)
	}

	db struct {
		conn *sqlx.DB
	}
	// Configuration for connection database.
	Configuration struct {
		Type     string
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	// Ctx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

// New - return new db.
func New(conn *sqlx.DB) (Storage, error) {
	db := &db{conn: conn}
	if err := db.migrate(); err != nil {
		return nil, errors.Wrapf(err, "failed to migrate")
	}

	return db, nil
}

// nolint:gochecknoglobals
var tables = []string{schemeTableBrands, schemeTableProducts}

func (db *db) migrate() error {
	_, err := db.conn.Exec(strings.Join(tables, ";"))

	return errors.Wrapf(err, "failed migrate tables")
}
