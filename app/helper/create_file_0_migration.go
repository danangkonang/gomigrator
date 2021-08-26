package helper

import (
	"fmt"
	"os"
)

var (
	zero_migration = "db/migration/0.go"
)

func CreateZeroMigration() {
	_, err := os.Stat(zero_migration)

	if os.IsNotExist(err) {
		var file, err = os.Create(zero_migration)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package migration\n\n"

		writeText += "import (\n"
		writeText += `	"database/sql"`
		writeText += "\n"
		writeText += `	"fmt"`
		writeText += "\n"
		writeText += `	"log"`
		writeText += "\n"
		writeText += `	"time"`
		writeText += "\n\n"
		writeText += `	_ "github.com/lib/pq"`
		writeText += "\n"
		writeText += ")\n\n"

		writeText += "type Migration struct{}\n\n"

		writeText += "type DB struct {\n"
		writeText += "	Db *sql.DB\n"
		writeText += "}\n\n"

		writeText += "func Connection() *DB {\n"
		writeText += "	connection := fmt.Sprintf(\n"
		writeText += `		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",`
		writeText += "\n"
		writeText += `		"localhost",`
		writeText += "\n"
		writeText += "		5432,\n"
		writeText += `		"postgres",`
		writeText += "\n"
		writeText += `		"postgres",`
		writeText += "\n"
		writeText += `		"default",`
		writeText += "\n"
		writeText += "	)\n"
		writeText += `	db, err := sql.Open("postgres", connection)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		log.Panic(err.Error())\n"
		writeText += "	}\n"
		writeText += "	db.SetMaxOpenConns(100)\n"
		writeText += "	db.SetMaxIdleConns(100)\n"
		writeText += "	db.SetConnMaxLifetime(5 * time.Minute)\n"
		writeText += "	return &DB{Db: db}\n"
		writeText += "}\n"

		_, err = file.WriteString(writeText)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()
	}
}
