package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Kelurahan() {
	db := config.Connect()
	db.Exec(`DROP TABLE kelurahan`)
	_, err := db.Exec(`CREATE TABLE kelurahan(
	id_kelurahan serial PRIMARY KEY,
	kelurahan VARCHAR (225) NOT NULL,
	kecamatan_id INTEGER NOT NULL,
	kabupaten_id INTEGER NOT NULL,
	provinsi_id INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatal("kelurahan", err)
	}
}
