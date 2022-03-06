package testing

import "sync"

var (
	data   = make(map[string]string)
	locker sync.RWMutex
)

func WriteToMap(k, v string) {
	locker.Lock()
	defer locker.Unlock()
	data[k] = v
}

func ReadFromMap(k string) string {
	locker.RLock()
	defer locker.RUnlock()
	return data[k]
}
