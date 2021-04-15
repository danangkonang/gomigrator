package helper

import (
	"fmt"
	"os"
)

func CreateMainFile(thisDir, dirMigration string) {
	path := dirMigration + "/migration.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package main\n\n"
		writeText += "import (\n"
		writeText += `	"os"`
		writeText += "\n\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/migration/app/helper"`
		writeText += "\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/migration/app/execusion"`
		writeText += "\n"
		writeText += ")\n\n"

		// function
		writeText += "func main() {\n"
		writeText += "	arrCmd := os.Args[1:]\n"
		writeText += "	if len(arrCmd) == 0 {\n"
		writeText += "		helper.PrintHelper()\n"
		writeText += "		return\n"
		writeText += "	}\n"
		writeText += "	runCmd()\n"
		writeText += "}\n"
		writeText += "\n"

		writeText += "func runCmd() {\n"
		writeText += "	usrCmd := os.Args[1]\n"
		writeText += "	switch usrCmd {\n"
		writeText += `	case "migration":`
		writeText += "\n"
		writeText += "		execusion.RuningMigration()\n"
		writeText += `	case "seeder":`
		writeText += "\n"
		writeText += "		execusion.RuningSeeder() \n"
		writeText += `	case "down":`
		writeText += "\n"
		writeText += "		execusion.DownTables() \n"
		writeText += `	case "drop":`
		writeText += "\n"
		writeText += "		execusion.DropTables() \n"
		writeText += "	default:\n"
		writeText += "		helper.PrintHelper()\n"
		writeText += "	}\n"
		writeText += "}\n"

		file.WriteString(writeText)
		defer file.Close()
	}
}
