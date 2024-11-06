package defaultLogger

import (
	"github.com/duongbui2002/core-package/core/constants"
	"github.com/duongbui2002/core-package/core/logger"
	"github.com/duongbui2002/core-package/core/logger/config"
	"github.com/duongbui2002/core-package/core/logger/logrous"
	"github.com/duongbui2002/core-package/core/logger/models"
)

var l logger.Logger

func initLogger() {
	l = logrous.NewLogrusLogger(
		&config.LogOptions{LogType: models.Logrus, CallerEnabled: false},
		constants.Dev,
	)

}

func GetLogger() logger.Logger {
	if l == nil {
		initLogger()
	}

	return l
}
