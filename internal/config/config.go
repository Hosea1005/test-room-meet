package config

import (
	"github.com/joho/godotenv"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"os"
)

func Logger() *zap.Logger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	return logger
}

func Environment() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			Logger().Error("error while loading .env file", zap.Error(err))
		}
	} else {
		Logger().Warn("running service without configuration from .env")
	}
}
