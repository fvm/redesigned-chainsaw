package main

import (
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day01"
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day02"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	var err error

	logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	err = day01.Solve()
	if err != nil {
		logger.Error("Something went wrong while solving day one", zap.Error(err))
	}

	err = day02.Solve()
	if err != nil {
		logger.Error("Something went wrong while solving day two", zap.Error(err))
	}
}
