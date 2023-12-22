package record

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

func TestRecordStringCache(t *testing.T) {
	newCache := func() *recordStringCache {
		return &recordStringCache{
			cacheString: make(map[string]string),
		}
	}
	getKeyFromIndex := func(t *testing.T, index int, m map[string]string) string {
		keys := []string{}
		for k, _ := range m {
			keys = append(keys, k)
		}
		if len(keys) < index {
			t.Errorf("Invalid index: %d < %d", len(keys), index)
			return ""
		}
		return keys[index]
	}

	t.Run(
		"CacheFormat",
		func(t *testing.T) {
			t.Run(
				"Initial",
				func(t *testing.T) {
					rbc := newCache()
					if len(rbc.cacheString) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheString)
					}
					strData := "foobar"
					nameData := "bar"
					rbc.CacheFormatString(nameData, strData)
					if len(rbc.cacheString) != 1 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheString)
					}
					if v, ok := rbc.cacheString[nameData]; ok {
						if strData != v {
							t.Errorf("Unexpected value: %q != %q", strData, v)
						}
					} else {
						t.Errorf("Unable to find key %s in the cache data: %#v", nameData, rbc.cacheString)
					}
				},
			)
			t.Run(
				"Multiple",
				func(t *testing.T) {
					rbc := newCache()
					if len(rbc.cacheString) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheString)
					}
					localMap := make(map[string]string)
					for i := 0; i < rand.Intn(1000); i++ {
						strData := fmt.Sprintf("foobar-%d", i)
						nameData := fmt.Sprintf("bar-%d", i)
						rbc.CacheFormatString(nameData, strData)
						localMap[nameData] = strData
					}

					if len(rbc.cacheString) != len(localMap) {
						t.Errorf("Unexpected data lengths: %d != %d", len(rbc.cacheString), len(localMap))
					}
					for k, kV := range localMap {
						if v, ok := rbc.cacheString[k]; ok {
							if kV != v {
								t.Errorf("Unexpected value: %q != %q", kV, v)
							}
						} else {
							t.Errorf("[%s] Unable to find key in the cache data: %#v", k, rbc.cacheString)
						}
					}
				},
			)
		},
	)
	t.Run(
		"GetCacheFormat",
		func(t *testing.T) {
			t.Run(
				"Empty",
				func(t *testing.T) {
					rbc := newCache()
					if len(rbc.cacheString) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheString)
					}
					_, err := rbc.GetCacheFormatString("non-existent-key")
					if !errors.Is(err, interfaces.ErrNoCachedRecord) {
						t.Errorf("Unexpected error received: %#v != %#v", err, interfaces.ErrNoCachedRecord)
					}
				},
			)
			t.Run(
				"Only-One",
				func(t *testing.T) {
					rbc := newCache()
					if len(rbc.cacheString) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheString)
					}
					strData := "foobar"
					nameData := "bar"
					rbc.CacheFormatString(nameData, strData)
					if len(rbc.cacheString) != 1 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheString)
					}
					extractedValue, err := rbc.GetCacheFormatString(nameData)
					if err != nil {
						t.Errorf("Unexpected error received: %#v", err)
					}
					if strData != extractedValue {
						t.Errorf("Unexpected value: %q != %q", strData, extractedValue)
					}
				},
			)
			t.Run(
				"From Many",
				func(t *testing.T) {
					rbc := newCache()
					if len(rbc.cacheString) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheString)
					}
					localMap := make(map[string]string)
					for i := 0; i < rand.Intn(1000); i++ {
						strData := fmt.Sprintf("foobar-%d", i)
						nameData := fmt.Sprintf("bar-%d", i)
						rbc.CacheFormatString(nameData, strData)
						localMap[nameData] = strData
					}

					if len(rbc.cacheString) != len(localMap) {
						t.Errorf("Unexpected data lengths: %d != %d", len(rbc.cacheString), len(localMap))
					}

					randEntry := getKeyFromIndex(
						t,
						rand.Intn(len(localMap)),
						localMap,
					)
					if len(randEntry) > 0 {
						extractedValue, err := rbc.GetCacheFormatString(randEntry)
						if err != nil {
							t.Errorf("Unexpected error received: %#v", err)
						}
						if localMap[randEntry] != extractedValue {
							t.Errorf("Unexpected value: %q != %q", localMap[randEntry], extractedValue)
						}
					}
				},
			)
		},
	)
}

