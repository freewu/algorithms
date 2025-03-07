package main

// 1. Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:
//     2 <= nums.length <= 10^4
//     -10^9 <= nums[i] <= 10^9
//     -10^9 <= target <= 10^9
//     Only one valid answer exists.

// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

// # 解题思路
// 最优的做法时间复杂度是 O(n)。
// 1 声明一个 map[值] = 下标
// 2 顺序扫描数组,
// 3 计算出 目标值 - 当前值 的 差值
// 4 查找 差值 是否能在 map 里找到
// 5 能找到 返回 [ 当前下标 , map[差值] ]
// 6 没找到  map[当前值] = 当前下票
// 7 回到第3步

import "fmt"

// O(n^2) 自己的解法
func twoSum(nums []int, target int) []int {
    n := len(nums)
    if n < 2 { return nil }
    // 如果数据只有两个值的情况
    if 2 == n && ((nums[0] + nums[1]) != target) { return nil }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            // 冒泡的方式遍历
            if (nums[i] + nums[j]) == target {
                return []int{ i, j }
            }
        }
    }
    return nil
}

// best speed solution O(n)
func twoSum1(nums []int, target int) []int {
    mp := make(map[int]int, len(nums)) // 优先设定好map长度,避免扩容产生性能波动
    for i, v := range nums {
        if index, ok := mp[target - v]; ok { // 少个中间变量
            return []int{ index, i }
        }
        mp[v] = i
    }
    return []int{}
}

func twoSum2(nums []int, target int) []int {
    mp := make(map[int]int) // 使用一个map 来存 map[值]= 位置
    for i := 0; i < len(nums); i++ {
        key := target - nums[i] // 得到 当前数的差值
        if _, ok := mp[key]; ok { // 差值存在
            return []int{ mp[key], i}
        }
        mp[nums[i]] = i //
    }
    return nil
}

func main() {  
    // Example 1:
    // Input: nums = [2,7,11,15], target = 9
    // Output: [0,1]
    // Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
    fmt.Printf("twoSum([]int{2,7,11,15}, 9) = %v\n",twoSum([]int{2,7,11,15}, 9)) // [0,1]
    // Example 2:
    // Input: nums = [3,2,4], target = 6
    // Output: [1,2]
    fmt.Printf("twoSum([]int{3,2,4}, 6) = %v\n",twoSum([]int{3,2,4}, 6)) // [1,2]
    // Example 3:
    // Input: nums = [3,3], target = 6
    // Output: [0,1]
    fmt.Printf("twoSum([]int{3,3}, 6) = %v\n",twoSum([]int{3,3}, 6)) // [0,1]

    fmt.Printf("twoSum([]int{1,2,3,4,5,6,7,8,9}, 6) = %v\n",twoSum([]int{1,2,3,4,5,6,7,8,9}, 6)) // [0 4]
    fmt.Printf("twoSum([]int{9,8,7,6,5,4,3,2,1}, 6) = %v\n",twoSum([]int{9,8,7,6,5,4,3,2,1}, 6)) // [4 8]

    fmt.Printf("twoSum1([]int{2,7,11,15}, 9) = %v\n",twoSum1([]int{2,7,11,15}, 9)) // [0,1]
    fmt.Printf("twoSum1([]int{3,2,4}, 6) = %v\n",twoSum1([]int{3,2,4}, 6)) // [1,2]
    fmt.Printf("twoSum1([]int{3,3}, 6) = %v\n",twoSum1([]int{3,3}, 6)) // [0,1]
    fmt.Printf("twoSum1([]int{1,2,3,4,5,6,7,8,9}, 6) = %v\n",twoSum1([]int{1,2,3,4,5,6,7,8,9}, 6)) // [1 3]
    fmt.Printf("twoSum1([]int{9,8,7,6,5,4,3,2,1}, 6) = %v\n",twoSum1([]int{9,8,7,6,5,4,3,2,1}, 6)) // [5 7]

    fmt.Printf("twoSum2([]int{2,7,11,15}, 9) = %v\n",twoSum2([]int{2,7,11,15}, 9)) // [0,1]
    fmt.Printf("twoSum2([]int{3,2,4}, 6) = %v\n",twoSum2([]int{3,2,4}, 6)) // [1,2]
    fmt.Printf("twoSum2([]int{3,3}, 6) = %v\n",twoSum2([]int{3,3}, 6)) // [0,1]
    fmt.Printf("twoSum2([]int{1,2,3,4,5,6,7,8,9}, 6) = %v\n",twoSum2([]int{1,2,3,4,5,6,7,8,9}, 6)) // [1 3]
    fmt.Printf("twoSum2([]int{9,8,7,6,5,4,3,2,1}, 6) = %v\n",twoSum2([]int{9,8,7,6,5,4,3,2,1}, 6)) // [5 7]
}