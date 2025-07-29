package main

// 2419. Longest Subarray With Maximum Bitwise AND
// You are given an integer array nums of size n.

// Consider a non-empty subarray from nums that has the maximum possible bitwise AND.

//     In other words, let k be the maximum value of the bitwise AND of any subarray of nums. 
//     Then, only subarrays with a bitwise AND equal to k should be considered.

// Return the length of the longest such subarray.

// The bitwise AND of an array is the bitwise AND of all the numbers in it.

// A subarray is a contiguous sequence of elements within an array.

// Example 1:
// Input: nums = [1,2,3,3,2,2]
// Output: 2
// Explanation:
// The maximum possible bitwise AND of a subarray is 3.
// The longest subarray with that value is [3,3], so we return 2.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 1
// Explanation:
// The maximum possible bitwise AND of a subarray is 4.
// The longest subarray with that value is [4], so we return 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"

func longestSubarray(nums []int) int {
    res, count, mx := 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        if v > mx { // 找到更大的
            mx, count, res = v, 1, 1
        } else if v == mx {
            count++
            res = max(res, count)
        } else {
            count = 0
        }
    }
    return res
}

func longestSubarray1(nums []int) int {
    mx := -1
    for i := 0; i < len(nums); i++ { // 找出最大的值
        if mx < nums[i]{
            mx = nums[i]
        }
    }
    res, count := 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == mx {
            count++ // 记录最大值的连续长度
        } else {
            count = 0
        }
        if res < count {
            res = count
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,3,2,2]
    // Output: 2
    // Explanation:
    // The maximum possible bitwise AND of a subarray is 3.
    // The longest subarray with that value is [3,3], so we return 2.
    fmt.Println(longestSubarray([]int{1,2,3,3,2,2})) // 2
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 1
    // Explanation:
    // The maximum possible bitwise AND of a subarray is 4.
    // The longest subarray with that value is [4], so we return 1.
    fmt.Println(longestSubarray([]int{1,2,3,4})) // 1

    fmt.Println(longestSubarray([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(longestSubarray([]int{9,8,7,6,5,4,3,2,1})) // 1

    fmt.Println(longestSubarray1([]int{1,2,3,3,2,2})) // 2
    fmt.Println(longestSubarray1([]int{1,2,3,4})) // 1
    fmt.Println(longestSubarray1([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(longestSubarray1([]int{9,8,7,6,5,4,3,2,1})) // 1
}