package main

// 3584. Maximum Product of First and Last Elements of a Subsequence
// You are given an integer array nums and an integer m.

// Return the maximum product of the first and last elements of any subsequence of nums of size m.

// Example 1:
// Input: nums = [-1,-9,2,3,-2,-3,1], m = 1
// Output: 81
// Explanation:
// The subsequence [-9] has the largest product of the first and last elements: -9 * -9 = 81. Therefore, the answer is 81.

// Example 2:
// Input: nums = [1,3,-5,5,6,-4], m = 3
// Output: 20
// Explanation:
// The subsequence [-5, 6, -4] has the largest product of the first and last elements.

// Example 3:
// Input: nums = [2,-1,2,-6,5,2,-5,7], m = 2
// Output: 35
// Explanation:
// The subsequence [5, 7] has the largest product of the first and last elements.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5
//     1 <= m <= nums.length

import "fmt"

func maximumProduct(nums []int, m int) int64 {
    first, last := 0, m - 1
    res, mn, mx := nums[0] * nums[last], nums[0], nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for ;last < len(nums); last, first = last+1, first+1 {
        mx = max(nums[first], mx)
        mn = min(nums[first], mn)
        res = max(res, max(mn * nums[last], mx * nums[last]))
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [-1,-9,2,3,-2,-3,1], m = 1
    // Output: 81
    // Explanation:
    // The subsequence [-9] has the largest product of the first and last elements: -9 * -9 = 81. Therefore, the answer is 81.
    fmt.Println(maximumProduct([]int{-1,-9,2,3,-2,-3,1}, 1)) // 81
    // Example 2:
    // Input: nums = [1,3,-5,5,6,-4], m = 3
    // Output: 20
    // Explanation:
    // The subsequence [-5, 6, -4] has the largest product of the first and last elements.
    fmt.Println(maximumProduct([]int{1,3,-5,5,6,-4}, 3)) // 20
    // Example 3:
    // Input: nums = [2,-1,2,-6,5,2,-5,7], m = 2
    // Output: 35
    // Explanation:
    // The subsequence [5, 7] has the largest product of the first and last elements.
    fmt.Println(maximumProduct([]int{2,-1,2,-6,5,2,-5,7}, 2)) // 35

    fmt.Println(maximumProduct([]int{1,2,3,4,5,6,7,8,9}, 2)) // 72
    fmt.Println(maximumProduct([]int{9,8,7,6,5,4,3,2,1}, 2)) // 72
}