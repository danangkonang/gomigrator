package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	dbUrl := GetDbUrl()
	connection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open(connection, dbUrl)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func GetDbUrl() string {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	var url = os.Getenv("DATABASE_URL")
	if url == "" {
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		host := os.Getenv("DB_HOST")
		localUrl := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		return localUrl
	} else {
		return url
	}
}
