package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(any(err))
	}
	zap.ReplaceGlobals(logger)
}
