package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

func setupLogger(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logLevel)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2011.01-21 15:30:21",
	})

	logrus.SetOutput(os.Stdout)
}
