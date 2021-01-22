package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

func InitLogger() *logrus.Logger {
	Log := logrus.New()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   os.Getenv("LOG_LOCATION"),
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     90, //days
		Level:      logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
		},
	})

	if err != nil {
		Log.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	Log.SetReportCaller(true)

	if os.Getenv("SERVER_ENV") == "PRODUCTION" {
		Log.AddHook(rotateFileHook)
	}

	return Log
}
