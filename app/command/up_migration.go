package command

import (
	"os"
	"os/exec"
	"strings"

	"github.com/danangkonang/migration-go-cli/app/model"
)

func UpMigration(up *model.UpDown) {
	// fmt.Println(up)
	// cmd := exec.Command("go", "run", "db/bin.go", "run", "migration", strings.Join(second, ","))
	cmd := exec.Command("go", "run", "db/bin.go", "up", "migration", "--tables="+strings.Join(up.Tables, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
