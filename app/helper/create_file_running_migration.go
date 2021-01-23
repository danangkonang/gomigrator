package helper

import (
	"fmt"
	"os"
)

func CreateRunningMigrationFile(thisDir, targetDir string) {
	path := targetDir + "/get_root_dir_name.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package execusion\n\n"
		writeText += "func ReadeFileInFolder() {\n"
		writeText += "}\n"

		file.WriteString(writeText)
		defer file.Close()
	}
}
