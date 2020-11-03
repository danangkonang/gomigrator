package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Sub_items() {
	db := config.Connect()
	db.Exec(`DROP TABLE sub_items`)
	_, err := db.Exec(`CREATE TABLE sub_items(
		id_sub_item serial PRIMARY KEY,
		category_id INTEGER NOT NULL,
		sub_category_id INTEGER NOT NULL,
		sub_item VARCHAR (225) NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
