package log

import (
	"context"
	"fmt"
	"github.com/juxuny/env"
	"github.com/juxuny/supervisor/trace"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"time"
)

var (
	logServerPrefix = env.GetString("LOG_SERVER_PREFIX")
)

const rpcLogTimeout = time.Millisecond * 10

type ColorFunc func(v interface{}) string

type Logger struct {
	fields           Fields // 输出固定参数
	disableCallStack bool
	level            log.Level
	ReportCaller     bool
	depth            int
	enableStdout     bool
}

type Fields log.Fields

func (f Fields) Get() log.Fields {
	return log.Fields(f)
}

// 给日志添加前缀字段，方便识别
func NewPrefix(moduleName string) *Logger {
	return New(Fields{
		"prefix": moduleName,
		"app":    env.GetString("LOG_SERVER_APP_NAME"),
	})
}

func New(fields ...Fields) *Logger {

	l := &Logger{}

	if len(fields) > 0 {
		l.fields = fields[0]
	}

	if env.GetString("LOG_LEVEL") == "debug" {
		l.SetLevel(log.DebugLevel)
		l.enableStdout = true
	} else {
		l.SetLevel(log.DebugLevel)
		l.enableStdout = false
	}
	l.ReportCaller = true
	return l
}

func (l *Logger) SetCallStackDepth(depth int) *Logger {
	l.depth = depth
	return l
}

func (l *Logger) SetLevel(level log.Level) {
	l.level = level
}

func (l *Logger) SetReportCaller(reportCaller bool) *Logger {
	l.ReportCaller = reportCaller
	return l
}

func (l *Logger) output(level string, v ...interface{}) string {
	var reqId = trace.GetReqId()
	var fieldList []string
	var hasUid bool
	for k, v := range l.fields {
		if strings.ToLower(k) == "uid" {
			hasUid = true
		}
		fieldList = append(fieldList, fmt.Sprintf("%s=%v", k, v))
	}
	if !hasUid {
		if uid, found := trace.GetUid(); found {
			fieldList = append(fieldList, fmt.Sprintf("uid=%d", uid))
		}
	}
	var levelOutput = level
	if level == "ERROR" {
		levelOutput = Color.Red(level)
	}
	var messages = []string{
		fmt.Sprintf("[%s]", time.Now().Format("200601-02 15:04:05.000")),
		fmt.Sprintf("<reqId=%s>", reqId),
		fmt.Sprintf("[%s]", levelOutput),
	}
	if l.ReportCaller {
		_, file, line, ok := runtime.Caller(2 + l.depth)
		if ok {
			position := Color.LightPurple(fmt.Sprintf("(%s:%d)", file, line))
			messages = append(messages, position)
		}
	}
	messages = append(messages, strings.Join(fieldList, " "))
	var str = strings.Join(messages, " ") + " "
	var joinSlice = func(values ...interface{}) string {
		ret := ""
		for _, item := range values {
			ret += fmt.Sprintf("%v ", item)
		}
		return strings.Trim(ret, " ")
	}
	str += joinSlice(v...)
	return str
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level < log.DebugLevel {
		return
	}
	message := l.output("DEBUG", v...)
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level < log.DebugLevel {
		return
	}
	message := l.output("DEBUG", fmt.Sprintf(format, v...))
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level < log.InfoLevel {
		return
	}
	message := l.output("INFO", v...)
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level < log.InfoLevel {
		return
	}
	message := l.output("INFO", fmt.Sprintf(format, v...))
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Print(v ...interface{}) {
	if l.level < log.InfoLevel {
		return
	}
	message := l.output("INFO", v...)
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Println(v ...interface{}) {
	if l.level < log.InfoLevel {
		return
	}
	message := l.output("INFO", v...)
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	if l.level < log.InfoLevel {
		return
	}
	message := l.output("INFO", fmt.Sprintf(format, v...))
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Warning(v ...interface{}) {
	if l.level < log.WarnLevel {
		return
	}
	message := l.output("WARN", v...)
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level < log.WarnLevel {
		return
	}
	message := l.output("WARN", v...)
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level < log.WarnLevel {
		return
	}
	message := l.output("WARN", fmt.Sprintf(format, v...))
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.level < log.ErrorLevel {
		return
	}
	message := l.output("ERROR", v...)
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level < log.ErrorLevel {
		return
	}
	message := l.output("ERROR", fmt.Sprintf(format, v...))
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.level < log.FatalLevel {
		return
	}
	message := l.output("ERROR", v...)
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.level < log.FatalLevel {
		return
	}
	message := l.output("ERROR", fmt.Sprintf(format, v...))
	if l.enableStdout {
		fmt.Println(message)
	}
	if rpcLogger != nil {
		ctx, cancel := context.WithTimeout(context.Background(), rpcLogTimeout)
		defer cancel()
		_ = rpcLogger.Add(ctx, logServerPrefix, message)
	}
}

// 日志输出时附带数据参数
func (l *Logger) WithMap(data Fields) *Logger {
	return l.Clone(data)
}

func (l *Logger) WithFields(f log.Fields) *Logger {
	return l.Clone(Fields(f))
}

// 复制实例
// fields 添加输出参数
func (l *Logger) Clone(fields ...Fields) *Logger {
	f := l.fields
	if len(fields) > 0 {
		for k, v := range fields[0] {
			f[k] = v
		}
	}
	return New(f)
}
