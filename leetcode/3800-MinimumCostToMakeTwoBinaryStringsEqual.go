package main

// 3800. Minimum Cost to Make Two Binary Strings Equal
// You are given two binary strings s and t, both of length n, and three positive integers flipCost, swapCost, and crossCost.

// You are allowed to apply the following operations any number of times (in any order) to the strings s and t:
//     1. Choose any index i and flip s[i] or t[i] (change '0' to '1' or '1' to '0'). 
//        The cost of this operation is flipCost.
//     2. Choose two distinct indices i and j, and swap either s[i] and s[j] or t[i] and t[j]. 
//        The cost of this operation is swapCost.
//     3. Choose an index i and swap s[i] with t[i]. 
//        The cost of this operation is crossCost.

// Return an integer denoting the minimum total cost needed to make the strings s and t equal.

// Example 1:
// Input: s = "01000", t = "10111", flipCost = 10, swapCost = 2, crossCost = 2
// Output: 16
// Explanation:
// We can perform the following operations:
// Swap s[0] and s[1] (swapCost = 2). After this operation, s = "10000" and t = "10111".
// Cross swap s[2] and t[2] (crossCost = 2). After this operation, s = "10100" and t = "10011".
// Swap s[2] and s[3] (swapCost = 2). After this operation, s = "10010" and t = "10011".
// Flip s[4] (flipCost = 10). After this operation, s = t = "10011".
// The total cost is 2 + 2 + 2 + 10 = 16.

// Example 2:
// Input: s = "001", t = "110", flipCost = 2, swapCost = 100, crossCost = 100
// Output: 6
// Explanation:
// Flipping all the bits of s makes the strings equal, and the total cost is 3 * flipCost = 3 * 2 = 6.

// Example 3:
// Input: s = "1010", t = "1010", flipCost = 5, swapCost = 5, crossCost = 5
// Output: 0
// Explanation:
// The strings are already equal, so no operations are required.

// Constraints:
//     n == s.length == t.length
//     1 <= n <= 10^5​​​​​​​
//     1 <= flipCost, swapCost, crossCost <= 10^9
//     s and t consist only of the characters '0' and '1'.

import "fmt"

func minimumCost(s string, t string, flipCost int, swapCost int, crossCost int) int64 {
    count := [2][2]int{}
    for i, v := range s {
        count[v & 1][t[i] & 1]++
    }
    a, b := count[0][1], count[1][0]
    if a > b {
        a, b = b, a
    }
    res1 := (a + b) * flipCost
    res2 := a * swapCost + (b - a)*  flipCost
    avg := (a + b) / 2
    res3 := (avg - a) * crossCost + avg * swapCost + (a + b) % 2 * flipCost
    return int64(min(res1, res2, res3))
}

func main() {
    // Example 1:
    // Input: s = "01000", t = "10111", flipCost = 10, swapCost = 2, crossCost = 2
    // Output: 16
    // Explanation:
    // We can perform the following operations:
    // Swap s[0] and s[1] (swapCost = 2). After this operation, s = "10000" and t = "10111".
    // Cross swap s[2] and t[2] (crossCost = 2). After this operation, s = "10100" and t = "10011".
    // Swap s[2] and s[3] (swapCost = 2). After this operation, s = "10010" and t = "10011".
    // Flip s[4] (flipCost = 10). After this operation, s = t = "10011".
    // The total cost is 2 + 2 + 2 + 10 = 16.
    fmt.Println(minimumCost("01000", "10111", 10, 2, 2)) // 16  
    // Example 2:
    // Input: s = "001", t = "110", flipCost = 2, swapCost = 100, crossCost = 100
    // Output: 6
    // Explanation:
    // Flipping all the bits of s makes the strings equal, and the total cost is 3 * flipCost = 3 * 2 = 6.
    fmt.Println(minimumCost("001", "110", 2, 100, 100)) // 6  
    // Example 3:
    // Input: s = "1010", t = "1010", flipCost = 5, swapCost = 5, crossCost = 5
    // Output: 0
    // Explanation:
    // The strings are already equal, so no operations are required.
    fmt.Println(minimumCost("1010", "1010", 5, 5, 5)) // 0  
}