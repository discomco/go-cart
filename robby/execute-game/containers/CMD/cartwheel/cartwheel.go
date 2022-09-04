package cartwheel

import (
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/spokes"
	"log"
)

func RunCarthweel(cfgPath string) error {
	dig := BuildCartWheel(cfgPath)
	var runner spokes.IApp
	var appLogger logger.IAppLogger
	err := dig.Invoke(func(
		r spokes.IApp,
		l logger.IAppLogger) {
		runner = r
		appLogger = l
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	appLogger.Fatal(runner.Run())
	return nil
}
