package main

// 462. Minimum Moves to Equal Array Elements II
// Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.
// In one move, you can increment or decrement an element of the array by 1.
// Test cases are designed so that the answer will fit in a 32-bit integer.

// Example 1:
// Input: nums = [1,2,3]
// Output: 2
// Explanation:
// Only two moves are needed (remember each move increments or decrements one element):
// [1,2,3]  =>  [2,2,3]  =>  [2,2,2]

// Example 2:
// Input: nums = [1,10,2,9]
// Output: 16
 
// Constraints:
//     n == nums.length
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"
import "sort"

// 想象成在一个坐标轴上，选一个点，使这个点到其他的点之间的距离之和最短
// 使到各个点的距离没有重合
func minMoves2(nums []int) int {
    sort.Ints(nums)
    res, n := 0, len(nums)
    for i := 0; i < n / 2; i++ {
        res += nums[n - i - 1] - nums[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 2
    // Explanation:
    // Only two moves are needed (remember each move increments or decrements one element):
    // [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
    fmt.Println(minMoves2([]int{1,2,3})) // 2
    // Example 2:
    // Input: nums = [1,10,2,9]
    // Output: 16
    fmt.Println(minMoves2([]int{1,10,2,9})) // 16
}