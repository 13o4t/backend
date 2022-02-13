package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(opts ...Option) *zap.SugaredLogger {
	opt := defaultOptions()
	for _, o := range opts {
		o.apply(&opt)
	}

	core := zapcore.NewCore(encoder(opt), writer(opt), level(opt))

	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func level(opt options) zapcore.Level {
	var lvl zapcore.Level
	lvl.Set(opt.level)
	return lvl
}

func encoder(opt options) zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(cfg)
}

func writer(opt options) zapcore.WriteSyncer {
	sync := &lumberjack.Logger{
		Filename:   opt.filename,
		MaxSize:    opt.maxSize,
		MaxBackups: opt.maxBackups,
		MaxAge:     opt.maxAge,
		LocalTime:  opt.localtime,
		Compress:   opt.compress,
	}
	return zapcore.AddSync(sync)
}
