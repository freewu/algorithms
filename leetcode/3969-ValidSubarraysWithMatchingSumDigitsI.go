package main

// 3969. Valid Subarrays With Matching Sum Digits I
// You are given an integer array nums and an integer digit x.

// A subarray nums[l..r] is considered valid if the sum of its elements satisfies both of the following conditions:
//     1. The first digit of the sum is equal to x.
//     2. The last digit of the sum is equal to x.

// Return the number of valid subarrays.

// Example 1:
// Input: nums = [1,100,1], x = 1
// Output: 4
// Explanation:
// The valid subarrays are:
// nums[0..0]: sum = 1
// nums[0..1]: sum = 1 + 100 = 101
// nums[1..2]: sum = 100 + 1 = 101
// nums[2..2]: sum = 1
// Thus, the answer is 4.

// Example 2:
// Input: nums = [1], x = 2
// Output: 0
// Explanation:
// The only subarray is nums[0..0] with a sum of 1, which does not satisfy the conditions.
// Thus, the answer is 0.

// Constraints:
//     1 <= nums.length <= 1500
//     1 <= nums[i] <= 10^9
//     1 <= x <= 9

import "fmt"

func countValidSubarrays(nums []int, x int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            last, first := sum % 10, sum
            for first >= 10 {
                first /= 10
            }
            if first == x && last == x {
                res++
            }
        }
    }
    return res
}

func countValidSubarrays1(nums []int, x int) int {
    res, n := 0, len(nums)
    for i := 0; i < n; i++ {
        sum := 0
        for j := i; j < n; j++ {
            sum += nums[j]
            v := sum
            if v % 10 != x {
                continue
            }
            for v >= 10 {
                v /= 10
            }
            if v == x {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,100,1], x = 1
    // Output: 4
    // Explanation:
    // The valid subarrays are:
    // nums[0..0]: sum = 1
    // nums[0..1]: sum = 1 + 100 = 101
    // nums[1..2]: sum = 100 + 1 = 101
    // nums[2..2]: sum = 1
    // Thus, the answer is 4.
    fmt.Println(countValidSubarrays([]int{1,100,1}, 1)) // 4
    // Example 2:
    // Input: nums = [1], x = 2
    // Output: 0
    // Explanation:
    // The only subarray is nums[0..0] with a sum of 1, which does not satisfy the conditions.
    // Thus, the answer is 0.
    fmt.Println(countValidSubarrays([]int{1}, 2)) // 0

    fmt.Println(countValidSubarrays([]int{1,2,3,4,5,6,7,8,9}, 2)) // 2
    fmt.Println(countValidSubarrays([]int{9,8,7,6,5,4,3,2,1}, 2)) // 2

    fmt.Println(countValidSubarrays1([]int{1,100,1}, 1)) // 4
    fmt.Println(countValidSubarrays1([]int{1}, 2)) // 0
    fmt.Println(countValidSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 2
    fmt.Println(countValidSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 2
}