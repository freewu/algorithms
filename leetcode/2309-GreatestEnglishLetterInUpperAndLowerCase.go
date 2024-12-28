package main

// 2309. Greatest English Letter in Upper and Lower Case
// Given a string of English letters s, 
// return the greatest English letter which occurs as both a lowercase and uppercase letter in s. 
// The returned letter should be in uppercase. If no such letter exists, return an empty string.

// An English letter b is greater than another letter a if b appears after a in the English alphabet.

// Example 1:
// Input: s = "lEeTcOdE"
// Output: "E"
// Explanation:
// The letter 'E' is the only letter to appear in both lower and upper case.

// Example 2:
// Input: s = "arRAzFif"
// Output: "R"
// Explanation:
// The letter 'R' is the greatest letter to appear in both lower and upper case.
// Note that 'A' and 'F' also appear in both lower and upper case, but 'R' is greater than 'F' or 'A'.

// Example 3:
// Input: s = "AbCdEfGhIjK"
// Output: ""
// Explanation:
// There is no letter that appears in both lower and upper case.

// Constraints:
//     1 <= s.length <= 1000
//     s consists of lowercase and uppercase English letters.

import "fmt"
import "math/bits"

func greatestLetter1(s string) string {
    mask := uint64(0)
    for _, c := range s {
        mask |= 1 << (c - 'A')
    }
    mask &= mask >> ('a'-'A')
    if mask == 0 { return "" }
    return string(byte('A' + bits.Len64(mask) - 1))
}

func greatestLetter(s string) string {
    mp, diff := make([]bool, 'z' + 1), 'a' - 'A'
    for _, v := range s {
        mp[v] = true
    }
    for i := 'Z'; i >= 'A'; i-- {
        if mp[i] && mp[i + diff] {
            return string(i)
        }
    }
    return ""
}

func main() {
    // Example 1:
    // Input: s = "lEeTcOdE"
    // Output: "E"
    // Explanation:
    // The letter 'E' is the only letter to appear in both lower and upper case.
    fmt.Println(greatestLetter("lEeTcOdE")) // "E"
    // Example 2:
    // Input: s = "arRAzFif"
    // Output: "R"
    // Explanation:
    // The letter 'R' is the greatest letter to appear in both lower and upper case.
    // Note that 'A' and 'F' also appear in both lower and upper case, but 'R' is greater than 'F' or 'A'.
    fmt.Println(greatestLetter("arRAzFif")) // "R"
    // Example 3:
    // Input: s = "AbCdEfGhIjK"
    // Output: ""
    // Explanation:
    // There is no letter that appears in both lower and upper case.
    fmt.Println(greatestLetter("AbCdEfGhIjK")) // ""

    fmt.Println(greatestLetter1("lEeTcOdE")) // "E"
    fmt.Println(greatestLetter1("arRAzFif")) // "R"
    fmt.Println(greatestLetter1("AbCdEfGhIjK")) // ""
}