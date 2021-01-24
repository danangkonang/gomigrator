package command

import (
	"log"
	"os"
)

func MigrationUndo() {
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
