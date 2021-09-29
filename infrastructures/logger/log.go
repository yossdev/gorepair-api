package logger

import (
	"gorepair-rest-api/config"
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

var once sync.Once

var Log *logrus.Logger

func init() {
	once.Do(func() {
		Log = newLogger()
	})
}

func newLogger() *logrus.Logger {
	Log := logrus.New()
	log.Println("Setup Logger")
	if config.Get().LogPath != "" {
		err := os.Mkdir(config.Get().LogPath, 0755)
		if err != nil {
			log.Println("Failed to create log path")
		}
		log.Println("Success to create log path")
	}

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   config.Get().LogPath,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     90, //days
		Level:      logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
			ForceColors:     true,
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

	Log.AddHook(rotateFileHook)

	return Log
}