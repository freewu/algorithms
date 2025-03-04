package main 

// 3316. Find Maximum Removals From Source String
// You are given a string source of size n, a string pattern that is a subsequence of source, 
// and a sorted integer array targetIndices that contains distinct numbers in the range [0, n - 1].

// We define an operation as removing a character at an index idx from source such that:
//     idx is an element of targetIndices.
//     pattern remains a subsequence of source after removing the character.

// Performing an operation does not change the indices of the other characters in source. 
// For example, if you remove 'c' from "acb", the character at index 2 would still be 'b'.

// Return the maximum number of operations that can be performed.

// Example 1:
// Input: source = "abbaa", pattern = "aba", targetIndices = [0,1,2]
// Output: 1
// Explanation:
// We can't remove source[0] but we can do either of these two operations:
// Remove source[1], so that source becomes "a_baa".
// Remove source[2], so that source becomes "ab_aa".

// Example 2:
// Input: source = "bcda", pattern = "d", targetIndices = [0,3]
// Output: 2
// Explanation:
// We can remove source[0] and source[3] in two operations.

// Example 3:
// Input: source = "dda", pattern = "dda", targetIndices = [0,1,2]
// Output: 0
// Explanation:
// We can't remove any character from source.

// Example 4:
// Input: source = "yeyeykyded", pattern = "yeyyd", targetIndices = [0,2,3,4]
// Output: 2
// Explanation:
// We can remove source[2] and source[3] in two operations.

// Constraints:
//     1 <= n == source.length <= 3 * 10^3
//     1 <= pattern.length <= n
//     1 <= targetIndices.length <= n
//     targetIndices is sorted in ascending order.
//     The input is generated such that targetIndices contains distinct elements in the range [0, n - 1].
//     source and pattern consist only of lowercase English letters.
//     The input is generated such that pattern appears as a subsequence in source.

import "fmt"

func maxRemovals(source string, pattern string, targetIndices []int) int {
    n, m := len(source), len(pattern)
    // Convert targetIndices to a set for fast lookup
    target := make(map[int]bool)
    for _, v := range targetIndices {
        target[v] = true
    }
    // Initialize dp array with INT_MIN values
    dp := make([]int, m + 1)
    for i := 0; i < m; i++ {
        dp[i] = -1 << 31 // INT_MIN equivalent
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp[m] = 0
    // Traverse source in reverse
    for i := n - 1; i >= 0; i-- {
        for j := 0; j <= m; j++ {
            // Increment dp[j] if i is in targetIndices (i.e., can be removed)
            if target[i] {
                dp[j]++
            }
            // Update dp[j] if there's a match between source[i] and pattern[j]
            if j < m && source[i] == pattern[j] {
                dp[j] = max(dp[j], dp[j+1])
            }
        }
    }
    return dp[0]
}

func maxRemovals1(source string, pattern string, targetIndices []int) int {
    m, n := len(pattern), len(source)
    target, mem := make([]int, n), make([]int, m + 1)
    for _, i := range targetIndices {
        target[i] = 1
    }
    for i := range mem {
        mem[i] = -1 << 31
    }
    mem[m] = 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        for j := 0; j <= m; j++ {
            mem[j] += target[i]
            if j < m && source[i] == pattern[j] {
                mem[j] = max(mem[j], mem[j + 1])
            }
        }
    }
    return mem[0]
}

func main() {
    // Example 1:
    // Input: source = "abbaa", pattern = "aba", targetIndices = [0,1,2]
    // Output: 1
    // Explanation:
    // We can't remove source[0] but we can do either of these two operations:
    // Remove source[1], so that source becomes "a_baa".
    // Remove source[2], so that source becomes "ab_aa".
    fmt.Println(maxRemovals("abbaa", "aba", []int{0,1,2})) // 1
    // Example 2:
    // Input: source = "bcda", pattern = "d", targetIndices = [0,3]
    // Output: 2
    // Explanation:
    // We can remove source[0] and source[3] in two operations.
    fmt.Println(maxRemovals("bcda", "d", []int{0,3})) // 2
    // Example 3:
    // Input: source = "dda", pattern = "dda", targetIndices = [0,1,2]
    // Output: 0
    // Explanation:
    // We can't remove any character from source.
    fmt.Println(maxRemovals("dda", "dda", []int{0,1,2})) // 0
    // Example 4:
    // Input: source = "yeyeykyded", pattern = "yeyyd", targetIndices = [0,2,3,4]
    // Output: 2
    // Explanation:
    // We can remove source[2] and source[3] in two operations.
    fmt.Println(maxRemovals("yeyeykyded", "yeyyd", []int{0,2,3,4})) // 2

    fmt.Println(maxRemovals("bluefrog", "bluef", []int{0,2,3,4})) // 0

    fmt.Println(maxRemovals1("abbaa", "aba", []int{0,1,2})) // 1
    fmt.Println(maxRemovals1("bcda", "d", []int{0,3})) // 2
    fmt.Println(maxRemovals1("dda", "dda", []int{0,1,2})) // 0
    fmt.Println(maxRemovals1("yeyeykyded", "yeyyd", []int{0,2,3,4})) // 2
    fmt.Println(maxRemovals1("bluefrog", "bluef", []int{0,2,3,4})) // 0
}