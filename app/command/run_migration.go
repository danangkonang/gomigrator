package command

import (
	"os"
	"os/exec"
	"strings"

	"github.com/danangkonang/migration-go-cli/app/helper"
)

func MigrationsRun() {
	switch os.Args[2] {
	case "--hepl":
		helper.MultyPrintHelper("run")
	case "-h":
		helper.MultyPrintHelper("run")
	case "migration":
		execMigration("run", "migration", os.Args[3:])
	case "-m":
		execMigration("run", "migration", os.Args[3:])
	case "seeder":
		execMigration("run", "seeder", os.Args[3:])
	case "-s":
		execMigration("run", "seeder", os.Args[3:])
	default:
		helper.ErrorCommand(os.Args[2])
	}
}

func execMigration(firstFlag, typeFlag string, second []string) {
	cmd := exec.Command("go", "run", "migration/main.go", firstFlag, typeFlag, strings.Join(second, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
