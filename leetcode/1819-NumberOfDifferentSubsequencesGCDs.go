package main

// 1819. Number of Different Subsequences GCDs
// You are given an array nums that consists of positive integers.

// The GCD of a sequence of numbers is defined as the greatest integer that divides all the numbers in the sequence evenly.
//     For example, the GCD of the sequence [4,6,16] is 2.

// A subsequence of an array is a sequence that can be formed by removing some elements (possibly none) of the array.
//     For example, [2,5,10] is a subsequence of [1,2,1,2,4,1,5,10].

// Return the number of different GCDs among all non-empty subsequences of nums.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/17/image-1.png"/>
// Input: nums = [6,10,3]
// Output: 5
// Explanation: The figure shows all the non-empty subsequences and their GCDs.
// The different GCDs are 6, 10, 3, 2, and 1.

// Example 2:
// Input: nums = [5,15,40,5,6]
// Output: 7

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 2 * 10^5

import "fmt"

func countDifferentSubsequenceGCDs(nums []int) int {
    dp := make([]int, 200001)
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for _, v := range nums {
        for j := 1; j * j <= v; j++ {
            if v % j == 0 {
                k := v / j
                for _, y := range []int{j, k} {
                    if dp[y] == 0 {
                        dp[y] = v
                    } else {
                        dp[y] = gcd(v, dp[y])
                    }
                }
            }
        }
    }
    res := 0
    for i := 1; i <= 200000; i++ {
        if dp[i] == i {
            res++
        }
    }
    return res
}

func countDifferentSubsequenceGCDs1(nums []int) int {
    res, mx := 0, 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for _, v := range nums { // 找到最大的值
        mx = max(mx, v)
    }
    visited := make(map[int]bool)
    for _, v := range nums {
        visited[v] = true
    }
    for i := 1; i <= mx; i++ {
        g := 0
        for j := i; j <= mx; j += i {
            if visited[j] {
                g = gcd(g, j)
                if g == i {
                    res++
                    break
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/17/image-1.png"/>
    // Input: nums = [6,10,3]
    // Output: 5
    // Explanation: The figure shows all the non-empty subsequences and their GCDs.
    // The different GCDs are 6, 10, 3, 2, and 1.
    fmt.Println(countDifferentSubsequenceGCDs([]int{6,10,3})) // 5
    // Example 2:
    // Input: nums = [5,15,40,5,6]
    // Output: 7
    fmt.Println(countDifferentSubsequenceGCDs([]int{5,15,40,5,6})) // 7

    fmt.Println(countDifferentSubsequenceGCDs1([]int{6,10,3})) // 5
    fmt.Println(countDifferentSubsequenceGCDs1([]int{5,15,40,5,6})) // 7
}