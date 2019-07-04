package db

import (
	"github.com/pkg/errors"
)

var tableBrand = struct {
	name           string
	columnID       string
	columnName     string
	columnIsDelete string
}{name: "brands", columnID: "brand_id", columnName: "name", columnIsDelete: "is_delete"}

const (
	schemeTableBrands = `create table if not exists brands
(
    brand_id        	 serial primary key,
    name     		 	 text not null unique,
	is_delete            boolean DEFAULT false,
	created_at 		 	 TIMESTAMPTZ NOT NULL DEFAULT NOW()
)`
)

type (
	NewBrand struct {
		Name string
	}
)

// BrandCreate - create new brand.
func (db *db) BrandCreate(ctx Ctx, brand NewBrand) (ID int, err error) {
	str := "INSERT INTO " + tableBrand.name + " (" + tableBrand.columnName + ") VALUES ($1) RETURNING " + tableBrand.columnID
	err = db.conn.QueryRowxContext(ctx, str, brand.Name).Scan(&ID)

	return ID, errors.Wrapf(err, "failed to set new brand")
}

// BrandDelete - remove brand.
func (db *db) BrandDelete(ctx Ctx, brandID int) (err error) {
	str := "UPDATE " + tableBrand.name + " SET is_delete=true WHERE " + tableBrand.columnID + "= $1"
	_, err = db.conn.ExecContext(ctx, str, brandID)

	return errors.Wrapf(err, "failed to remove brand")
}
