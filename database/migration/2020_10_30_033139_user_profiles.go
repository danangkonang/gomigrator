package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func User_profiles() {
	db := config.Connect()
	db.Exec(`DROP TABLE user_profile`)
	_, err := db.Exec(`CREATE TABLE user_profile(
	id_profile serial PRIMARY KEY,
	user_id VARCHAR (32) NOT NULL REFERENCES users ON DELETE CASCADE,
	gender CHAR NULL,
	telephone CHAR (13) NULL,
	image VARCHAR (40) NULL,
	address VARCHAR (225) NULL,
	kelurahan CHAR(4) NULL,
	kecamatan CHAR(3) NULL,
	kabupaten CHAR(2) NULL,
	provinsi CHAR NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal("user_profiles", err.Error())
	}
}
