package helper

import (
	"time"
)

func GetTime() string {
	t := time.Now()
	waktu := t.Format("2006_01_02_150405")
	return waktu
}
