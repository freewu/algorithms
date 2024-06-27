package main

// 677. Map Sum Pairs
// Design a map that allows you to do the following:
//     Maps a string key to a given value.
//     Returns the sum of the values that have a key with a prefix equal to a given string.

// Implement the MapSum class:
//     MapSum() 
//         Initializes the MapSum object.
//     void insert(String key, int val) 
//         Inserts the key-val pair into the map. 
//         If the key already existed, the original key-value pair will be overridden to the new one.
//     int sum(string prefix) 
//         Returns the sum of all the pairs' value whose key starts with the prefix.

// Example 1:
// Input
// ["MapSum", "insert", "sum", "insert", "sum"]
// [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
// Output
// [null, null, 3, null, 5]
// Explanation
// MapSum mapSum = new MapSum();
// mapSum.insert("apple", 3);  
// mapSum.sum("ap");           // return 3 (apple = 3)
// mapSum.insert("app", 2);    
// mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)

// Constraints:
//     1 <= key.length, prefix.length <= 50
//     key and prefix consist of only lowercase English letters.
//     1 <= val <= 1000
//     At most 50 calls will be made to insert and sum.

import "fmt"

type MapSum struct {
    dict map[string]int
    set map[string]int
}

func Constructor() MapSum {
    return MapSum{ make(map[string]int), make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int)  {
    diff := val
    if v, ok := this.set[key]; ok {
        diff -= v
    }
    this.set[key] = val
    for i := 1; i <= len(key); i++ {
        this.dict[key[:i]] += diff
    }
}

func (this *MapSum) Sum(prefix string) int {
    return this.dict[prefix]
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

func main() {
    // MapSum mapSum = new MapSum();
    obj := Constructor()
    // mapSum.insert("apple", 3);  
    obj.Insert("apple", 3)
    fmt.Println(obj)
    // mapSum.sum("ap");           // return 3 (apple = 3)
    fmt.Println(obj.Sum("ap")) // 3
    // mapSum.insert("app", 2);    
    obj.Insert("app", 2)
    fmt.Println(obj)
    // mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)
    fmt.Println(obj.Sum("ap")) // 5
}