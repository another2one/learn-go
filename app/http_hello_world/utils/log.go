package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"learn-go/common/tool"
	"os"
)

var (
	Logger  *zap.SugaredLogger
	LogPath = tool.ProjectPath + "app/http_hello_world/http/logs"
)

func InitLogger() {
	log := zap.New(getCore(), zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = log.Sugar()
	defer Logger.Sync()
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(config)
}

func getCore() zapcore.Core {
	fileO1, err := os.OpenFile(LogPath+"info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("create file err: " + err.Error())
	}
	fileO2, err := os.OpenFile(LogPath+"error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("create file err: " + err.Error())
	}
	return zapcore.NewTee(
		zapcore.NewCore(getEncoder(), zapcore.AddSync(fileO2), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.ErrorLevel
		})),
		zapcore.NewCore(getEncoder(), zapcore.AddSync(fileO1), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl != zapcore.ErrorLevel
		})),
	)
}
