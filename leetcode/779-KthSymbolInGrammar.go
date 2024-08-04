package main

// 779. K-th Symbol in Grammar
// We build a table of n rows (1-indexed). We start by writing 0 in the 1st row. 
// Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.
//     For example, for n = 3, the 1st row is 0, the 2nd row is 01, and the 3rd row is 0110.

// Given two integer n and k, return the kth (1-indexed) symbol in the nth row of a table of n rows.

// Example 1:
// Input: n = 1, k = 1
// Output: 0
// Explanation: row 1: 0

// Example 2:
// Input: n = 2, k = 1
// Output: 0
// Explanation: 
// row 1: 0
// row 2: 01

// Example 3:
// Input: n = 2, k = 2
// Output: 1
// Explanation: 
// row 1: 0
// row 2: 01

// Constraints:
//     1 <= n <= 30
//     1 <= k <= 2n - 1

import "fmt"

// 递归
func kthGrammar(n int, k int) int {
    if n == 1 && k == 1 {
        return 0
    }
    prev_k := k / 2 + k % 2
    prev := kthGrammar(n - 1, prev_k)
    if k % 2 == 1 {
        return prev
    } else {
        return prev^1
    }
}

func kthGrammar1(n int, k int) int {
    if n == 1 && k == 1 {
        return 0
    }
    prev := kthGrammar1(n - 1, (k + 1) / 2)
    if k % 2 == 0 {
        return 1 - prev
    } else {
        return prev
    }
}

func main() {
    // Example 1:
    // Input: n = 1, k = 1
    // Output: 0
    // Explanation: row 1: 0
    fmt.Println(kthGrammar(1, 1)) // 0
    // Example 2:
    // Input: n = 2, k = 1
    // Output: 0
    // Explanation: 
    // row 1: 0
    // row 2: 01
    fmt.Println(kthGrammar(2, 1)) // 0
    // Example 3:
    // Input: n = 2, k = 2
    // Output: 1
    // Explanation: 
    // row 1: 0
    // row 2: 01
    fmt.Println(kthGrammar(2, 2)) // 1

    fmt.Println(kthGrammar1(1, 1)) // 0
    fmt.Println(kthGrammar1(2, 1)) // 0
    fmt.Println(kthGrammar1(2, 2)) // 1
}