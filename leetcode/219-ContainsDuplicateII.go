package main

// 219. Contains Duplicate II
// Given an integer array nums and an integer k, 
// return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

// Example 1:
// Input: nums = [1,2,3,1], k = 3
// Output: true

// Example 2:
// Input: nums = [1,0,1,1], k = 1
// Output: true

// Example 3:
// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9
//     0 <= k <= 10^5

// 如果数组里面有重复数字，并且重复数字的下标差值小于等于 K 就输出 true，
// 如果没有重复数字或者下标差值超过了 k，则输出 flase。

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
    if len(nums) <= 1 {
        return false
    }
    if k <= 0 {
        return false
    }
    record := make(map[int]bool, len(nums))
    for i, n := range nums {
        if _, found := record[n]; found {
            return true
        }
        record[n] = true
        // 如果存在重复 且 下标差值 <= k
        // 满足 nums[i] == nums[j] 且 abs(i - j) <= k
        if len(record) == k + 1 { // 如果 下标差值超过了 k
            //fmt.Printf("i = %v,i - k = %v,nums[i-k] = %v\n",i,i - k, nums[i-k])
            delete(record, nums[i-k]) // map 的长度如果超过了 K 以后就删除掉 i-k 的那个元素，这样一直维护 map 里面只有 K 个元素
        }
    }
    return false
}

// best solution
func containsNearbyDuplicate1(nums []int, k int) bool {
    mp := make(map[int]int, len(nums))
    for i, v := range nums {
        // 如果存在重复 且 下标差值 <= k
        // 满足 nums[i] == nums[j] 且 abs(i - j) <= k
        if index, ok := mp[v]; ok && i - index <= k {
            return true
        }
        mp[v] = i
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,1], k = 3
    // Output: true
    fmt.Printf("containsNearbyDuplicate([]int{ 1,2,3,1 },3) = %v\n",containsNearbyDuplicate([]int{ 1,2,3,1 },3)) // true
    // Example 2:
    // Input: nums = [1,0,1,1], k = 1
    // Output: true
    fmt.Printf("containsNearbyDuplicate([]int{ 1,0,1,1},1) = %v\n",containsNearbyDuplicate([]int{ 1,0,1,1},1)) // true
    // Example 3:
    // Input: nums = [1,2,3,1,2,3], k = 2
    // Output: false
    fmt.Printf("containsNearbyDuplicate([]int{ 1,2,3,1,2,3},2) = %v\n",containsNearbyDuplicate([]int{ 1,2,3,1,2,3},2)) // false

    fmt.Printf("containsNearbyDuplicate1([]int{ 1,2,3,1 },3) = %v\n",containsNearbyDuplicate1([]int{ 1,2,3,1 },3)) // true
    fmt.Printf("containsNearbyDuplicate1([]int{ 1,0,1,1},1) = %v\n",containsNearbyDuplicate1([]int{ 1,0,1,1},1)) // true
    fmt.Printf("containsNearbyDuplicate1([]int{ 1,2,3,1,2,3},2) = %v\n",containsNearbyDuplicate1([]int{ 1,2,3,1,2,3},2)) // false
}
