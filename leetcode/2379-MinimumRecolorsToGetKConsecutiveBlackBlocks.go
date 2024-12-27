package main

// 2379. Minimum Recolors to Get K Consecutive Black Blocks
// You are given a 0-indexed string blocks of length n, where blocks[i] is either 'W' or 'B', representing the color of the ith block. 
// The characters 'W' and 'B' denote the colors white and black, respectively.

// You are also given an integer k, which is the desired number of consecutive black blocks.

// In one operation, you can recolor a white block such that it becomes a black block.

// Return the minimum number of operations needed such that there is at least one occurrence of k consecutive black blocks.

// Example 1:
// Input: blocks = "WBBWWBBWBW", k = 7
// Output: 3
// Explanation:
// One way to achieve 7 consecutive black blocks is to recolor the 0th, 3rd, and 4th blocks
// so that blocks = "BBBBBBBWBW". 
// It can be shown that there is no way to achieve 7 consecutive black blocks in less than 3 operations.
// Therefore, we return 3.

// Example 2:
// Input: blocks = "WBWBBBW", k = 2
// Output: 0
// Explanation:
// No changes need to be made, since 2 consecutive black blocks already exist.
// Therefore, we return 0.

// Constraints:
//     n == blocks.length
//     1 <= n <= 100
//     blocks[i] is either 'W' or 'B'.
//     1 <= k <= n

import "fmt"

func minimumRecolors(blocks string, k int) int {
    res := 0
    for i := 0; i < k; i++ {
        if blocks[i] == 'W' {
            res++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    tmp := res
    for i := 0; i < len(blocks) - k; i++ {
        if blocks[i] != blocks[i + k] {
            if blocks[i] == 'W' {
                tmp--
            } else {
                tmp++
            }
        }
        res = min(res, tmp)
    }
    return res
}

func minimumRecolors1(blocks string, k int) int {
    res, count := 1 << 31, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, c := range blocks {
        if c == 'W' {
            count++
        }
        if i < k-1 { continue }
        res = min(res, count)
        if blocks[i - k + 1] == 'W' {
            count--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: blocks = "WBBWWBBWBW", k = 7
    // Output: 3
    // Explanation:
    // One way to achieve 7 consecutive black blocks is to recolor the 0th, 3rd, and 4th blocks
    // so that blocks = "BBBBBBBWBW". 
    // It can be shown that there is no way to achieve 7 consecutive black blocks in less than 3 operations.
    // Therefore, we return 3.
    fmt.Println(minimumRecolors("WBBWWBBWBW", 7)) // 3
    // Example 2:
    // Input: blocks = "WBWBBBW", k = 2
    // Output: 0
    // Explanation:
    // No changes need to be made, since 2 consecutive black blocks already exist.
    // Therefore, we return 0.
    fmt.Println(minimumRecolors("WBWBBBW", 2)) // 0

    fmt.Println(minimumRecolors1("WBBWWBBWBW", 7)) // 3
    fmt.Println(minimumRecolors1("WBWBBBW", 2)) // 0
}