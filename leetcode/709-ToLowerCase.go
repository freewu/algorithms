package main

// 709. To Lower Case
// Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.

// Example 1:
// Input: s = "Hello"
// Output: "hello"

// Example 2:
// Input: s = "here"
// Output: "here"

// Example 3:
// Input: s = "LOVELY"
// Output: "lovely"
 
// Constraints:
//     1 <= s.length <= 100
//     s consists of printable ASCII characters.

import "fmt"
import "strings"

func toLowerCase(s string) string {
    res := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        if s[i] >= 'A' && s[i] <= 'Z' {
            res[i] = s[i] + 32
        } else {
            res[i] = s[i]
        }
    }
    return string(res)
}

// use strings lib
func toLowerCase1(s string) string {
    return strings.ToLower(s)
}

func main() {
    //fmt.Println('a' - 'A')
    fmt.Println(toLowerCase("Hello")) // hello
    fmt.Println(toLowerCase("here")) // here
    fmt.Println(toLowerCase("LOVELY")) // LOVELY

    fmt.Println(toLowerCase1("Hello")) // hello
    fmt.Println(toLowerCase1("here")) // here
    fmt.Println(toLowerCase1("LOVELY")) // LOVELY
}