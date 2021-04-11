package command

import (
	"os"
	"os/exec"
)

func Down(typeFlag string) {
	// fmt.Println(typeFlag)
	cmd := exec.Command("go", "run", "migration/migration.go", typeFlag)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
