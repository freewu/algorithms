package main

// 984. String Without AAA or BBB
// Given two integers a and b, return any string s such that:
//     s has length a + b and contains exactly a 'a' letters, and exactly b 'b' letters,
//     The substring 'aaa' does not occur in s, and
//     The substring 'bbb' does not occur in s.

// Example 1:
// Input: a = 1, b = 2
// Output: "abb"
// Explanation: "abb", "bab" and "bba" are all correct answers.

// Example 2:
// Input: a = 4, b = 1
// Output: "aabaa"

// Constraints:
//     0 <= a, b <= 100
//     It is guaranteed such an s exists for the given a and b.

import "fmt"

// greedy
func strWithout3a3b(a int, b int) string {
    res, first, second := make([]byte, a + b), byte('a'), byte('b')
    if a < b {
        first, second = second, first
        a, b = b, a
    }
    i := 0
    for a != b && a >= 2 && b >= 1 { // 使用 aab 模式填充
        res[i], res[i+1], res[i+2] = first, first, second
        a -= 2
        b -= 1
        i += 3
    }
    for a != 0 || b != 0 { // 使用 ab 模式填充
        if a != 0 {
            res[i] = first
            i+=1
            a -= 1
        }
        if b != 0 {
            res[i] = second
            i+=1
            b-=1
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: a = 1, b = 2
    // Output: "abb"
    // Explanation: "abb", "bab" and "bba" are all correct answers.
    fmt.Println(strWithout3a3b(1,2)) // "abb"
    // Example 2:
    // Input: a = 4, b = 1
    // Output: "aabaa"
    fmt.Println(strWithout3a3b(4,1)) // "aabaa"
}