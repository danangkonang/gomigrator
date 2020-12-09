package migration

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Messages() {
	db := config.Connect()
	db.Exec(`DROP TABLE messages`)
	_, err := db.Exec(`CREATE TABLE messages(
	id_message serial PRIMARY KEY,
	msg_from VARCHAR (32) NOT NULL,
	msg_to VARCHAR (32) NOT NULL,
	message TEXT NOT NULL,
	is_read BOOLEAN,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal("messages", err)
	}
}
