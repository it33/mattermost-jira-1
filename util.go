package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
)

func initLog(path string, verbose bool) *logrus.Logger {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Failed to open/create log -", err.Error())
		os.Exit(1)
	}

	f, err := os.OpenFile(absPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Failed to open/create log -", err.Error())
		os.Exit(1)
	}

	multi := io.MultiWriter(f, os.Stdout)

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(multi)
	log := logrus.New()
	log.Infof("Logs: %s", absPath)
	return log
}
