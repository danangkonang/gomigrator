package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Users() {
	db := config.Connect()
	db.Exec(`DROP TABLE users`)
	_, err := db.Exec(`CREATE TABLE users(
	user_id CHAR(32) NOT NULL,
	email VARCHAR (225) NOT NULL,
	password VARCHAR (225) NOT NULL,
	username VARCHAR (225) NULL,
	user_role INTEGER NOT NULL,
	is_active BOOLEAN,
	is_verify BOOLEAN,
	gender VARCHAR (225) NULL,
	telephone VARCHAR (225) NULL,
	avatar VARCHAR (225) NULL,
	address VARCHAR (225) NULL,
	kabupaten INTEGER NULL,
	kecamatan INTEGER NULL,
	kelurahan INTEGER NULL,
	provinsi INTEGER NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
