package helper

import (
	"fmt"
	"os"
)

func CreateGetTime(rootDir, targetDir string) {
	path := targetDir + "/get_time.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package helper\n\n"
		writeText += `import "time"`
		writeText += "\n\n"
		// function
		writeText += "func GetTime() string {\n"
		writeText += `	t := time.Now()`
		writeText += "\n"
		writeText += `	waktu := t.Format("2006-01-02 15:04:05")`
		writeText += "\n"
		writeText += "	return waktu\n"
		writeText += "}\n"

		file.WriteString(writeText)
		defer file.Close()
	}
}
