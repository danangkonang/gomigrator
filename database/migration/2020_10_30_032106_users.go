package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Users() {
	db := config.Connect()
	db.Exec(`DROP TABLE users`)
	_, err := db.Exec(`CREATE TABLE users(
	user_id CHAR(32) NOT NULL PRIMARY KEY,
	name VARCHAR (50) NOT NULL,
	email VARCHAR (225) NOT NULL,
	password VARCHAR (225) NOT NULL,
	user_role CHAR NOT NULL,
	is_active BOOLEAN,
	is_verify BOOLEAN,
	token VARCHAR (225) NULL,
	token_exp INTEGER NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
	)`)
	if err != nil {
		log.Fatal("users", err)
	}
}
