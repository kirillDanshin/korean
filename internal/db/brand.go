package db

import (
	"github.com/pkg/errors"
)

// nolint
var tableBrand = struct {
	name           string
	columnID       string
	columnName     string
	columnIsDelete string
}{name: "brands", columnID: "brand_id", columnName: "brand_name", columnIsDelete: "is_delete"}

const (
	schemeTableBrands = `create table if not exists brands
(
    brand_id        	 serial primary key,
    brand_name     		 text not null,
	is_delete            boolean DEFAULT false,
	created_at 		 	 TIMESTAMPTZ NOT NULL DEFAULT NOW()
);  CREATE UNIQUE INDEX if not exists brands_unique_name ON brands(brand_name) WHERE is_delete = false;`
)

type (
	// NewBrand - model for creating new brand in storage.
	NewBrand struct {
		Name string
	}

	Brand struct {
		Id   int    `db:"brand_id"`
		Name string `db:"brand_name"`
	}
)

// BrandCreate - create new brand.
func (db *db) BrandCreate(ctx Ctx, brand NewBrand) ([]Brand, error) {
	str := "INSERT INTO " + tableBrand.name + " (" + tableBrand.columnName + ") " + "VALUES ($1) "
	_, err := db.conn.ExecContext(ctx, str, brand.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set new brand")
	}

	return db.getBrands(ctx)
}

func (db *db) getBrands(ctx Ctx) ([]Brand, error) {
	var brands []Brand
	str := "SELECT " + tableBrand.columnID + "," + tableBrand.columnName + " FROM " + tableBrand.name +
		" WHERE " + tableBrand.columnIsDelete + " <> $1"
	err := db.conn.SelectContext(ctx, &brands, str, true)

	return brands, errors.Wrapf(err, "failed to get brands")
}

// BrandDelete - remove brand.
func (db *db) BrandDelete(ctx Ctx, brandID int) (err error) {
	str := "UPDATE " + tableBrand.name + " SET is_delete=true WHERE " + tableBrand.columnID + "= $1"
	_, err = db.conn.ExecContext(ctx, str, brandID)

	return errors.Wrapf(err, "failed to remove brand")
}

func (db *db) GetBrands(ctx Ctx) ([]Brand, error) {
	return db.getBrands(ctx)
}
