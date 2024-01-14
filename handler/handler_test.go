package handler

import (
	"bytes"
	"testing"

	"github.com/ClockwerksSoftware/golog/formatter"
	"github.com/ClockwerksSoftware/golog/internal"
)

func TestHandler(t *testing.T) {
	t.Run(
		"New",
		func(t *testing.T) {
			lh := New()
			if lh == nil {
				t.Errorf("Unexpectedly did not receive a handler instance")
			}
			name := lh.Name()
			if name != basicHandlerName {
				t.Errorf("Unexpected name found: %q != %q", name, basicHandlerName)
			}
			children := lh.ChildHandlers()
			if children == nil {
				t.Errorf("Unexpectedly did not get a valid empty list of child handlers")
			} else if len(children) != 0 {
				t.Errorf("Unexpectedly found a non-empty list of child handlers: %#v", children)
			}

			filters := lh.Filters()
			if filters == nil {
				t.Errorf("Unexpectedly did not get a valid empty list of filters")
			} else if len(filters) != 0 {
				t.Errorf("Unexpectedly found a non-empty list of filters: %#v", filters)
			}

			formatter := lh.Formatter()
			if formatter == nil {
				t.Errorf("Unexpectedly did not find a formatter")
			}

			writer := lh.Output()
			if writer == nil {
				t.Errorf("Unexpectedly did not find an output writer")
			}
		},
	)
	t.Run(
		"Name",
		func(t *testing.T) {
			lh := New()
			if lh == nil {
				t.Errorf("Unexpectedly did not receive a handler instance")
			}
			name := lh.Name()
			if name != basicHandlerName {
				t.Errorf("Unexpected name found: %q != %q", name, basicHandlerName)
			}
		},
	)
	t.Run(
		"Children",
		func(t *testing.T) {
			lh := New()
			if lh == nil {
				t.Errorf("Unexpectedly did not receive a handler instance")
			}
			children := lh.ChildHandlers()
			if children == nil {
				t.Errorf("Unexpectedly did not get a valid empty list of child handlers")
			} else if len(children) != 0 {
				t.Errorf("Unexpectedly found a non-empty list of child handlers: %#v", children)
			}

			childLh := New()
			lh.AddChildHandler(childLh)
			children = lh.ChildHandlers()
			if children == nil {
				t.Errorf("Unexpectedly did not get a valid list of child handlers")
			} else if len(children) != 1 {
				t.Errorf("Unexpectedly found an unexpected list of child handlers: %#v", children)
			}
		},
	)
	t.Run(
		"Filters",
		func(t *testing.T) {
			lh := New()
			if lh == nil {
				t.Errorf("Unexpectedly did not receive a handler instance")
			}
			filters := lh.Filters()
			if filters == nil {
				t.Errorf("Unexpectedly did not get a valid empty list of filters")
			} else if len(filters) != 0 {
				t.Errorf("Unexpectedly found a non-empty list of filters: %#v", filters)
			}
			newFilter := &internal.MockFilter{}
			lh.AddFilter(newFilter)
			filters = lh.Filters()
			if filters == nil {
				t.Errorf("Unexpectedly did not get a valid list of filters")
			} else if len(filters) != 1 {
				t.Errorf("Unexpectedly found an unexpected list of filters: %#v", filters)
			}
		},
	)
	t.Run(
		"Formatter",
		func(t *testing.T) {
			lh := New()
			if lh == nil {
				t.Errorf("Unexpectedly did not receive a handler instance")
			}
			formatter := lh.Formatter()
			if formatter == nil {
				t.Errorf("Unexpectedly did not find a formatter")
			}
			lh.SetFormatter(nil)
			if lh.Formatter() != nil {
				t.Errorf("Unexpectedly unable to set the formatter to nil - %#v", lh.Formatter())
			}
			lh.SetFormatter(formatter)
			if formatter == nil {
				t.Errorf("Unexpectedly did not find a formatter")
			}
		},
	)
	t.Run(
		"Output",
		func(t *testing.T) {
			lh := New()
			if lh == nil {
				t.Errorf("Unexpectedly did not receive a handler instance")
			}
			writer := lh.Output()
			if writer == nil {
				t.Errorf("Unexpectedly did not find an output writer")
			}
			lh.SetOutput(nil)
			if lh.Output() != nil {
				t.Errorf("Unexpectedly unable to set the output to nil - %#v", lh.Output())
			}
			lh.SetOutput(writer)
			if writer == nil {
				t.Errorf("Unexpectedly did not find an output writer")
			}
		},
	)
	t.Run(
		"Handle",
		func(t *testing.T) {
			configureHandler := func() (*logHandler, *bytes.Buffer) {
				lh := New()
				buffer := &bytes.Buffer{}
				lh.SetOutput(buffer)
				return lh.(*logHandler), buffer
			}
			t.Run(
				"Nil Log Record",
				func(t *testing.T) {
					lh, lhOutput := configureHandler()
					lh.Handle(nil)

					out := lhOutput.String()
					if len(out) > 0 {
						t.Errorf("Unexpected output received: %q", out)
					}
				},
			)
			t.Run(
				"Filtered Out",
				func(t *testing.T) {
					lh, lhOutput := configureHandler()
					lfilter := &internal.MockFilter{
						AllowRecord: false,
					}
					lh.AddFilter(lfilter)

					lr := internal.NewMockRecord()
					lh.Handle(lr)
					out := lhOutput.String()
					if len(out) > 0 {
						t.Errorf("Unexpected output received: %q", out)
					}
				},
			)
			t.Run(
				"Without Formatter",
				func(t *testing.T) {
					lh, lhOutput := configureHandler()
					lh.SetFormatter(nil)

					lr := internal.NewMockRecord()
					lr.MockRawMessage = "raw formatter - %s %s %s"
					lr.MockRawMessageArgs = []any{
						"raw",
						"message",
						"args",
					}

					lh.Handle(lr)
					expectedOutput := "raw formatter - raw message args"
					out := lhOutput.String()
					if out != expectedOutput {
						t.Errorf("Unexpected output received: %q != %q", out, expectedOutput)
					}
				},
			)
			t.Run(
				"With Formatter",
				func(t *testing.T) {
					lh, lhOutput := configureHandler()

					lr := internal.NewMockRecord()
					lr.MockRawMessage = "raw formatter - %#v"
					lr.MockRawMessageArgs = []any{
						"raw",
						"message",
						"args",
					}

					lh.Handle(lr)
					basicFormatter := formatter.NewLogFormatter()
					expectedOutput := basicFormatter.FormatString(lr)
					out := lhOutput.String()
					if out != expectedOutput {
						t.Errorf("Unexpected output received: %q != %q", out, expectedOutput)
					}
				},
			)
			t.Run(
				"Nil Output Stream",
				func(t *testing.T) {
					lh, _ := configureHandler()
					lh.SetFormatter(nil)
					lh.SetOutput(nil)

					lr := internal.NewMockRecord()
					lr.MockRawMessage = "raw formatter - %#v"
					lr.MockRawMessageArgs = []any{
						"raw",
						"message",
						"args",
					}

					lh.Handle(lr)
					// fails if it crashes!
				},
			)
			t.Run(
				"Child Handler",
				func(t *testing.T) {
					lh, _ := configureHandler()
					lh.SetFormatter(nil)
					lh.SetOutput(nil)

					for i := 0; i < 10; i++ {
						lhN, _ := configureHandler()
						lhN.SetFormatter(nil)
						lhN.SetOutput(nil)
						lh.AddChildHandler(lhN)
					}

					lr := internal.NewMockRecord()
					lr.MockRawMessage = "raw formatter - %#v"
					lr.MockRawMessageArgs = []any{
						"raw",
						"message",
						"args",
					}

					lh.Handle(lr)
					// fails if it crashes!
				},
			)
		},
	)
}
