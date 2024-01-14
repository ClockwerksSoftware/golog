package loggers

import (
	"fmt"
	"testing"

	"github.com/ClockwerksSoftware/golog/interfaces"
	"github.com/ClockwerksSoftware/golog/internal"
	"github.com/ClockwerksSoftware/golog/record"
)

func TestChildLogger(t *testing.T) {
	t.Run(
		"construct",
		func(t *testing.T) {
			expectedName := "foo"
			cl := &coreLogger{
				name: expectedName,
				handlers: make([]interfaces.Handler, 0),
				childLoggers: make(map[string]interfaces.Log),
			}
			ccl := NewChildLogger(expectedName, cl).(*childLogger)

			if ccl.Name() != expectedName {
				t.Errorf("Unexpected `name` received: %q != %q", cl.Name(), expectedName)
			}
			if ccl.root != cl {
				t.Errorf("Unexpected root found: %#v != %#v", ccl.root, cl)
			}
		},
	)
	t.Run(
		"logCalls",
		func(t *testing.T) {
			type TestScenarioParameters struct {
				beginFn func(t *testing.T)
				doFn func(t *testing.T, cl *childLogger)
				expectFn func(t *testing.T, output string)
			}
			type TestScenario struct {
				name string
				setupFn func(t *testing.T) TestScenarioParameters
			}
			TestScenarios := []TestScenario{
				{
					name: "Info",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "some message"
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Info(msg)
							},
							expectFn: func(t *testing.T, output string) {
								if output != msg {
									t.Errorf("Unexpected result: %q != %q", output, msg)
								}
							},
						}
					},
				},
				{
					name: "Infof",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "some message - %s, %d"
						args := []any{"foo", 100}
						expectedMsg := fmt.Sprintf(msg, args...)
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Infof(msg, args...)
							},
							expectFn: func(t *testing.T, output string) {
								if output != expectedMsg {
									t.Errorf("Unexpected result: %q != %q", output, expectedMsg)
								}
							},
						}
					},
				},
				{
					name: "Warning",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "warn some message"
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Warning(msg)
							},
							expectFn: func(t *testing.T, output string) {
								if output != msg {
									t.Errorf("Unexpected result: %q != %q", output, msg)
								}
							},
						}
					},
				},
				{
					name: "Warningf",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "warn some message - %s, %d"
						args := []any{"foo", 100}
						expectedMsg := fmt.Sprintf(msg, args...)
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Warningf(msg, args...)
							},
							expectFn: func(t *testing.T, output string) {
								if output != expectedMsg {
									t.Errorf("Unexpected result: %q != %q", output, expectedMsg)
								}
							},
						}
					},
				},
				{
					name: "Error",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "error some message"
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Error(msg)
							},
							expectFn: func(t *testing.T, output string) {
								if output != msg {
									t.Errorf("Unexpected result: %q != %q", output, msg)
								}
							},
						}
					},
				},
				{
					name: "Errorf",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "error some message - %s, %d"
						args := []any{"foo", 100}
						expectedMsg := fmt.Sprintf(msg, args...)
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Errorf(msg, args...)
							},
							expectFn: func(t *testing.T, output string) {
								if output != expectedMsg {
									t.Errorf("Unexpected result: %q != %q", output, expectedMsg)
								}
							},
						}
					},
				},
				{
					name: "Debug",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "debug some message"
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Debug(msg)
							},
							expectFn: func(t *testing.T, output string) {
								if output != msg {
									t.Errorf("Unexpected result: %q != %q", output, msg)
								}
							},
						}
					},
				},
				{
					name: "Debugf",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "debug some message - %s, %d"
						args := []any{"foo", 100}
						expectedMsg := fmt.Sprintf(msg, args...)
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Debugf(msg, args...)
							},
							expectFn: func(t *testing.T, output string) {
								if output != expectedMsg {
									t.Errorf("Unexpected result: %q != %q", output, expectedMsg)
								}
							},
						}
					},
				},
				{
					name: "Critical",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "critical some message"
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Critical(msg)
							},
							expectFn: func(t *testing.T, output string) {
								if output != msg {
									t.Errorf("Unexpected result: %q != %q", output, msg)
								}
							},
						}
					},
				},
				{
					name: "Criticalf",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "critical some message - %s, %d"
						args := []any{"foo", 100}
						expectedMsg := fmt.Sprintf(msg, args...)
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Criticalf(msg, args...)
							},
							expectFn: func(t *testing.T, output string) {
								if output != expectedMsg {
									t.Errorf("Unexpected result: %q != %q", output, expectedMsg)
								}
							},
						}
					},
				},
				{
					name: "Log",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "log some message"
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Log(record.GetLogLevel(record.DEBUG), msg)
							},
							expectFn: func(t *testing.T, output string) {
								if output != msg {
									t.Errorf("Unexpected result: %q != %q", output, msg)
								}
							},
						}
					},
				},
				{
					name: "Logf",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "log some message - %s, %d"
						args := []any{"foo", 100}
						expectedMsg := fmt.Sprintf(msg, args...)
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								cl.Logf(record.GetLogLevel(record.ERROR), msg, args...)
							},
							expectFn: func(t *testing.T, output string) {
								if output != expectedMsg {
									t.Errorf("Unexpected result: %q != %q", output, expectedMsg)
								}
							},
						}
					},
				},
				{
					name: "ProcessRecord",
					setupFn: func(t *testing.T) TestScenarioParameters {
						msg := "log some message - %s, %d"
						args := []any{"foo", 100}
						expectedMsg := fmt.Sprintf(msg, args...)
						return TestScenarioParameters{
							beginFn: func(t *testing.T) {},
							doFn: func(t *testing.T, cl *childLogger) {
								r := record.NewLogRecord("foo", record.DEBUG, msg, args...)
								cl.ProcessRecord(r)
							},
							expectFn: func(t *testing.T, output string) {
								if output != expectedMsg {
									t.Errorf("Unexpected result: %q != %q", output, expectedMsg)
								}
							},
						}
					},
				},
			}

			for _, scenario := range TestScenarios {
				scenario := scenario
				t.Run(
					scenario.name,
					func(t *testing.T) {
						cl := &coreLogger{
							name:  t.Name(),
							handlers: make([]interfaces.Handler, 0),
							childLoggers: make(map[string]interfaces.Log),
						}
						ccl := NewChildLogger(t.Name(), cl).(*childLogger)
						mh := internal.NewMockHandler(t.Name())
						cl.AddHandler(mh)

						parameters := scenario.setupFn(t)

						parameters.beginFn(t)
						parameters.doFn(t, ccl)
						parameters.expectFn(t, mh.Buffer.String())
					},
				)
			}
		},
	)
	t.Run(
		"addHandler",
		func(t *testing.T) {
			cl := &coreLogger{
				name:  t.Name(),
				handlers: make([]interfaces.Handler, 0),
				childLoggers: make(map[string]interfaces.Log),
			}
			ccl := NewChildLogger(t.Name(), cl).(*childLogger)
			mh := internal.NewMockHandler(t.Name())
			if len(cl.handlers) != 0 {
				t.Errorf("Unexpected handler found: %#v", cl.handlers)
			}
			ccl.AddHandler(mh)
			if len(cl.handlers) != 1 {
				t.Errorf("Unexpected handlers found: (len != 1) - %#v", cl.handlers)
			}
		},
	)
}
