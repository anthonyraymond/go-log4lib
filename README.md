# Why log4lib
Log4lib allow library to add log to the application logger output.

Most of the time library should not log, and instead return errors. But there is some time where sending a WARN and keep going is better. Log4lib is meant to do this.
It provides a set of wrapper for the most commonly used log library.

# How to use
### As a library developer
Library developer should only install `github.com/anthonyraymond/go-log4lib`
```shell script
go get github.com/anthonyraymond/go-log4lib
```

>project/log.go
```go
package thelibrary

import (
    "github.com/anthonyraymond/go-log4lib"
    "github.com/anthonyraymond/log4libwrapper"
)

// Default logger logs to console by default, but user can replace the logger using the SetLibraryLogger()
var logger = log4libwrapper.Default()

func SetLibraryLogger(logger log4lib.LibLogger) {
	logger = logger
}

func GetLibLogger() log4lib.LibLogger {
	return logger
}
```

>project/mylibraryfile.go
```go
package thelibrary

func doSomething() {
    GetLibLogger().Info("hello :) that's a log")
}
```

### As a library user
Library users should only install wrapper and **must not install** `github.com/anthonyraymond/go-log4lib`

```shell script
go get github.com/anthonyraymond/log4libwrapper-zap
or
go get github.com/anthonyraymond/log4libwrapper-logrus
or
go get github.com/anthonyraymond/log4libwrapper-golog
```


Zap:
```go
package myapp

import (
    "github.com/anthonyraymond/log4libwrapper-zap"
    "go.uber.org/zap"
    "thelibrary"
)

func init() {
    var logger *zap.Logger = zap.NewProduction()
    thelibrary.SetLibraryLogger(log4libwrapper.WrapZapLogger(logger.Sugar()))
}

```

logrus:
TODO

golog:
TODO