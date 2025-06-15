package main

// 3581. Count Odd Letters from Number
// You are given an integer n perform the following steps:
//     1. Convert each digit of n into its lowercase English word (e.g., 4 → "four", 1 → "one").
//     2. Concatenate those words in the original digit order to form a string s.

// Return the number of distinct characters in s that appear an odd number of times.

// Example 1:
// Input: n = 41
// Output: 5
// Explanation:
// 41 → "fourone"
// Characters with odd frequencies: 'f', 'u', 'r', 'n', 'e'. Thus, the answer is 5.

// Example 2:
// Input: n = 20
// Output: 5
// Explanation:
// 20 → "twozero"
// Characters with odd frequencies: 't', 'w', 'z', 'e', 'r'. Thus, the answer is 5.

// Constraints:
//     1 <= n <= 10^9

import "fmt"

func countOddLetters(n int) int {
    dict := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" }
    res, mp := 0, make([]int,26)
    for true {
        v := n % 10
        for _, b :=  range dict[v] {
            mp[b - 'a']++
        }
        n /= 10
        if n == 0 { break }
    }
    for _,v := range mp {
        if v % 2 == 1 { 
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 41
    // Output: 5
    // Explanation:
    // 41 → "fourone"
    // Characters with odd frequencies: 'f', 'u', 'r', 'n', 'e'. Thus, the answer is 5.
    fmt.Println(countOddLetters(41)) // 5
    // Example 2:
    // Input: n = 20
    // Output: 5
    // Explanation:
    // 20 → "twozero"
    // Characters with odd frequencies: 't', 'w', 'z', 'e', 'r'. Thus, the answer is 5.
    fmt.Println(countOddLetters(20)) // 5

    fmt.Println(countOddLetters(1024)) // 6
    fmt.Println(countOddLetters(999_999_999)) // 2
    fmt.Println(countOddLetters(1_000_000_000)) // 3
}