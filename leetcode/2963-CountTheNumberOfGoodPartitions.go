package main

// 2963. Count the Number of Good Partitions
// You are given a 0-indexed array nums consisting of positive integers.

// A partition of an array into one or more contiguous subarrays is called good if no two subarrays contain the same number.

// Return the total number of good partitions of nums.

// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 8
// Explanation: The 8 possible good partitions are: ([1], [2], [3], [4]), ([1], [2], [3,4]), ([1], [2,3], [4]), ([1], [2,3,4]), ([1,2], [3], [4]), ([1,2], [3,4]), ([1,2,3], [4]), and ([1,2,3,4]).

// Example 2:
// Input: nums = [1,1,1,1]
// Output: 1
// Explanation: The only possible good partition is: ([1,1,1,1]).

// Example 3:
// Input: nums = [1,2,1,3]
// Output: 2
// Explanation: The 2 possible good partitions are: ([1,2,1], [3]) and ([1,2,1,3]).

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func numberOfGoodPartitions(nums []int) int {
    res, mod := 1, 1_000_000_007
    max := func (x, y int) int { if x > y { return x; }; return y; }
    subArrayCount := func (arr []int) int{
        mp := make(map[int]int)
        for i, v := range arr{
            mp[v] = i
        }
        res := 0
        for i := 0; i < len(arr); res++ {
            last := mp[arr[i]]
            for j := i + 1; j < last; j++ {
                last = max(mp[arr[j]], last)
            }
            i = last + 1
        }
        return res
    }
    for i := subArrayCount(nums); i > 1; i-- {
        res = (res * 2) % mod
    }
    return res
}

func numberOfGoodPartitions1(nums []int) int {
    mp := make(map[int]int)
    for i, v := range nums {
        mp[v] = i
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, mx := 1, 0
    for i, v := range nums[:len(nums) - 1] { // 少统计最后一段区间
        mx = max(mx, mp[v])
        if mx == i { // 区间无法延长
            res = res * 2 % 1_000_000_007
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 8
    // Explanation: The 8 possible good partitions are: ([1], [2], [3], [4]), ([1], [2], [3,4]), ([1], [2,3], [4]), ([1], [2,3,4]), ([1,2], [3], [4]), ([1,2], [3,4]), ([1,2,3], [4]), and ([1,2,3,4]).
    fmt.Println(numberOfGoodPartitions([]int{1,2,3,4})) // 8
    // Example 2:
    // Input: nums = [1,1,1,1]
    // Output: 1
    // Explanation: The only possible good partition is: ([1,1,1,1]).
    fmt.Println(numberOfGoodPartitions([]int{1,1,1,1})) // 1
    // Example 3:
    // Input: nums = [1,2,1,3]
    // Output: 2
    // Explanation: The 2 possible good partitions are: ([1,2,1], [3]) and ([1,2,1,3]).
    fmt.Println(numberOfGoodPartitions([]int{1,2,1,3})) // 2

    fmt.Println(numberOfGoodPartitions([]int{1,2,3,4,5,6,7,8,9})) // 256
    fmt.Println(numberOfGoodPartitions([]int{9,8,7,6,5,4,3,2,1})) // 256

    fmt.Println(numberOfGoodPartitions1([]int{1,2,3,4})) // 8
    fmt.Println(numberOfGoodPartitions1([]int{1,1,1,1})) // 1
    fmt.Println(numberOfGoodPartitions1([]int{1,2,1,3})) // 2
    fmt.Println(numberOfGoodPartitions1([]int{1,2,3,4,5,6,7,8,9})) // 256
    fmt.Println(numberOfGoodPartitions1([]int{9,8,7,6,5,4,3,2,1})) // 256
}