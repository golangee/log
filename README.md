# log [![GoDoc](https://godoc.org/github.com/golangee/log?status.svg)](http://godoc.org/github.com/golangee/log) 
There is more to it, than picking the next logger and call it a day, to get logging right.
First, you have to be clear about your intention. What do you really need, logging, monitoring or
tracing?

## about logging
Logging is about tracking events and their related data. Think about administrators
as your customer. In the old days, log files where managed entirely by
the application itself, e.g. by appending to existing files and rotating when required. Today,
most operating systems provide advanced logging facilities attached to stdout. On Linux, this is
likely the systemd journal. Thus, there is usually no need anymore to handle this in the applications'
layer.

However, the most important finding about logging is, that this is NOT the correct place to (solely) 
inform about alerts, warnings or panic situations: consider, that nobody will read your logs. A common way
to collect, inspect and filter log data is using [elastic search and kibana](https://www.elastic.co/de/).

## about tracing
Tracing is about following the programs' flow, typically from a user perspective. This
may be accomplished by heavy and very verbose logging instruments, but indeed tools like 
[Jaeger](https://github.com/jaegertracing/) are more suited for this.

## about monitoring
Monitoring is about instrumenting, collecting, aggregating and analyzing metrics to understand
the behavior of your application and system, especially under load. Well known tools are 
[Prometheus](https://prometheus.io/) and [Grafana](https://grafana.com/).

## how to log right?
If you are still here and are sure, that you need to log information, keep the following
rules of thumb in mind:
* avoid any vendor lock-in into existing logging frameworks by linking directly to them.
* realize, that log levels do not make sense. Real *errors* will cause your application to halt
just like a *panic* or an *os.Exit*. Anything else is an information, and
the evaluation is in the eye of the beholder. It is worth to read the article from
[Dave Cheney](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).
* logging hurts performance, and in high performance code, the argument propagation will still
cause allocations and even worse - *escaping* to the heap. To enable verbose logs, 
guard all according calls with compile time flags, so that the actual calls can get eliminated entirely.
* use a standard scheme for your fields, like the 
[ECS field reference](https://www.elastic.co/guide/en/ecs/current/ecs-field-reference.html).


This is our recommended logging interface:
```go
type Logger interface {
    // Println processes and prints the arguments as fields. The interpretation and formatting depends on the
    // concrete implementation and may range from fmt.Println over log.Println to a full structured logger.
    // Implementations are encouraged to type-switch on each field.
    Println(fields ...interface{})
}
```

This is our recommended guard:
```go
// +build !debug

package log

// Debug is a build tag determined at build time, so that the compiler can remove dead code.
const Debug = false
```

## Available implementations
* [logrus](https://github.com/sirupsen/logrus) (★ 15k)
* [zap](https://github.com/uber-go/zap) (★ 10k)
* [zerolog](https://github.com/rs/zerolog) (★ 3k)
* [apex](https://github.com/apex/log) (★ 1k)


## usage
If you want **zero** dependencies, just copy the *recommended logging interface* and provide injection
capabilities through factory methods. Alternatively, this package provides a standard factory and 
three simple but ready-to use logger implementations. To get the best of both worlds, we recommend to 
just start with the dependency, and optimize later by setting a factory to any of the 
implementations above, when required. Note, that the default logger is *simple.PrintColored*, if started
from within your IDE and otherwise *simple.PrintStructured*. However, you can change it using 
*SetDefault* to whatever you like. 

There are also the following default special treatments:
* 

```go
package main

import (
  "github.com/golangee/log"
  "github.com/golangee/log/ecs"
  "context"
)

func main(){
    // prints from IDE: 2020-12-14T10:46:37+01:00 my.logger hello world
    // prints in prod: {"@timestamp":"2020-12-14T10:46:37+01:00","log.logger":"my.logger","message":"hello world"}
    log.Println("hello world")

    // prints from IDE: 2020-12-14T10:46:37+01:00 my.logger INFO auto message https://automatic.url automatic error *errors.errorString
    // prints in prod: {"@timestamp":"2020-12-14T10:46:37+01:00","error.message":"automatic error","error.type":"*errors.errorString","log.level":"info","log.logger":"my.logger","message":"auto message","url.path":"https://automatic.url"}
    log.Println("info", "auto message", "https://automatic.url", fmt.Errorf("automatic error"))
	
    // how to create custom logger setup
    myLogger := log.NewLogger(ecs.Log("my.logger"))

    // prints from IDE: 2020-11-20T15:26:07+01:00 my.logger hello
    // prints in prod: {"@timestamp":"2020-11-20T15:26:07+01:00","log.level":"trace","log.logger":"my.logger","message":"hello"}
    myLogger.Println(ecs.Msg("hello")) 
    myLogger.Println(ecs.Msg("world"))

    // guard verbose and/or expensive logs
    if log.Debug{
    	myLogger.Println(ecs.Msg("info point 1"), ecs.ErrStack()) 
    }

    // in your http middleware you should use the context
    reqCtx := log.WithLogger(context.Background(), log.WithFields(myLogger, ecs.Log("my.request.logger")))
    log.FromContext(reqCtx).Println(ecs.Msg("from a request"))
}
```

