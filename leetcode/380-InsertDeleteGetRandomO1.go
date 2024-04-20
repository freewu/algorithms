package main

// 380. Insert Delete GetRandom O(1)
// Implement the RandomizedSet class:
//     RandomizedSet() Initializes the RandomizedSet object.
//     bool insert(int val) 
//         Inserts an item val into the set if not present. 
//         Returns true if the item was not present, false otherwise.
//     bool remove(int val) 
//         Removes an item val from the set if present. 
//         Returns true if the item was present, false otherwise.
//     int getRandom() 
//         Returns a random element from the current set of elements(it's guaranteed that at least one element exists when this method is called). 
//         Each element must have the same probability of being returned.

// You must implement the functions of the class such that each function works in average O(1) time complexity.

// Example 1:
// Input
// ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
// [[], [1], [2], [2], [], [1], [2], []]
// Output
// [null, true, false, true, 2, true, false, 2]

// Explanation
// RandomizedSet randomizedSet = new RandomizedSet();
// randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
// randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
// randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
// randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
// randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
// randomizedSet.insert(2); // 2 was already in the set, so return false.
// randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
 
// Constraints:
//     -2^31 <= val <= 2^31 - 1
//     At most 2 * 10^5 calls will be made to insert, remove, and getRandom.
//     There will be at least one element in the data structure when getRandom is called.

import "fmt"
import "math/rand"

type RandomizedSet struct {
    m map[int]int
    nums []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        m: make(map[int]int),
        nums: make([]int, 0),
    }
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.m[val];ok {
        return false
    }
    this.m[val] = len(this.nums)
    this.nums = append(this.nums, val)
    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.m[val];ok {
        // 将切片中最后一个元素放到被删除元素的位置
        idx := this.m[val]
        last := this.nums[len(this.nums)-1]
        this.nums[idx]=last
        // 更新最后一个元素的下标为被删除元素的下标
        this.m[last] = idx
        // 移除数组的最后一个元素
        this.nums = this.nums[:len(this.nums)-1]
        // 删除map中被删除元素的下标索引
        delete(this.m, val)
        return true
    }
    return false
}

func (this *RandomizedSet) GetRandom() int {
    r := rand.Intn(len(this.nums))
    return this.nums[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
    // RandomizedSet randomizedSet = new RandomizedSet();
    obj := Constructor()
    fmt.Println(obj)
    // randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
    fmt.Println(obj.Insert(1)) // true
    fmt.Println(obj)
    // randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
    fmt.Println(obj.Remove(2)) // false
    fmt.Println(obj)
    // randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
    fmt.Println(obj.Insert(2)) // true
    fmt.Println(obj)
    // randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
    fmt.Println(obj.GetRandom()) // 1 or 2
    fmt.Println(obj)
    // randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
    fmt.Println(obj.Remove(1)) // true
    fmt.Println(obj)
    // randomizedSet.insert(2); // 2 was already in the set, so return false.
    fmt.Println(obj.Insert(2)) // false
    fmt.Println(obj)
    // randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
    fmt.Println(obj.GetRandom()) // 2
}