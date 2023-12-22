package formatter

import (
	"fmt"
	"testing"

	"github.com/ClockwerksSoftware/golog/internal"
)

func TestBasicFormatter(t *testing.T) {
	compareBytes := func(t *testing.T, b1 []byte, b2 []byte) {
		for i, v := range b1 {
			if b2[i] != v {
				t.Errorf("Unexpected byte value at index %d: %v != %v", i, v, b2[i])
			}
		}
	}
	t.Run(
		"New",
		func(t *testing.T) {
			lf := NewLogFormatter()
			if lf == nil {
				t.Errorf("Failed to allocate formatter")
			}
		},
	)
	t.Run(
		"Name",
		func(t *testing.T) {
			lf := NewLogFormatter()
			if lf == nil {
				t.Errorf("Failed to allocate formatter")
			}
			extractedName := lf.Name()
			if extractedName != basicFormatterName {
				t.Errorf("Unexpected name found: %q != %q", extractedName, basicFormatterName)
			}
		},
	)
	t.Run(
		"Format",
		func(t *testing.T) {
			t.Run(
				"Cached",
				func(t *testing.T) {
					lf := NewLogFormatter()
					if lf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					cachedValue := "some cached value"
					mr := internal.NewMockRecord()
					mr.CacheFormatString(lf.Name(), cachedValue)
					expectedValue := []byte(cachedValue)
					extractedValue := lf.Format(mr)
					compareBytes(t, expectedValue, extractedValue)
				},
			)
			t.Run(
				"Noncached",
				func(t *testing.T) {
					lf := NewLogFormatter()
					if lf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					mr := internal.NewMockRecord()
					tValue, _ := mr.Time().MarshalText()
					strValue := fmt.Sprintf(
						"%s %s %s",
						mr.Level().String(),
						tValue,
						fmt.Sprintf(
							mr.RawMessage(),
							mr.RawMessageArgs(),
						),
					)
					expectedValue := []byte(strValue)
					extractedValue := lf.Format(mr)
					compareBytes(t, expectedValue, extractedValue)
				},
			)
		},
	)
	t.Run(
		"FormatString",
		func(t *testing.T) {
			t.Run(
				"Cached",
				func(t *testing.T) {
					lf := NewLogFormatter()
					if lf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					expectedValue := "some cached value"
					mr := internal.NewMockRecord()
					mr.CacheFormatString(lf.Name(), expectedValue)
					extractedValue := lf.FormatString(mr)
					if expectedValue != extractedValue {
						t.Errorf("Unexpected value: %q != %q", expectedValue, extractedValue)
					}
				},
			)
			t.Run(
				"Noncached",
				func(t *testing.T) {
					lf := NewLogFormatter()
					if lf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					mr := internal.NewMockRecord()
					tValue, _ := mr.Time().MarshalText()
					expectedValue := fmt.Sprintf(
						"%s %s %s",
						mr.Level().String(),
						tValue,
						fmt.Sprintf(
							mr.RawMessage(),
							mr.RawMessageArgs(),
						),
					)
					extractedValue := lf.FormatString(mr)
					if expectedValue != extractedValue {
						t.Errorf("Unexpected value: %q != %q", expectedValue, extractedValue)
					}
				},
			)
		},
	)
}
