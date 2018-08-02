package loom

import (
	"io"
	"log"
	"os"
	"strings"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

type Logger interface {
	With(keyvals ...interface{}) Logger
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Warn(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	Log(keyvals ...interface{}) error
}

type GoKitLogger struct {
	kitlog.Logger
}

func NewFilter(next kitlog.Logger, options ...kitlevel.Option) Logger {
	return &GoKitLogger{kitlevel.NewFilter(next, options...)}
}

var (
	AllowDebug = kitlevel.AllowDebug
	AllowInfo  = kitlevel.AllowInfo
	AllowWarn  = kitlevel.AllowWarn
	AllowError = kitlevel.AllowError
	Allow      = func(level string) kitlevel.Option {
		switch level {
		case "debug":
			return AllowDebug()
		case "info":
			return AllowInfo()
		case "warn":
			return AllowWarn()
		case "error":
			return AllowError()
		default:
			return nil
		}
	}
)

var msgKey = "_msg"

// Info logs a message at level Info.
func (l *GoKitLogger) Info(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Info(l)
	kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...)
}

// Log dont call this directly, its for compatibility for gokit
func (l *GoKitLogger) Log(keyvals ...interface{}) error {
	//	fmt.Printf("In Log -%v\n", keyvals...)
	//	fmt.Printf("In Log(inner) -%v\n", l.Logger)
	return l.Logger.Log(keyvals...)
}

// Debug logs a message at level Debug.
func (l *GoKitLogger) Debug(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Debug(l)
	kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...)
}

// Error logs a message at level Error.
func (l *GoKitLogger) Error(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Error(l)
	kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...)
}

// Warn logs a message at level Debug.
func (l *GoKitLogger) Warn(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Warn(l)
	kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...)
}

func (l *GoKitLogger) With(keyvals ...interface{}) Logger {
	return &GoKitLogger{kitlog.With(l.Logger, keyvals...)}
}

func NewLoomLogger(loomLogLevel, dest string) Logger {
	w := MakeFileLoggerWriter(loomLogLevel, dest)
	logTr := func(w io.Writer) kitlog.Logger {
		fmtLogger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(w))
		return kitlog.With(fmtLogger, "ts", kitlog.DefaultTimestampUTC)
	}
	return MakeLoomLogger(loomLogLevel, w, logTr)
}

func MakeLoomLogger(logLevel string, w io.Writer, tr func(w io.Writer) kitlog.Logger) *GoKitLogger {
	loggerFunc := func(w io.Writer) *GoKitLogger {
		baseLogger := kitlog.With(tr(w), "module", "loom")
		return &GoKitLogger{
			NewFilter(baseLogger, Allow(logLevel)),
		}
	}
	return loggerFunc(w)
}

func MakeFileLoggerWriter(loomLogLevel, dest string) io.Writer {
	if dest == "" {
		dest = "file://-"
	}
	destParts := strings.Split(dest, "://")
	if len(destParts) != 2 {
		log.Fatalf("Invalid log destination specification %s", dest)
	}
	// Keep below code as reference for syslog later
	//	case "syslog":
	//		sysLog, err := syslog.Dial("tcp", destParts[1],
	//			syslogLevel(loomLogLevel)|syslog.LOG_DAEMON, "loom")
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		Default = kitsyslog.NewSyslogLogger(sysLog, loggerFunc, kitsyslog.PrioritySelectorOption)
	if destParts[1] == "-" {
		return os.Stderr
	} else {
		log.Printf("Sending logs to file %s\n", destParts[1])
		f, err := os.OpenFile(destParts[1], os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(err)
		}
		return f
	}
}
