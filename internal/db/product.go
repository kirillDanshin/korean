package db

import (
	"database/sql"
	"github.com/pkg/errors"
)

// nolint
var tableProduct = struct {
	name              string
	columnID          string
	columnBrandID     string
	columnName        string
	columnDescription string
	columnApply       string
	columnPrice       string
	columnAvatar      string
	columnIsDelete    string
}{
	name:              "products",
	columnID:          "product_id",
	columnBrandID:     "brand_id",
	columnName:        "name",
	columnDescription: "description",
	columnApply:       "apply",
	columnPrice:       "price",
	columnAvatar:      "avatar",
	columnIsDelete:    "is_delete",
}

const (
	schemeTableProducts = `create table if not exists products
(
    product_id        	 serial primary key,
	brand_id 			 int references brands (brand_id),
    name     		 	 text not null,
	description			 text not null,
	apply  				 text,
	price				 integer not null,
	avatar				 text,
	is_delete            boolean DEFAULT false,
	created_at 		 	 TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	set_at 		 	 	 TIMESTAMPTZ NOT NULL DEFAULT NOW()
)`
)

type (
	// Product - makeup unit.
	Product struct {
		ID          int            `db:"product_id"`
		Name        string         `db:"name"`
		Description string         `db:"description"`
		Apply       string         `db:"apply"`
		Price       int            `db:"price"`
		Avatar      sql.NullString `db:"avatar"`
		Brand       string         `db:"brand_name"`
	}
	// NewProduct - model for creating new product in storage.
	NewProduct struct {
		Name        string
		Description string
		Apply       string
		Price       int
		BrandID     int
	}
)

// BrandCreate - create new brand.
func (db *db) ProductCreate(ctx Ctx, product NewProduct) (*Product, error) {
	str := "INSERT INTO " + tableProduct.name +
		" (" + tableProduct.columnName + "," +
		tableProduct.columnBrandID + "," +
		tableProduct.columnDescription + "," +
		tableProduct.columnApply + "," +
		tableProduct.columnPrice +
		" ) " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING " + tableProduct.columnID

	var id int
	err := db.conn.QueryRowxContext(ctx, str,
		product.Name,
		product.BrandID,
		product.Description,
		product.Apply,
		product.Price,
	).Scan(&id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set new brand")
	}

	p, err := db.getProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (db *db) ProductByID(ctx Ctx, id int) (*Product, error) {
	return db.getProductByID(ctx, id)
}

// TODO refactoring sql query
func (db *db) getProductByID(ctx Ctx, id int) (*Product, error) {
	p := &Product{}
	str := "SELECT " + tableProduct.name + "." + tableProduct.columnID +
		"," + tableProduct.name + "." + tableProduct.columnPrice +
		"," + tableProduct.name + "." + tableProduct.columnApply +
		"," + tableProduct.name + "." + tableProduct.columnDescription +
		"," + tableProduct.name + "." + tableProduct.columnName +
		"," + tableProduct.name + "." + tableProduct.columnAvatar +
		"," + tableBrand.name + "." + tableBrand.columnName + " AS brand_name" +
		" FROM " + tableProduct.name +
		" LEFT JOIN " + tableBrand.name +
		" ON " + tableProduct.name + ".brand_id" +
		" = " + tableBrand.name + ".brand_id" +
		" WHERE product_id = $1"

	if err := db.conn.GetContext(ctx, p, str, id); err != nil {
		return nil, errors.Wrapf(err, "failed to getting product by id")
	}

	return p, nil
}

// ProductDelete - remove brand.
func (db *db) ProductDelete(ctx Ctx, productID int) (err error) {
	str := "DELETE FROM " + tableProduct.name + " WHERE " + tableProduct.columnID + "= $1"
	_, err = db.conn.ExecContext(ctx, str, productID)

	return errors.Wrapf(err, "failed to remove brand")
}

type Params struct {
	BrandID int
	Query   string
	Pagination
}

type Pagination struct {
	Limit  int
	Offset int
}

func (db *db) ListProduct(ctx Ctx, params Params) ([]Product, error) {
	switch {
	case params.BrandID > 0:
		return db.productsByBrand(ctx, params.BrandID, params.Pagination)
	case params.Query != "":
		return nil, nil
	default:
		return db.products(ctx, params.Pagination)
	}
}

func (db *db) products(ctx Ctx, pagination Pagination) ([]Product, error) {
	str := "SELECT " + tableProduct.name + "." + tableProduct.columnID +
		"," + tableProduct.name + "." + tableProduct.columnPrice +
		"," + tableProduct.name + "." + tableProduct.columnApply +
		"," + tableProduct.name + "." + tableProduct.columnDescription +
		"," + tableProduct.name + "." + tableProduct.columnName +
		"," + tableProduct.name + "." + tableProduct.columnAvatar +
		"," + tableBrand.name + "." + tableBrand.columnName + " AS brand_name" +
		" FROM " + tableProduct.name +
		" LEFT JOIN " + tableBrand.name +
		" ON " + tableProduct.name + ".brand_id" +
		" = " + tableBrand.name + ".brand_id" +
		" LIMIT $1 OFFSET $2"

	return db.getProducts(ctx, str, pagination.Limit, pagination.Offset)
}

// TODO refactoring sql query
func (db *db) productsByBrand(ctx Ctx, brandID int, pagination Pagination) ([]Product, error) {
	str := "SELECT " + tableProduct.name + "." + tableProduct.columnID +
		"," + tableProduct.name + "." + tableProduct.columnPrice +
		"," + tableProduct.name + "." + tableProduct.columnApply +
		"," + tableProduct.name + "." + tableProduct.columnDescription +
		"," + tableProduct.name + "." + tableProduct.columnName +
		"," + tableProduct.name + "." + tableProduct.columnAvatar +
		"," + tableBrand.name + "." + tableBrand.columnName + " AS brand_name" +
		" FROM " + tableProduct.name +
		" LEFT JOIN " + tableBrand.name +
		" ON " + tableProduct.name + ".brand_id" +
		" = " + tableBrand.name + ".brand_id" +
		" WHERE " + tableBrand.name + "." + tableBrand.columnID + " = $1 LIMIT $2 OFFSET $3"

	return db.getProducts(ctx, str, brandID, pagination.Limit, pagination.Offset)
}

func (db *db) getProducts(ctx Ctx, str string, args ...interface{}) ([]Product, error) {
	var products []Product
	if err := db.conn.SelectContext(ctx, &products, str, args...); err != nil {
		return nil, errors.Wrapf(err, "failed to getting products by since")
	}

	return products, nil
}
