package unilog4logrus

import (
	"context"
	"reflect"
	"testing"

	logrus "github.com/sirupsen/logrus"

	log "github.com/blugnu/go-logspy"
	"github.com/blugnu/unilog"
)

func TestLogrusAdapter(t *testing.T) {
	// ARRANGE
	logger := logrus.New()
	logger.SetOutput(log.Sink())
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	sut := &adapter{logger}

	testcases := []struct {
		name   string
		fn     func(string)
		output string
	}{
		{name: "trace", fn: func(s string) { sut.Emit(unilog.Trace, s) }, output: "level=trace msg=\"entry text\"\n"},
		{name: "debug", fn: func(s string) { sut.Emit(unilog.Debug, s) }, output: "level=debug msg=\"entry text\"\n"},
		{name: "info", fn: func(s string) { sut.Info(s) }, output: "level=info msg=\"entry text\"\n"},
		{name: "warn", fn: func(s string) { sut.Warn(s) }, output: "level=warning msg=\"entry text\"\n"},
		{name: "error", fn: func(s string) { sut.Emit(unilog.Error, s) }, output: "level=error msg=\"entry text\"\n"},
		{name: "fatal", fn: func(s string) { sut.Emit(unilog.Fatal, s) }, output: "level=fatal msg=\"entry text\"\n"},
		{name: "debug and error", fn: func(s string) { sut.Emit(unilog.Debug, s); sut.Emit(unilog.Error, s) }, output: "level=debug msg=\"entry text\"\nlevel=error msg=\"entry text\"\n"},
		{name: "withfield", fn: func(s string) {
			a := sut
			b := sut.WithField("field", "data")

			t.Run("returns new logger", func(t *testing.T) {
				wanted := true
				got := a != b
				if wanted != got {
					t.Errorf("wanted %v, got %v", wanted, got)
				}
			})

			a.Emit(unilog.Info, s)
			b.Emit(unilog.Info, s)
		}, output: "level=info msg=\"entry text\"\nlevel=info msg=\"entry text\" field=data\n"},
		{name: "newentry", fn: func(s string) {
			a := sut
			b := sut.NewEntry()

			t.Run("returns new logger", func(t *testing.T) {
				wanted := true
				got := a != b
				if wanted != got {
					t.Errorf("wanted %v, got %v", wanted, got)
				}
			})

			a.Emit(unilog.Info, s)
			b.Emit(unilog.Info, s)
		}, output: "level=info msg=\"entry text\"\nlevel=info msg=\"entry text\"\n"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			defer log.Reset()
			// ACT
			tc.fn("entry text")

			// ASSERT
			wanted := tc.output
			got := log.String()
			if wanted != got {
				t.Errorf("\nwanted %q\ngot    %q", wanted, got)
			}
		})
	}
}

func TestLogrusEntryAdapter(t *testing.T) {
	// ARRANGE
	logger := logrus.New()
	logger.SetOutput(log.Sink())
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	sut := &entryAdapter{logrus.NewEntry(logger)}

	testcases := []struct {
		name   string
		fn     func(string)
		output string
	}{
		{name: "debug", fn: func(s string) { sut.Emit(unilog.Debug, s) }, output: "level=debug msg=\"entry text\"\n"},
		{name: "info", fn: func(s string) { sut.Info(s) }, output: "level=info msg=\"entry text\"\n"},
		{name: "warn", fn: func(s string) { sut.Warn(s) }, output: "level=warning msg=\"entry text\"\n"},
		{name: "error", fn: func(s string) { sut.Emit(unilog.Error, s) }, output: "level=error msg=\"entry text\"\n"},
		{name: "debug and error", fn: func(s string) { sut.Emit(unilog.Debug, s); sut.Emit(unilog.Error, s) }, output: "level=debug msg=\"entry text\"\nlevel=error msg=\"entry text\"\n"},
		{name: "withfield", fn: func(s string) {
			a := sut
			b := sut.WithField("field", "data")

			t.Run("returns new logger", func(t *testing.T) {
				wanted := true
				got := a != b
				if wanted != got {
					t.Errorf("wanted %v, got %v", wanted, got)
				}
			})

			a.Emit(unilog.Info, s)
			b.Emit(unilog.Info, s)
		}, output: "level=info msg=\"entry text\"\nlevel=info msg=\"entry text\" field=data\n"},
		{name: "newentry", fn: func(s string) {
			a := sut
			b := sut.NewEntry()

			t.Run("returns new logger", func(t *testing.T) {
				wanted := true
				got := a != b
				if wanted != got {
					t.Errorf("wanted %v, got %v", wanted, got)
				}
			})

			a.Emit(unilog.Info, s)
			b.Emit(unilog.Info, s)
		}, output: "level=info msg=\"entry text\"\nlevel=info msg=\"entry text\"\n"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			defer log.Reset()

			// ACT
			tc.fn("entry text")

			// ASSERT
			wanted := tc.output
			got := log.String()
			if wanted != got {
				t.Errorf("\nwanted %q\ngot    %q", wanted, got)
			}
		})
	}
}

func TestLogger(t *testing.T) {
	// ARRANGE
	ctx := context.Background()
	log := &logrus.Logger{}

	// ACT
	result := Logger(ctx, log)

	// ASSERT
	wanted := unilog.UsingAdapter(ctx, &adapter{log})
	got := result
	if !reflect.DeepEqual(wanted, got) {
		t.Errorf("\nwanted %#v\ngot    %#v", wanted, got)
	}
}
