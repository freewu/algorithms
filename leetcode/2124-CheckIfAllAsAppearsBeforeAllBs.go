package main

// 2124. Check if All A's Appears Before All B's
// Given a string s consisting of only the characters 'a' and 'b', 
// return true if every 'a' appears before every 'b' in the string. 
// Otherwise, return false.

// Example 1:
// Input: s = "aaabbb"
// Output: true
// Explanation:
// The 'a's are at indices 0, 1, and 2, while the 'b's are at indices 3, 4, and 5.
// Hence, every 'a' appears before every 'b' and we return true.

// Example 2:
// Input: s = "abab"
// Output: false
// Explanation:
// There is an 'a' at index 2 and a 'b' at index 1.
// Hence, not every 'a' appears before every 'b' and we return false.

// Example 3:
// Input: s = "bbb"
// Output: true
// Explanation:
// There are no 'a's, hence, every 'a' appears before every 'b' and we return true.

// Constraints:
//     1 <= s.length <= 100
//     s[i] is either 'a' or 'b'.

import "fmt"

func checkString(s string) bool {
    flag := false
    for _, v := range s {
        if v == 'a' {
            if flag {
                return false
            }
        } else {
            flag = true
        }
    }
    return true
}

func checkString1(s string) bool {
    i, n := 0, len(s)
    for i < n && s[i] == 'a' {
        i++
    }
    for i < n && s[i] == 'b' {
        i++
    }
    return i == n
}

func main() {
    // Example 1:
    // Input: s = "aaabbb"
    // Output: true
    // Explanation:
    // The 'a's are at indices 0, 1, and 2, while the 'b's are at indices 3, 4, and 5.
    // Hence, every 'a' appears before every 'b' and we return true.
    fmt.Println(checkString("aaabbb")) // true
    // Example 2:
    // Input: s = "abab"
    // Output: false
    // Explanation:
    // There is an 'a' at index 2 and a 'b' at index 1.
    // Hence, not every 'a' appears before every 'b' and we return false.
    fmt.Println(checkString("abab")) // false
    // Example 3:
    // Input: s = "bbb"
    // Output: true
    // Explanation:
    // There are no 'a's, hence, every 'a' appears before every 'b' and we return true.
    fmt.Println(checkString("bbb")) // true

    fmt.Println(checkString("bluefrog")) // true
    fmt.Println(checkString("leetcode")) // true
    fmt.Println(checkString("freewu")) // true

    fmt.Println(checkString1("aaabbb")) // true
    fmt.Println(checkString1("abab")) // false
    fmt.Println(checkString1("bbb")) // true
}