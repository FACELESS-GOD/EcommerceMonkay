package Middleware

import (
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"
)

type CustomLogger struct {
	logger *zap.Logger
}

func NewLoggerDev() *CustomLogger {
	logger := zap.Must(zap.NewDevelopment())

	return &CustomLogger{
		logger: logger,
	}

}

func NewLoggerProd() *CustomLogger {
	logger := zap.Must(zap.NewProduction())

	return &CustomLogger{
		logger: logger,
	}

}

func (log *CustomLogger) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(Writer http.ResponseWriter, Req *http.Request) {

			defer func() {
				if err := recover(); err != nil {
					defer log.logger.Sync()
					Writer.WriteHeader(http.StatusInternalServerError)

					log.logger.Log(
						zap.ErrorLevel, string(debug.Stack()))

				}
			}()
			next.ServeHTTP(Writer, Req)
		})
}
