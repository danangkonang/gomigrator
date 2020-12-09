package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Categories() {
	db := config.Connect()
	db.Exec(`DROP TABLE categories`)
	_, err := db.Exec(`CREATE TABLE categories(
		id_category serial PRIMARY KEY,
		category VARCHAR (225) NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal("categories", err)
	}
}
