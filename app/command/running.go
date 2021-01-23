package command

import (
	"log"
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

func TipeRun() {
	fullCommand := os.Args[1:]
	if len(fullCommand) == 1 {
		helper.PrintHelper()
		return
	}
	subCommand := os.Args[2]
	switch subCommand {
	case "migration":
		// execusion.RuningMigration()
		break
	case "-m":
		// execusion.RuningMigration()
		break
	case "seeder":
		// execusion.RunSeeder()
		break
	case "-s":
		// execusion.RunSeeder()
		break
	default:
		helper.PrintHelper()
	}
}

func TypeBack() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	remove(dir)

}

func remove(dir string) {
	err := os.RemoveAll(dir + "/migration/")
	if err != nil {
		log.Fatal(err)
	}
}
