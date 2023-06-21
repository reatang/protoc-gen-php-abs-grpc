package log

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// 初始化日志功能，从opt参数中获取日志配置
func ConfigureLogging(paramsMap map[string]string) error {
	if paramsMap["logtostderr"] == "true" {
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: true,
		})
		log.SetOutput(os.Stderr)
		log.Debugf("Logging configured completed, logging has been enabled")
		levelStr := paramsMap["loglevel"]
		if levelStr != "" {
			level, err := log.ParseLevel(levelStr)
			if err != nil {
				return errors.Wrapf(err, "error parsing log level %s", levelStr)
			}

			log.SetLevel(level)
		} else {
			log.SetLevel(log.InfoLevel)
		}
	}

	return nil
}
