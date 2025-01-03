package main

// 3105. Longest Strictly Increasing or Strictly Decreasing Subarray
// You are given an array of integers nums. 
// Return the length of the longest subarray of nums which is either strictly increasing or strictly decreasing.

// Example 1:
// Input: nums = [1,4,3,3,2]
// Output: 2
// Explanation:
// The strictly increasing subarrays of nums are [1], [2], [3], [3], [4], and [1,4].
// The strictly decreasing subarrays of nums are [1], [2], [3], [3], [4], [3,2], and [4,3].
// Hence, we return 2.

// Example 2:
// Input: nums = [3,3,3,3]
// Output: 1
// Explanation:
// The strictly increasing subarrays of nums are [3], [3], [3], and [3].
// The strictly decreasing subarrays of nums are [3], [3], [3], and [3].
// Hence, we return 1.

// Example 3:
// Input: nums = [3,2,1]
// Output: 3
// Explanation:
// The strictly increasing subarrays of nums are [3], [2], and [1].
// The strictly decreasing subarrays of nums are [3], [2], [1], [3,2], [2,1], and [3,2,1].
// Hence, we return 3.

// Constraints:
//     1 <= nums.length <= 50
//     1 <= nums[i] <= 50

import "fmt"

func longestMonotonicSubarray(nums []int) int {
    n := len(nums)
    if n == 0 || n == 1 { return n }
    res, inc, dec := 1, 1, 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i - 1] {
            inc++
            dec = 1
        } else if nums[i] < nums[i - 1] {
            dec++
            inc = 1
        } else {
            inc, dec = 1, 1
        }
        if inc > res {
            res = inc
        }
        if dec > res {
            res = dec
        }
    }
    return res
}

func longestMonotonicSubarray1(nums []int) int {
    res, i, n := 1, 0, len(nums)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i < n - 1{
        if nums[i+1] == nums[i] {
            i++
            continue
        }
        i0 := i
        inc := nums[i+1] > nums[i]
        i += 2
        for i < n && nums[i] != nums[i-1] && nums[i] > nums[i-1] == inc {
            i++
        }
        res = max(res, i - i0)
        i--
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,3,3,2]
    // Output: 2
    // Explanation:
    // The strictly increasing subarrays of nums are [1], [2], [3], [3], [4], and [1,4].
    // The strictly decreasing subarrays of nums are [1], [2], [3], [3], [4], [3,2], and [4,3].
    // Hence, we return 2.
    fmt.Println(longestMonotonicSubarray([]int{1,4,3,3,2})) // 2
    // Example 2:
    // Input: nums = [3,3,3,3]
    // Output: 1
    // Explanation:
    // The strictly increasing subarrays of nums are [3], [3], [3], and [3].
    // The strictly decreasing subarrays of nums are [3], [3], [3], and [3].
    // Hence, we return 1.
    fmt.Println(longestMonotonicSubarray([]int{3,3,3,3})) // 1
    // Example 3:
    // Input: nums = [3,2,1]
    // Output: 3
    // Explanation:
    // The strictly increasing subarrays of nums are [3], [2], and [1].
    // The strictly decreasing subarrays of nums are [3], [2], [1], [3,2], [2,1], and [3,2,1].
    // Hence, we return 3.
    fmt.Println(longestMonotonicSubarray([]int{3,2,1})) // 3

    fmt.Println(longestMonotonicSubarray([]int{9,8,7,6,5,4,3,2,1})) // 9
    fmt.Println(longestMonotonicSubarray([]int{1,2,3,4,5,6,7,8,9})) // 9

    fmt.Println(longestMonotonicSubarray1([]int{1,4,3,3,2})) // 2
    fmt.Println(longestMonotonicSubarray1([]int{3,3,3,3})) // 1
    fmt.Println(longestMonotonicSubarray1([]int{3,2,1})) // 3
    fmt.Println(longestMonotonicSubarray1([]int{9,8,7,6,5,4,3,2,1})) // 9
    fmt.Println(longestMonotonicSubarray1([]int{1,2,3,4,5,6,7,8,9})) // 9
}