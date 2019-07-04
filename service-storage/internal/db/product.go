package db

import (
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

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
	Product struct {
		ID          int            `db:"product_id"`
		Name        string         `db:"name"`
		Description string         `db:"description"`
		Apply       string         `db:"apply"`
		Price       int            `db:"price"`
		Avatar      sql.NullString `db:"avatar"`
		Brand       string         `db:"brand_name"`
	}

	NewProduct struct {
		Name        string
		Description string
		Apply       string
		Price       int
		BrandID     int
	}
)

// BrandCreate - create new brand.
func (db *db) ProductCreate(ctx Ctx, product NewProduct) (ID int, err error) {
	str := "INSERT INTO " + tableProduct.name +
		" (" + tableProduct.columnName + "," +
		tableProduct.columnBrandID + "," +
		tableProduct.columnDescription + "," +
		tableProduct.columnApply + "," +
		tableProduct.columnPrice +
		" ) " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING " + tableProduct.columnID

	err = db.conn.QueryRowxContext(ctx, str,
		product.Name,
		product.BrandID,
		product.Description,
		product.Apply,
		product.Price,
	).Scan(&ID)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to set new brand")
	}

	p, err := db.getProductByID(ctx, ID)
	if err != nil {
		return 0, err
	}

	go func() { db.setProduct <- p }()

	return ID, nil
}

// TODO refactoring
func (db *db) getProductByID(ctx Ctx, id int) (Product, error) {
	p := Product{}
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

	return p, errors.Wrapf(db.conn.GetContext(ctx, &p, str, id), "failed to getting product by id")
}

// ProductDelete - remove brand.
func (db *db) ProductDelete(ctx Ctx, productID int) (err error) {
	str := "DELETE FROM " + tableProduct.name + " WHERE " + tableProduct.columnID + "= $1"
	_, err = db.conn.ExecContext(ctx, str, productID)
	if err != nil {
		return errors.Wrapf(err, "failed to remove brand")
	}

	go func() { db.delProduct <- productID }()

	return nil
}

func (db *db) GetProducts(ctx Ctx, since time.Time, ch chan<- Product) error {
	var products []Product
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
		" WHERE set_at > $1"

	if err := db.conn.SelectContext(ctx, &products, str, since); err != nil {
		return errors.Wrapf(err, "failed to getting products by since")
	}

	go func() {
		for i := range products {
			ch <- products[i]
		}
	}()

	return nil
}
