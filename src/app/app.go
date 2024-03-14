package app

import (
	"github.com/nazudis/disbursement/src/config"

	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Logger *logrus.Logger
)

func init() {
	logFile := viper.GetString(config.LogFile)
	currentTime := time.Now().Format("2006-01-02")
	f, err := os.OpenFile(fmt.Sprintf("%s-%s.log", logFile, currentTime), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	Logger.SetOutput(f)

	err = syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr: %v", err)
	}
}
