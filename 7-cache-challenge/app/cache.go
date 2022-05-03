package app

import "sync"

type Cache struct {
	db       *database
	fsWorker *filesystemWorker
}

func NewCache(fsBasePath string) *Cache {
	cache := &Cache{
		db:       newDatabase(),
		fsWorker: newFilesystemWorker(fsBasePath),
	}

	go cache.fsWorker.start()

	var wg sync.WaitGroup
	wg.Add(1)
	go cache.load(&wg)
	wg.Wait()

	return cache
}

func (cache *Cache) load(wg *sync.WaitGroup) {
	defer wg.Done()
	recordsToSave := make(chan record, 10)
	go cache.fsWorker.load(recordsToSave)
	for record := range recordsToSave {
		cache.Store(record.key, record.value)
	}
}

func (cache *Cache) Store(key, val string) {
	go func(cache *Cache) {
		cache.db.put(key, val)
		cache.fsWorker.persist(key, val)
	}(cache)
}

func (cache *Cache) Get(key string) (string, bool) {
	return cache.db.get(key)
}

func (cache *Cache) Remove(key string) {
	go func(cache *Cache) {
		cache.db.remove(key)
	}(cache)
}
