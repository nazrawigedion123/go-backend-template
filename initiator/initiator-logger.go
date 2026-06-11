package initiator

import (
	"os"

	"github.com/nazrawigedion123/go-backend-template/platform/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewConsoleLoggerWithLevel() *zap.Logger {
	lvl := zapcore.Level(viper.GetInt("logger.level"))

	// Console encoder (human-readable with colors)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Optional: Use JSON encoder for production-like output
	// consoleEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// Console core → logs everything >= lvl
	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= lvl
		}),
	)

	return zap.New(consoleCore, zap.AddCaller(), zap.AddCallerSkip(1))
}

func InitLogger() logger.Logger {
	lg := NewConsoleLoggerWithLevel()
	return logger.New(lg)
}
