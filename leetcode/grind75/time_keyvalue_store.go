package grind75

// Time Based Key-Value Store

// Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

// Implement the TimeMap class:

// TimeMap() Initializes the object of the data structure.
// void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
// String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".

// Example 1:

// Input
// ["TimeMap", "set", "get", "get", "set", "get", "get"]
// [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
// Output
// [null, null, "bar", "bar", null, "bar2", "bar2"]

// Explanation
// TimeMap timeMap = new TimeMap();
// timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
// timeMap.get("foo", 1);         // return "bar"
// timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
// timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
// timeMap.get("foo", 4);         // return "bar2"
// timeMap.get("foo", 5);         // return "bar2"

type timeValue struct {
	Time  int
	Value string
}

type TimeMap struct {
	KeyStore map[string][]timeValue
}

func Constructor() TimeMap {
	return TimeMap{
		KeyStore: map[string][]timeValue{},
	}
}

func (this *TimeMap) Set(key, value string, timestamp int) {
	if _, ok := this.KeyStore[key]; !ok {
		this.KeyStore[key] = []timeValue{{Time: timestamp, Value: value}}
		return
	}
	index := this.binarySearch(key, timestamp)
	newValue := timeValue{Time: timestamp, Value: value}

	if index <= len(this.KeyStore[key])-1 && this.KeyStore[key][index].Time == timestamp {
		this.KeyStore[key][index] = newValue
		return
	}

	values := this.KeyStore[key]
	values = append(values[:index], append([]timeValue{newValue}, values[index:]...)...)
	this.KeyStore[key] = values
}

func (this TimeMap) binarySearch(key string, timestamp int) int {
	left, rigth := 0, len(this.KeyStore[key])-1
	for left <= rigth {
		mid := left + (rigth-left)/2
		time := this.KeyStore[key][mid].Time
		if time == timestamp {
			return mid
		}
		if time > timestamp {
			rigth = mid - 1
		} else {
			left = mid + 1
		}
	}
	if rigth < 0 {
		return -1
	}
	return left
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.KeyStore) == 0 {
		return ""
	}
	if _, found := this.KeyStore[key]; !found {
		return ""
	}
	index := this.binarySearch(key, timestamp)
	if index == -1 {
		return ""
	}
	if index == len(this.KeyStore[key]) {
		return this.KeyStore[key][index-1].Value
	}
	if this.KeyStore[key][index].Time != timestamp {
		return this.KeyStore[key][index-1].Value
	}
	return this.KeyStore[key][index].Value
}
