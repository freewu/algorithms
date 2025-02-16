package main

// 面试题 16.18. Pattern Matching LCCI
// You are given two strings, pattern and value. 
// The pattern string consists of just the letters a and b, describing a pattern within a string. 
// For example, the string catcatgocatgo matches the pattern aabab (where cat is a and go is b). 
// It also matches patterns like a, ab, and b. 
// Write a method to determine if value matches pattern. a and b cannot be the same string.

// Example 1:
// Input:  pattern = "abba", value = "dogcatcatdog"
// Output:  true

// Example 2:
// Input:  pattern = "abba", value = "dogcatcatfish"
// Output:  false

// Example 3:
// Input:  pattern = "aaaa", value = "dogcatcatdog"
// Output:  false

// Example 4:
// Input:  pattern = "abba", value = "dogdogdogdog"
// Output:  true
// Explanation:  "a"="dogdog",b=""，vice versa.

// Note:
//     0 <= len(pattern) <= 1000
//     0 <= len(value) <= 1000
//     pattern only contains "a" and "b", value only contains lowercase letters.

import "fmt"

func patternMatching(pattern string, value string) bool {
    a, b := 0, 0
    for i := 0; i < len(pattern); i++ {
        if pattern[i] == 'a' {
            a++
        } else {
            b++
        }
    }
    if a < b {
        a, b = b, a
        tmp := ""
        for i := 0; i < len(pattern); i++ {
            if pattern[i] == 'a' {
                tmp += "b"
            } else {
                tmp += "a"
            }
        }
        pattern = tmp
    }
    if len(value) == 0 { return b == 0 }
    if len(pattern) == 0 { return false }
    for na := 0; a * na <= len(value); na++ {
        rest := len(value) - a * na
        if (b == 0 && rest == 0) || (b != 0 && rest % b == 0) {
            nb := 0
            if b == 0 {
                nb = 0
            } else {
                nb = rest / b
            }
            pos, correct := 0, true
            va, vb := "", ""
            for i := 0; i < len(pattern); i++ {
                if pattern[i] == 'a' {
                    sub := value[pos:pos + na]
                    if len(va) == 0 {
                        va = sub
                    } else if va != sub {
                        correct = false
                        break
                    }
                    pos += na
                } else {
                    sub := value[pos:pos + nb]
                    if len(vb) == 0 {
                        vb = sub
                    } else if vb != sub {
                        correct = false
                        break
                    }
                    pos += nb
                }
            }
            if correct && va != vb {
                return true
            }
        } 
    }
    return false
}

func main() {
    // Example 1:
    // Input:  pattern = "abba", value = "dogcatcatdog"
    // Output:  true
    fmt.Println(patternMatching("abba", "dogcatcatdog")) // true
    // Example 2:
    // Input:  pattern = "abba", value = "dogcatcatfish"
    // Output:  false
    fmt.Println(patternMatching("abba", "dogcatcatfish")) // false
    // Example 3:
    // Input:  pattern = "aaaa", value = "dogcatcatdog"
    // Output:  false
    fmt.Println(patternMatching("aaaa", "dogcatcatdog")) // false
    // Example 4:
    // Input:  pattern = "abba", value = "dogdogdogdog"
    // Output:  true
    // Explanation:  "a"="dogdog",b=""，vice versa.
    fmt.Println(patternMatching("abba", "dogdogdogdog")) // true

    fmt.Println(patternMatching("abab", "bluefrogbluefrog")) // true
    fmt.Println(patternMatching("abba", "leetcodeleetcode")) // true
}