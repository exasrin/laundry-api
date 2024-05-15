package middleware

import (
	"go-api-enigma/config"
	"go-api-enigma/model"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

func LogRequestMiddleware(log *logrus.Logger) gin.HandlerFunc {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)

	startTime := time.Now()

	return func(ctx *gin.Context) {
		ctx.Next()

		endTime := time.Since(startTime)
		requestLog := model.RequestLog{
			StartTime:  startTime,
			EndTime:    endTime,
			StatusCode: ctx.Writer.Status(),
			ClientIp:   ctx.ClientIP(),
			Method:     ctx.Request.Method,
			Path:       ctx.Request.URL.Path,
			UserAgent:  ctx.Request.UserAgent(),
		}

		switch {
		case ctx.Writer.Status() >= 500:
			log.Error(requestLog)
		case ctx.Writer.Status() >= 400:
			log.Warn(requestLog)
		default:
			log.Info(requestLog)
		}
	}
}
