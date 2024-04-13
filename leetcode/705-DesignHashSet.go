package main

// 705. Design HashSet
// Design a HashSet without using any built-in hash table libraries.
// Implement MyHashSet class:
//     void add(key) Inserts the value key into the HashSet.
//     bool contains(key) Returns whether the value key exists in the HashSet or not.
//     void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.
    
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
 
// Constraints:
//     0 <= key <= 10^6
//     At most 10^4 calls will be made to add, remove, and contains.

import "fmt"

// // 使用数组来处理
// type MyHashSet struct {
//     data []int
// }

// func Constructor() MyHashSet {
//     m := MyHashSet{}
//     m.data = make([]int, 1000001) // 0 <= key <= 10^6
//     return m
// }

// func (this *MyHashSet) Add(key int)  {
//     this.data[key] = 1
// }

// func (this *MyHashSet) Remove(key int)  {
//     this.data[key] = 0
// }

// func (this *MyHashSet) Contains(key int) bool {
//     if this.data[key] == 1 {
//         return true
//     }
//     return false
// }

// 使用一个二维数组来存
type MyHashSet struct {
    set [][]int
}

func Constructor() MyHashSet {
    return MyHashSet{
        set: make([][]int, 1000),
    }
}


func (this *MyHashSet) Add(key int)  {
    h := key % 1000
    if len(this.set[h]) == 0 {
        this.set[h] = make([]int, 0, 10)
        this.set[h] = append(this.set[h], key)
        return
    }
    for _, v := range this.set[h] {
        if v == key {
            return
        }
    }
    this.set[h] = append(this.set[h], key)
}


func (this *MyHashSet) Remove(key int)  {
    h := key % 1000
    if len(this.set[h]) == 0 {
        return
    }
    for i, v := range this.set[h] {
        if v == key {
            this.set[h][i], this.set[h][len(this.set[h])-1] = this.set[h][len(this.set[h])-1], this.set[h][i]
            this.set[h] = this.set[h][:len(this.set[h])-1]
            return
        }
    }
}

func (this *MyHashSet) Contains(key int) bool {
    h := key % 1000
    if len(this.set[h]) == 0 {
        return false
    }
    for _, v := range this.set[h] {
        if v == key {
            return true
        }
    }
    return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {
    // MyHashSet myHashSet = new MyHashSet();
    obj := Constructor()
    // myHashSet.add(1);      // set = [1]
    obj.Add(1)
    //fmt.Println(obj)
    // myHashSet.add(2);      // set = [1, 2]
    obj.Add(2)
    //fmt.Println(obj)
    // myHashSet.contains(1); // return True
    fmt.Println(obj.Contains(1)) // true
    // myHashSet.contains(3); // return False, (not found)
    fmt.Println(obj.Contains(3)) // false
    // myHashSet.add(2);      // set = [1, 2]
    obj.Add(2)
    //fmt.Println(obj)
    // myHashSet.contains(2); // return True
    fmt.Println(obj.Contains(2)) // true
    // myHashSet.remove(2);   // set = [1]
    obj.Remove(2)
    //fmt.Println(obj)
    // myHashSet.contains(2); // return False, (already removed)
    fmt.Println(obj.Contains(2)) // false
}