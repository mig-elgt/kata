package grind75

// 146. LRU Cache

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

// Constraints:

// 1 <= capacity <= 3000
// 0 <= key <= 104
// 0 <= value <= 105
// At most 2 * 105 calls will be made to get and put.

type LRUNode struct {
	Key, Val int
	Before   *LRUNode
	Next     *LRUNode
}

type LRUCache struct {
	Capacity int
	Head     *LRUNode
	Tail     *LRUNode
	Keys     map[int]*LRUNode
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{Capacity: capacity, Keys: map[int]*LRUNode{}}
}

func (this *LRUCache) Get(key int) int {
	if _, found := this.Keys[key]; !found {
		return -1
	}
	return this.Keys[key].Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, found := this.Keys[key]; !found {
		if len(this.Keys) < this.Capacity {
			node := &LRUNode{Key: key, Val: value}
			if this.Head == nil {
				this.Head = node
				this.Tail = node
			} else {
				node.Before = this.Tail
				this.Tail.Next = node
				this.Tail = node
			}
			this.Keys[key] = node
		} else {
			node := &LRUNode{Key: key, Val: value}
			this.Keys[key] = node
			// Delete key from LRU Cache
			lru := this.Tail.Key
			delete(this.Keys, lru)
			// Deleete Tail node
			this.Tail.Before.Next = nil
			this.Tail = this.Tail.Before
			// Add new node before Head
			node.Next = this.Head
			this.Head.Before = node
			this.Head = node
		}
	} else {
		this.Keys[key].Val = value
	}
}
