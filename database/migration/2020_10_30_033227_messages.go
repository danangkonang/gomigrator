package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Messages() {
	db := config.Connect()
	db.Exec(`DROP TABLE messages`)
	_, err := db.Exec(`CREATE TABLE messages(
	id_product serial PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
