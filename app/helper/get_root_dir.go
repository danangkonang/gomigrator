package helper

import (
	"os"
	"strings"
)

func MyRootDir() string {
	// dir, _ := os.Getwd()
	dir, _ := os.Getwd()
	s := strings.Split(dir, "/")
	mydir := len(s) - 1
	thisDir := s[mydir]
	return thisDir
}
