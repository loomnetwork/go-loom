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
	err := kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...)
	if err != nil {
		errLogger := kitlevel.Error(l)
		kitlog.With(errLogger, msgKey, msg).Log("err", err)
	}
}

// Debug logs a message at level Debug.
func (l *Logger) Debug(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Debug(l)
	if err := kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...); err != nil {
		errLogger := kitlevel.Error(l)
		errLogger.Log("err", err)
	}
}

// Error logs a message at level Error.
func (l *Logger) Error(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Error(l)
	lWithMsg := kitlog.With(lWithLevel, msgKey, msg)
	if err := lWithMsg.Log(keyvals...); err != nil {
		lWithMsg.Log("err", err)
	}
}

// Warn logs a message at level Debug.
func (l *Logger) Warn(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Warn(l)
	if err := kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...); err != nil {
		errLogger := kitlevel.Error(l)
		kitlog.With(errLogger, msgKey, msg).Log("err", err)
	}
}

func NewLoomLogger(loomLogLevel, dest string) *Logger {
	w := MakeFileLoggerWriter(loomLogLevel, dest)
	logTr := func(w io.Writer) kitlog.Logger {
		return kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(w))
	}
	return MakeLoomLogger(w, logTr)
}

func MakeLoomLogger(w io.Writer, tr func(w io.Writer) kitlog.Logger) *Logger {
	loggerFunc := func(w io.Writer) *Logger {
		baseLogger := kitlog.With(tr(w), "module", "loom")
		return &Logger{
			NewFilter(baseLogger, AllowDebug()),
		}
	}
	return loggerFunc(w)
}

func MakeFileLoggerWriter(loomLogLevel, dest string) io.Writer {
	if dest == "" {
		dest = "file://loom.log"
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
	log.Printf("Sending logs to %s\n", dest)
	if destParts[1] == "-" {
		return os.Stderr
	} else {
		f, err := os.OpenFile(destParts[1], os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(err)
		}
		return f
	}
}
