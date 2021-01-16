package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func News() {
	db := config.Connect()
	db.Exec(`DROP TABLE news`)
	_, err := db.Exec(`CREATE TABLE news(
		id_news CHAR (32) PRIMARY KEY,
		user_id CHAR (32) NOT NULL,
		news_name VARCHAR (225) NOT NULL,
		images VARCHAR (225) NOT NULL,
		description TEXT NOT NULL,
		provinsi_id INTEGER NOT NULL,
		kabupaten_id INTEGER NOT NULL,
		kecamatan_id INTEGER NOT NULL,
		kelurahan_id INTEGER NOT NULL,
		address TEXT NULL,
		longitude DECIMAL NULL,
		langitude DECIMAL NULL,
		viewed INTEGER NULL,
		active BOOLEAN,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
