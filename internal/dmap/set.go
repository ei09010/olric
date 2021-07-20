package dmap

import (
	"github.com/buraksezer/olric/internal/protocol"
)

// Expire: expire,
// Flags:  flags,
// Length: length,

// Put sets the value for the given key. It overwrites any previous value
// for that key and it's thread-safe. The key has to be string. value type
// is arbitrary. It is safe to modify the contents of the arguments after
// Put returns but not before.

func (mc *MemCacheMap) Set(key string, value interface{}, expire int64, length int64, memFlags int32) error {

	env, err := mc.DMap.prepareAndSerialize(protocol.OpMemCachedSet, key, value, nilTimeout, 0, length, memFlags, expire)
	if err != nil {
		return err
	}
	return mc.DMap.put(env)
}
