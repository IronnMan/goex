package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goex/pkg/app"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var Logger *zap.Logger

// InitLogger log initialization
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	// get logs written to media
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	// set log level, see config/log.go file for details
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("Log initialization error, log level setting is wrong. Please modify the log.level configuration item in the config/log.go file")
	}

	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	Logger = zap.New(core,
		zap.AddCaller(),                   // calling file and line number, internally uses runtime.Caller
		zap.AddCallerSkip(1),              // encapsulate one layer, call the file to remove one layer (runtime.Call(1))
		zap.AddStacktrace(zap.ErrorLevel), // the stacktrace will only be displayed when there is an Error
	)

	zap.ReplaceGlobals(Logger)
}

func getEncoder() zapcore.Encoder {
	// log format rules
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if app.IsLocal() {
		// keyword highlighting for terminal output
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// locally set the built-in Console decoder (support stacktrace new line)
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// online environment uses JSON encoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder custom friendly time format
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter logging medium
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {

	if logType == "daily" {
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	if app.IsLocal() {
		// local development terminal prints and records files
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		// the production environment only records files
		return zapcore.AddSync(lumberJackLogger)
	}
}
