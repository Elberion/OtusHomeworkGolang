package hw04lrucache

import "sync"

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
	mu       sync.Mutex
}

type Item struct {
	Key
	Value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if val, found := l.items[key]; found {
		l.queue.MoveToFront(val)
		val.Value.(*Item).Value = value
		return true
	}

	if l.capacity == l.queue.Len() {
		if lastItem := l.queue.Back(); lastItem != nil {
			l.queue.Remove(lastItem)
			delete(l.items, lastItem.Value.(*Item).Key)
		}
	}
	item := &Item{key, value}
	l.items[key] = l.queue.PushFront(item)
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if val, found := l.items[key]; found {
		l.queue.MoveToFront(val)
		return val.Value.(*Item).Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
