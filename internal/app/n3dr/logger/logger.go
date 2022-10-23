package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/030/n3dr/internal/app/n3dr/project"
	"github.com/sirupsen/logrus"
)

func File(packageName string, log *logrus.Logger) error {
	level := os.Getenv("N3DR_LOG_LEVEL")
	if level != "" {
		h, err := project.Home()
		if err != nil {
			return err
		}
		dir := filepath.Join(h, "log")
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		f := filepath.Join(dir, time.Now().Format("20060102150405111")+".log")
		log.Infof("writing package: '%s' log to file: '%s'", packageName, f)
		file, err := os.OpenFile(filepath.Clean(f), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err == nil {
			log.SetFormatter(&logrus.JSONFormatter{})
			log.SetReportCaller(true)

			switch level {
			case "trace":
				log.SetLevel(logrus.TraceLevel)
			case "debug":
				log.SetLevel(logrus.DebugLevel)
			case "info":
				// default level is info
			default:
				return fmt.Errorf("log level: '%s' is invalid as it should be 'trace', 'debug' or 'info'", level)
			}

			log.Out = file
		} else {
			log.Info("failed to log to file, using default stderr")
		}
	}
	return nil
}
