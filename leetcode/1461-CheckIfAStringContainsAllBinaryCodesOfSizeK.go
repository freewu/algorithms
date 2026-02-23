package main

// 1461. Check If a String Contains All Binary Codes of Size K
// Given a binary string s and an integer k, return true if every binary code of length k is a substring of s. 
// Otherwise, return false.

// Example 1:
// Input: s = "00110110", k = 2
// Output: true
// Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They can be all found as substrings at indices 0, 1, 3 and 2 respectively.

// Example 2:
// Input: s = "0110", k = 1
// Output: true
// Explanation: The binary codes of length 1 are "0" and "1", it is clear that both exist as a substring. 

// Example 3:
// Input: s = "0110", k = 2
// Output: false
// Explanation: The binary code "00" is of length 2 and does not exist in the array.

// Constraints:
//     1 <= s.length <= 5 * 10^5
//     s[i] is either '0' or '1'.
//     1 <= k <= 20

import "fmt"

func hasAllCodes(s string, k int) bool {
    n, set := len(s), make(map[string]bool)
    if n < k { return false }
    for i := 0; i <= n - k; i++ {
        set[s[i:i+k]] = true
        if len(set) == 1 << k {
            return true
        }
    }
    return len(set) == 1 << k
}

func hasAllCodes1(s string, k int) bool {
    need := 1 << k
    visited, mask, curr := make([]bool, 1 << k), need - 1, 0
    // Check if the length of substrings of length k is less than the required number i.e 1<<k-1
    if len(s) - k < mask {
        return false
    }
    for i := 0; i < len(s); i++ {
        curr = ((curr << 1) | int(s[i]-'0')) & mask
        // Build the number associated with the initial substring
        if i < k-1 { continue }
        if !visited[curr] {
            need--
            visited[curr] = true
            if need == 0 { return true }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: s = "00110110", k = 2
    // Output: true
    // Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They can be all found as substrings at indices 0, 1, 3 and 2 respectively.
    fmt.Println(hasAllCodes("00110110", 2)) // true
    // Example 2:
    // Input: s = "0110", k = 1
    // Output: true
    // Explanation: The binary codes of length 1 are "0" and "1", it is clear that both exist as a substring. 
    fmt.Println(hasAllCodes("0110", 1)) // true
    // Example 3:
    // Input: s = "0110", k = 2
    // Output: false
    // Explanation: The binary code "00" is of length 2 and does not exist in the array.
    fmt.Println(hasAllCodes("0110", 2)) // false

    fmt.Println(hasAllCodes("1111111111", 2)) // false
    fmt.Println(hasAllCodes("0000000000", 2)) // false
    fmt.Println(hasAllCodes("1010101010", 2)) // false
    fmt.Println(hasAllCodes("0101010101", 2)) // false
    fmt.Println(hasAllCodes("1111100000", 2)) // false
    fmt.Println(hasAllCodes("0000011111", 2)) // false

    fmt.Println(hasAllCodes1("00110110", 2)) // true
    fmt.Println(hasAllCodes1("0110", 1)) // true
    fmt.Println(hasAllCodes1("0110", 2)) // false
    fmt.Println(hasAllCodes1("1111111111", 2)) // false
    fmt.Println(hasAllCodes1("0000000000", 2)) // false
    fmt.Println(hasAllCodes1("1010101010", 2)) // false
    fmt.Println(hasAllCodes1("0101010101", 2)) // false
    fmt.Println(hasAllCodes1("1111100000", 2)) // false
    fmt.Println(hasAllCodes1("0000011111", 2)) // false
}