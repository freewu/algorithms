package main

// 981. Time Based Key-Value Store
// Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.
// Implement the TimeMap class
//     TimeMap() Initializes the object of the data structure.
//     void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
//     String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".  

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
 
// Constraints:
//     1 <= key.length, value.length <= 100
//     key and value consist of lowercase English letters and digits.
//     1 <= timestamp <= 10^7
//     All the timestamps timestamp of set are strictly increasing.
//     At most 2 * 105 calls will be made to set and get.

import "fmt"
import "sort"

type Value struct {
    value string
    timestamp int
}

type TimeMap struct {
    data map[string][]Value
}

func Constructor() TimeMap {
    return TimeMap{make(map[string][]Value)}
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _,ok := this.data[key]; !ok {
        this.data[key] = make([]Value,0)
    }
    this.data[key] = append(this.data[key],Value{value,timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
    i := sort.Search(len(this.data[key]), func (j int) bool {
        return this.data[key][j].timestamp > timestamp
    })  
    if i == 0 {
        return ""
    }
    return this.data[key][i-1].value
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
    // TimeMap timeMap = new TimeMap();
    obj := Constructor();
    // timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
    obj.Set("foo", "bar", 1);
    fmt.Println(obj)
    // timeMap.get("foo", 1);         // return "bar"
    fmt.Println(obj.Get("foo", 1)) // bar
    // timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
    fmt.Println(obj.Get("foo", 3)) // bar
    // timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
    obj.Set("foo", "bar2", 4);
    fmt.Println(obj)
    // timeMap.get("foo", 4);         // return "bar2"
    fmt.Println(obj.Get("foo", 4)) // bar2
    // timeMap.get("foo", 5);         // return "bar2"
    fmt.Println(obj.Get("foo", 5)) // bar2
    // timeMap.get("foo", 1);         // return "bar"
    fmt.Println(obj.Get("foo", 1)) // bar
}