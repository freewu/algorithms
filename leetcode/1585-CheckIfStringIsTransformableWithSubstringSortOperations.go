package main

// 1585. Check If String Is Transformable With Substring Sort Operations
// Given two strings s and t, transform string s into string t using the following operation any number of times:

//     Choose a non-empty substring in s and sort it in place so the characters are in ascending order.
//         For example, applying the operation on the underlined substring in "14234" results in "12344".

// Return true if it is possible to transform s into t. Otherwise, return false.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "84532", t = "34852"
// Output: true
// Explanation: You can transform s into t using the following sort operations:
// "84532" (from index 2 to 3) -> "84352"
// "84352" (from index 0 to 2) -> "34852"

// Example 2:
// Input: s = "34521", t = "23415"
// Output: true
// Explanation: You can transform s into t using the following sort operations:
// "34521" -> "23451"
// "23451" -> "23415"

// Example 3:
// Input: s = "12345", t = "12435"
// Output: false

// Constraints:
//     s.length == t.length
//     1 <= s.length <= 10^5
//     s and t consist of only digits.

import "fmt"

func isTransformable(s string, t string) bool {
    sr, tr, sc, tc := make([][]int, 10), make([][]int, 10), make([]int, 10), make([]int, 10)
    for i := range sr {
        sr[i], tr[i] = make([]int, 10), make([]int, 10)
    }
    equal := true
    for i := 0; i < len(s); i++ {
        ss, tt := int(s[i] - '0'), int(t[i] - '0')
        if equal {
            if tt > ss { // cannot produce t > s
                return false
            } else if tt < ss {
                equal = false
            }
        }
        // running frequencies of digits in s & t
        sc[ss]++
        tc[tt]++
        // running inversions in s & t
        for j := 9; j > ss; j-- { sr[ss][j] += sc[j] }
        for j := 9; j > tt; j-- { tr[tt][j] += tc[j] }
    }
    for i := 0; i < 9; i++ {
        if sc[i] != tc[i] { return false } // cannot modify digit count
        for j := i + 1; j <= 9; j++ {
            if tr[i][j] > sr[i][j] { return false }  // cannot add inversions
        }
    }
    return true
}

func isTransformable1(s string, t string) bool {
    n, m := len(s), len(t)
    if n != m {
        return false
    }
    g := make([][]int, 10)
    for i, ch := range s {
        idx := int(ch) - int('0')
        g[idx] = append(g[idx], i)
    }
    for i := 0; i < n; i++ {
        dig := int(t[i]) - int('0')
        if len(g[dig]) == 0 {
            return false
        }
        for j := 0; j < dig; j++ {
            if len(g[j]) > 0 && g[j][0] < g[dig][0] {
                return false
            }
        }
        g[dig] = g[dig][1:]
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "84532", t = "34852"
    // Output: true
    // Explanation: You can transform s into t using the following sort operations:
    // "84532" (from index 2 to 3) -> "84352"
    // "84352" (from index 0 to 2) -> "34852"
    fmt.Println(isTransformable("84532", "34852")) // true
    // Example 2:
    // Input: s = "34521", t = "23415"
    // Output: true
    // Explanation: You can transform s into t using the following sort operations:
    // "34521" -> "23451"
    // "23451" -> "23415"
    fmt.Println(isTransformable("34521", "23415")) // true
    // Example 3:
    // Input: s = "12345", t = "12435"
    // Output: false
    fmt.Println(isTransformable("12345", "12435")) // false

    fmt.Println(isTransformable1("84532", "34852")) // true
    fmt.Println(isTransformable1("34521", "23415")) // true
    fmt.Println(isTransformable1("12345", "12435")) // false
}