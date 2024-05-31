package main

// 956. Tallest Billboard
// You are installing a billboard and want it to have the largest height. 
// The billboard will have two steel supports, one on each side. 
// Each steel support must be an equal height.

// You are given a collection of rods that can be welded together. 
// For example, if you have rods of lengths 1, 2, and 3, you can weld them together to make a support of length 6.

// Return the largest possible height of your billboard installation. 
// If you cannot support the billboard, return 0.

// Example 1:
// Input: rods = [1,2,3,6]
// Output: 6
// Explanation: We have two disjoint subsets {1,2,3} and {6}, which have the same sum = 6.

// Example 2:
// Input: rods = [1,2,3,4,5,6]
// Output: 10
// Explanation: We have two disjoint subsets {2,3,5} and {4,6}, which have the same sum = 10.

// Example 3:
// Input: rods = [1,2]
// Output: 0
// Explanation: The billboard cannot be supported, so we return 0.
 
// Constraints:
//     1 <= rods.length <= 20
//     1 <= rods[i] <= 1000
//     sum(rods[i]) <= 5000

import "fmt"

func tallestBillboard(rods []int) int {
    sum := 0
    for _, v := range rods { // 计算总和
        sum += v
    }
    dp, curr := make([]int, sum + 1), make([]int, sum + 1) 
    for i,_ := range dp { // init dp 
        dp[i] = -1
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp[0] = 0
    for _, v := range rods {
        copy(curr, dp)
        for i := 0; i <= sum - v; i++ {
            if curr[i] < 0 { 
                continue 
            }
            dp[i + v] = max(dp[i + v], curr[i])
            dp[abs(i - v)] = max(dp[abs(i - v)], curr[i] + min(v, i))
        }
    }
    return dp[0]
}

func main() {
    // Example 1:
    // Input: rods = [1,2,3,6]
    // Output: 6
    // Explanation: We have two disjoint subsets {1,2,3} and {6}, which have the same sum = 6.
    fmt.Println(tallestBillboard([]int{1,2,3,6})) // 6
    // Example 2:
    // Input: rods = [1,2,3,4,5,6]
    // Output: 10
    // Explanation: We have two disjoint subsets {2,3,5} and {4,6}, which have the same sum = 10.
    fmt.Println(tallestBillboard([]int{1,2,3,4,5,6})) // 10
    // Example 3:
    // Input: rods = [1,2]
    // Output: 0
    // Explanation: The billboard cannot be supported, so we return 0.
    fmt.Println(tallestBillboard([]int{1,2})) // 0
}