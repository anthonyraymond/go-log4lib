package log4libwrapper

import (
	"fmt"
	"github.com/anthonyraymond/go-log4lib"
	"log"
	"os"
)

type builtinLoggerWrapper struct {
	delegateDebug  *log.Logger
	delegateInfo  *log.Logger
	delegateWarn  *log.Logger
	delegateError *log.Logger
	delegatePanic *log.Logger
	delegateFatal *log.Logger
}

func (w *builtinLoggerWrapper) Debug(args ...interface{}) {
	_ = w.delegateDebug.Output(2, fmt.Sprintln(args...))
}

func (w *builtinLoggerWrapper) Info(args ...interface{}) {
	_ = w.delegateInfo.Output(2, fmt.Sprintln(args...))
}

func (w *builtinLoggerWrapper) Warn(args ...interface{}) {
	_ = w.delegateWarn.Output(2, fmt.Sprintln(args...))
}

func (w *builtinLoggerWrapper) Error(args ...interface{}) {
	_ = w.delegateError.Output(2, fmt.Sprintln(args...))
}

func (w *builtinLoggerWrapper) Panic(args ...interface{}) {
	s := fmt.Sprint(args...)
	_ = w.delegatePanic.Output(2, s)
	panic(s)
}

func (w *builtinLoggerWrapper) Fatal(args ...interface{}) {
	_ = w.delegateFatal.Output(2, fmt.Sprint(args...))
	os.Exit(1)
}

func (w *builtinLoggerWrapper) Debugf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	_ = w.delegateDebug.Output(2, fmt.Sprintln(s))
}

func (w *builtinLoggerWrapper) Infof(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	_ = w.delegateInfo.Output(2, fmt.Sprintln(s))
}

func (w *builtinLoggerWrapper) Warnf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	_ = w.delegateWarn.Output(2, fmt.Sprintln(s))
}

func (w *builtinLoggerWrapper) Errorf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	_ = w.delegateError.Output(2, fmt.Sprintln(s))
}

func (w *builtinLoggerWrapper) Panicf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	_ = w.delegatePanic.Output(2, fmt.Sprintln(s))
	panic(s)
}

func (w *builtinLoggerWrapper) Fatalf(template string, args ...interface{}) {
	s := fmt.Sprintf(template, args...)
	_ = w.delegateFatal.Output(2, fmt.Sprintln(s))
	os.Exit(1)
}

func WrapBuiltinLogger(debugLogger *log.Logger, infoLogger *log.Logger, warnLogger *log.Logger, errorLogger *log.Logger, panicLogger *log.Logger, fatalLogger *log.Logger) log4lib.LibLogger {
	if debugLogger == nil {
		panic(fmt.Errorf("cannot create a BuiltInLoggerWrapper with a nil debugLogger"))
	}
	if infoLogger == nil {
		panic(fmt.Errorf("cannot create a BuiltInLoggerWrapper with a nil infoLogger"))
	}
	if warnLogger == nil {
		panic(fmt.Errorf("cannot create a BuiltInLoggerWrapper with a nil warnLogger"))
	}
	if errorLogger == nil {
		panic(fmt.Errorf("cannot create a BuiltInLoggerWrapper with a nil errorLogger"))
	}
	if panicLogger == nil {
		panic(fmt.Errorf("cannot create a BuiltInLoggerWrapper with a nil panicLogger"))
	}
	if fatalLogger == nil {
		panic(fmt.Errorf("cannot create a BuiltInLoggerWrapper with a nil fatalLogger"))
	}
	return &builtinLoggerWrapper{
		delegateDebug:  debugLogger,
		delegateInfo:  infoLogger,
		delegateWarn:  warnLogger,
		delegateError: errorLogger,
		delegatePanic: panicLogger,
		delegateFatal: fatalLogger,
	}
}

func Default() log4lib.LibLogger {
	return WrapBuiltinLogger(
		log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "WARN\t", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "PANIC\t", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "FATAL\t", log.Ldate|log.Ltime|log.Lshortfile),
	)
}
