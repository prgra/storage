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

func (d *db) Push(key string, i interface{}) {
	d.mutex.Lock()
	d.data[key] = i
	d.mutex.Unlock()
}

func (d *db) Get(key string) (interface{}, bool) {
	d.mutex.RLock()
	// Здесь если не проверять на существование можно ускороить из - за отсутвия аллокации памяти
	res, ok := d.data[key]
	d.mutex.RUnlock()
	return res, ok
}
