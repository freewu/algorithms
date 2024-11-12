package main

// 1848. Minimum Distance to the Target Element
// Given an integer array nums (0-indexed) and two integers target and start, 
// find an index i such that nums[i] == target and abs(i - start) is minimized. 
// Note that abs(x) is the absolute value of x.

// Return abs(i - start).

// It is guaranteed that target exists in nums.

// Example 1:
// Input: nums = [1,2,3,4,5], target = 5, start = 3
// Output: 1
// Explanation: nums[4] = 5 is the only value equal to target, so the answer is abs(4 - 3) = 1.

// Example 2:
// Input: nums = [1], target = 1, start = 0
// Output: 0
// Explanation: nums[0] = 1 is the only value equal to target, so the answer is abs(0 - 0) = 0.

// Example 3:
// Input: nums = [1,1,1,1,1,1,1,1,1,1], target = 1, start = 0
// Output: 0
// Explanation: Every value of nums is 1, but nums[0] minimizes abs(i - start), which is abs(0 - 0) = 0.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^4
//     0 <= start < nums.length
//     target is in nums.

import "fmt"

func getMinDistance(nums []int, target int, start int) int {
    res := 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            v := abs(i - start)
            if v < res {
                res = v
            }
        }
    }
    return res
}

func getMinDistance1(nums []int, target int, start int) int {
    res := 1 << 31
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, v := range nums {
        if v == target {
            res = min(res, abs(i - start))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], target = 5, start = 3
    // Output: 1
    // Explanation: nums[4] = 5 is the only value equal to target, so the answer is abs(4 - 3) = 1.
    fmt.Println(getMinDistance([]int{1,2,3,4,5}, 5, 3)) // 1
    // Example 2:
    // Input: nums = [1], target = 1, start = 0
    // Output: 0
    // Explanation: nums[0] = 1 is the only value equal to target, so the answer is abs(0 - 0) = 0.
    fmt.Println(getMinDistance([]int{1}, 1, 0)) // 0
    // Example 3:
    // Input: nums = [1,1,1,1,1,1,1,1,1,1], target = 1, start = 0
    // Output: 0
    // Explanation: Every value of nums is 1, but nums[0] minimizes abs(i - start), which is abs(0 - 0) = 0.
    fmt.Println(getMinDistance([]int{1,1,1,1,1,1,1,1,1,1}, 1, 0)) // 0

    fmt.Println(getMinDistance1([]int{1,2,3,4,5}, 5, 3)) // 1
    fmt.Println(getMinDistance1([]int{1}, 1, 0)) // 0
    fmt.Println(getMinDistance1([]int{1,1,1,1,1,1,1,1,1,1}, 1, 0)) // 0
}