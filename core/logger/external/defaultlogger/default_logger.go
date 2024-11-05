package defaultLogger

import (
	"github.com/duongbuidinh600/core-package/core/constants"
	"github.com/duongbuidinh600/core-package/core/logger"
	"github.com/duongbuidinh600/core-package/core/logger/config"
	"github.com/duongbuidinh600/core-package/core/logger/logrous"
	"github.com/duongbuidinh600/core-package/core/logger/models"
	"os"
)

var l logger.Logger

func initLogger() {
	logType := os.Getenv("LogConfig_LogType")

	switch logType {

	case "Logrus":
		l = logrous.NewLogrusLogger(
			&config.LogOptions{LogType: models.Logrus, CallerEnabled: false},
			constants.Dev,
		)
		break
	default:
	}
}

func GetLogger() logger.Logger {
	if l == nil {
		initLogger()
	}

	return l
}
