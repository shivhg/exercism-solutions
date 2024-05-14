package singleton_pattern

import (
	"container/list"
	"math/rand"
	"sync"
)

type randomGenerator struct {
}

var instance *randomGenerator
var once sync.Once

func newGenerator() *randomGenerator {
	once.Do(func() {
		instance = &randomGenerator{}
	})

	return instance
}

func (randomGenerator) generateRandomInt() int {
	return rand.Int()
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	ll       *list.List
}
type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}
	if c.ll.Len() == c.capacity {
		backElem := c.ll.Back()
		c.ll.Remove(backElem)
		delete(c.cache, backElem.Value.(*entry).key)
	}
	newEntry := &entry{key, value}
	elem := c.ll.PushFront(newEntry)
	c.cache[key] = elem
}
