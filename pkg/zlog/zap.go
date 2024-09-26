package zlog

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var encoder zapcore.Encoder

type zapLogger struct {
	*zap.SugaredLogger
}

var l Logger = &zapLogger{}

func init() {
	encoder = getZapEncoder()
	writeSyncer := getZapWriter(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	l = logger.Sugar()
}

func newDefaultEncoderConfig() *zapcore.EncoderConfig {
	return &zapcore.EncoderConfig{
		// Key决定日志要输出什么内容
		TimeKey:          "time",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "msg",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,     // 定义每行之间的分隔符, 默认为换行
		EncodeLevel:      zapcore.LowercaseLevelEncoder, // 定义level信息输出的格式, 本质是一个函数
		EncodeTime:       zapcore.ISO8601TimeEncoder,    // 定义time信息输出的格式
		EncodeDuration:   zapcore.StringDurationEncoder, // 序列化时间格式
		EncodeCaller:     zapcore.ShortCallerEncoder,    // 定义代码调用的格式
		ConsoleSeparator: " ",
	}
}

func getZapEncoder() zapcore.Encoder {
	encoderConfig := newDefaultEncoderConfig()
	// Custom encoder to change the log format
	encoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeCaller = func(ec zapcore.EntryCaller, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(ec.TrimmedPath() + ":")
	}
	encoder = zapcore.NewConsoleEncoder(*encoderConfig)
	return encoder
}

func getZapWriter(writer io.Writer) zapcore.WriteSyncer {
	return zapcore.AddSync(writer)
}

func SetLogOutput(writers ...io.Writer) {
	encoder = getZapEncoder()
	var cores []zapcore.Core

	for _, writer := range writers {
		fileWriter := getZapWriter(writer)
		core := zapcore.NewCore(encoder, fileWriter, zap.DebugLevel)
		cores = append(cores, core)
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller())
	l = logger.Sugar()
}
