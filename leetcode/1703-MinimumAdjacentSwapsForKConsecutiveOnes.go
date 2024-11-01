package main

// 1703. Minimum Adjacent Swaps for K Consecutive Ones
// You are given an integer array, nums, and an integer k. nums comprises of only 0's and 1's. 
// In one move, you can choose two adjacent indices and swap their values.

// Return the minimum number of moves required so that nums has k consecutive 1's.

// Example 1:
// Input: nums = [1,0,0,1,0,1], k = 2
// Output: 1
// Explanation: In 1 move, nums could be [1,0,0,0,1,1] and have 2 consecutive 1's.

// Example 2:
// Input: nums = [1,0,0,0,0,0,1,1], k = 3
// Output: 5
// Explanation: In 5 moves, the leftmost 1 can be shifted right until nums = [0,0,0,0,0,1,1,1].

// Example 3:
// Input: nums = [1,1,0,1], k = 2
// Output: 0
// Explanation: nums already has 2 consecutive 1's.

// Constraints:
//     1 <= nums.length <= 10^5
//     nums[i] is 0 or 1.
//     1 <= k <= sum(nums)

import "fmt"

func minMoves(nums []int, k int) int {
    gaps := []int{}
    for i, last := 0, -1; i < len(nums); i++ {
        if nums[i] == 1 {
            if last > -1 {
                gaps = append(gaps, i - 1 - last)
            }
            last = i
        }
    }
    lsum, rsum, wlsum, wrsum := 0, 0, 0, 0
    for i := k / 2 - 1; i >= 0; i-- {
        lsum += gaps[i]// lsum = 3 + 0
        wlsum += lsum; // wlsum = 1 * 3 + 2 * 0
    }
    for i := k / 2; i < k - 1; i++ {
        rsum += gaps[i] // rsum = 2 + 5
        wrsum += rsum // wrsum = 2 * 2 + 1 * 5
    }
    res := wlsum + wrsum
    for p, q, r := 0, k / 2, k - 1; r < len(gaps); r++ {
        wlsum += (k / 2) * gaps[q] - lsum
        lsum += gaps[q] - gaps[p]
        rsum += gaps[r] - gaps[q]
        wrsum += rsum-((k-1)/2)* gaps[q]
        res = min(res, wlsum + wrsum)
        p++
        q++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,0,0,1,0,1], k = 2
    // Output: 1
    // Explanation: In 1 move, nums could be [1,0,0,0,1,1] and have 2 consecutive 1's.
    fmt.Println(minMoves([]int{1,0,0,1,0,1}, 2)) // 1
    // Example 2:
    // Input: nums = [1,0,0,0,0,0,1,1], k = 3
    // Output: 5
    // Explanation: In 5 moves, the leftmost 1 can be shifted right until nums = [0,0,0,0,0,1,1,1].
    fmt.Println(minMoves([]int{1,0,0,0,0,0,1,1}, 3)) // 5
    // Example 3:
    // Input: nums = [1,1,0,1], k = 2
    // Output: 0
    // Explanation: nums already has 2 consecutive 1's.
    fmt.Println(minMoves([]int{1,1,0,1}, 2)) // 0
}