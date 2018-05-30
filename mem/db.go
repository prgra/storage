package storage

import "sync"

type db struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

var stor db

func init() {
	stor.mutex.Lock()
	stor.data = make(map[string]interface{})
	stor.mutex.Unlock()
}

// Push :: setter
func Push(key string, i interface{}) {
	stor.mutex.Lock()
	stor.data[key] = i
	stor.mutex.Unlock()
}

// Get :: getter
func Get(key string) (interface{}, bool) {
	stor.mutex.RLock()
	// Здесь если не проверять на существование можно ускороить из - за отсутвия аллокации памяти
	res, ok := stor.data[key]
	stor.mutex.RUnlock()
	return res, ok
}
