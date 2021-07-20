package client

import (
	"github.com/buraksezer/olric/internal/protocol"
)

// DMap provides methods to access distributed maps on Olric cluster.
type MemCacheMap struct {
	*DMap
}

// Expire    int64
// Flags     int32
// Length    int64

// better description tha this
// Put sets the value for the given key. It overwrites any previous value for that key and it's thread-safe.
// It is safe to modify the contents of the arguments after Put returns but not before.

// check noreply flag situation
func (mc *MemCacheMap) Set(key string, value interface{}, expire int64, flags int32, length int64) error {
	data, err := mc.DMap.serializer.Marshal(value)
	if err != nil {
		return err
	}
	req := protocol.NewDMapMessage(protocol.OpMemCachedSet)
	req.SetDMap(mc.DMap.name)
	req.SetKey(key)
	req.SetValue(data)
	req.SetExtra(protocol.SetExtra{
		Expire: expire,
		Flags:  flags,
		Length: length,
	})
	resp, err := mc.DMap.request(req)
	if err != nil {
		return err
	}
	return checkStatusCode(resp)
}
