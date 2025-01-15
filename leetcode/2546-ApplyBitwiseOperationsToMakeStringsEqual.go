package main

// 2546. Apply Bitwise Operations to Make Strings Equal
// You are given two 0-indexed binary strings s and target of the same length n. 
// You can do the following operation on s any number of times:
//     1. Choose two different indices i and j where 0 <= i, j < n.
//     2. Simultaneously, replace s[i] with (s[i] OR s[j]) and s[j] with (s[i] XOR s[j]).

// For example, if s = "0110", you can choose i = 0 and j = 2, 
// then simultaneously replace s[0] with (s[0] OR s[2] = 0 OR 1 = 1), 
// and s[2] with (s[0] XOR s[2] = 0 XOR 1 = 1), so we will have s = "1110".

// Return true if you can make the string s equal to target, or false otherwise.

// Example 1:
// Input: s = "1010", target = "0110"
// Output: true
// Explanation: We can do the following operations:
// - Choose i = 2 and j = 0. We have now s = "0010".
// - Choose i = 2 and j = 1. We have now s = "0110".
// Since we can make s equal to target, we return true.

// Example 2:
// Input: s = "11", target = "00"
// Output: false
// Explanation: It is not possible to make s equal to target with any number of operations.

// Constraints:
//     n == s.length == target.length
//     2 <= n <= 10^5
//     s and target consist of only the digits 0 and 1.

import "fmt"
import "strings"

func makeStringsEqual(s string, target string) bool {
    return strings.Contains(s, "1") == strings.Contains(target, "1")
}

func makeStringsEqual1(s string, target string) bool {
    // 1, 1 -> 1, 0
    // 0, 1 -> 1, 1
    // 0, 0 -> 0, 0
    if s == target { return true }
    s1, t1 := 0, 0
    for i := range s {
        if s[i] == '1' { s1++ }
        if target[i] == '1' { t1++ }
    }
    if s1 == 0 || t1 == 0 { return false }
    return true
}

func main() {
    // Example 1:
    // Input: s = "1010", target = "0110"
    // Output: true
    // Explanation: We can do the following operations:
    // - Choose i = 2 and j = 0. We have now s = "0010".
    // - Choose i = 2 and j = 1. We have now s = "0110".
    // Since we can make s equal to target, we return true.
    fmt.Println(makeStringsEqual("1010", "0110")) // true
    // Example 2:
    // Input: s = "11", target = "00"
    // Output: false
    // Explanation: It is not possible to make s equal to target with any number of operations.
    fmt.Println(makeStringsEqual("11", "00")) // false

    fmt.Println(makeStringsEqual1("1010", "0110")) // true
    fmt.Println(makeStringsEqual1("11", "00")) // false
}