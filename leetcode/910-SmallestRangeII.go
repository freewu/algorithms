package main

// 910. Smallest Range II
// You are given an integer array nums and an integer k.

// For each index i where 0 <= i < nums.length, change nums[i] to be either nums[i] + k or nums[i] - k.

// The score of nums is the difference between the maximum and minimum elements in nums.

// Return the minimum score of nums after changing the values at each index.

// Example 1:
// Input: nums = [1], k = 0
// Output: 0
// Explanation: The score is max(nums) - min(nums) = 1 - 1 = 0.

// Example 2:
// Input: nums = [0,10], k = 2
// Output: 6
// Explanation: Change nums to be [2, 8]. The score is max(nums) - min(nums) = 8 - 2 = 6.

// Example 3:
// Input: nums = [1,3,6], k = 3
// Output: 3
// Explanation: Change nums to be [4, 6, 3]. The score is max(nums) - min(nums) = 6 - 3 = 3.

// Constraints:
//     1 <= nums.length <= 10^4
//     0 <= nums[i] <= 10^4
//     0 <= k <= 10^4

import "fmt"
import "sort"

func smallestRangeII(nums []int, k int) int {
    if len(nums) <= 1 {
        return 0
    }
    sort.Ints(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := nums[len(nums)-1] - nums[0]
    for i := 0; i < len(nums)-1; i++ {
        l, r := min(nums[0]+k, nums[i+1]-k),  max(nums[i]+k, nums[len(nums)-1]-k)
        res = min(res, r - l)
    }
    return res
}

func smallestRangeII1(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    res := nums[n-1] - nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n-1; i++ {
        mx := max(nums[i]+k, nums[n-1]-k)
        mn := min(nums[0]+k, nums[i+1]-k)
        res = min(res, abs(mx - mn))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1], k = 0
    // Output: 0
    // Explanation: The score is max(nums) - min(nums) = 1 - 1 = 0.
    fmt.Println(smallestRangeII([]int{1}, 0)) // 0
    // Example 2:
    // Input: nums = [0,10], k = 2
    // Output: 6
    // Explanation: Change nums to be [2, 8]. The score is max(nums) - min(nums) = 8 - 2 = 6.
    fmt.Println(smallestRangeII([]int{0,10}, 2)) // 6
    // Example 3:
    // Input: nums = [1,3,6], k = 3
    // Output: 3
    // Explanation: Change nums to be [4, 6, 3]. The score is max(nums) - min(nums) = 6 - 3 = 3.
    fmt.Println(smallestRangeII([]int{1,3,6}, 3)) // 3

    fmt.Println(smallestRangeII1([]int{1}, 0)) // 0
    fmt.Println(smallestRangeII1([]int{0,10}, 2)) // 6
    fmt.Println(smallestRangeII1([]int{1,3,6}, 3)) // 3
}