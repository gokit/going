package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Logger option for configurations
type Config struct {
	Level      string                 `json:"level"`      // 日志级别：debug、info、warn、error、dpanic、panic、fatal
	File       string                 `json:"file"`       // 日志文件路径，不存在则报错
	MaxSize    int                    `json:"maxSize"`    // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int                    `json:"maxBackups"` // 日志文件最多保存多少个备份
	MaxAge     int                    `json:"maxAge"`     // 文件最多保存多少天
	Compress   bool                   `json:"compress"`   // 是否生成压缩文件
	Localtime  bool                   `json:"localtime"`  // 备份文件名称是否使用本地时间格式化，否则使用 UTC 时间
	Fields     map[string]interface{} `json:"fields"`     // 初始化字段
}

// Logger is a custom logger
type Logger struct {
	*zap.Logger
}

// NewZapLogger create a new zap.Logger
func NewLogger(config *Config) (*Logger, error) {

	hook := lumberjack.Logger{
		Filename:   config.File,       // 日志文件路径，不存在则报错
		MaxSize:    config.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: config.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     config.MaxAge,     // 文件最多保存多少天
		Compress:   config.Compress,   // 是否生成压缩文件
		LocalTime:  config.Localtime,  // 备份文件名称是否使用本地时间格式化，否则使用 UTC 时间
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	err := atomicLevel.UnmarshalText([]byte(viper.GetString("logger.level")))

	if err != nil {
		return nil, err
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 获取初始化字段
	var fs []zap.Field

	for k, v := range viper.GetStringMap("logger.fields") {
		fs = append(fs, zap.Any(k, v))
	}

	// 设置初始化字段
	fields := zap.Fields(fs...)

	// 构造日志
	logger := &Logger{
		zap.New(core, caller, development, fields),
	}

	return logger, nil
}
