package formatter

import (
	"fmt"
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/ClockwerksSoftware/golog/interfaces"
	"github.com/ClockwerksSoftware/golog/internal"
	"github.com/ClockwerksSoftware/golog/record"
)


func TestJsonFormatter(t *testing.T) {
	makeJsonValue := func(mr interfaces.Record) ([]byte, error) {
		data := map[string]any {}
		data["name"] = mr.Name()
		data["message"] = fmt.Sprintf(
			mr.RawMessage(),
			mr.RawMessageArgs(),
		)
		data["level"] = mr.Level().String()
		tempTime, _ := mr.Time().MarshalText()
		data["time"] = string(tempTime)
		for _, attr := range mr.Attributes() {
			data[attr.Key()] = attr.Value()
		}
		
		return json.Marshal(data)
	}
	compareBytes := func(t *testing.T, b1 []byte, b2 []byte) {
		if len(b1) != len(b2) {
			t.Errorf("Unexpected byte data: %#v != %#v", b1, b2)
			return
		}
		for i, v := range b1 {
			if b2[i] != v {
				t.Errorf("Unexpected byte value at index %d: %v != %v", i, v, b2[i])
			}
		}
	}
		/*
	compareJson := func(t *testing.T, original interfaces.Record, value string) {
		t.Logf("encoded: %q", value)
		type data struct {
			name string `json:"name"`
			message string `json:"message"`
			level string `json:"level"`
			time string `json:"time"`
		}
	}
		*/
	t.Run(
		"New",
		func(t *testing.T) {
			ljf := NewLogJsonFormatter()
			if ljf == nil {
				t.Errorf("Failed to allocate formatter")
			}
		},
	)
	t.Run(
		"Name",
		func(t *testing.T) {
			ljf := NewLogJsonFormatter()
			if ljf == nil {
				t.Errorf("Failed to allocate formatter")
			}
			extractedName := ljf.Name()
			if extractedName != jsonFormatterName {
				t.Errorf("Unexpected name found: %q != %q", extractedName, jsonFormatterName)
			}
		},
	)
	t.Run(
		"Format",
		func(t *testing.T) {
			t.Run(
				"Cached",
				func(t *testing.T) {
					ljf := NewLogJsonFormatter()
					if ljf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					cachedValue := []byte("some cached value")
					mr := internal.NewMockRecord()
					mr.CacheFormat(ljf.Name(), cachedValue)
					expectedValue := []byte(cachedValue)
					extractedValue := ljf.Format(mr)
					compareBytes(t, expectedValue, extractedValue)
				},
			)
			t.Run(
				"Noncached",
				func(t *testing.T) {
					ljf := NewLogJsonFormatter()
					if ljf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					mr := internal.NewMockRecord()
					expectedValue, jsonDataErr := makeJsonValue(mr)
					if jsonDataErr != nil {
						t.Errorf("Unable to generate expected value: %#v", jsonDataErr)
					}
					extractedValue := ljf.Format(mr)
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
					ljf := NewLogJsonFormatter()
					if ljf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					expectedValue := []byte("some cached value")
					mr := internal.NewMockRecord()
					mr.CacheFormat(ljf.Name(), expectedValue)
					extractedValue := ljf.FormatString(mr)
					if string(expectedValue) != extractedValue {
						t.Errorf("Unexpected value: %q != %q", string(expectedValue), extractedValue)
					}
				},
			)
			t.Run(
				"Noncached",
				func(t *testing.T) {
					ljf := NewLogJsonFormatter()
					if ljf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					mr := internal.NewMockRecord()
					jsonByteData, jsonDataErr := makeJsonValue(mr)
					if jsonDataErr != nil {
						t.Errorf("Unable to generate expected value: %#v", jsonDataErr)
					}
					expectedValue := string(jsonByteData)
					extractedValue := ljf.FormatString(mr)
					if expectedValue != extractedValue {
						t.Errorf("Unexpected value: %q != %q", expectedValue, extractedValue)
					}
				},
			)
			t.Run(
				"Noncached with attributes",
				func(t *testing.T) {
					ljf := NewLogJsonFormatter()
					if ljf == nil {
						t.Errorf("Failed to allocate formatter")
					}
					mr := internal.NewMockRecord()
					localArgs := make([]interfaces.Attribute, 0)
					rand_args_count := rand.Intn(1000)
					for i := 0; i < rand_args_count; i++ {
						iV := record.NewLogAttr(
							fmt.Sprintf("log-attr-%d", i),
							i,
						)
						localArgs = append(localArgs, iV)
					}
					mr.AddAttributes(localArgs...)

					jsonByteData, jsonDataErr := makeJsonValue(mr)
					if jsonDataErr != nil {
						t.Errorf("Unable to generate expected value: %#v", jsonDataErr)
					}
					expectedValue := string(jsonByteData)
					extractedValue := ljf.FormatString(mr)
					if expectedValue != extractedValue {
						t.Errorf("Unexpected value: %q != %q", expectedValue, extractedValue)
					}
				},
			)
		},
	)
}
