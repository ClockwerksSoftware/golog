package record

import (
	"math/rand"
	"testing"
)

func TestLogRecordLocation(t *testing.T) {
	t.Run(
		"NewLocation",
		func(t *testing.T) {
			lrl := NewLogRecordLocation()
			if lrl == nil {
				t.Errorf("Failed to allocate a new log record location object")
			}
			plrl := lrl.(*LogRecordLocation)
			if plrl.filename != "" {
				t.Errorf("Unexpected filename set: %q", plrl.filename)
			}
			if plrl.line != 0 {
				t.Errorf("Unexpected line number set: %d != %d", plrl.line, 0)
			}
			if plrl.ok {
				t.Errorf("Unexpected valid location")
			}
			if plrl.stack != "" {
				t.Errorf("Unexpected stack value set: %q", plrl.stack)
			}
		},
	)
	t.Run(
		"Filename",
		func(t *testing.T) {
			lrl := NewLogRecordLocation()
			if lrl == nil {
				t.Errorf("Failed to allocate a new log record location object")
			}
			plrl := lrl.(*LogRecordLocation)
			if plrl.filename != "" {
				t.Errorf("Unexpected filename set: %q", plrl.filename)
			}
			newFilename := "some-filename"
			plrl.SetFilename(newFilename)
			if plrl.filename != newFilename {
				t.Errorf("Unexpected filename set: %q != %q", plrl.filename, newFilename)
			}
			extractedFilename := lrl.Filename()
			if extractedFilename != newFilename {
				t.Errorf("Unexpected filename set: %q != %q", extractedFilename, newFilename)
			}
		},
	)
	t.Run(
		"Line",
		func(t *testing.T) {
			lrl := NewLogRecordLocation()
			if lrl == nil {
				t.Errorf("Failed to allocate a new log record location object")
			}
			plrl := lrl.(*LogRecordLocation)
			if plrl.line != 0 {
				t.Errorf("Unexpected line number set: %d != %d", plrl.line, 0)
			}
			newLine := rand.Int()
			plrl.SetLine(newLine)
			if plrl.line != newLine {
				t.Errorf("Unexpected line set: %d != %d", plrl.line, newLine)
			}
			extractedLine := lrl.Line()
			if extractedLine != newLine {
				t.Errorf("Unexpected line set: %d != %d", extractedLine, newLine)
			}
		},
	)
	t.Run(
		"Valid",
		func(t *testing.T) {
			lrl := NewLogRecordLocation()
			if lrl == nil {
				t.Errorf("Failed to allocate a new log record location object")
			}
			plrl := lrl.(*LogRecordLocation)
			if plrl.ok {
				t.Errorf("Unexpected valid location")
			}
			plrl.SetValid(true)
			if !plrl.ok {
				t.Errorf("Unexpected invalid location (direct)")
			}
			extractedValid := lrl.Valid()
			if !extractedValid {
				t.Errorf("Unexpected invalid location (interface)")
			}
		},
	)
	t.Run(
		"Stack",
		func(t *testing.T) {
			lrl := NewLogRecordLocation()
			if lrl == nil {
				t.Errorf("Failed to allocate a new log record location object")
			}
			plrl := lrl.(*LogRecordLocation)
			if plrl.stack != "" {
				t.Errorf("Unexpected stack set: %q", plrl.stack)
			}
			newStack := "some-stack"
			plrl.SetStack(newStack)
			if plrl.stack != newStack {
				t.Errorf("Unexpected stack set: %q != %q", plrl.stack, newStack)
			}
			extractedStack := lrl.Stack()
			if extractedStack != newStack {
				t.Errorf("Unexpected stack set: %q != %q", extractedStack, newStack)
			}
		},
	)
}
