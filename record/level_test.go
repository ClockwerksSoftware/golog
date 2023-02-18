package record

import (
    //"fmt"
    "testing"
)

func TestLogLevel(t *testing.T) {
    t.Run(
        "String",
        func(t *testing.T) {
            type testParameters struct {
                value LogLevelInt
                expectedValue string
                beginFn func(t *testing.T)
                cleanFn func(t *testing.T)
            }
            type testScenario struct {
                name string
                setup func(t *testing.T) testParameters
            }
            testScenarios := []testScenario{
                {
                    name: "Critical",
                    setup: func(t *testing.T) testParameters {
                        return testParameters{
                            value: CRITICAL,
                            expectedValue: "CRITICAL",
                            beginFn: func(t *testing.T) {/* nothing to do */},
                            cleanFn: func(t *testing.T) {/* nothing to do */},
                        }
                    },
                },
                {
                    name: "Unknown",
                    setup: func(t *testing.T) testParameters {
                        return testParameters{
                            value: 100,
                            expectedValue: "UNKNOWN",
                            beginFn: func(t *testing.T) {/* nothing to do */},
                            cleanFn: func(t *testing.T) {/* nothing to do */},
                        }
                    },
                },
                {
                    name: "custom",
                    setup: func(t *testing.T) testParameters {
                        myLevel := LogLevelInt(200)
                        myLevelString := "beatrice"
                        return testParameters{
                            value: myLevel,
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
                            value: 200,
                            expectedValue: "UNKNOWN",
                            beginFn: func(t *testing.T) {/* nothing to do */},
                            cleanFn: func(t *testing.T) {/* nothing to do */},
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
                            level:parameters.value,
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
