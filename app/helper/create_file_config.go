package helper

import (
	"fmt"
	"os"
)

func CreateConfigFile(dirConfig string) {
	path := dirConfig + "/config.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package config\n\n"
		writeText += "import (\n"
		writeText += `	"database/sql"`
		writeText += "\n"
		writeText += `	"fmt"`
		writeText += "\n"
		writeText += `	"os"`
		writeText += "\n\n"
		writeText += `	_ "github.com/go-sql-driver/mysql"`
		writeText += "\n"
		writeText += `	"github.com/joho/godotenv"`
		writeText += "\n"
		writeText += `	_ "github.com/lib/pq"`
		writeText += "\n"
		writeText += ")\n\n"
		writeText += "func Connect() *sql.DB {\n"
		writeText += "	dbUrl := GetDbUrl()\n"
		writeText += `	connection := os.Getenv("DB_CONNECTION")`
		writeText += "\n"
		writeText += "	db, err := sql.Open(connection, dbUrl)\n"

		writeText += "	if err != nil {\n"
		writeText += "		panic(err.Error())\n"
		writeText += "	}\n"
		writeText += "	err = db.Ping()\n"
		writeText += "	if err != nil {\n"
		writeText += "		panic(err.Error())\n"
		writeText += "	}\n"
		writeText += "	return db\n"
		writeText += "}\n"
		writeText += "\n"

		writeText += "func GetDbUrl() string {\n"
		writeText += "	err := godotenv.Load()\n"
		writeText += "	if err != nil {\n"
		writeText += "		panic(err.Error())\n"
		writeText += "	}\n"
		writeText += `	var url = os.Getenv("DATABASE_URL")`
		writeText += "\n"
		writeText += `	if url == "" {`
		writeText += "\n"
		writeText += `		user := os.Getenv("DB_USER")`
		writeText += "\n"
		writeText += `		password := os.Getenv("DB_PASSWORD")`
		writeText += "\n"
		writeText += `		dbname := os.Getenv("DB_NAME")`
		writeText += "\n"
		writeText += `		port := os.Getenv("DB_PORT")`
		writeText += "\n"
		writeText += `		host := os.Getenv("DB_HOST")`
		writeText += "\n"
		writeText += `		localUrl := fmt.Sprintf("host=%s port=%s user=%s "+`
		writeText += "\n"
		writeText += `			"password=%s dbname=%s sslmode=disable",`
		writeText += "\n"
		writeText += `			host, port, user, password, dbname)`
		writeText += "\n"
		writeText += "		return localUrl\n"
		writeText += "	} else {\n"
		writeText += "		return url\n"
		writeText += "	}\n"

		writeText += "}\n"

		file.WriteString(writeText)
		defer file.Close()
	}
	// makeDirDatabase()
}
