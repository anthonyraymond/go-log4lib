package log4libwrapper

import (
	"fmt"
	"github.com/anthonyraymond/go-log4lib"
	"log"
	"os"
)

type builtinLoggerWrapper struct {
	delegateInfo  *log.Logger
	delegateWarn  *log.Logger
	delegateError *log.Logger
	delegatePanic *log.Logger
	delegateFatal *log.Logger
}

func (w *builtinLoggerWrapper) Info(args ...interface{}) {
	w.delegateInfo.Println(args...)
}

func (w *builtinLoggerWrapper) Warn(args ...interface{}) {
	w.delegateWarn.Println(args...)
}

func (w *builtinLoggerWrapper) Error(args ...interface{}) {
	w.delegateError.Println(args...)
}

func (w *builtinLoggerWrapper) Panic(args ...interface{}) {
	w.delegatePanic.Panic(args...)
}

func (w *builtinLoggerWrapper) Fatal(args ...interface{}) {
	w.delegateFatal.Fatal(args...)
}

func (w *builtinLoggerWrapper) Infof(template string, args ...interface{}) {
	w.delegateInfo.Println(fmt.Sprintf(template, args...))
}

func (w *builtinLoggerWrapper) Warnf(template string, args ...interface{}) {
	w.delegateWarn.Println(fmt.Sprintf(template, args...))
}

func (w *builtinLoggerWrapper) Errorf(template string, args ...interface{}) {
	w.delegateError.Println(fmt.Sprintf(template, args...))
}

func (w *builtinLoggerWrapper) Panicf(template string, args ...interface{}) {
	w.delegatePanic.Println(fmt.Sprintf(template, args...))
}

func (w *builtinLoggerWrapper) Fatalf(template string, args ...interface{}) {
	w.delegateFatal.Println(fmt.Sprintf(template, args...))
}

func NewBuiltInLoggerWrapper(infoLogger *log.Logger, warnLogger *log.Logger, errorLogger *log.Logger, panicLogger *log.Logger, fatalLogger *log.Logger) log4lib.LibLogger {
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
		delegateInfo:  infoLogger,
		delegateWarn:  warnLogger,
		delegateError: errorLogger,
		delegatePanic: panicLogger,
		delegateFatal: fatalLogger,
	}
}

func Default() log4lib.LibLogger {
	return NewBuiltInLoggerWrapper(
		log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "WARN:", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "PANIC:", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "FATAL:", log.Ldate|log.Ltime|log.Lshortfile),
	)
}
