package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"koala/src/util/configUtil"
	"os"
)

var Logger *zap.Logger

func InitZapLogger() {
	logFile := configUtil.GetLogFile()
	logColor := configUtil.GetLogColor()

	maxSize := configUtil.GetLogMaxSize()
	maxBackups := configUtil.GetLogMaxBackups()
	maxAge := configUtil.GetLogMaxAge()

	hook := lumberjack.Logger{
		//Filename:   "./logs/spikeProxy1.log", // 日志文件路径
		Filename:   logFile,    // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   true,       // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		NameKey:  "logger",
		//CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      //默认/n分行
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 编码器
		EncodeTime:     zapcore.RFC3339TimeEncoder,     // RFC3339 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		//EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		//EncodeName: zapcore.FullNameEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	if logColor {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		//zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(GetZapLogLevel()), // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	//development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("initfiled", "init"))
	//构造日志
	Logger = zap.New(core, caller)


	Logger.Info("logger init success.")

	//Example
	//logger.Info("无法获取网址",
	// zap.String("url", "http://www.baidu.com"),
	// zap.Int("attempt", 3),
	// zap.Duration("backoff", time.Second))
}

func GetZapLogLevel() zapcore.Level {
	var level zapcore.Level
	logLevel := configUtil.GetLogLevel()
	switch logLevel {
	case "info":
		level = zap.InfoLevel
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	}

	return level
}


func Debug(msg string, fields ...zapcore.Field) {
	Logger.Debug(msg, fields...)
}
func Info(msg string, fields ...zapcore.Field) {
	Logger.Info(msg, fields...)
}
func Warn(msg string, fields ...zapcore.Field) {
	Logger.Warn(msg, fields...)
}
func Error(msg string, fields ...zapcore.Field) {
	Logger.Error(msg, fields...)
}
func Panic(msg string, fields ...zapcore.Field) {
	Logger.Panic(msg, fields...)
}
func Fatal(msg string, fields ...zapcore.Field) {
	Logger.Fatal(msg, fields...)
}
