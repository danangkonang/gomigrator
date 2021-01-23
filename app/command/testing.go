package command

import (
	"os"
	"os/exec"
)

func TestingComand() {
	// var output1, err = exec.Command("go", "run", "migration/main.go").Output()
	// // fmt.Printf("%s", string(output1))
	// fmt.Println(err)
	// fmt.Println(output1)
	cmd := exec.Command("go", "run", "migration/migration.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal("%s\n", err)
	// }
}
