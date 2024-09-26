package zlog

type Logger interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(template string, args ...interface{})

	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(template string, args ...interface{})

	Panic(args ...interface{})
	Panicln(args ...interface{})
	Panicf(template string, args ...interface{})
}

func GetLogger() Logger {
	return l
}

func SetLogger(logger Logger) {
	l = logger
}

func Info(args ...interface{}) {
	l.Info(args...)
}

func Infoln(args ...interface{}) {
	l.Infoln(args...)
}

func Infof(template string, args ...interface{}) {
	l.Infof(template, args...)
}

func Debug(args ...interface{}) {
	l.Debug(args...)
}

func Debugln(args ...interface{}) {
	l.Debugln(args...)
}
func Debugf(template string, args ...interface{}) {
	l.Debugf(template, args...)
}

func Panic(args ...interface{}) {
	l.Panic(args...)
}

func Panicln(args ...interface{}) {
	l.Panicln(args...)
}

func Panicf(template string, args ...interface{}) {
	l.Panicf(template, args...)
}
