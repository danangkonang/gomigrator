package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Products() {
	db := config.Connect()
	db.Exec(`DROP TABLE products`)
	_, err := db.Exec(`CREATE TABLE products(
		id_product CHAR (32) PRIMARY KEY,
		user_id CHAR (32) NOT NULL,
		product_name VARCHAR (225) NOT NULL,
		images VARCHAR (225) NOT NULL,
		price NUMERIC NOT NULL CHECK (price > 0),
		discount numeric,
		description TEXT NOT NULL,
		brand VARCHAR (225) NULL,
		address TEXT NOT NULL,
		condition BOOLEAN,
		stock INTEGER NOT NULL,
		category_id INTEGER NOT NULL,
		weight INTEGER NOT NULL,
		weight_unit CHAR (20) NOT NULL,
		active BOOLEAN,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		CHECK (discount > 0 AND price > discount)
	)`)
	if err != nil {
		log.Fatal("products", err)
	}
}
