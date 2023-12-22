package record

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

func TestRecordByteCache(t *testing.T) {
	newCache := func() *recordByteCache {
		return &recordByteCache{
			cacheBytes: make(map[string][]byte),
		}
	}
	compareBytes := func(t *testing.T, b1 []byte, b2 []byte) {
		for i, v := range b1 {
			if b2[i] != v {
				t.Errorf("Unexpected byte value at index %d: %v != %v", i, v, b2[i])
			}
		}
	}
	getKeyFromIndex := func(t *testing.T, index int, m map[string][]byte) string {
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
					if len(rbc.cacheBytes) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheBytes)
					}
					byteData := []byte("foobar")
					nameData := "bar"
					rbc.CacheFormat(nameData, byteData)
					if len(rbc.cacheBytes) != 1 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheBytes)
					}
					if v, ok := rbc.cacheBytes[nameData]; ok {
						compareBytes(t, byteData, v)
					} else {
						t.Errorf("Unable to find key %s in the cache data: %#v", nameData, rbc.cacheBytes)
					}
				},
			)
			t.Run(
				"Multiple",
				func(t *testing.T) {
					rbc := newCache()
					if len(rbc.cacheBytes) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheBytes)
					}
					localMap := make(map[string][]byte)
					for i := 0; i < rand.Intn(1000); i++ {
						byteData := []byte(
							fmt.Sprintf("foobar-%d", i),
						)
						nameData := fmt.Sprintf("bar-%d", i)
						rbc.CacheFormat(nameData, byteData)
						localMap[nameData] = byteData
					}

					if len(rbc.cacheBytes) != len(localMap) {
						t.Errorf("Unexpected data lengths: %d != %d", len(rbc.cacheBytes), len(localMap))
					}
					for k, kV := range localMap {
						if v, ok := rbc.cacheBytes[k]; ok {
							compareBytes(t, kV, v)
						} else {
							t.Errorf("[%s] Unable to find key in the cache data: %#v", k, rbc.cacheBytes)
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
					if len(rbc.cacheBytes) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheBytes)
					}
					_, err := rbc.GetCacheFormat("non-existent-key")
					if !errors.Is(err, interfaces.ErrNoCachedRecord) {
						t.Errorf("Unexpected error received: %#v != %#v", err, interfaces.ErrNoCachedRecord)
					}
				},
			)
			t.Run(
				"Only-One",
				func(t *testing.T) {
					rbc := newCache()
					if len(rbc.cacheBytes) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheBytes)
					}
					byteData := []byte("foobar")
					nameData := "bar"
					rbc.CacheFormat(nameData, byteData)
					if len(rbc.cacheBytes) != 1 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheBytes)
					}
					extractedValue, err := rbc.GetCacheFormat(nameData)
					if err != nil {
						t.Errorf("Unexpected error received: %#v", err)
					}
					compareBytes(t, byteData, extractedValue)
				},
			)
			t.Run(
				"From Many",
				func(t *testing.T) {
					rbc := newCache()
					if len(rbc.cacheBytes) != 0 {
						t.Errorf("Unexpected data in the cache: %#v", rbc.cacheBytes)
					}
					localMap := make(map[string][]byte)
					for i := 0; i < rand.Intn(1000); i++ {
						byteData := []byte(
							fmt.Sprintf("foobar-%d", i),
						)
						nameData := fmt.Sprintf("bar-%d", i)
						rbc.CacheFormat(nameData, byteData)
						localMap[nameData] = byteData
					}

					if len(rbc.cacheBytes) != len(localMap) {
						t.Errorf("Unexpected data lengths: %d != %d", len(rbc.cacheBytes), len(localMap))
					}

					randEntry := getKeyFromIndex(
						t,
						rand.Intn(len(localMap)),
						localMap,
					)
					if len(randEntry) > 0 {
						extractedValue, err := rbc.GetCacheFormat(randEntry)
						if err != nil {
							t.Errorf("Unexpected error received: %#v", err)
						}
						compareBytes(t, localMap[randEntry], extractedValue)
					}
				},
			)
		},
	)
}
