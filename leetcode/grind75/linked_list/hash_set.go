package linked_list

// 705. Design HashSet

// Design a HashSet without using any built-in hash table libraries.

// Implement MyHashSet class:

// void add(key) Inserts the value key into the HashSet.
// bool contains(key) Returns whether the value key exists in the HashSet or not.
// void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.

// Example 1:

// Input
// ["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
// [[], [1], [2], [1], [3], [2], [2], [2], [2]]
// Output
// [null, null, null, true, false, null, true, null, false]

// Explanation
// MyHashSet myHashSet = new MyHashSet();
// myHashSet.add(1);      // set = [1]
// myHashSet.add(2);      // set = [1, 2]
// myHashSet.contains(1); // return True
// myHashSet.contains(3); // return False, (not found)
// myHashSet.add(2);      // set = [1, 2]
// myHashSet.contains(2); // return True
// myHashSet.remove(2);   // set = [1]
// myHashSet.contains(2); // return False, (already removed)

type ListNode struct {
	Key  int
	Next *ListNode
}

type MyHashSet struct {
	BucketSize int
	Set        []*ListNode
}

func NewMyHashSet() MyHashSet {
	return MyHashSet{
		BucketSize: 10000,
		Set:        make([]*ListNode, 10000),
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % this.BucketSize
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	h := this.hash(key)
	this.Set[h] = &ListNode{
		Key:  key,
		Next: this.Set[h],
	}
}

func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	node := this.Set[h]
	if node == nil {
		return
	}
	if node.Key == key {
		this.Set[h] = node.Next
	} else {
		for node.Next != nil {
			if node.Next.Key == key {
				node.Next = node.Next.Next
				return
			}
			node = node.Next
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	node := this.Set[h]
	for node != nil {
		if node.Key == key {
			return true
		}
		node = node.Next
	}
	return false
}
