package Middleware

import (
	"net/http"

	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Logger"
)

type CustomLogger struct {
	logger Logger.LoggerInterface
}

func NewCustomLogger(logInterface Logger.LoggerInterface) *CustomLogger {

	return &CustomLogger{
		logger: logInterface,
	}

}

func (log *CustomLogger) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(Writer http.ResponseWriter, Req *http.Request) {

			defer func() {
				if err := recover(); err != nil {
					Writer.WriteHeader(http.StatusInternalServerError)
					log.logger.Log()
				}
			}()
			next.ServeHTTP(Writer, Req)
		})
}
