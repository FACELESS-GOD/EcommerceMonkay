package Logger

import (
	"runtime/debug"

	"go.uber.org/zap"
)

type LoggerInterface interface {
	Log()
}

type Prod struct {
	logger *zap.Logger
}

func NewLoggerProd() *Prod {
	logger := zap.Must(zap.NewProduction())
	return &Prod{
		logger: logger,
	}
}

func (prod *Prod) Log() {

	defer prod.logger.Sync()
	prod.logger.Log(zap.ErrorLevel, string(debug.Stack()))

}

type Dev struct {
	logger *zap.Logger
}

func NewLoggerDev() *Dev {
	logger := zap.Must(zap.NewDevelopment())
	return &Dev{
		logger: logger,
	}
}

func (dev *Dev) Log() {

}
