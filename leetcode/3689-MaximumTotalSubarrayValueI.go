package main

// 3689. Maximum Total Subarray Value I
// You are given an integer array nums of length n and an integer k.

// You need to choose exactly k non-empty subarrays nums[l..r] of nums. 
// Subarrays may overlap, and the exact same subarray (same l and r) can be chosen more than once.

// The value of a subarray nums[l..r] is defined as: max(nums[l..r]) - min(nums[l..r]).

// The total value is the sum of the values of all chosen subarrays.

// Return the maximum possible total value you can achieve.

// Example 1:
// Input: nums = [1,3,2], k = 2
// Output: 4
// Explanation:
// One optimal approach is:
// Choose nums[0..1] = [1, 3]. The maximum is 3 and the minimum is 1, giving a value of 3 - 1 = 2.
// Choose nums[0..2] = [1, 3, 2]. The maximum is still 3 and the minimum is still 1, so the value is also 3 - 1 = 2.
// Adding these gives 2 + 2 = 4.

// Example 2:
// Input: nums = [4,2,5,1], k = 3
// Output: 12
// Explanation:
// One optimal approach is:
// Choose nums[0..3] = [4, 2, 5, 1]. The maximum is 5 and the minimum is 1, giving a value of 5 - 1 = 4.
// Choose nums[0..3] = [4, 2, 5, 1]. The maximum is 5 and the minimum is 1, so the value is also 4.
// Choose nums[2..3] = [5, 1]. The maximum is 5 and the minimum is 1, so the value is again 4.
// Adding these gives 4 + 4 + 4 = 12.

// Constraints:
//     1 <= n == nums.length <= 5 * 10^​​​​​​​4
//     0 <= nums[i] <= 10^9
//     1 <= k <= 10^5

import "fmt"

func maxTotalValue(nums []int, k int) int64 {
    getMaxAndMin := func(nums []int) (int64, int64) {
        inf := int64(1 << 31)
        mx, mn := -inf, inf
        for _, v := range nums {
            mx = max(mx, int64(v))
            mn = min(mn, int64(v))
        }
        return mx, mn
    }
    mx, mn := getMaxAndMin(nums)
    return (mx - mn) * int64(k)
}

func maxTotalValue1(nums []int, k int) int64 {
    mn, mx := 1 << 31, 0
    for _, v := range nums {
        mn, mx = min(mn, v), max(mx,v)
    }
    return int64(mx - mn) * int64(k)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2], k = 2
    // Output: 4
    // Explanation:
    // One optimal approach is:
    // Choose nums[0..1] = [1, 3]. The maximum is 3 and the minimum is 1, giving a value of 3 - 1 = 2.
    // Choose nums[0..2] = [1, 3, 2]. The maximum is still 3 and the minimum is still 1, so the value is also 3 - 1 = 2.
    // Adding these gives 2 + 2 = 4.
    fmt.Println(maxTotalValue([]int{1,3,2}, 2)) // 4
    // Example 2:
    // Input: nums = [4,2,5,1], k = 3
    // Output: 12
    // Explanation:
    // One optimal approach is:
    // Choose nums[0..3] = [4, 2, 5, 1]. The maximum is 5 and the minimum is 1, giving a value of 5 - 1 = 4.
    // Choose nums[0..3] = [4, 2, 5, 1]. The maximum is 5 and the minimum is 1, so the value is also 4.
    // Choose nums[2..3] = [5, 1]. The maximum is 5 and the minimum is 1, so the value is again 4.
    // Adding these gives 4 + 4 + 4 = 12.
    fmt.Println(maxTotalValue([]int{4,2,5,1}, 3)) // 12

    fmt.Println(maxTotalValue([]int{1,2,3,4,5,6,7,8,9}, 2)) // 16
    fmt.Println(maxTotalValue([]int{9,8,7,6,5,4,3,2,1}, 2)) // 16

    fmt.Println(maxTotalValue1([]int{1,3,2}, 2)) // 4
    fmt.Println(maxTotalValue1([]int{4,2,5,1}, 3)) // 12
    fmt.Println(maxTotalValue1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 16
    fmt.Println(maxTotalValue1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 16
}