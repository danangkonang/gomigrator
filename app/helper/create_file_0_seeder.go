package helper

import (
	"fmt"
	"os"
)

var (
	zero_seeder = "db/seeder/0.go"
)

func CreateZeroSeeder() {
	_, err := os.Stat(zero_seeder)

	if os.IsNotExist(err) {
		var file, err = os.Create(zero_seeder)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package seeder\n\n"

		writeText += "type Seeder struct{}\n"

		_, err = file.WriteString(writeText)
		if isErrors(err) {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()
	}
}
