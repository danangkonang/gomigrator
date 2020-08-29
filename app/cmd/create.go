package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var dir = "database"

func CreateTableMigrate(name string) {
	path := dir + "/" + helper.GetTime() + "-" + name + ".go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		file.WriteString("package database\n")
		defer file.Close()
	}
	fmt.Println("file berhasil dibuat", path)
}

func CmdType(usrCmd string) {
	subCmd := os.Args[2:]
	fullCmd := os.Args[1:]
	if len(fullCmd) == 1 {
		helper.PrintHelper()
		return
	}
	arrSubCmd := strings.Join(subCmd, "")
	newArrSubCmd := strings.Split(arrSubCmd, ":")

	firstKey := newArrSubCmd[0]
	secondKey := newArrSubCmd[1]
	switch firstKey {
	case "create":
		isMigrate(secondKey)
		break
	default:
		fmt.Println("default")
	}
	fmt.Println(firstKey, secondKey)

}

func isMigrate(secondKey string) {
	switch secondKey {
	case "migration":
		fmt.Println("masuk create migration")
		// isMigration(secondKey)
		break
	default:
		fmt.Println("default")
	}
}
