package unilog4logrus

import (
	"context"

	"github.com/blugnu/unilog"
	"github.com/sirupsen/logrus"
)

// Configuration interface provides convenience functions for
// configuring some aspects of a logrus logger or adapter.
//
// Most configuration should normally be performed using the
// normal mechanisms provided by logrus itself.
type Configuration interface {
	SetLevel(unilog.Level) // SetLevel enables the logging level to be configured in terms of `unilog.Level`.
}

// Logger returns a `Logger` that wraps a specified `logrus.Logger` and
// a supplied `Context`.
func Logger(ctx context.Context, log *logrus.Logger) (unilog.Logger, Configuration) {
	a := &adapter{log}
	return unilog.UsingAdapter(ctx, a), a
}

// logrusLevel is a simple map of unilog.Level to corresponding
// logrus.Level constants
var logrusLevel = map[unilog.Level]logrus.Level{
	unilog.Trace: logrus.TraceLevel,
	unilog.Debug: logrus.DebugLevel,
	unilog.Info:  logrus.InfoLevel,
	unilog.Warn:  logrus.WarnLevel,
	unilog.Error: logrus.ErrorLevel,
	unilog.Fatal: logrus.FatalLevel,
}

// adapter implements the unilog.Adapter interface encapsulating
// a `logrus.Logger`.
//
// The majority of logging will use this adapter to initialise
// new entries which for `logrus` are implemented by a separate
// `entryAdapter`.
type adapter struct {
	*logrus.Logger
}

func (log *adapter) Emit(level unilog.Level, s string) {
	log.Logger.Log(logrusLevel[level], s)
}

func (log *adapter) NewEntry() unilog.Adapter {
	entry := logrus.NewEntry(log.Logger)
	return &entryAdapter{entry}
}

func (log *adapter) SetLevel(level unilog.Level) {
	log.Logger.SetLevel(logrusLevel[level])
}

func (log *adapter) WithField(name string, value any) unilog.Adapter {
	return &entryAdapter{log.Logger.WithField(name, value)}
}

// entryAdapter implements the unilog.Adapter interface encapsulating
// an `logrus.Entry`.
type entryAdapter struct {
	*logrus.Entry
}

func (log *entryAdapter) Emit(level unilog.Level, s string) {
	log.Entry.Log(logrusLevel[level], s)
}

func (log *entryAdapter) NewEntry() unilog.Adapter {
	entry := logrus.NewEntry(log.Logger)
	return &entryAdapter{entry}
}

func (log *entryAdapter) WithField(name string, value any) unilog.Adapter {
	return &entryAdapter{log.Entry.WithField(name, value)}
}
