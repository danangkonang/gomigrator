package helper

import (
	"fmt"
	"os"
)

func CreateEnvFile() {
	path := ".env"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		file.WriteString(`DB_CONNECTION=postgres
DB_HOST=localhost
DB_TYPE=postgres
DB_PORT=5432
DB_NAME=testing
DB_PASSWORD=postgres
DB_USER=postgres
`)
		defer file.Close()
	}
	// fmt.Println("file berhasil dibuat", path)
	// createConfigFile()

}
