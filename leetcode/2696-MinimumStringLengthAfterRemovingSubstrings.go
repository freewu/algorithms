package main

// 2696. Minimum String Length After Removing Substrings
// You are given a string s consisting only of uppercase English letters.
// You can apply some operations to this string where, in one operation, you can remove any occurrence of one of the substrings "AB" or "CD" from s.
// Return the minimum possible length of the resulting string that you can obtain.
// Note that the string concatenates after removing the substring and could produce new "AB" or "CD" substrings.

// Example 1:
// Input: s = "ABFCACDB"
// Output: 2
// Explanation: We can do the following operations:
// - Remove the substring "ABFCACDB", so s = "FCACDB".
// - Remove the substring "FCACDB", so s = "FCAB".
// - Remove the substring "FCAB", so s = "FC".
// So the resulting length of the string is 2.
// It can be shown that it is the minimum length that we can obtain.

// Example 2:
// Input: s = "ACBBD"
// Output: 5
// Explanation: We cannot do any operations on the string so the length remains the same.

// Constraints:
//     1 <= s.length <= 100
//     s consists only of uppercase English letters.

import "fmt"

// func minLength(s string) int {
//     res := len(s)
//     for i := 0; i < len(s); i += 2 {
//         // 可以从 s 中删除 任一个 "AB" 或 "CD" 子字符串
//         if s[i] == 'A' && s[i + 1] == 'B' { res -= 2 }
//         if s[i] == 'C' && s[i + 1] == 'D' { res -= 2 }
//     }
//     return res
// }

func minLength(s string) int {
    ab, cd, res, found  := "AB", "CD", len(s) ,true
    for found {
        found = false
        newS := ""
        i := 0
        for i < len(s) {
            if i < len(s)-1 && (s[i:i+2] == ab || s[i:i+2] == cd) {
                i += 2
                found = true
            } else {
                newS += string(s[i])
                i++
            }
        }
        s = newS
        if len(newS) < res {
            res = len(newS)
        }
    }
    return res
}

func minLength1(s string) int {
    for i := 1; i < len(s); i++ {
        if (s[i]=='B' && s[i-1]=='A') || (s[i]=='D' && s[i-1]=='C') {
            if i == 1 {
                s = s[2:]
                i = 0
            } else {
                s = s[:i-1]+s[i+1:]
                i = i - 2
            }
        }
    }
    return len(s)
}

func main() {
    // Example 1:
    // Input: s = "ABFCACDB"
    // Output: 2
    // Explanation: We can do the following operations:
    // - Remove the substring "ABFCACDB", so s = "FCACDB".
    // - Remove the substring "FCACDB", so s = "FCAB".
    // - Remove the substring "FCAB", so s = "FC".
    // So the resulting length of the string is 2.
    // It can be shown that it is the minimum length that we can obtain.
    fmt.Println(minLength("ABFCACDB")) // 2
    // Example 2:
    // Input: s = "ACBBD"
    // Output: 5
    // Explanation: We cannot do any operations on the string so the length remains the same.
    fmt.Println(minLength("ACBBD")) // 5

    fmt.Println(minLength1("ABFCACDB")) // 2
    fmt.Println(minLength1("ACBBD")) // 5
}