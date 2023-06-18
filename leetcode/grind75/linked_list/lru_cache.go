// 146. LRU Cache

// Companies
// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

// Implement the LRUCache class:

// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// int get(int key) Return the value of the key if the key exists, otherwise return -1.
// void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity.

// Example 1:

// Input
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// Output
// [null, null, null, 1, null, -1, null, -1, 3, 4]

// Explanation
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // cache is {1=1}
// lRUCache.put(2, 2); // cache is {1=1, 2=2}
// lRUCache.get(1);    // return 1
// lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
// lRUCache.get(2);    // returns -1 (not found)
// lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
// lRUCache.get(1);    // return -1 (not found)
// lRUCache.get(3);    // return 3
// lRUCache.get(4);    // return 4

package linked_list

type nodeLRU struct {
	before *nodeLRU
	next   *nodeLRU
	value  int
}

type listLRU struct {
	Head *nodeLRU
	Tail *nodeLRU
}

func (this *listLRU) PushBack(node *nodeLRU) {
	if this.Head == nil {
		this.Head = node
	} else {
		this.Tail.next = node
	}
	this.Tail = node
}

func (this *listLRU) PushFront(node *nodeLRU) {
	if this.Head == nil {
		this.Head = node
		this.Tail = node
	} else {
		node.next = this.Head
		this.Head.before = node
		this.Head = node
	}
}

func (this *listLRU) RemoveTail() {
	this.Tail = this.Tail.before
	this.Tail.next = nil
}

func (this *listLRU) Remove(node *nodeLRU) {
	if node == this.Tail {
		this.Tail = node.before
		node.before.next = nil
	} else {
		node.before.next = node.next
		node.next.before = node.before
	}
	node.before = nil
	node.next = nil
}

type LRUCache struct {
	capacity int
	cache    map[int]*nodeLRU
	list     *listLRU
}

func ConstructorLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    map[int]*nodeLRU{},
		list:     &listLRU{},
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		currNode := this.cache[key]
		if currNode != this.list.Head {
			this.list.Remove(currNode)
			this.list.PushFront(currNode)
		}
		return this.cache[key].value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		if len(this.cache) == this.capacity {
			this.list.RemoveTail()
		}
		node := &nodeLRU{value: value}
		this.list.PushFront(node)
		this.cache[key] = node
	} else {
		currNode := this.cache[key]
		if currNode != this.list.Head {
			this.list.Remove(currNode)
			this.list.PushFront(currNode)
		}
		this.cache[key].value = value
	}
}
