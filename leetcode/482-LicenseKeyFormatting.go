package main

// 482. License Key Formatting
// You are given a license key represented as a string s that consists of only alphanumeric characters and dashes. 
// The string is separated into n + 1 groups by n dashes. You are also given an integer k.

// We want to reformat the string s such that each group contains exactly k characters, 
// except for the first group, which could be shorter than k but still must contain at least one character. 
// Furthermore, there must be a dash inserted between two groups, and you should convert all lowercase letters to uppercase.

// Return the reformatted license key.

// Example 1:
// Input: s = "5F3Z-2e-9-w", k = 4
// Output: "5F3Z-2E9W"
// Explanation: The string s has been split into two parts, each part has 4 characters.
// Note that the two extra dashes are not needed and can be removed.

// Example 2:
// Input: s = "2-5g-3-J", k = 2
// Output: "2-5G-3J"
// Explanation: The string s has been split into three parts, each part has 2 characters except the first part as it could be shorter as mentioned above.
 
// Constraints:
//     1 <= s.length <= 10^5
//     s consists of English letters, digits, and dashes '-'.
//     1 <= k <= 10^4

import "fmt"
import "strings"
import "bytes"

func licenseKeyFormatting(s string, k int) string {
    s = strings.ToUpper(strings.Replace(s, "-", "", -1)) // 大写并替换掉 - 
    if len(s) <= k { // 不够分组
        return s
    }
    f := len(s) % k 
    var res bytes.Buffer
    if f != 0 {
        res.WriteString(s[:f] + "-") 
    }
    for i:= f + k; i < len(s); i += k {
        res.WriteString(s[f:i])
        res.WriteString("-")
        f = i
    }
    res.WriteString(s[f:])
    return res.String()
}

func main() {
    // Example 1:
    // Input: s = "5F3Z-2e-9-w", k = 4
    // Output: "5F3Z-2E9W"
    // Explanation: The string s has been split into two parts, each part has 4 characters.
    // Note that the two extra dashes are not needed and can be removed.
    fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4)) // "5F3Z-2E9W"
    // Example 2:
    // Input: s = "2-5g-3-J", k = 2
    // Output: "2-5G-3J"
    // Explanation: The string s has been split into three parts, each part has 2 characters except the first part as it could be shorter as mentioned above.
    fmt.Println(licenseKeyFormatting("2-5g-3-J", 2)) // "2-5G-3J"
}