package main

// 4105. Repeated String Match II
// You are given two strings a and b​​​​​​​ consisting of lowercase English letters, and two integers x and y.

// Return true if the string formed by repeating b exactly y times (b * y) is a substring of the string formed by repeating a exactly x times (a * x). 
// Otherwise, return false.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: a = "ab", b = "abab", x = 2, y = 1
// Output: true
// Explanation:
//     a * x = "ab" * 2 = "abab"
//     b * y = "abab" * 1 = "abab"
//     They are identical, so b * y is a substring of a * x. Thus, the answer is true.

// Example 2:
// Input: a = "abc", b = "bc", x = 2, y = 3
// Output: false
// Explanation:
//     a * x = "abc" * 2 = "abcabc"
//     b * y = "bc" * 3 = "bcbcbc"
//     The answer is false as b * y = "bcbcbc" is not a substring of a * x = "abcabc".

// Example 3:
// Input: a = "abcd", b = "cdab", x = 2, y = 1
// Output: true
// Explanation:
//     a * x = "abcd" * 2 = "abcdabcd"
//     b * y = "cdab" appears in a * x = "abcdabcd" at indices [2..5]. Thus, the answer is true.

// Constraints:
//     1 <= a.length <= 100
//     1 <= b.length <= 100
//     1 <= x <= 10^{9}
//     1 <= y <= 10^{9}​​​​​​​
//     a and b consist of lowercase English letters.

import "fmt"
import "strings"

func repeatedStringMatch(a, b string, x, y int) bool {
    n, m := len(a), len(b)
    an, bn := n * x, m * y
    if bn > an {
        return false
    }
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x int, y int) int { return (x * y) / gcd(x, y) }
    mx := min(n, an - bn + 1)
    firsts := min(lcm(n, m), bn)
    A := strings.Repeat(a, (mx+firsts-2)/n+1)
    B := strings.Repeat(b, firsts/m)
    return strings.Contains(A, B)
}

func main() {
    // Example 1:
    // Input: a = "ab", b = "abab", x = 2, y = 1
    // Output: true
    // Explanation:
    //     a * x = "ab" * 2 = "abab"
    //     b * y = "abab" * 1 = "abab"
    //     They are identical, so b * y is a substring of a * x. Thus, the answer is true.
    fmt.Println(repeatedStringMatch("ab", "abab", 2, 1)) // true 
    // Example 2:
    // Input: a = "abc", b = "bc", x = 2, y = 3
    // Output: false
    // Explanation:
    //     a * x = "abc" * 2 = "abcabc"
    //     b * y = "bc" * 3 = "bcbcbc"
    //     The answer is false as b * y = "bcbcbc" is not a substring of a * x = "abcabc".
    fmt.Println(repeatedStringMatch("abc", "bc", 2, 3)) // false
    // Example 3:
    // Input: a = "abcd", b = "cdab", x = 2, y = 1
    // Output: true
    // Explanation:
    //     a * x = "abcd" * 2 = "abcdabcd"
    //     b * y = "cdab" appears in a * x = "abcdabcd" at indices [2..5]. Thus, the answer is true.
    fmt.Println(repeatedStringMatch("abcd", "cdab", 2, 1)) // true
}