package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Provinsi() {
	db := config.Connect()
	db.Exec(`DROP TABLE provinsi`)
	_, err := db.Exec(`CREATE TABLE provinsi(
		id_provinsi serial PRIMARY KEY,
		provinsi VARCHAR (225) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
