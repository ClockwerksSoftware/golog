package record

import (
	"fmt"

	"github.com/ClockwerksSoftware/golog/interfaces"
)

type recordStringCache struct {
	cacheString map[string]string
}

func (rsc *recordStringCache) CacheFormatString(formatName string, formattedRecord string) {
	rsc.cacheString[formatName] = formattedRecord
}

func (rsc *recordStringCache) GetCacheFormatString(formatName string) (string, error) {
	if v, ok := rsc.cacheString[formatName]; ok {
		return v, nil
	}

	return "", fmt.Errorf(
		"%w: no cached record with name %q",
		interfaces.ErrNoCachedRecord,
		formatName,
	)
}
