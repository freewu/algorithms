package main

// 3659. Partition Array Into K-Distinct Groups
// You are given an integer array nums and an integer k.

// Your task is to determine whether it is possible to partition all elements of nums into one or more groups such that:
//     1. Each group contains exactly k distinct elements.
//     2. Each element in nums must be assigned to exactly one group.

// Return true if such a partition is possible, otherwise return false.

// Example 1:
// Input: nums = [1,2,3,4], k = 2
// Output: true
// Explanation:
// One possible partition is to have 2 groups:
// Group 1: [1, 2]
// Group 2: [3, 4]
// Each group contains k = 2 distinct elements, and all elements are used exactly once.

// Example 2:
// Input: nums = [3,5,2,2], k = 2
// Output: true
// Explanation:
// One possible partition is to have 2 groups:
// Group 1: [2, 3]
// Group 2: [2, 5]
// Each group contains k = 2 distinct elements, and all elements are used exactly once.

// Example 3:
// Input: nums = [1,5,2,3], k = 3
// Output: false
// Explanation:
// We cannot form groups of k = 3 distinct elements using all values exactly once.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     ​​​​​​​1 <= k <= nums.length

import "fmt"
import "slices"

// map
func partitionArray(nums []int, k int) bool {
    mx, n := 0, len(nums)
    if n % k > 0 {
        return false
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v]++
        mx = max(mx, mp[v])
    }
    return mx <= n / k
}

// array
func partitionArray1(nums []int, k int) bool {
    mx, n := 0, len(nums)
    if n % k > 0 {
        return false
    }
    count := make([]int, slices.Max(nums) + 1)
    for _, v := range nums {
        count[v]++
        mx = max(mx, count[v])
    }
    return mx <= n / k
}

func partitionArray2(nums []int, k int) bool {
    n := len(nums)
    if n % k > 0 {
        return false
    }
    count := make([]int, slices.Max(nums)+1)
    for _, v := range nums {
        count[v]++
        if count[v] > n / k { // 立即返回
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4], k = 2
    // Output: true
    // Explanation:
    // One possible partition is to have 2 groups:
    // Group 1: [1, 2]
    // Group 2: [3, 4]
    // Each group contains k = 2 distinct elements, and all elements are used exactly once.
    fmt.Println(partitionArray([]int{1,2,3,4}, 2)) // true
    // Example 2:
    // Input: nums = [3,5,2,2], k = 2
    // Output: true
    // Explanation:
    // One possible partition is to have 2 groups:
    // Group 1: [2, 3]
    // Group 2: [2, 5]
    // Each group contains k = 2 distinct elements, and all elements are used exactly once.
    fmt.Println(partitionArray([]int{3,5,2,2}, 2)) // true
    // Example 3:
    // Input: nums = [1,5,2,3], k = 3
    // Output: false
    // Explanation:
    // We cannot form groups of k = 3 distinct elements using all values exactly once.
    fmt.Println(partitionArray([]int{1,5,2,3}, 3)) // false

    fmt.Println(partitionArray([]int{1,2,3,4,5,6,7,8,9}, 3)) // true
    fmt.Println(partitionArray([]int{9,8,7,6,5,4,3,2,1}, 3)) // true

    fmt.Println(partitionArray1([]int{1,2,3,4}, 2)) // true
    fmt.Println(partitionArray1([]int{3,5,2,2}, 2)) // true
    fmt.Println(partitionArray1([]int{1,5,2,3}, 3)) // false
    fmt.Println(partitionArray1([]int{1,2,3,4,5,6,7,8,9}, 3)) // true
    fmt.Println(partitionArray1([]int{9,8,7,6,5,4,3,2,1}, 3)) // true

    fmt.Println(partitionArray2([]int{1,2,3,4}, 2)) // true
    fmt.Println(partitionArray2([]int{3,5,2,2}, 2)) // true
    fmt.Println(partitionArray2([]int{1,5,2,3}, 3)) // false
    fmt.Println(partitionArray2([]int{1,2,3,4,5,6,7,8,9}, 3)) // true
    fmt.Println(partitionArray2([]int{9,8,7,6,5,4,3,2,1}, 3)) // true
}