package main

// 3350. Adjacent Increasing Subarrays Detection II
// Given an array nums of n integers, 
// your task is to find the maximum value of k for which there exist two adjacent subarrays of length k each, 
// such that both subarrays are strictly increasing. 
// Specifically, check if there are two subarrays of length k starting at indices a and b (a < b), where:
//     1. Both subarrays nums[a..a + k - 1] and nums[b..b + k - 1] are strictly increasing.
//     2. The subarrays must be adjacent, meaning b = a + k.

// Return the maximum possible value of k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [2,5,7,8,9,2,3,4,3,1]
// Output: 3
// Explanation:
// The subarray starting at index 2 is [7, 8, 9], which is strictly increasing.
// The subarray starting at index 5 is [2, 3, 4], which is also strictly increasing.
// These two subarrays are adjacent, and 3 is the maximum possible value of k for which two such adjacent strictly increasing subarrays exist.

// Example 2:
// Input: nums = [1,2,3,4,4,4,4,5,6,7]
// Output: 2
// Explanation:
// The subarray starting at index 0 is [1, 2], which is strictly increasing.
// The subarray starting at index 2 is [3, 4], which is also strictly increasing.
// These two subarrays are adjacent, and 2 is the maximum possible value of k for which two such adjacent strictly increasing subarrays exist.

// Constraints:
//     2 <= nums.length <= 2 * 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func maxIncreasingSubarrays(nums []int) int {
    res, mx, up, n := 0, 0, 1, len(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        if nums[i] > nums[i - 1] {
            up++
        } else {
            mx, up = up, 1
        }
        res = max(res,  max(up / 2, min(mx, up)))
    }
    return res
}

func maxIncreasingSubarrays1(nums []int) int {
    res, pre, count := 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        count++
        if i == len(nums) - 1 || v >= nums[i + 1] {
            res = max(res, max(count / 2, min(pre, count)))
            pre, count = count, 0
        }
    }
    return res
}

func maxIncreasingSubarrays2(nums []int) int {
    res, pre, cur, n := 0, 1, 1, len(nums)
    for i := 1; i < n; i++ {
        for ; i < n && nums[i] > nums[i-1]; i++ {
            cur++
        }
        if cur >= 2 {
            res = max(res, cur / 2)
        }
        res = max(res, min(cur, pre))
        pre = cur
        cur = 1
    }
    return res  
}

func main() {
    // Example 1:
    // Input: nums = [2,5,7,8,9,2,3,4,3,1]
    // Output: 3
    // Explanation:
    // The subarray starting at index 2 is [7, 8, 9], which is strictly increasing.
    // The subarray starting at index 5 is [2, 3, 4], which is also strictly increasing.
    // These two subarrays are adjacent, and 3 is the maximum possible value of k for which two such adjacent strictly increasing subarrays exist.
    fmt.Println(maxIncreasingSubarrays([]int{2,5,7,8,9,2,3,4,3,1})) // 3
    // Example 2:
    // Input: nums = [1,2,3,4,4,4,4,5,6,7]
    // Output: 2
    // Explanation:
    // The subarray starting at index 0 is [1, 2], which is strictly increasing.
    // The subarray starting at index 2 is [3, 4], which is also strictly increasing.
    // These two subarrays are adjacent, and 2 is the maximum possible value of k for which two such adjacent strictly increasing subarrays exist.
    fmt.Println(maxIncreasingSubarrays([]int{1,2,3,4,4,4,4,5,6,7})) // 2

    fmt.Println(maxIncreasingSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(maxIncreasingSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 1

    fmt.Println(maxIncreasingSubarrays1([]int{2,5,7,8,9,2,3,4,3,1})) // 3
    fmt.Println(maxIncreasingSubarrays1([]int{1,2,3,4,4,4,4,5,6,7})) // 2
    fmt.Println(maxIncreasingSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(maxIncreasingSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // 1

    fmt.Println(maxIncreasingSubarrays2([]int{2,5,7,8,9,2,3,4,3,1})) // 3
    fmt.Println(maxIncreasingSubarrays2([]int{1,2,3,4,4,4,4,5,6,7})) // 2
    fmt.Println(maxIncreasingSubarrays2([]int{1,2,3,4,5,6,7,8,9})) // 4
    fmt.Println(maxIncreasingSubarrays2([]int{9,8,7,6,5,4,3,2,1})) // 1
}