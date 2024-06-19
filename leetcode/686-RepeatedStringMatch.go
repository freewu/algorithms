package main

// 686. Repeated String Match
// Given two strings a and b, return the minimum number of times you should repeat string a so that string b is a substring of it. 
// If it is impossible for b​​​​​​ to be a substring of a after repeating it, return -1.

// Notice: string "abc" repeated 0 times is "", repeated 1 time is "abc" and repeated 2 times is "abcabc".

// Example 1:
// Input: a = "abcd", b = "cdabcdab"
// Output: 3
// Explanation: We return 3 because by repeating a three times "abcdabcdabcd", b is a substring of it.
// a:  abcd abcd abcd
// b:    cd abcd ab

// Example 2:
// Input: a = "a", b = "aa"
// Output: 2

// Constraints:
//     1 <= a.length, b.length <= 10^4
//     a and b consist of lowercase English letters.

import "fmt"
import "strings"

func repeatedStringMatch(a string, b string) int {
    count, aa := 0, a
    for {
        if strings.Contains(aa, b) {
            return count + 1
        } else if len(a) * count > len(b) {
            return -1
        }
        aa += a
        count++
    }
}

// best solution
func repeatedStringMatch1(a string, b string) int {
    if strings.Contains(a, b) {
        return 1
    }
    idx := strings.Index(b, a)
    if idx == -1 {
        if strings.Contains(a+a, b) {
            return 2
        }
        return -1
    }
    count, ls, rs := 1, "", ""
    if idx != 0 {
        ls = b[:idx]
    }
    if idx+len(a) < len(b) {
        rs = b[idx+len(a):]
    }
    if ls != "" {
        if len(a) > len(ls) && ls == a[len(a)-len(ls):] {
            count++
        } else {
            return -1
        }
    }
    if rs != "" {
        c := len(rs) / len(a)
        sb := strings.Builder{}
        for i := 0; i < c; i++ {
            sb.WriteString(a)
        }
        temp := sb.String()
        if !strings.Contains(rs, temp) {
            return -1
        }
        count += c
        if len(temp) != len(rs) {
            if a[:len(rs)-len(temp)]!=rs[len(temp):] {
                return -1
            }
            count++
        }
    }
    return count
}

func main() {
    // Example 1:
    // Input: a = "abcd", b = "cdabcdab"
    // Output: 3
    // Explanation: We return 3 because by repeating a three times "abcdabcdabcd", b is a substring of it.
    // a:  abcd abcd abcd
    // b:    cd abcd ab
    fmt.Println(repeatedStringMatch("abcd","cdabcdab"))
    // Example 2:
    // Input: a = "a", b = "aa"
    // Output: 2
    fmt.Println(repeatedStringMatch("a","aa"))

    fmt.Println(repeatedStringMatch1("abcd","cdabcdab"))
    fmt.Println(repeatedStringMatch1("a","aa"))
}