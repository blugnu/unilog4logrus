<div align="center" style="margin-bottom:20px">
  <!-- <img src=".assets/banner.png" alt="logger" /> -->
  <div align="center">
    <a href="https://github.com/blugnu/unilog4logrus/actions/workflows/qa.yml"><img alt="build-status" src="https://github.com/blugnu/unilog4logrus/actions/workflows/qa.yml/badge.svg?branch=master&style=flat-square"/></a>
    <a href="https://goreportcard.com/report/github.com/blugnu/unilog4logrus" ><img alt="go report" src="https://goreportcard.com/badge/github.com/blugnu/unilog4logrus"/></a>
    <a><img alt="go version >= 1.14" src="https://img.shields.io/github/go-mod/go-version/blugnu/unilog4logrus?style=flat-square"/></a>
    <a href="https://github.com/blugnu/unilog4logrus/blob/master/LICENSE"><img alt="MIT License" src="https://img.shields.io/github/license/blugnu/unilog4logrus?color=%234275f5&style=flat-square"/></a>
    <a href="https://coveralls.io/github/blugnu/unilog4logrus?branch=master"><img alt="coverage" src="https://img.shields.io/coveralls/github/blugnu/unilog4logrus?style=flat-square"/></a>
    <a href="https://pkg.go.dev/github.com/blugnu/unilog4logrus"><img alt="docs" src="https://pkg.go.dev/badge/github.com/blugnu/unilog4logrus"/></a>
  </div>
</div>

<br>

# unilog4logrus

Implements a [unilog](https://github.com/blugnu/unilog) `Adapter` to emit logs using [logrus](https://github.com/sirupsen/logrus).

## How To use This Adapter

1. Configure your `logrus` logger in whatever way suits your project
2. Initialise a `unilog.Logger` by calling `unilog4logrus.Logger()`, supplying an initial context and the configured [logrus](https://github.com/sirupsen/logrus) logger
3. Pass the `unilog.Logger` into any modules that support a [unilog](https://github.com/blugnu/unilog) Logger`
4. Emit logs from your project using the `unilog.Logger`
5. Enjoy reading your logs!


#### Example: Using unilog with logrus

```golang
var logger unilog.Logger

func main() {
  // Configure a logrus logger
	lr := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

  // Get a unilog encapsulating the logrus logger
	logger = unilog4logrus.Logger(context.Background(), lr)

  // Pass logger into the `foo` module, which supports unilog
  foo.Logger = logger

  // Do some logging ourselves
  log := logger.NewEntry()
  log.Info("logging initialised")

  // Any logs written by SetupTheFoo() will use the same logger as 'log'
  if err := foo.SetupTheFoo(); err != nil {
    log.FatalError(err)
  }

  // ... etc
}
```