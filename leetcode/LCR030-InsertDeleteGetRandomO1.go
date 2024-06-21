package main

// LCR 030. O(1) 时间插入、删除和获取随机元素
// 设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构：
//     insert(val)：当元素 val 不存在时返回 true ，并向集合中插入该项，否则返回 false 。
//     remove(val)：当元素 val 存在时返回 true ，并从集合中移除该项，否则返回 false 。
//     getRandom：随机返回现有集合中的一项。每个元素应该有 相同的概率 被返回。

// 示例 :
// 输入: inputs = ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
// [[], [1], [2], [2], [], [1], [2], []]
// 输出: [null, true, false, true, 2, true, false, 2]
// 解释:
// RandomizedSet randomSet = new RandomizedSet();  // 初始化一个空的集合
// randomSet.insert(1); // 向集合中插入 1 ， 返回 true 表示 1 被成功地插入
// randomSet.remove(2); // 返回 false，表示集合中不存在 2 
// randomSet.insert(2); // 向集合中插入 2 返回 true ，集合现在包含 [1,2] 
// randomSet.getRandom(); // getRandom 应随机返回 1 或 2 
// randomSet.remove(1); // 从集合中移除 1 返回 true 。集合现在包含 [2] 
// randomSet.insert(2); // 2 已在集合中，所以返回 false 
// randomSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 

// 提示：
//     -2^31 <= val <= 2^31 - 1
//     最多进行 2 * 10^5 次 insert ， remove 和 getRandom 方法调用
//     当调用 getRandom 方法时，集合中至少有一个元素

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