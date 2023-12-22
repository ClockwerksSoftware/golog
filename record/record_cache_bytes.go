package record

import (
	"fmt"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

type recordByteCache struct {
	cacheBytes  map[string][]byte
}

func (rbc *recordByteCache) CacheFormat(formatName string, formattedRecord []byte) {
	rbc.cacheBytes[formatName] = formattedRecord
}

func (rbc *recordByteCache) GetCacheFormat(formatName string) ([]byte, error) {
	if v, ok := rbc.cacheBytes[formatName]; ok {
		return v, nil
	}

	return nil, fmt.Errorf(
		"%w: no cached record with name %q",
		interfaces.ErrNoCachedRecord,
		formatName,
	)
}
