package main

// 2168. Unique Substrings With Equal Digit Frequency
// Given a digit string s, return the number of unique substrings of s where every digit appears the same number of times.

// Example 1:
// Input: s = "1212"
// Output: 5
// Explanation: The substrings that meet the requirements are "1", "2", "12", "21", "1212".
// Note that although the substring "12" appears twice, it is only counted once.

// Example 2:
// Input: s = "12321"
// Output: 9
// Explanation: The substrings that meet the requirements are "1", "2", "3", "12", "23", "32", "21", "123", "321".

// Constraints:
//     1 <= s.length <= 1000
//     s consists of digits.

import "fmt"

func equalDigitFrequency(s string) int {
    n, mp := len(s), make(map[string]bool)
    for i := 0; i < n; i++ {
        arr, mx, sum := [10]int{}, 1, 1
        arr[int(s[i]-'0')] = 1
        for j := i + 1; j <= n; j++ {
            if (j - i) % sum == 0 && (j - i) / sum == mx { mp[s[i:j]] = true }
            if j < n {
                index := int(s[j] - '0')
                if arr[index] == 0 { sum++ }
                arr[index]++
                if arr[index] > mx { mx = arr[index] }
            }
        }
    }
    return len(mp)
}

func main() {
    // Example 1:
    // Input: s = "1212"
    // Output: 5
    // Explanation: The substrings that meet the requirements are "1", "2", "12", "21", "1212".
    // Note that although the substring "12" appears twice, it is only counted once.
    fmt.Println(equalDigitFrequency("1212")) // 5 
    // Example 2:
    // Input: s = "12321"
    // Output: 9
    // Explanation: The substrings that meet the requirements are "1", "2", "3", "12", "23", "32", "21", "123", "321".
    fmt.Println(equalDigitFrequency("12321")) // 9
}