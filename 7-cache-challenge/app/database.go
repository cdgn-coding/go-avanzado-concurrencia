package app

import "sync"

type record struct {
	key   string
	value string
}

type database struct {
	data map[string]string
	mu   sync.RWMutex
}

func newDatabase() *database {
	return &database{
		data: make(map[string]string),
	}
}

func (database *database) put(key, val string) {
	database.mu.Lock()
	database.data[key] = val
	database.mu.Unlock()
}

func (database *database) get(key string) (string, bool) {
	database.mu.RLock()
	res, err := database.data[key]
	database.mu.RUnlock()
	return res, err
}

func (database *database) remove(key string) {
	database.mu.Lock()
	delete(database.data, key)
	database.mu.Unlock()
}
