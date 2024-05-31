package main

// 414. Third Maximum Number
// Given an integer array nums, return the third distinct maximum number in this array. 
// If the third maximum does not exist, return the maximum number.

// Example 1:
// Input: nums = [3,2,1]
// Output: 1
// Explanation:
// The first distinct maximum is 3.
// The second distinct maximum is 2.
// The third distinct maximum is 1.

// Example 2:
// Input: nums = [1,2]
// Output: 2
// Explanation:
// The first distinct maximum is 2.
// The second distinct maximum is 1.
// The third distinct maximum does not exist, so the maximum (2) is returned instead.

// Example 3:
// Input: nums = [2,2,3,1]
// Output: 1
// Explanation:
// The first distinct maximum is 3.
// The second distinct maximum is 2 (both 2's are counted together since they have the same value).
// The third distinct maximum is 1.

// Constraints:
//     1 <= nums.length <= 10^4
//     -2^31 <= nums[i] <= 2^31 - 1
 
// Follow up: Can you find an O(n) solution?

import "fmt"

func thirdMax(nums []int) int {
    inf := -1 << 63
    a, b, c := inf, inf, inf
    for _, v := range nums {
        if v > a { // 找了比目前最大的还大的数,依次传
            c, b, a = b, a, v
        } else if v < a && v > b { // 找到目前比第二还大的数
            c, b = b, v
        } else if v < b && v > c { // 第三大的数
            c = v
        }
    }
    if c == inf {
        return a
    }
    return c
}

func main() {
    // Example 1:
    // Input: nums = [3,2,1]
    // Output: 1
    // Explanation:
    // The first distinct maximum is 3.
    // The second distinct maximum is 2.
    // The third distinct maximum is 1.
    fmt.Println(thirdMax([]int{3,2,1})) // 1
    // Example 2:
    // Input: nums = [1,2]
    // Output: 2
    // Explanation:
    // The first distinct maximum is 2.
    // The second distinct maximum is 1.
    // The third distinct maximum does not exist, so the maximum (2) is returned instead.
    fmt.Println(thirdMax([]int{1,2})) // 2
    // Example 3:
    // Input: nums = [2,2,3,1]
    // Output: 1
    // Explanation:
    // The first distinct maximum is 3.
    // The second distinct maximum is 2 (both 2's are counted together since they have the same value).
    // The third distinct maximum is 1.
    fmt.Println(thirdMax([]int{2,2,3,1})) // 1
}