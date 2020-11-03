package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func User_profiles() {
	db := config.Connect()
	db.Exec(`DROP TABLE user_profiles`)
	_, err := db.Exec(`CREATE TABLE user_profiles(
	id_product serial PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
