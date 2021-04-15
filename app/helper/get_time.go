package helper

import (
	"os"
	"strings"
	"time"
)

func GetTime() string {
	t := time.Now()
	waktu := t.Format("2006_01_02_150405")
	return waktu
}

func RootDir() string {
	dir, _ := os.Getwd()
	s := strings.Split(dir, "/")
	mydir := len(s) - 1
	thisDir := s[mydir]
	return thisDir
}
