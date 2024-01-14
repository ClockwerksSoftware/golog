package record

import (
	"errors"
	"testing"

	"github.com/ClockwerksSoftware/golog/common"
)

func TestLogLevel(t *testing.T) {
	t.Run(
		"new",
		func(t *testing.T) {
			l := NewLogLevel()
			if l == nil {
				t.Errorf("failed to generate a new log level instance")
			}
		},
	)
	t.Run(
		"get",
		func(t *testing.T) {
			desiredLogLevel := LogLevelInt(DEBUG)
			l := GetLogLevel(desiredLogLevel).(*LogLevel)
			if l == nil {
				t.Logf("failed to create LogLevel instance")
			} else {
				if l.level != desiredLogLevel {
					t.Logf("Unexpected log level found: %d != %d", l.level, desiredLogLevel)
				}
			}
		},
	)
	t.Run(
		"Int",
		func(t *testing.T) {
			desiredLogLevel := DEBUG
			l := &LogLevel{
				level: LogLevelInt(desiredLogLevel),
			}
			if l.Int() != int(desiredLogLevel) {
				t.Logf("Unexpected log level found: %d != %d", l.level, desiredLogLevel)
			}
		},
	)
	t.Run(
		"AddLogLevel",
		func(t *testing.T) {
			type testResult struct {
				err error
			}
			type testParameters struct {
				level     LogLevelInt
				name      string
				overwrite bool
				beginFn   func(t *testing.T)
				expectFn  func(t *testing.T, tr *testResult)
				cleanFn   func(t *testing.T)
			}
			type testScenario struct {
				name  string
				setup func(t *testing.T) testParameters
			}
			testScenarios := []testScenario{
				{
					name: "existingNoOverWrite",
					setup: func(t *testing.T) testParameters {
						return testParameters{
							level:     LogLevelInt(DEBUG),
							name:      "DEBUG",
							overwrite: false,
							beginFn:   func(t *testing.T) { /* Nothing to do */ },
							expectFn: func(t *testing.T, tr *testResult) {
								if !errors.Is(tr.err, common.ErrLogLevelAlreadyExists) {
									t.Errorf("Unexpected error received: %v != %v", tr.err, common.ErrLogLevelAlreadyExists)
								}
							},
							cleanFn: func(t *testing.T) { /* Nothing to do */ },
						}
					},
				},
				{
					name: "existingOverWrite",
					setup: func(t *testing.T) testParameters {
						myLevelInt := LogLevelInt(DEBUG)
						return testParameters{
							level:     myLevelInt,
							name:      "BUGED",
							overwrite: true,
							beginFn:   func(t *testing.T) { /* Nothing to do */ },
							expectFn: func(t *testing.T, tr *testResult) {
								if tr.err != nil {
									t.Errorf("Unexpected error received: %v", tr.err)
								}
							},
							cleanFn: func(t *testing.T) {
								// reset it
								levelMap[myLevelInt] = "DEBUG"
							},
						}
					},
				},
				{
					name: "new",
					setup: func(t *testing.T) testParameters {
						myLevelInt := LogLevelInt(120)
						return testParameters{
							level:     LogLevelInt(myLevelInt),
							name:      "BUGED",
							overwrite: true,
							beginFn:   func(t *testing.T) { /* Nothing to do */ },
							expectFn: func(t *testing.T, tr *testResult) {
								if tr.err != nil {
									t.Errorf("Unexpected error received: %v", tr.err)
								}
							},
							cleanFn: func(t *testing.T) {
								// reset it
								delete(levelMap, myLevelInt)
							},
						}
					},
				},
			}

			for _, scenario := range testScenarios {
				t.Run(
					scenario.name,
					func(t *testing.T) {
						parameters := scenario.setup(t)
						parameters.beginFn(t)
						parameters.expectFn(
							t,
							&testResult{
								err: AddLogLevel(
									parameters.level,
									parameters.name,
									parameters.overwrite,
								),
							},
						)
						parameters.cleanFn(t)
					},
				)
			}
		},
	)
	t.Run(
		"String",
		func(t *testing.T) {
			type testParameters struct {
				value         LogLevelInt
				expectedValue string
				beginFn       func(t *testing.T)
				cleanFn       func(t *testing.T)
			}
			type testScenario struct {
				name  string
				setup func(t *testing.T) testParameters
			}
			testScenarios := []testScenario{
				{
					name: "Critical",
					setup: func(t *testing.T) testParameters {
						return testParameters{
							value:         CRITICAL,
							expectedValue: "CRITICAL",
							beginFn:       func(t *testing.T) { /* nothing to do */ },
							cleanFn:       func(t *testing.T) { /* nothing to do */ },
						}
					},
				},
				{
					name: "Unknown",
					setup: func(t *testing.T) testParameters {
						return testParameters{
							value:         100,
							expectedValue: "UNKNOWN",
							beginFn:       func(t *testing.T) { /* nothing to do */ },
							cleanFn:       func(t *testing.T) { /* nothing to do */ },
						}
					},
				},
				{
					name: "custom",
					setup: func(t *testing.T) testParameters {
						myLevel := LogLevelInt(200)
						myLevelString := "beatrice"
						return testParameters{
							value:         myLevel,
							expectedValue: myLevelString,
							beginFn: func(t *testing.T) {
								_ = AddLogLevel(myLevel, myLevelString, false)
							},
							cleanFn: func(t *testing.T) {
								if _, ok := levelMap[myLevel]; ok {
									delete(levelMap, myLevel)
								}
							},
						}
					},
				},
				{
					name: "customRemoval",
					setup: func(t *testing.T) testParameters {
						return testParameters{
							value:         200,
							expectedValue: "UNKNOWN",
							beginFn:       func(t *testing.T) { /* nothing to do */ },
							cleanFn:       func(t *testing.T) { /* nothing to do */ },
						}
					},
				},
			}

			// NOTE: Do not parallize as the this test modifies the levelMap package global variable
			for _, scenario := range testScenarios {
				t.Run(
					scenario.name,
					func(t *testing.T) {
						parameters := scenario.setup(t)
						parameters.beginFn(t)
						lvl := &LogLevel{
							level: parameters.value,
						}
						s := lvl.String()
						if s != parameters.expectedValue {
							t.Errorf("Unexpected value: %s != %s", s, parameters.expectedValue)
						}
						parameters.cleanFn(t)
					},
				)
			}
		},
	)
}
