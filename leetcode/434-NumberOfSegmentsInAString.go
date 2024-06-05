package main

// 434. Number of Segments in a String
// Given a string s, return the number of segments in the string.
// A segment is defined to be a contiguous sequence of non-space characters.

// Example 1:
// Input: s = "Hello, my name is John"
// Output: 5
// Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
// Example 2:

// Input: s = "Hello"
// Output: 1

// Constraints:
//     0 <= s.length <= 300
//     s consists of lowercase and uppercase English letters, digits, or one of the following characters "!@#$%^&*()_+-=',.:".
//     The only space character in s is ' '.

import "fmt"
import "regexp"
import "strings"

func countSegments2(s string) int {
    return len(strings.Fields(s))
}

func countSegments1(s string) int {
    return len(regexp.MustCompile(`([^\s]+)`).FindAll([]byte(s), -1))
}

func countSegments(s string) int {
    cnt := 0
    for i := range s {
        if i != 0 && s[i-1] != ' ' && s[i] == ' '{
            cnt++
        }
    }
    if len(s) > 0 && s[len(s)-1] != ' ' {
        cnt++
    }
    return cnt
}

// func countSegments(s string) int {
//     if len(s) == 0 {
//         return 0
//     }
//     res, space := 0, 0
//     for i := 0; i < len(s); i++ {
//         if s[i] == ' ' {
//             if i != 0 && s[i - 1] != ' ' {
//                 res++
//             }
//             space++
//         }
//     }
//     if len(s) == space { // 全是空格的情况
//         return 0
//     }
//     return res + 1
// }

func main() {
    // Example 1:
    // Input: s = "Hello, my name is John"
    // Output: 5
    // Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
    // Example 2:
    fmt.Println(countSegments("Hello, my name is John")) // 5
    // Input: s = "Hello"
    // Output: 1
    fmt.Println(countSegments("Hello")) // 1
    fmt.Println(countSegments("                ")) // 0
    fmt.Println(countSegments("a")) // 1

    fmt.Println(countSegments1("Hello, my name is John")) // 5
    fmt.Println(countSegments1("Hello")) // 1
    fmt.Println(countSegments1("                ")) // 0
    fmt.Println(countSegments1("a")) // 1

    fmt.Println(countSegments2("Hello, my name is John")) // 5
    fmt.Println(countSegments2("Hello")) // 1
    fmt.Println(countSegments2("                ")) // 0
    fmt.Println(countSegments2("a")) // 1
}