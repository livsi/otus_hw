package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if _, ok := c.items[key]; ok {
		c.items[key].Value = cacheItem{key, value}
		c.queue.MoveToFront(c.items[key])
		return true
	}

	c.items[key] = c.queue.PushFront(cacheItem{key, value})

	if c.queue.Len() > c.capacity {
		rmKey := c.queue.Back().Value.(cacheItem)
		delete(c.items, rmKey.key)
		c.queue.Remove(c.queue.Back())
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if _, ok := c.items[key]; ok {
		c.queue.MoveToFront(c.items[key])
		val := c.items[key].Value.(cacheItem)
		return val.value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = nil
	c.queue = NewList()
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
