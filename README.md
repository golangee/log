# log [![GoDoc](https://godoc.org/github.com/golangee/log?status.svg)](http://godoc.org/github.com/golangee/log)
This is **NOT** another logger but a simple, clean and potentially
dependency free logging facade. You can use it either by copy'n
paste the interface definitions or use it just as a usual dependency.

It follows more or less the thoughts 
from [Dave Cheney](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) and 
[zap](https://github.com/uber-go/zap/blob/master/FAQ.md#why-arent-logger-and-sugaredlogger-interfaces).

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
Abstractions (especially interfaces in Go) come with a cost.
However, you may use the *log.Debug* compile time flag, to guard
verbose developer messages. Also, there are a few helper methods
defined to comply to the [ECS](https://www.elastic.co/guide/en/ecs/current/ecs-field-reference.html).

```go
// Field is an alias to a key/value tuple, to break dependency.
type Field = struct {
	Key string
	Val interface{}
}

// Logger provides the abstraction for logging.
// It is kept as simple as possible and to avoid recursive dependency
// cycles by using other interfaces or concrete types. It deliberately breaks with the conventional logger APIs, due
// to the following considerations:
//  * there are verbose developer specific logs which are not important or even bad for life system. You mostly
//    even want that there is no cost for the log parameter propagation, which would cause even more harm like
//    escaping values and heap-pressure, even if disabled. The only way to avoid this, is a guarded
//    compile time constant evaluation.
//  * anything else which is so important, that a developer is not sure to turn off in production, must not be guarded.
//    Instead it is up to the administrator or software engineer to filter through the log in a structured way. Any
//    error which does not kill your application, is just another kind of information.
//  * any error which does break your post-variants and you cannot continue, should be logged (again, still info)
//    either bail out with a runtime panic or an explicit os.Exit.
type Logger interface {
	Info(fields ...Field)
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

