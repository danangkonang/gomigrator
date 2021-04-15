package command

import (
	"os"
	"os/exec"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

func MigrationRun() {
	fullCommand := os.Args[1:]
	if len(fullCommand) == 1 {
		helper.PrintHelper()
		return
	}
	subCommand := os.Args[2]
	switch subCommand {
	case "migration":
		runExecusin("migration")
	case "-m":
		runExecusin("migration")
	case "seeder":
		runExecusin("seeder")
	case "-s":
		runExecusin("seeder")
	default:
		helper.ErrorCommand(os.Args[2])
	}
}

func runExecusin(typeFlag string) {
	cmd := exec.Command("go", "run", "migration/migration.go", typeFlag)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
