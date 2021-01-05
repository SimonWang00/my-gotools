package log

import (
	"fmt"
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	NewLogger(&Options{
		LogFileDir: "logs",
		AppName:    "myblog",
		Level:      "debug",
	})
	log := GetLogger()
	for i := 0; i < 1; i++ {
		time.Sleep(time.Second / 2)
		log.Info(fmt.Sprint("test log ", i) )
		log.Debug(fmt.Sprint("debug log ", i))
		//log.Warn(fmt.Sprint("Info log ", i))
	}
}
