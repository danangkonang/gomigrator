package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Sub_categories() {
	db := config.Connect()
	db.Exec(`DROP TABLE sub_categories`)
	_, err := db.Exec(`CREATE TABLE sub_categories(
		id_sub_category serial PRIMARY KEY,
		category_id INTEGER NOT NULL,
		sub_category VARCHAR (225) NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal("sub_categories", err)
	}
}
