package middlewares

import (
	"context"
	"github.com/teris-io/shortid"
	"go.uber.org/zap"
	"net/http"
)

const requestIDKey = "requestID"

func RequestIDMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID, _ := shortid.Generate()
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)

			w.Header().Set("X-Request-ID", requestID)

			// Log the request ID using zap
			logger.Info("Processing request", zap.String("request_id", requestID))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}
