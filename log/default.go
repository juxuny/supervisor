package log

var defaultLogger = NewPrefix("default").SetCallStackDepth(1)

func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}
func Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func Println(v ...interface{}) {
	defaultLogger.Println(v...)
}

func Print(v ...interface{}) {
	defaultLogger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}
func Warnf(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v...)
}
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}
