package linked_list

/*

LFU Cache

Design and implement a data structure for a Least Frequently Used (LFU) cache.

Implement the LFUCache class:

LFUCache(int capacity) Initializes the object with the capacity of the data structure.
int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
void put(int key, int value) Update the value of the key if present, or inserts the key if not already present. When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item. For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.
To determine the least frequently used key, a use counter is maintained for each key in the cache. The key with the smallest use counter is the least frequently used key.

When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter for a key in the cache is incremented either a get or put operation is called on it.

The functions get and put must each run in O(1) average time complexity.

Example

Input
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

Explanation
// cnt(x) = the use counter for key x
// cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // return 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // return -1 (not found)
lfu.get(3);      // return 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // return -1 (not found)
lfu.get(3);      // return 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // return 4
                 // cache=[4,3], cnt(4)=2, cnt(3)=3
*/

import (
	"math"
)

type cacheValue struct {
	data  int
	count int
}

type node struct {
	value int
	next  *node
}

type keyList struct {
	head *node
	tail *node
}

func (this *keyList) Pop() *node {
	aux := this.head
	this.head = this.head.next
	return aux
}

func (this *keyList) Delete(key int) *node {
	prev, aux := this.head, this.head
	var found bool
	for aux != nil {
		if aux.value == key {
			found = true
			break
		}
		prev = aux
		aux = aux.next
	}
	if !found {
		return nil
	}
	if aux == this.head {
		this.head = aux.next
		if aux == this.tail {
			this.tail = aux.next
		}
		return aux
	}
	if aux == this.tail {
		this.tail = prev
	}
	prev.next = aux.next
	return aux
}

func (this *keyList) Push(key int) {
	n := &node{value: key}
	if this.head == nil {
		this.head, this.tail = n, n
		return
	}
	this.tail.next, this.tail = n, n
}

func (this *keyList) IsEmpty() bool {
	return this.head == nil
}

type keyCounter struct {
	counter map[int]*keyList
}

func (this *keyCounter) increase(key, count int) {
	if _, ok := this.counter[count]; !ok {
		list := &keyList{}
		list.Push(key)
		this.counter[count] = list
	} else {
		this.counter[count].Push(key)
	}
	this.deleteKey(key, count)
}

func (this *keyCounter) deleteKey(key, count int) {
	if count > 1 {
		this.counter[count-1].Delete(key)
		if this.counter[count-1].IsEmpty() {
			delete(this.counter, count-1)
		}
	}
}

func (this *keyCounter) getLFV() int {
	minKey := math.MaxInt
	for key := range this.counter {
		if key < minKey {
			minKey = key
		}
	}
	node := this.counter[minKey].Pop()
	if this.counter[minKey].IsEmpty() {
		delete(this.counter, minKey)
	}
	return node.value
}

type LFUCache struct {
	capacity   int
	isReaches  bool
	cache      map[int]cacheValue
	keyCounter keyCounter
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    map[int]cacheValue{},
		keyCounter: keyCounter{
			counter: map[int]*keyList{},
		},
	}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.keyCounter.increase(key, v.count+1)
		cv := this.cache[key]
		cv.count++
		this.cache[key] = cv
		return this.cache[key].data
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if len(this.cache) == this.capacity {
		if _, ok := this.cache[key]; !ok {
			lfu := this.keyCounter.getLFV()
			delete(this.cache, lfu)
		}
	}
	if v, ok := this.cache[key]; ok {
		this.keyCounter.increase(key, v.count+1)
		cv := this.cache[key]
		cv.data = value
		cv.count++
		this.cache[key] = cv
	} else {
		this.keyCounter.increase(key, 1)
		this.cache[key] = cacheValue{
			data:  value,
			count: 1,
		}
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(keyF
 * obj.Put(key,value);
 */
