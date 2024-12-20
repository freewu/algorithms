package main

// 3284. Sum of Consecutive Subarrays
// We call an array arr of length n consecutive if one of the following holds:
//     arr[i] - arr[i - 1] == 1 for all 1 <= i < n.
//     arr[i] - arr[i - 1] == -1 for all 1 <= i < n.

// The value of an array is the sum of its elements.

// For example, [3, 4, 5] is a consecutive array of value 12 and [9, 8] is another of value 17. 
// While [3, 4, 3] and [8, 6] are not consecutive.

// Given an array of integers nums, return the sum of the values of all consecutive subarrays.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Note that an array of length 1 is also considered consecutive.

// Example 1:
// Input: nums = [1,2,3]
// Output: 20
// Explanation:
// The consecutive subarrays are: [1], [2], [3], [1, 2], [2, 3], [1, 2, 3].
// Sum of their values would be: 1 + 2 + 3 + 3 + 5 + 6 = 20.

// Example 2:
// Input: nums = [1,3,5,7]
// Output: 16
// Explanation:
// The consecutive subarrays are: [1], [3], [5], [7].
// Sum of their values would be: 1 + 3 + 5 + 7 = 16.

// Example 3:
// Input: nums = [7,6,1,2]
// Output: 32
// Explanation:
// The consecutive subarrays are: [7], [6], [1], [2], [7, 6], [1, 2].
// Sum of their values would be: 7 + 6 + 1 + 2 + 13 + 3 = 32.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func getSum(nums []int) int {
    res, s, t, f, g, mod := nums[0], nums[0], nums[0], 1, 1, 1_000_000_007
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 1; i < len(nums); i++ {
        x, y := nums[i-1], nums[i]
        if y - x == 1 {
            f++
            s += f * y
            res = (res + s) % mod
        } else {
            f, s = 1, y
        }
        if y - x == -1 {
            g++
            t += g * y
            res = (res + t) % mod
        } else {
            g, t = 1, y
        }
        if abs(y - x) != 1 {
            res = (res + y) % mod
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 20
    // Explanation:
    // The consecutive subarrays are: [1], [2], [3], [1, 2], [2, 3], [1, 2, 3].
    // Sum of their values would be: 1 + 2 + 3 + 3 + 5 + 6 = 20.
    fmt.Println(getSum([]int{1,2,3})) // 20
    // Example 2:
    // Input: nums = [1,3,5,7]
    // Output: 16
    // Explanation:
    // The consecutive subarrays are: [1], [3], [5], [7].
    // Sum of their values would be: 1 + 3 + 5 + 7 = 16.
    fmt.Println(getSum([]int{1,3,5,7})) // 16
    // Example 3:
    // Input: nums = [7,6,1,2]
    // Output: 32
    // Explanation:
    // The consecutive subarrays are: [7], [6], [1], [2], [7, 6], [1, 2].
    // Sum of their values would be: 7 + 6 + 1 + 2 + 13 + 3 = 32.
    fmt.Println(getSum([]int{7,6,1,2})) // 32
}