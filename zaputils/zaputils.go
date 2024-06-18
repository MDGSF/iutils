package zaputils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLog(traceID, spanID string) (*zap.Logger, *zap.SugaredLogger) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config.EncoderConfig),
		zapcore.Lock(os.Stdout),
		zap.InfoLevel,
	)

	traceField := zap.Field{Key: "trace", Type: zapcore.StringType, String: traceID}
	spanField := zap.Field{Key: "span", Type: zapcore.StringType, String: spanID}

	logger := zap.New(core, zap.AddCaller(), zap.Fields(traceField, spanField))
	sugar := logger.Sugar()
	return logger, sugar
}
