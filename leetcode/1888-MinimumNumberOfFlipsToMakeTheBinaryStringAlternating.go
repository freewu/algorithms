package main

// 1888. Minimum Number of Flips to Make the Binary String Alternating
// You are given a binary string s. 
// You are allowed to perform two types of operations on the string in any sequence:
//     Type-1: Remove the character at the start of the string s and append it to the end of the string.
//     Type-2: Pick any character in s and flip its value, i.e., if its value is '0' it becomes '1' and vice-versa.

// Return the minimum number of type-2 operations you need to perform such that s becomes alternating.

// The string is called alternating if no two adjacent characters are equal.

// For example, the strings "010" and "1010" are alternating, while the string "0100" is not.

// Example 1:
// Input: s = "111000"
// Output: 2
// Explanation: Use the first operation two times to make s = "100011".
// Then, use the second operation on the third and sixth elements to make s = "101010".

// Example 2:
// Input: s = "010"
// Output: 0
// Explanation: The string is already alternating.

// Example 3:
// Input: s = "1110"
// Output: 1
// Explanation: Use the second operation on the second element to make s = "1010".

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either '0' or '1'.

import "fmt"

// sliding window 
func minFlips(s string) int {
    res, ns0, ns1 := 1 << 31, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(s) * 2; i++ {
        if i < len(s) {
            if (i % 2 == 0 && s[i] == '0') || (i % 2 != 0 && s[i] == '1') {  // check every position and check wether need to do Type-2 pick
                ns1++
            } else {
                ns0++
            }
            continue
        }
        iter := i % len(s) // below is conditions of expanding (sliding window) characters
        if (iter % 2 == 0 && s[iter] == '0') || (iter % 2 != 0 && s[iter] == '1') { // first item of sliding window need to be removed, so need to 'REVERT' Type-2 pick
            ns1-- 
        } else {
            ns0--
        }
        if (i % 2 == 0 && s[iter] == '0') || (i % 2 != 0 && s[iter] == '1') { // check every position and check wether need to do Type-2 pick
            ns1++
        } else {
            ns0++
        } 
        res = min(res, min(ns0, ns1))
    }
    return res
}

func minFlips1(s string) int {
    n, count, str := len(s), 0, "01"
    for i := 0; i < n; i++ {
        if s[i] != str[i&1] {
            count++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := min(count, n - count)
    s += s
    for i := 0; i < n; i++ {
        if s[i] != str[i & 1] {
            count--
        }
        if s[i] != str[(i + n) & 1] {
            count++
        }
        res = min(res, min(count, n - count))
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "111000"
    // Output: 2
    // Explanation: Use the first operation two times to make s = "100011".
    // Then, use the second operation on the third and sixth elements to make s = "101010".
    fmt.Println(minFlips("111000")) // 2
    // Example 2:
    // Input: s = "010"
    // Output: 0
    // Explanation: The string is already alternating.
    fmt.Println(minFlips("010")) // 0
    // Example 3:
    // Input: s = "1110"
    // Output: 1
    // Explanation: Use the second operation on the second element to make s = "1010".
    fmt.Println(minFlips("1110")) // 1

    fmt.Println(minFlips1("111000")) // 2
    fmt.Println(minFlips1("010")) // 0
    fmt.Println(minFlips1("1110")) // 1
}