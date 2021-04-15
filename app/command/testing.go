package command

import (
	"os"
	"os/exec"
)

func TestingComand() {
	cmd := exec.Command("go", "run", "migration/migration.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
