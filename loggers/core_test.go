package loggers

import (
	"fmt"
	"testing"

	"github.com/ClockwerksSoftware/golog/interfaces"
	"github.com/ClockwerksSoftware/golog/internal"
	"github.com/ClockwerksSoftware/golog/record"
)



func TestCoreLogger(t *testing.T) {
	t.Run(
		"construct",
		func(t *testing.T) {
			expectedName := "foo"
			cl := &coreLogger{
				name: expectedName,
				handlers: make([]interfaces.Handler, 0),
				childLoggers: make(map[string]interfaces.Log),
			}

			if cl.Name() != expectedName {
				t.Errorf("Unexpected `name` received: %q != %q", cl.Name(), expectedName)
			}
			if len(cl.handlers) != 0 {
				t.Errorf("Unexpected handlers found: %#v", cl.handlers)
			}
			if len(cl.childLoggers) != 0 {
				t.Errorf("Unexpected child loggers found: %#v", cl.childLoggers)
			}
		},
	)
	t.Run(
		"logCalls",
		func(t *testing.T) {
			t.Run(
				"unformatted",
				func(t *testing.T) {
					type TestScenarioParameters struct {
						beginFn func(t *testing.T)
						doFn func(t *testing.T, cl *coreLogger)
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
									doFn: func(t *testing.T, cl *coreLogger) {
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
								mh := internal.NewMockHandler(t.Name())
								cl.AddHandler(mh)

								parameters := scenario.setupFn(t)

								parameters.beginFn(t)
								parameters.doFn(t, cl)
								parameters.expectFn(t, mh.Buffer.String())
							},
						)
					}
				},
			)
		},
	)
}
