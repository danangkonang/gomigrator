package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Kabupaten() {
	db := config.Connect()
	db.Exec(`DROP TABLE kabupaten`)
	_, err := db.Exec(`CREATE TABLE kabupaten(
	id_kabupaten serial PRIMARY KEY,
	kabupaten VARCHAR (225) NOT NULL,
	provinsi_id INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatal("kabupaten", err)
	}
}
