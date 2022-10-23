package artifacts

import (
	"github.com/030/n3dr/internal/app/n3dr/logger"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	if err := logger.File("internal/app/n3dr/artifactsv2/artifacts", log); err != nil {
		panic(err)
	}
}
