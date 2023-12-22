package record

import (
	"runtime"
	"strings"
	"testing"
)

func TestRecord(t *testing.T) {
	t.Run(
		"New",
		func(t *testing.T) {
			expectedName := t.Name()
			var expectedLevel LogLevelInt = INFO
			expectedMessage := "test log message"
			expectedArgs := make([]any, 0)
			var stackData [2 << 10]byte
			// reduce the call depth by 1 to get the matching call stack data
			_, expectedFilename, expectedLine, expectedValid := runtime.Caller(logRecordCallerDepth - 1)
			stackLength := runtime.Stack(stackData[:], false)
			r := NewLogRecord(
				expectedName,
				expectedLevel,
				expectedMessage,
				expectedArgs...
			)
			expectedStack := string(stackData[:stackLength])
			// the expected stack will differ based on the location data of the
			// line number of the `runtime.Stack` call above and its equivalent
			// call within `NewLogRecord`.
			// two sections of the stack are extracted:
			// 1. the first part up to this file's filename, after the goroutine data
			// 2. the second part after the addressing data following the first part
			// these can then be searched for in the stack data that is
			// extracted by `NewLogRecord` to verify it's geting the correct stack data
			offsetIndex := strings.Index(expectedStack, "record_test.go")
			comparableStackFirst :=  expectedStack[23:offsetIndex+len("record_test.go")]
			offsetIndex += strings.Index(expectedStack[offsetIndex:], "\n")
			comparableStack := expectedStack[offsetIndex:]

			if r == nil {
				t.Errorf("Unable to allocate new log record")
			}
			pr := r.(*logRecord)
			if pr.name != expectedName {
				t.Errorf("Unexpected name found: %q != %q", pr.name, expectedName)
			}
			if pr.level.Int() != int(expectedLevel) {
				t.Errorf("Unexpected level found: %d != %d", pr.level.Int(), expectedLevel)
			}
			if pr.message != expectedMessage {
				t.Errorf("Unexpected message found: %q != %q", pr.message, expectedMessage)
			}
			if len(pr.attrs) != len(expectedArgs) {
				t.Errorf("Unexpected arg length: %d != %d", len(pr.attrs), len(expectedArgs))
			}
			for i, v := range expectedArgs {
				if pr.attrs[i] != v {
					t.Errorf("Unexpected arg found at index %d: %#v != %#v", i, pr.attrs[i], v)
				}
			}
			extractedFilename := pr.location.Filename()
			extractedLine := pr.location.Line()
			extractedValid := pr.location.Valid()
			if extractedFilename != expectedFilename {
				t.Errorf("Unexpected Location Filename: %q != %q", extractedFilename, expectedFilename)
			}
			if extractedLine != expectedLine {
				t.Errorf("Unexpected Location Line: %d != %d", extractedLine, expectedLine)
			}
			if extractedValid != expectedValid {
				t.Errorf("Unexpected Location Validity: %t != %t", extractedValid, expectedValid)
			}
			extractedStack := pr.location.Stack()
			if !strings.Contains(extractedStack, comparableStackFirst) {
				t.Errorf("Did not find overlapping stack data: %q does not contain %q", extractedStack, comparableStackFirst)
			}
			if !strings.Contains(extractedStack, comparableStack) {
				t.Errorf("Did not find overlapping stack data: %q does not contain %q", extractedStack, comparableStack)
			}
		},
	)
}
