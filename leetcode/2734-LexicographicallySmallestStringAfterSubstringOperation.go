package main

// 2734. Lexicographically Smallest String After Substring Operation
// Given a string s consisting of lowercase English letters. Perform the following operation:
//     Select any non-empty substring then replace every letter of the substring with the preceding letter of the English alphabet. 
//     For example, 'b' is converted to 'a', and 'a' is converted to 'z'.

// Return the lexicographically smallest string after performing the operation.

// Example 1:
// Input: s = "cbabc"
// Output: "baabc"
// Explanation:
// Perform the operation on the substring starting at index 0, and ending at index 1 inclusive.

// Example 2:
// Input: s = "aa"
// Output: "az"
// Explanation:
// Perform the operation on the last letter.

// Example 3:
// Input: s = "acbbc"
// Output: "abaab"
// Explanation:
// Perform the operation on the substring starting at index 1, and ending at index 4 inclusive.

// Example 4:
// Input: s = "leetcode"
// Output: "kddsbncd"
// Explanation:
// Perform the operation on the entire string.

// Constraints:
//     1 <= s.length <= 3 * 10^5
//     s consists of lowercase English letters

import "fmt"
import "strings"

func smallestString(s string) string {
    arr := strings.Split(s, "a")
    index := 0
    for index < len(arr) && len(arr[index]) == 0 {
        index++
    }
    if index == len(arr) {
        arr[index-1] = "z"
        arr = arr[1:]
        return strings.Join(arr, "a")
    }
    res := make([]byte, len(arr[index]))
    for i := range arr[index] {
        res[i] = arr[index][i] - 1
    }
    arr[index] = string(res)
    return strings.Join(arr, "a")
}

func smallestString1(s string) string {
    arr, n := []byte(s), len(s)
    for i, v := range arr {
        if v > 'a' {
            for j := i; j < n; j++ {
                if arr[j] == 'a' {
                    return string(arr)
                }
                arr[j]--
            }
            return string(arr)
        }
    }
    arr[n-1] = 'z'
    return string(arr)
}

func main() {
    // Example 1:
    // Input: s = "cbabc"
    // Output: "baabc"
    // Explanation:
    // Perform the operation on the substring starting at index 0, and ending at index 1 inclusive.
    fmt.Println(smallestString("cbabc")) // "baabc"
    // Example 2:
    // Input: s = "aa"
    // Output: "az"
    // Explanation:
    // Perform the operation on the last letter.
    fmt.Println(smallestString("aa")) // "az"
    // Example 3:
    // Input: s = "acbbc"
    // Output: "abaab"
    // Explanation:
    // Perform the operation on the substring starting at index 1, and ending at index 4 inclusive.
    fmt.Println(smallestString("acbbc")) // "abaab"
    // Example 4:
    // Input: s = "leetcode"
    // Output: "kddsbncd"
    // Explanation:
    // Perform the operation on the entire string.
    fmt.Println(smallestString("leetcode")) // "kddsbncd"

    fmt.Println(smallestString1("cbabc")) // "baabc"
    fmt.Println(smallestString1("aa")) // "az"
    fmt.Println(smallestString1("acbbc")) // "abaab"
    fmt.Println(smallestString1("leetcode")) // "kddsbncd"
}