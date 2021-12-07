package main

import (
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day01"
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day02"
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day03"
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day04"
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day05"
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day06"
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day07"
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
		logger.Error("Something went wrong while solving day one", zap.String("message", err.Error()))
	}

	err = day02.Solve()
	if err != nil {
		logger.Error("Something went wrong while solving day two", zap.String("message", err.Error()))
	}

	err = day03.Solve()
	if err != nil {
		logger.Error("Something went wrong while solving day three", zap.String("message", err.Error()))
	}

	err = day04.Solve()
	if err != nil {
		logger.Error("Something went wrong while solving day four", zap.String("message", err.Error()))
	}

	err = day05.Solve()
	if err != nil {
		logger.Error("Something went wrong while solving day five", zap.String("message", err.Error()))
	}

	err = day06.Solve()
	if err != nil {
		logger.Error("Something went wrong while solving day six", zap.String("message", err.Error()))
	}

	err = day07.Solve()
	if err != nil {
		logger.Error("Something went wrong while solving day seven", zap.String("message", err.Error()))
	}
}
