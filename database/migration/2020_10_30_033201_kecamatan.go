package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Kecamatan() {
	db := config.Connect()
	db.Exec(`DROP TABLE kecamatan`)
	_, err := db.Exec(`CREATE TABLE kecamatan(
	id_kecamatan serial PRIMARY KEY,
	kecamatan VARCHAR (225) NOT NULL,
	kabupaten_id INTEGER NOT NULL,
	provinsi_id INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatal("kecamatan", err)
	}
}
