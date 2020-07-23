package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"koala/src/log"
	"koala/src/util/configUtil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var Logger *zap.Logger

// InitLogger 初始化Logger
func InitRouteLogger() (err error) {
	routerLogFile := configUtil.GetRouterLogFile()

	maxSize := configUtil.GetLogMaxSize()
	maxBackups := configUtil.GetLogMaxBackups()
	maxAge := configUtil.GetLogMaxAge()

	writeSyncer := getLogWriter(routerLogFile, maxSize, maxBackups, maxAge)
	encoder := getEncoder()

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writeSyncer),
		zap.NewAtomicLevelAt(log.GetZapLogLevel()), // 日志级别
		)

	Logger = zap.New(core, zap.AddCaller())
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(&lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		//cost := time.Since(start)
		latency := time.Since(start)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				//zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("latency", latency),
			)
		}
	}
}

//func GinLogger(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		start := time.Now()
//		// some evil middlewares modify this values
//		path := c.Request.URL.Path
//		query := c.Request.URL.RawQuery
//		c.Next()
//
//		end := time.Now()
//		latency := end.Sub(start)
//		if utc {
//			end = end.UTC()
//		}
//
//		if len(c.Errors) > 0 {
//			// Append error field if this is an erroneous request.
//			for _, e := range c.Errors.Errors() {
//				logger.Error(e)
//			}
//		} else {
//			logger.Info(path,
//				zap.Int("status", c.Writer.Status()),
//				zap.String("method", c.Request.Method),
//				zap.String("path", path),
//				zap.String("query", query),
//				zap.String("ip", c.ClientIP()),
//				zap.String("user-agent", c.Request.UserAgent()),
//				zap.String("time", end.Format(timeFormat)),
//				zap.Duration("latency", latency),
//			)
//		}
//	}
//}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
