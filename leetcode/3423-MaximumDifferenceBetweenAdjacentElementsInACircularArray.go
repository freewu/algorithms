package main

// 3423. Maximum Difference Between Adjacent Elements in a Circular Array
// Given a circular array nums, find the maximum absolute difference between adjacent elements.

// Note: In a circular array, the first and last elements are adjacent.

// Example 1:
// Input: nums = [1,2,4]
// Output: 3
// Explanation:
// Because nums is circular, nums[0] and nums[2] are adjacent. 
// They have the maximum absolute difference of |4 - 1| = 3.

// Example 2:
// Input: nums = [-5,-10,-5]
// Output: 5
// Explanation:
// The adjacent elements nums[0] and nums[1] have the maximum absolute difference of |-5 - (-10)| = 5.

// Constraints:
//     2 <= nums.length <= 100
//     -100 <= nums[i] <= 100

import "fmt"

func maxAdjacentDistance(nums []int) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res := abs(nums[0] - nums[len(nums) - 1])
    for i := 0; i < len(nums) - 1; i++ {
        diff := abs(nums[i] - nums[i + 1])
        if diff > res {
            res = diff
        }
    } 
    return res
}

func maxAdjacentDistance1(nums []int) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := abs(nums[0] - nums[len(nums) - 1])
    for i := 1; i < len(nums); i++ {
        res = max(res, abs(nums[i] - nums[i - 1]))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,4]
    // Output: 3
    // Explanation:
    // Because nums is circular, nums[0] and nums[2] are adjacent. 
    // They have the maximum absolute difference of |4 - 1| = 3.
    fmt.Println(maxAdjacentDistance([]int{1,2,4})) // 3
    // Example 2:
    // Input: nums = [-5,-10,-5]
    // Output: 5
    // Explanation:
    // The adjacent elements nums[0] and nums[1] have the maximum absolute difference of |-5 - (-10)| = 5.
    fmt.Println(maxAdjacentDistance([]int{-5,-10,-5})) // 5

    fmt.Println(maxAdjacentDistance([]int{3,2,-5,-3})) // 7
    fmt.Println(maxAdjacentDistance([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(maxAdjacentDistance([]int{9,8,7,6,5,4,3,2,1})) // 8

    fmt.Println(maxAdjacentDistance1([]int{1,2,4})) // 3
    fmt.Println(maxAdjacentDistance1([]int{-5,-10,-5})) // 5
    fmt.Println(maxAdjacentDistance1([]int{3,2,-5,-3})) // 7
    fmt.Println(maxAdjacentDistance1([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(maxAdjacentDistance1([]int{9,8,7,6,5,4,3,2,1})) // 8
}