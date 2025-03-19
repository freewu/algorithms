package main

// 1799. Maximize Score After N Operations
// You are given nums, an array of positive integers of size 2 * n. You must perform n operations on this array.
// In the ith operation (1-indexed), you will:
//     Choose two elements, x and y.
//     Receive a score of i * gcd(x, y).
//     Remove x and y from nums.

// Return the maximum score you can receive after performing n operations.
// The function gcd(x, y) is the greatest common divisor of x and y.

// Example 1:
// Input: nums = [1,2]
// Output: 1
// Explanation: The optimal choice of operations is:
// (1 * gcd(1, 2)) = 1

// Example 2:
// Input: nums = [3,4,6,8]
// Output: 11
// Explanation: The optimal choice of operations is:
// (1 * gcd(3, 6)) + (2 * gcd(4, 8)) = 3 + 8 = 11

// Example 3:
// Input: nums = [1,2,3,4,5,6]
// Output: 14
// Explanation: The optimal choice of operations is:
// (1 * gcd(1, 5)) + (2 * gcd(2, 4)) + (3 * gcd(3, 6)) = 1 + 4 + 9 = 14
 
// Constraints:
//     1 <= n <= 7
//     nums.length == 2 * n
//     1 <= nums[i] <= 10^6

import "fmt"

func maxScore(nums []int) int {
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, q, dp := 0, [][2]int{ {1, 0}}, make([]int, 1 << len(nums))
    for len(q) > 0 {
        op, mask := q[0][0], q[0][1]
        q = q[1:]
        if op-1 == len(nums) / 2 {
            res = max(res, dp[mask])
            continue
        }
        for i := len(nums) - 1; i >= 0; i-- {
            for j := i + 1; j < len(nums); j++ {
                newMask := mask | (1 << i) | (1 << j)
                if mask&(1<<i) != 0 || mask&(1<<j) != 0 {
                    continue
                }
                if dp[newMask] == 0 {
                    q = append(q, [2]int{op + 1, newMask})
                }
                dp[newMask] = max(dp[newMask], op * gcd(nums[i], nums[j]) + dp[mask])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2]
    // Output: 1
    // Explanation: The optimal choice of operations is:
    // (1 * gcd(1, 2)) = 1
    fmt.Println(maxScore([]int{1,2})) // 1
    // Example 2:
    // Input: nums = [3,4,6,8]
    // Output: 11
    // Explanation: The optimal choice of operations is:
    // (1 * gcd(3, 6)) + (2 * gcd(4, 8)) = 3 + 8 = 11
    fmt.Println(maxScore([]int{3,4,6,8})) // 11
    // Example 3:
    // Input: nums = [1,2,3,4,5,6]
    // Output: 14
    // Explanation: The optimal choice of operations is:
    // (1 * gcd(1, 5)) + (2 * gcd(2, 4)) + (3 * gcd(3, 6)) = 1 + 4 + 9 = 14
    fmt.Println(maxScore([]int{1,2,3,4,5,6})) // 14

    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9})) // 30
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1})) // 30
}