package main

// 2937. Make Three Strings Equal
// You are given three strings: s1, s2, and s3. 
// In one operation you can choose one of these strings and delete its rightmost character. 
// Note that you cannot completely empty a string.

// Return the minimum number of operations required to make the strings equal. 
// If it is impossible to make them equal, return -1.

// Example 1:
// Input: s1 = "abc", s2 = "abb", s3 = "ab"
// Output: 2
// Explanation: Deleting the rightmost character from both s1 and s2 will result in three equal strings.

// Example 2:
// Input: s1 = "dac", s2 = "bac", s3 = "cac"
// Output: -1
// Explanation: Since the first letters of s1 and s2 differ, they cannot be made equal.

// Constraints:
//     1 <= s1.length, s2.length, s3.length <= 100
//     s1, s2 and s3 consist only of lowercase English letters.

import "fmt"

func findMinimumOperations(s1 string, s2 string, s3 string) int {
    i := 0
    for ; i < len(s1) && i < len(s2) && i < len(s3); i++ {
        if !(s1[i] == s2[i] && s2[i] == s3[i]) { 
            break 
        }
    }
    if i > 0 {
        return len(s1) + len(s2) + len(s3) - (i * 3)
    }
    return -1
}

func main() {
    // Example 1:
    // Input: s1 = "abc", s2 = "abb", s3 = "ab"
    // Output: 2
    // Explanation: Deleting the rightmost character from both s1 and s2 will result in three equal strings.
    fmt.Println(findMinimumOperations("abc", "abb", "ab")) // 2
    // Example 2:
    // Input: s1 = "dac", s2 = "bac", s3 = "cac"
    // Output: -1
    // Explanation: Since the first letters of s1 and s2 differ, they cannot be made equal.
    fmt.Println(findMinimumOperations("dac", "bac", "cac")) // -1

    fmt.Println(findMinimumOperations("bluefrog", "leetcode", "abcdefgh")) // -1
}