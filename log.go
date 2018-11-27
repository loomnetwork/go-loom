package loom

import (
	"io"
	"log"
	"os"
	"strings"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

type Logger struct {
	kitlog.Logger
}

func NewFilter(next kitlog.Logger, options ...kitlevel.Option) *Logger {
	return &Logger{kitlevel.NewFilter(next, options...)}
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
func (l *Logger) Info(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Info(l)
	args := append([]interface{}{msgKey, msg}, keyvals...)
	if err := kitlog.With(lWithLevel).Log(args...); err != nil {
		errLogger := kitlevel.Error(l)
		kitlog.With(errLogger, msgKey, msg).Log("err", err)
	}
}

// Debug logs a message at level Debug.
func (l *Logger) Debug(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Debug(l)
	args := append([]interface{}{msgKey, msg}, keyvals...)
	if err := kitlog.With(lWithLevel).Log(args...); err != nil {
		errLogger := kitlevel.Error(l)
		errLogger.Log("err", err)
	}
}

// Error logs a message at level Error.
func (l *Logger) Error(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Error(l)
	args := append([]interface{}{msgKey, msg}, keyvals...)
	if err := kitlog.With(lWithLevel).Log(args...); err != nil {
		args = append([]interface{}{"err", err}, args...)
		lWithLevel.Log(args)
	}
}

// Warn logs a message at level Debug.
func (l *Logger) Warn(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Warn(l)
	args := append([]interface{}{msgKey, msg}, keyvals...)
	if err := kitlog.With(lWithLevel).Log(args...); err != nil {
		errLogger := kitlevel.Error(l)
		kitlog.With(errLogger, msgKey, msg).Log("err", err)
	}
}

func NewLoomLogger(loomLogLevel, dest string) *Logger {
	w := MakeFileLoggerWriter(loomLogLevel, dest)
	logTr := func(w io.Writer) kitlog.Logger {
		fmtLogger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(w))
		return kitlog.With(fmtLogger, "ts", kitlog.DefaultTimestampUTC)
	}
	return MakeLoomLogger(loomLogLevel, w, logTr)
}

func MakeLoomLogger(logLevel string, w io.Writer, tr func(w io.Writer) kitlog.Logger) *Logger {
	loggerFunc := func(w io.Writer) *Logger {
		if len(logLevel) == 0 {
			logLevel = "info"
		}
		baseLogger := kitlog.With(tr(w), "module", "loom")
		return &Logger{
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
		log.Printf("Sending logs to %s\n", dest)
		f, err := os.OpenFile(destParts[1], os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(err)
		}
		return f
	}
}
