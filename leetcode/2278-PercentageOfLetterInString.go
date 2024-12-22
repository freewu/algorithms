package main

// 2278. Percentage of Letter in String
// Given a string s and a character letter, 
// return the percentage of characters in s that equal letter rounded down to the nearest whole percent.

// Example 1:
// Input: s = "foobar", letter = "o"
// Output: 33
// Explanation:
// The percentage of characters in s that equal the letter 'o' is 2 / 6 * 100% = 33% when rounded down, so we return 33.

// Example 2:
// Input: s = "jjjj", letter = "k"
// Output: 0
// Explanation:
// The percentage of characters in s that equal the letter 'k' is 0%, so we return 0.

// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters.
//     letter is a lowercase English letter.

import "fmt"

func percentageLetter(s string, letter byte) int {
    mp := make([]int, 26)
    for _, v := range s {
        mp[v - 'a']++
    }
    v := mp[letter - 'a']
    if v == 0 { return 0 }
    return v * 100 / len(s)
}

func percentageLetter1(s string, letter byte) int {
    n, count := len(s), 0
    for i := 0; i < n; i++ {
        if s[i] == letter {
            count++
        }
    }
    return 100 * count / n
}

func main() {
    // Example 1:
    // Input: s = "foobar", letter = "o"
    // Output: 33
    // Explanation:
    // The percentage of characters in s that equal the letter 'o' is 2 / 6 * 100% = 33% when rounded down, so we return 33.
    fmt.Println(percentageLetter("foobar", 'o')) // 33
    // Example 2:
    // Input: s = "jjjj", letter = "k"
    // Output: 0
    // Explanation:
    // The percentage of characters in s that equal the letter 'k' is 0%, so we return 0.
    fmt.Println(percentageLetter("jjjj", 'k')) // 0

    fmt.Println(percentageLetter1("foobar", 'o')) // 33
    fmt.Println(percentageLetter1("jjjj", 'k')) // 0
}