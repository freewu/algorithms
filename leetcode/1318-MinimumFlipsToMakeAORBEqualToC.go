package main 

// 1318. Minimum Flips to Make a OR b Equal to c
// Given 3 positives numbers a, b and c. Return the minimum flips required in some bits of a and b to make ( a OR b == c ). (bitwise OR operation).
// Flip operation consists of change any single bit 1 to 0 or change the bit 0 to 1 in their binary representation.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/01/06/sample_3_1676.png" />
// Input: a = 2, b = 6, c = 5
// Output: 3
// Explanation: After flips a = 1 , b = 4 , c = 5 such that (a OR b == c)

// Example 2:
// Input: a = 4, b = 2, c = 7
// Output: 1

// Example 3:
// Input: a = 1, b = 2, c = 3
// Output: 0
 
// Constraints:
//     1 <= a <= 10^9
//     1 <= b <= 10^9
//     1 <= c <= 10^9

import "fmt"

func minFlips(a int, b int, c int) int {
    res := 0
    for a > 0 || b > 0 || c > 0 {
        ai, bi, ci := a%2, b%2, c%2
        a, b, c = a/2, b/2, c/2
        if ai|bi != ci {
            res++
            if ci == 0 && ai == 1 && bi == 1 {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/01/06/sample_3_1676.png" />
    // Input: a = 2, b = 6, c = 5
    // Output: 3
    // Explanation: After flips a = 1 , b = 4 , c = 5 such that (a OR b == c)
    fmt.Println(minFlips(2, 6, 5)) // 3
    // Example 2:
    // Input: a = 4, b = 2, c = 7
    // Output: 1
    fmt.Println(minFlips(4, 2, 7)) // 1
    // Example 3:
    // Input: a = 1, b = 2, c = 3
    // Output: 0
    fmt.Println(minFlips(1, 2, 3)) // 0
}