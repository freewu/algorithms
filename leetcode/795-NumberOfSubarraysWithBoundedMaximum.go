package main

// 795. Number of Subarrays with Bounded Maximum
// Given an integer array nums and two integers left and right, 
// return the number of contiguous non-empty subarrays such that the value of the maximum array element in that subarray is in the range [left, right].

// The test cases are generated so that the answer will fit in a 32-bit integer.

// Example 1:
// Input: nums = [2,1,4,3], left = 2, right = 3
// Output: 3
// Explanation: There are three subarrays that meet the requirements: [2], [2, 1], [3].

// Example 2:
// Input: nums = [2,9,2,5,6], left = 2, right = 8
// Output: 7

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9
//     0 <= left <= right <= 10^9

import "fmt"

func numSubarrayBoundedMax(nums []int, left int, right int) int {
    res, cur, start, last := 0, 0, 0, -1
    for i := 0; i < len(nums); i++ {
        if nums[i] < left {
            cur += i - start + 1
            cur -= i - last
        } else if nums[i] >= left && nums[i] <= right {
            last = i
            cur += i - start + 1
        } else {
            res += cur
            cur = 0
            start = i + 1
            last = i
        }
    }
    return res + cur
}

func main() {
    // Example 1:
    // Input: nums = [2,1,4,3], left = 2, right = 3
    // Output: 3
    // Explanation: There are three subarrays that meet the requirements: [2], [2, 1], [3].
    fmt.Println(numSubarrayBoundedMax([]int{2,1,4,3},2,3)) // 3
    // Example 2:
    // Input: nums = [2,9,2,5,6], left = 2, right = 8
    // Output: 7
    fmt.Println(numSubarrayBoundedMax([]int{2,9,2,5,6},2,8)) // 7
}