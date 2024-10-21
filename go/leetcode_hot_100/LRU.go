package leetcodehot100

import "container/list"

type LRUCache struct {
	size     int
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
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (c *LRUCache) Get(key int) int {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.size++
	}
	for c.capacity != 0 && c.size > c.capacity {
		c.RemoveOldest()
	}
}

func (c *LRUCache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.size--
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */