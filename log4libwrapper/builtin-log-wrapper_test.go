package log4libwrapper

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestBuiltinLoggerWrapper_MustLog(t *testing.T) {
	logReceiver := &strings.Builder{}
	builtinLoggerDebug := log.New(io.MultiWriter(logReceiver, os.Stdout), "debug\t", 0)
	builtinLoggerInfo := log.New(io.MultiWriter(logReceiver, os.Stdout), "info\t", 0)
	builtinLoggerWarn := log.New(io.MultiWriter(logReceiver, os.Stdout), "warn\t", 0)
	builtinLoggerError := log.New(io.MultiWriter(logReceiver, os.Stdout), "error\t", 0)
	builtinLoggerPanic := log.New(io.MultiWriter(logReceiver, os.Stdout), "panic\t", 0)
	builtinLoggerFatal := log.New(io.MultiWriter(logReceiver, os.Stdout), "fatal\t", 0)

	logger := WrapBuiltinLogger(builtinLoggerDebug, builtinLoggerInfo, builtinLoggerWarn, builtinLoggerError, builtinLoggerPanic, builtinLoggerFatal)
	logger.Debug("coucou")
	logger.Info("coucou")
	logger.Warn("coucou")
	logger.Error("coucou")
	logger.Debug("coucou", "joe", "la bidouille")
	logger.Info("coucou", "joe", "la bidouille")
	logger.Warn("coucou", "joe", "la bidouille")
	logger.Error("coucou", "joe", "la bidouille")
	logger.Debugf("coucou %s :)", "joe")
	logger.Infof("coucou %s :)", "joe")
	logger.Warnf("coucou %s :)", "joe")
	logger.Errorf("coucou %s :)", "joe")
	logger.Debugf("coucou %s %d :)", "joe", 12)
	logger.Infof("coucou %s %d :)", "joe", 12)
	logger.Warnf("coucou %s %d :)", "joe", 12)
	logger.Errorf("coucou %s %d :)", "joe", 12)

	expected := []string{
		"debug\tcoucou",
		"info\tcoucou",
		"warn\tcoucou",
		"error\tcoucou",
		"debug\tcoucou joe la bidouille",
		"info\tcoucou joe la bidouille",
		"warn\tcoucou joe la bidouille",
		"error\tcoucou joe la bidouille",
		"debug\tcoucou joe :)",
		"info\tcoucou joe :)",
		"warn\tcoucou joe :)",
		"error\tcoucou joe :)",
		"debug\tcoucou joe 12 :)",
		"info\tcoucou joe 12 :)",
		"warn\tcoucou joe 12 :)",
		"error\tcoucou joe 12 :)",
		"",
	}

	if !strings.EqualFold(strings.Join(expected, "\n"), logReceiver.String()) {
		t.Fatalf("log output is not correct, expected:\n%s\nactual:\n%s", strings.Join(expected, "\n"), logReceiver.String())
	}
}


func TestWrapZapLogger_CallerMustBeUnwrapped(t *testing.T) {

	// when the caller is added to the zap output we don't want to see the Info, Warn, ... wrapping method from zapLoggerWrapper, we want the real caller
	logReceiver := &strings.Builder{}
	builtinLogger := log.New(io.MultiWriter(logReceiver, os.Stdout), "", log.Lshortfile)

	WrapBuiltinLogger(builtinLogger, builtinLogger, builtinLogger, builtinLogger, builtinLogger, builtinLogger).Info("coucou")


	if !strings.Contains(logReceiver.String(), "builtin-log-wrapper_test.go") {
		t.Fatal("caller is now correct, the wrapper should be ignored and the caller must be the real calling function (here go-log4libwrapper-zap/log4libwrapper_test.go)")
	}
}