package main

// 1674. Minimum Moves to Make Array Complementary
// You are given an integer array nums of even length n and an integer limit. 
// In one move, you can replace any integer from nums with another integer between 1 and limit, inclusive.

// The array nums is complementary if for all indices i (0-indexed), nums[i] + nums[n - 1 - i] equals the same number. 
// For example, the array [1,2,3,4] is complementary because for all indices i, nums[i] + nums[n - 1 - i] = 5.

// Return the minimum number of moves required to make nums complementary.

// Example 1:
// Input: nums = [1,2,4,3], limit = 4
// Output: 1
// Explanation: In 1 move, you can change nums to [1,2,2,3] (underlined elements are changed).
// nums[0] + nums[3] = 1 + 3 = 4.
// nums[1] + nums[2] = 2 + 2 = 4.
// nums[2] + nums[1] = 2 + 2 = 4.
// nums[3] + nums[0] = 3 + 1 = 4.
// Therefore, nums[i] + nums[n-1-i] = 4 for every i, so nums is complementary.

// Example 2:
// Input: nums = [1,2,2,1], limit = 2
// Output: 2
// Explanation: In 2 moves, you can change nums to [2,2,2,2]. You cannot change any number to 3 since 3 > limit.

// Example 3:
// Input: nums = [1,2,1,2], limit = 2
// Output: 0
// Explanation: nums is already complementary.

// Constraints:
//     n == nums.length
//     2 <= n <= 10^5
//     1 <= nums[i] <= limit <= 10^5
//     n is even.

import "fmt"

// line sweep
func minMoves(nums []int, limit int) int {
    n := len(nums)
    arr := make([]int, 2 * limit + 2)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // line sweep , moves needed goes from 2->1->0->1->2
    // here the starting value is 0, so minimum value will go to -n in case of 2 moves for all n/2 pairs
    // so initial sum value need to be kept equal to n
    for i := 0; i < n/2; i++ {
        a, b := nums[i], nums[n - i - 1]
        arr[min(a, b) + 1]--
        arr[a + b]--
        arr[a + b + 1]++
        arr[max(a, b) + limit + 1]++
    }
    // if we calculate prefix sum now, then pre[i] tells total number of moves for pair sum i
    // but we want minimum moves, so take minimum
    res, moves := 1 << 31, n
    for i := 2; i <= limit * 2; i++{
        moves += arr[i]
        res = min(res, moves)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,4,3], limit = 4
    // Output: 1
    // Explanation: In 1 move, you can change nums to [1,2,2,3] (underlined elements are changed).
    // nums[0] + nums[3] = 1 + 3 = 4.
    // nums[1] + nums[2] = 2 + 2 = 4.
    // nums[2] + nums[1] = 2 + 2 = 4.
    // nums[3] + nums[0] = 3 + 1 = 4.
    // Therefore, nums[i] + nums[n-1-i] = 4 for every i, so nums is complementary.
    fmt.Println(minMoves([]int{1,2,4,3}, 4)) // 1
    // Example 2:
    // Input: nums = [1,2,2,1], limit = 2
    // Output: 2
    // Explanation: In 2 moves, you can change nums to [2,2,2,2]. You cannot change any number to 3 since 3 > limit.
    fmt.Println(minMoves([]int{1,2,2,1}, 2)) // 2
    // Example 3:
    // Input: nums = [1,2,1,2], limit = 2
    // Output: 0
    // Explanation: nums is already complementary.
    fmt.Println(minMoves([]int{1,2,1,2}, 2)) // 0
}