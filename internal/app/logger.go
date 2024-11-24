package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) {
	switch env {
	case envLocal:
		logrus.SetLevel(logrus.DebugLevel)
	case envDev:
		logrus.SetLevel(logrus.DebugLevel)
	case envProd:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2012-01-24 15:30:21",
	})

	logrus.SetOutput(os.Stdout)
}
