package unilog4logrus

import (
	"context"

	"github.com/blugnu/unilog"
	"github.com/sirupsen/logrus"
)

// Logger returns a `Logger` that wraps a specified `logrus.Logger` and
// a supplied `Context`.
func Logger(ctx context.Context, log *logrus.Logger) unilog.Logger {
	return unilog.UsingAdapter(ctx, &adapter{log})
}

var logrusLevel = map[unilog.Level]logrus.Level{
	unilog.Trace: logrus.TraceLevel,
	unilog.Debug: logrus.DebugLevel,
	unilog.Info:  logrus.InfoLevel,
	unilog.Warn:  logrus.WarnLevel,
	unilog.Error: logrus.ErrorLevel,
	unilog.Fatal: logrus.FatalLevel,
}

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

func (log *adapter) WithField(name string, value any) unilog.Adapter {
	return &entryAdapter{log.Logger.WithField(name, value)}
}

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
