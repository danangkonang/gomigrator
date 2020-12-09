package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Notif() {
	db := config.Connect()
	db.Exec(`DROP TABLE notif`)
	_, err := db.Exec(`CREATE TABLE notif(
	id_notif serial PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
