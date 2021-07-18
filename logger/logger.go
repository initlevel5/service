package logger

type Logger interface {
    Init(...Option)
    Log(level Level, v ...interface{})
    Options() Options
}

type Option func(*Options)

func Init(opts ...Option) error {}

func Log(level Level, v ...interface{}) {}
