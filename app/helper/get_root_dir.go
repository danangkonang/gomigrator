package helper

import (
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

const (
	RED   = "\033[91m"
	RESET = "\033[0m"
)

func MyRootDir() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		dir, _ := os.Getwd()
		s := strings.Split(dir, "/")
		thisDir := s[len(s)-1]
		return thisDir
	}
	modName := modfile.ModulePath(goModBytes)
	arr := strings.Split(modName, "/")
	module := arr[len(arr)-1]
	return module
}

// func GetModuleName() string {
// 	goModBytes, err := ioutil.ReadFile("go.mod")
// 	if err != nil {
// 		exitf(func() {}, 1, "%+v\n", err)
// 	}

// 	modName := modfile.ModulePath(goModBytes)
// 	// fmt.Fprintf(os.Stdout, "modName=%+v\n", modName)

// 	return modName
// }

// func exitf(beforeExitFunc func(), code int, format string, args ...interface{}) {
// 	beforeExitFunc()
// 	fmt.Fprintf(os.Stderr, RED+format+RESET, args...)
// 	os.Exit(code)
// }
