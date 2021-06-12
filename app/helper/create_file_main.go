package helper

import (
	"fmt"
	"os"
)

func CreateMainFile(thisDir, dirMigration string) {
	path := dirMigration + "/main.go"
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
		writeText += "\n"
		writeText += `	"strings"`
		writeText += "\n\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/migration/app/helper"`
		writeText += "\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/migration/app/execusion"`
		writeText += "\n"
		writeText += ")\n\n"

		// function
		writeText += "func main() {\n"
		writeText += "	if len(os.Args[1:]) == 0 {\n"
		writeText += "		helper.PrintHelper()\n"
		writeText += "		return\n"
		writeText += "	}\n"
		writeText += "	runCmd()\n"
		writeText += "}\n"
		writeText += "\n"

		writeText += "func runCmd() {\n"
		writeText += "	switch os.Args[1] {\n"
		writeText += `	case "run":`
		writeText += "\n"
		writeText += "		migrationOrSeeder()"
		writeText += "\n"

		writeText += `	case "reset":`
		writeText += "\n"

		writeText += "		var t execusion.Tables"
		writeText += "\n"
		writeText += `		if os.Args[2] != "" {`
		writeText += "\n"
		writeText += `			t.NameTable = strings.Split(os.Args[2], ",")`
		writeText += "\n"
		writeText += "		}"
		writeText += "\n"
		writeText += "		execusion.ResetTables(&t)"
		writeText += "\n"

		writeText += `	case "drop":`
		writeText += "\n"

		writeText += "		var t execusion.Tables"
		writeText += "\n"
		writeText += `		if os.Args[2] != "" {`
		writeText += "\n"
		writeText += `			t.NameTable = strings.Split(os.Args[2], ",")`
		writeText += "\n"
		writeText += "		}"
		writeText += "\n"
		writeText += "		execusion.DropTables(&t)"
		writeText += "\n"

		writeText += "	default:\n"
		writeText += "		helper.PrintHelper()\n"
		writeText += "	}\n"
		writeText += "}\n"
		writeText += "\n"

		writeText += "func migrationOrSeeder() {"
		writeText += "\n"
		writeText += "	switch os.Args[2] {"
		writeText += "\n"
		writeText += `	case "migration":`
		writeText += "\n"
		writeText += "		var t execusion.Tables"
		writeText += "\n"
		writeText += `		if os.Args[3] == "" {`
		writeText += "\n"
		writeText += "			execusion.RuningMigration(&t)"
		writeText += "\n"
		writeText += "		} else {"
		writeText += "\n"
		writeText += `			t.NameTable = strings.Split(os.Args[3], ",")`
		writeText += "\n"
		writeText += "			execusion.RuningMigration(&t)"
		writeText += "\n"
		writeText += "		}"
		writeText += "\n"
		writeText += `	case "seeder":`
		writeText += "\n"
		writeText += "		var t execusion.Tables"
		writeText += "\n"
		writeText += `		if os.Args[3] == "" {`
		writeText += "\n"
		writeText += "			execusion.RuningSeeder(&t)"
		writeText += "\n"
		writeText += "		} else {"
		writeText += "\n"
		writeText += `			t.NameTable = strings.Split(os.Args[3], ",")`
		writeText += "\n"
		writeText += "			execusion.RuningSeeder(&t)"
		writeText += "\n"
		writeText += "		}"
		writeText += "\n"
		writeText += "	default:"
		writeText += "\n"
		writeText += "		helper.PrintHelper()"
		writeText += "\n"
		writeText += "	}"
		writeText += "\n"
		writeText += "}"
		writeText += "\n"

		file.WriteString(writeText)
		defer file.Close()
	}

}
