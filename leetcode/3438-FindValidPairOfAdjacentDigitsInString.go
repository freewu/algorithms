package main

// 3438. Find Valid Pair of Adjacent Digits in String
// You are given a string s consisting only of digits.
// A valid pair is defined as two adjacent digits in s such that:
//     1. The first digit is not equal to the second.
//     2. Each digit in the pair appears in s exactly as many times as its numeric value.

// Return the first valid pair found in the string s when traversing from left to right. 
// If no valid pair exists, return an empty string.

// Example 1:
// Input: s = "2523533"
// Output: "23"
// Explanation:
// Digit '2' appears 2 times and digit '3' appears 3 times. Each digit in the pair "23" appears in s exactly as many times as its numeric value. Hence, the output is "23".

// Example 2:
// Input: s = "221"
// Output: "21"
// Explanation:
// Digit '2' appears 2 times and digit '1' appears 1 time. Hence, the output is "21".

// Example 3:
// Input: s = "22"
// Output: ""
// Explanation:
// There are no valid adjacent pairs.

// Constraints:
//     2 <= s.length <= 100
//     s only consists of digits from '1' to '9'.

import "fmt"

func findValidPair(s string) string {
    freq := make([]int, 10)
    for _, v := range s {
        freq[v - '0']++
    }
    for i := 0; i < len(s) - 1; i++ {
        first, second := int(s[i] - '0'), int(s[i + 1] - '0')
        if first != second && freq[first] == first && freq[second] == second {
            return string(s[i]) + string(s[i + 1])
        }
    }
    return ""
}

func findValidPair1(s string) string {
    count := [10]int{}
    for _, v := range s {
        count[v - '0']++
    }
    for i := 1; i < len(s); i++ {
        x, y := int(s[i - 1] - '0'), int(s[i] - '0')
        if x != y && count[x] == x && count[y] == y {
            return s[i-1 : i+1]
        }
    }
    return ""
}

func main() {
    // Example 1:
    // Input: s = "2523533"
    // Output: "23"
    // Explanation:
    // Digit '2' appears 2 times and digit '3' appears 3 times. Each digit in the pair "23" appears in s exactly as many times as its numeric value. Hence, the output is "23".
    fmt.Println(findValidPair("2523533")) // "23"
    // Example 2:
    // Input: s = "221"
    // Output: "21"
    // Explanation:
    // Digit '2' appears 2 times and digit '1' appears 1 time. Hence, the output is "21".
    fmt.Println(findValidPair("221")) // "21"
    // Example 3:
    // Input: s = "22"
    // Output: ""
    // Explanation:
    // There are no valid adjacent pairs.
    fmt.Println(findValidPair("22")) // ""

    fmt.Println(findValidPair("123456789")) // ""
    fmt.Println(findValidPair("987654321")) // ""

    fmt.Println(findValidPair1("2523533")) // "23"
    fmt.Println(findValidPair1("221")) // "21"
    fmt.Println(findValidPair1("22")) // ""
    fmt.Println(findValidPair1("123456789")) // ""
    fmt.Println(findValidPair1("987654321")) // ""
}