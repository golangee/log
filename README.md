# log [![GoDoc](https://godoc.org/github.com/golangee/sql?status.svg)](http://godoc.org/github.com/golangee/sql)
This is **NOT** another logger but a simple, clean and potentially
dependency free logging facade. You can use it either by copy'n
paste the interface definitions or use it just as a usual dependency.

It follows more or less the recommendation 
from [zap](https://github.com/uber-go/zap/blob/master/FAQ.md#why-arent-logger-and-sugaredlogger-interfaces).

However, even if this works out-of-the-box,
you should use an implementation for it and not
the build-in trivial logger implementation. Available implementations:
* [log-zap](https://github.com/golangee/log-zap)

## Alternatives (or implementations)
* [logrus](https://github.com/sirupsen/logrus) (★ 15k)
* [zap](https://github.com/uber-go/zap) (★ 10k)
* [zerolog](https://github.com/rs/zerolog) (★ 3k)
* [apex](https://github.com/apex/log) (★ 1k)


## interface design
It just provides structured logging, anything else can be
rendered from it. The interface is designed for easy integration
and type safety but not for achieving the best performance.
Abstractions often come with a cost. 

```go
// Field is an alias to a key/value tuple, to break dependency.
type Field = struct {
	Key string
	Val interface{}
}

// Logger provides the different log levels and field definitions.
// It is kept as simple as possible and to avoid recursive dependency
// cycles by returns interfaces or concrete types.
type Logger interface {
	Trace(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}
```

## usage

```go
package main

import (
  "github.com/golangee/log"
  "github.com/golangee/log-zap"
)

func main(){
 zap.Configure()
 log.New().Debug("hello world", log.Obj("id", 5))
}
```

