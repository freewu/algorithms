package main

// 2522. Partition String Into Substrings With Values at Most K
// You are given a string s consisting of digits from 1 to 9 and an integer k.

// A partition of a string s is called good if:
//     1. Each digit of s is part of exactly one substring.
//     2. The value of each substring is less than or equal to k.

// Return the minimum number of substrings in a good partition of s. 
// If no good partition of s exists, return -1.

// Note that:
//     1. The value of a string is its result when interpreted as an integer. 
//        For example, the value of "123" is 123 and the value of "1" is 1.
//     2. A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "165462", k = 60
// Output: 4
// Explanation: We can partition the string into substrings "16", "54", "6", and "2". Each substring has a value less than or equal to k = 60.
// It can be shown that we cannot partition the string into less than 4 substrings.

// Example 2:
// Input: s = "238182", k = 5
// Output: -1
// Explanation: There is no good partition for this string.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is a digit from '1' to '9'.
//     1 <= k <= 10^9

import "fmt"
import "strconv"

func minimumPartition1(s string, k int) int {
    res, start, end := 1, 0, 1
    for end <= len(s) {
        v, _ := strconv.Atoi(s[start:end])
        if v > k {
            if start + 1 == end { return -1 }
            start = end - 1
            res++
        } else {
            end++
        }
    }
    return res
}

func minimumPartition(s string, k int) int {
    res, num := 1, 0
    for i := 0; i < len(s); i++ {
        d := int(s[i] - '0')
        if d > k { return -1 }
        num = num * 10 + d
        if num > k {
            num = d
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "165462", k = 60
    // Output: 4
    // Explanation: We can partition the string into substrings "16", "54", "6", and "2". Each substring has a value less than or equal to k = 60.
    // It can be shown that we cannot partition the string into less than 4 substrings.
    fmt.Println(minimumPartition("165462", 60)) // 4
    // Example 2:
    // Input: s = "238182", k = 5
    // Output: -1
    // Explanation: There is no good partition for this string.
    fmt.Println(minimumPartition("238182", 5)) // -1

    fmt.Println(minimumPartition1("165462", 60)) // 4
    fmt.Println(minimumPartition1("238182", 5)) // -1
}