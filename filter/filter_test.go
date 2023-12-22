package filter

import (
	"testing"

	"github.com/ClockwerksSoftware/golog/interfaces"
	"github.com/ClockwerksSoftware/golog/record"
)

func TestFilter(t *testing.T) {
	t.Run(
		"new",
		func(t *testing.T) {
			f := NewLogFilter().(*logFilter)
			if f.name != nameLevelFilter {
				t.Errorf("unexpected filter name: %s != %s", f.name, nameLevelFilter)
			}
			expectedLogLevel := record.GetLogLevel(record.DEBUG)
			if f.level.Int() != expectedLogLevel.Int() {
				t.Errorf("Unexpected default log level: %v != %v", f.level, expectedLogLevel)
			}
		},
	)
	t.Run(
		"name",
		func(t *testing.T) {
			f := NewLogFilter()
			name := f.Name()
			if name != nameLevelFilter {
				t.Errorf("Unexpected filter name: %s != %s", name, nameLevelFilter)
			}
		},
	)
	t.Run(
		"setLevel",
		func(t *testing.T) {
			f := NewLogFilter().(*logFilter)
			expectedLogLevel := record.GetLogLevel(record.DEBUG)
			if f.level.Int() != expectedLogLevel.Int() {
				t.Errorf("Unexpected default log level: %v != %v", f.level, expectedLogLevel)
			}
			newLogLevel := record.GetLogLevel(record.WARNING)
			f.SetLevel(newLogLevel)
			if f.level.Int() != newLogLevel.Int() {
				t.Errorf("Filter did not update the log level: %v != %v", f.level, newLogLevel)
			}
		},
	)
	t.Run(
		"filter",
		func(t *testing.T) {
			type recordEntry struct {
				level   interfaces.Level
				rec     interfaces.Record
				allowed bool
			}
			processRecords := []recordEntry{
				{
					level: record.GetLogLevel(record.WARNING),
					rec: record.NewLogRecord(
						"foo",
						record.WARNING,
						"foobar",
					),
					allowed: true,
				},
				{
					level: record.GetLogLevel(record.DEBUG),
					rec: record.NewLogRecord(
						"foo",
						record.CRITICAL,
						"foobar",
					),
					allowed: false,
				},
				{
					level: record.GetLogLevel(record.DEBUG),
					rec: record.NewLogRecord(
						"foo",
						record.WARNING,
						"foobar",
					),
					allowed: true,
				},
			}

			f := NewLogFilter()
			for _, entry := range processRecords {
				f.(*logFilter).SetLevel(entry.level)
				r, recordAllowed := f.Filter(entry.rec)
				if recordAllowed != entry.allowed {
					t.Errorf("Unexpected filtering: Filter Result - %v, Expected Result: %v", recordAllowed, entry.allowed)
				}
				if entry.allowed && r == nil {
					t.Errorf("Unexpectedly did not receive the record back")
				}
				if !entry.allowed && r != nil {
					t.Errorf("Unexpectedly received the record back")
				}
			}
		},
	)
}
