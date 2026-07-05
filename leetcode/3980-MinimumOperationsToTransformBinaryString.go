package main

// 3980. Minimum Operations to Transform Binary String
// You are given two binary strings s1 and s2 of the same length n.

// You can perform the following operations on s1 any number of times, in any order:
//     1. Choose an index i such that s1[i] == '0', and change it to '1'.
//     2. Choose an index i such that 0 <= i < n - 1, and both s1[i] and s1[i + 1] are '1'. Change both characters to '0'.

// Return the minimum number of operations required to make s1 equal to s2. If it is impossible, return -1.

// Example 1:
// Input: s1 = "11", s2 = "00"
// Output: 1
// Explanation:
// Change indices 0 and 1 from '1' to '0' in one operation, so "11" becomes "00". Thus, the answer is 1.

// Example 2:
// Input: s1 = "01", s2 = "10"
// Output: 3
// Explanation:
// Change index 0 from '0' to '1', so "01" becomes "11".
// Change indices 0 and 1 from '1' to '0', so "11" becomes "00".
// Change index 0 from '0' to '1', so "00" becomes "10".
// Thus, the answer is 3.

// Example 3:
// Input: s1 = "1", s2 = "0"
// Output: -1
// Explanation:
// The first operation cannot change '1' to '0', and the second operation requires two adjacent characters. Therefore, it is impossible.

// Constraints:
//     1 <= n == s1.length == s2.length <= 10^5
//     s1 and s2 consist only of '0' and '1'.

import "fmt"

func minOperations(s1 string, s2 string) int {
    ops, ones, n := 0, 0, len(s1)
    if n == 1 {
        return int(s2[0]) - int(s1[0])
    }
    for i := range n {
        if s1[i] == '1' && s2[i] == '0' {
            ones++
        } else {
            ops += ones/2 + ones%2*2 + int(s1[i] ^ s2[i])     
            ones = 0
        }
    }
    return ops + ones / 2 + ones % 2 * 2
}

func minOperations1(s1 string, s2 string) int {
    res, n := 0, len(s1)
    if n == 1 && s1 == "1" && s2 == "0" {
        return -1
    }
    for i := 0; i < n; i++ {
        if s1[i] == s2[i] {
            continue
        }
        if s1[i] == '0' && s2[i] == '1' {
            res++
            continue
        }
        j := i
        for i+1 < n && s1[i+1] == '1' && s2[i+1] == '0' {
            i++
        }
        c := i-j+1
        res += c / 2
        if c % 2 == 1 {
            res += 2
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s1 = "11", s2 = "00"
    // Output: 1
    // Explanation:
    // Change indices 0 and 1 from '1' to '0' in one operation, so "11" becomes "00". Thus, the answer is 1.
    fmt.Println(minOperations("11", "00")) // 1 
    // Example 2:
    // Input: s1 = "01", s2 = "10"
    // Output: 3
    // Explanation:
    // Change index 0 from '0' to '1', so "01" becomes "11".
    // Change indices 0 and 1 from '1' to '0', so "11" becomes "00".
    // Change index 0 from '0' to '1', so "00" becomes "10".
    // Thus, the answer is 3.
    fmt.Println(minOperations("01", "10")) // 3
    // Example 3:
    // Input: s1 = "1", s2 = "0"
    // Output: -1
    // Explanation:
    // The first operation cannot change '1' to '0', and the second operation requires two adjacent characters. Therefore, it is impossible.
    fmt.Println(minOperations("1", "0")) // -1

    fmt.Println(minOperations("111", "000")) // 3
    fmt.Println(minOperations("111", "111")) // 0
    fmt.Println(minOperations("000", "000")) // 0
    fmt.Println(minOperations("000", "111")) // 3
    fmt.Println(minOperations("111", "001")) // 1
    fmt.Println(minOperations("111", "110")) // 2
    fmt.Println(minOperations("000", "001")) // 1
    fmt.Println(minOperations("000", "110")) // 2

    fmt.Println(minOperations1("11", "00")) // 1 
    fmt.Println(minOperations1("01", "10")) // 3
    fmt.Println(minOperations1("1", "0")) // -1
    fmt.Println(minOperations1("111", "000")) // 3
    fmt.Println(minOperations1("111", "111")) // 0
    fmt.Println(minOperations1("000", "000")) // 0
    fmt.Println(minOperations1("000", "111")) // 3
    fmt.Println(minOperations1("111", "001")) // 1
    fmt.Println(minOperations1("111", "110")) // 2
    fmt.Println(minOperations1("000", "001")) // 1
    fmt.Println(minOperations1("000", "110")) // 2
}