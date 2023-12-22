package record

import (
	"testing"
)

func TestLogAttr(t *testing.T) {
	t.Run(
		"Key",
		func(t *testing.T) {
			keyName := "foo"
			la := LogAttr{
				key: keyName,
			}

			result := la.Key()
			if result != keyName {
				t.Errorf("Unexpected Keyname: '%s' != '%s'", result, keyName)
			}
		},
	)
	t.Run(
		"Value",
		func(t *testing.T) {
			value := "bar"
			la := LogAttr{
				value: value,
			}

			result := la.Value()
			switch v := result.(type) {
			case string:
				if v != value {
					t.Errorf("Unexpected value: '%s' != '%s'", result, value)
				}
			default:
				t.Errorf("Unexpected value type %v != string", v)
			}
		},
	)
}
