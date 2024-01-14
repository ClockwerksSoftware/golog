package loggers

import (
	"testing"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

func TestGlobalMethods(t *testing.T) {
	t.Run(
		"DefaultLogger",
		func(t *testing.T) {
			if root == nil {
				t.Fatalf("init did not run and initialize the root logger")
			}
			dl := DefaultLogger()
			if dl != root {
				t.Errorf("Default logger did not return the appropriate logger: %#v != %#v", dl, root)
			}

			rl := GetLogger(rootLoggerName)
			ml := GetLogger(alternateLoggerName)

			if rl != root {
				t.Errorf("%s did not return the root logger: %#v != %#v", rootLoggerName, rl, root) 
			}
			if ml != root {
				t.Errorf("%s did not return the root logger: %#v != %#v", alternateLoggerName, ml, root) 
			}
		},
	)
	t.Run(
		"GetLogger",
		func(t *testing.T) {
			expectedName := "foobar"

			if len(root.childLoggers) != 0 {
				t.Errorf("Unexpected existing child loggers: %#v", root.childLoggers)
			}
			l := GetLogger(expectedName)
			if l.Name() != expectedName {
				t.Errorf("Child logger did not have the expected name: %q != %q", l.Name(), expectedName)
			}
			if len(root.childLoggers) != 1 {
				t.Errorf("Unexpected existing child loggers: %#v", root.childLoggers)
			}
			l2 := GetLogger(expectedName)
			if l2 != l {
				t.Errorf("Repetitiv call with same name did not yield the same logger: %#v != !%#v", l, l2)
			}
			if len(root.childLoggers) != 1 {
				t.Errorf("Unexpected existing child loggers: %#v", root.childLoggers)
			}

			// cleanup
			root.childLoggers = make(map[string]interfaces.Log)
		},
	)
	t.Run(
		"RemoveLogger",
		func(t *testing.T) {
			if len(root.childLoggers) != 0 {
				t.Errorf("Unexpected existing child loggers: %#v", root.childLoggers)
			}
			l := GetLogger(t.Name())
			if len(root.childLoggers) != 1 {
				t.Errorf("Unexpected existing child loggers: %#v", root.childLoggers)
			}
			RemoveLogger(l)
			if len(root.childLoggers) != 0 {
				t.Errorf("Unexpected existing child loggers: %#v", root.childLoggers)
			}
		},
	)
}
