package record

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"testing"

	"github.com/ClockwerksSoftware/golog/interfaces"
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
			if len(pr.attrs) != 0 {
				t.Errorf("Unexpected attrs found - Count: %d, Values: %#v", len(pr.attrs), pr.attrs)
			}
			if len(pr.messageArgs) != len(expectedArgs) {
				t.Errorf("Unexpected arg length: %d != %d", len(pr.attrs), len(expectedArgs))
			}
			for i, v := range expectedArgs {
				if pr.messageArgs[i] != v {
					t.Errorf("Unexpected arg found at index %d: %#v != %#v", i, pr.messageArgs[i], v)
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
	t.Run(
		"Location",
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

			ploc := r.Location()

			extractedFilename := pr.location.Filename()
			extractedLine := pr.location.Line()
			extractedValid := pr.location.Valid()
			if extractedFilename != expectedFilename {
				t.Errorf("Unexpected Location Filename: %q != %q", extractedFilename, expectedFilename)
			}
			if ploc.Filename() != expectedFilename {
				t.Errorf("Unexpected Location Filename: %q != %q", ploc.Filename(), expectedFilename)
			}

			if extractedLine != expectedLine {
				t.Errorf("Unexpected Location Line: %d != %d", extractedLine, expectedLine)
			}
			if ploc.Line() != expectedLine {
				t.Errorf("Unexpected Location Line: %d != %d", ploc.Line(), expectedLine)
			}

			if extractedValid != expectedValid {
				t.Errorf("Unexpected Location Validity: %t != %t", extractedValid, expectedValid)
			}
			if ploc.Valid() != expectedValid {
				t.Errorf("Unexpected Location Validity: %t != %t", ploc.Valid(), expectedValid)
			}

			extractedStack := pr.location.Stack()
			plocStack := ploc.Stack()
			if !strings.Contains(extractedStack, comparableStackFirst) {
				t.Errorf("Did not find overlapping stack data: %q does not contain %q", extractedStack, comparableStackFirst)
			}
			if !strings.Contains(extractedStack, comparableStack) {
				t.Errorf("Did not find overlapping stack data: %q does not contain %q", extractedStack, comparableStack)
			}
			if plocStack != extractedStack {
				t.Errorf("Unexpected stack found: %q != %q", plocStack, extractedStack)
			}
		},
	)
	t.Run(
		"Name",
		func(t *testing.T) {
			expectedName := t.Name()
			var expectedLevel LogLevelInt = INFO
			expectedMessage := "test log message"
			expectedArgs := make([]any, 0)
			r := NewLogRecord(
				expectedName,
				expectedLevel,
				expectedMessage,
				expectedArgs...
			)
			if r == nil {
				t.Errorf("Unable to allocate new log record")
			}
			pr := r.(*logRecord)
			if pr.name != expectedName {
				t.Errorf("Unexpected name found: %q != %q", pr.name, expectedName)
			}
			if r.Name() != expectedName {
				t.Errorf("Unexpected name found: %q != %q", r.Name(), expectedName)
			}
		},
	)
	t.Run(
		"Level",
		func(t *testing.T) {
			expectedName := t.Name()
			var expectedLevel LogLevelInt = INFO
			expectedMessage := "test log message"
			expectedArgs := make([]any, 0)
			r := NewLogRecord(
				expectedName,
				expectedLevel,
				expectedMessage,
				expectedArgs...
			)
			if r == nil {
				t.Errorf("Unable to allocate new log record")
			}
			pr := r.(*logRecord)
			if pr.level.Int() != int(expectedLevel) {
				t.Errorf("Unexpected level found: %d != %d", pr.level.Int(), expectedLevel)
			}
			if r.Level().Int() != int(expectedLevel) {
				t.Errorf("Unexpected level found: %d != %d", r.Level().Int(), expectedLevel)
			}
		},
	)
	t.Run(
		"Time",
		func(t *testing.T) {
			expectedName := t.Name()
			var expectedLevel LogLevelInt = INFO
			expectedMessage := "test log message"
			expectedArgs := make([]any, 0)
			r := NewLogRecord(
				expectedName,
				expectedLevel,
				expectedMessage,
				expectedArgs...
			)
			if r == nil {
				t.Errorf("Unable to allocate new log record")
			}
			pr := r.(*logRecord)
			extractedTime := r.Time()
			if extractedTime != pr.datetime {
				t.Errorf("Unexpected time found: %#v != %#v", extractedTime, pr.datetime)
			}
		},
	)
	t.Run(
		"RawMessage",
		func(t *testing.T) {
			expectedName := t.Name()
			var expectedLevel LogLevelInt = INFO
			expectedMessage := "test log message"
			expectedArgs := make([]any, 0)
			r := NewLogRecord(
				expectedName,
				expectedLevel,
				expectedMessage,
				expectedArgs...
			)
			if r == nil {
				t.Errorf("Unable to allocate new log record")
			}
			pr := r.(*logRecord)
			extractedMessage := r.RawMessage()
			if pr.message != expectedMessage {
				t.Errorf("Unexpected message found: %q != %q", pr.message, expectedMessage)
			}
			if extractedMessage != expectedMessage {
				t.Errorf("Unexpected message found: %q != %q", extractedMessage, expectedMessage)
			}
		},
	)
	t.Run(
		"RawMessageArgs",
		func(t *testing.T) {
			expectedName := t.Name()
			var expectedLevel LogLevelInt = INFO
			expectedMessage := "test log message"
			expectedArgs := []any{
				"foo",
				"bar",
				123,
				false,
				true,
			}
			r := NewLogRecord(
				expectedName,
				expectedLevel,
				expectedMessage,
				expectedArgs...
			)
			if r == nil {
				t.Errorf("Unable to allocate new log record")
			}
			pr := r.(*logRecord)
			extractedArgs := r.RawMessageArgs()
			if len(pr.attrs) != 0 {
				t.Errorf("Unexpected attributes found: %#v", pr.attrs)
			}
			if len(pr.messageArgs) != len(expectedArgs) {
				t.Errorf("Unexpected arg length: %d != %d", len(pr.messageArgs), len(expectedArgs))
			}
			if len(extractedArgs) != len(expectedArgs) {
				t.Errorf("Unexpected arg length: %d != %d", len(extractedArgs), len(expectedArgs))
			}
			for i, v := range expectedArgs {
				if pr.messageArgs[i] != v {
					t.Errorf("Unexpected arg found at index %d: %#v != %#v", i, pr.messageArgs[i], v)
				}
				if extractedArgs[i] != v {
					t.Errorf("Unexpected arg found at index %d: %#v != %#v", i, extractedArgs[i], v)
				}
			}
		},
	)
	t.Run(
		"Attributes",
		func(t *testing.T) {
			expectedName := t.Name()
			var expectedLevel LogLevelInt = INFO
			expectedMessage := "test log message"
			expectedArgs := []any{}
			r := NewLogRecord(
				expectedName,
				expectedLevel,
				expectedMessage,
				expectedArgs...
			)
			if r == nil {
				t.Errorf("Unable to allocate new log record")
			}
			pr := r.(*logRecord)
			if len(pr.attrs) != 0 {
				t.Errorf("Unexpected attributes found: %#v", pr.attrs)
			}

			localArgs := make([]interfaces.Attribute, 0)
			rand_args_count := rand.Intn(1000)
			for i := 0; i < rand_args_count; i++ {
				iV := NewLogAttr(
					fmt.Sprintf("log-attr-%d", i),
					i,
				)
				localArgs = append(localArgs, iV)
			}
			r.AddAttributes(localArgs...)

			extractedAttrs := r.Attributes()
			if len(pr.attrs) != len(localArgs) {
				t.Errorf("Unexpected attribute count found: %d != %d", len(pr.attrs), len(localArgs))
			}
			if len(extractedAttrs) != len(localArgs) {
				t.Errorf("Unexpected attribute count found: %d != %d", len(extractedAttrs), len(localArgs))
			}
			compareAttribute := func(a1 interfaces.Attribute, a2 interfaces.Attribute) bool {
				if a1.Key() != a2.Key() {
					return false
				}
				if a1.Value() != a2.Value() {
					return false
				}
				return true
			}

			for i, v := range localArgs {
				if !compareAttribute(pr.attrs[i], v) {
					t.Errorf("Unexpected arg found index %d: %#v != %#v", i, pr.attrs[i], v)
				}
				if !compareAttribute(extractedAttrs[i], v) {
					t.Errorf("Unexpected arg found index %d: %#v != %#v", i, extractedAttrs[i], v)
				}
			}
		},
	)
}
