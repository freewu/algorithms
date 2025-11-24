package main

// 3758. Convert Number Words to Digits
// You are given a string s consisting of lowercase English letters. 
// s may contain valid concatenated English words representing the digits 0 to 9, without spaces.

// Your task is to extract each valid number word in order and convert it to its corresponding digit, producing a string of digits.

// Parse s from left to right. At each position:
//     1. If a valid number word starts at the current position, append its corresponding digit to the result and advance by the length of that word.
//     2. Otherwise, skip exactly one character and continue parsing.

// Return the resulting digit string. If no number words are found, return an empty string.

// Example 1:
// Input: s = "onefourthree"
// Output: "143"
// Explanation:
// Parsing from left to right, extract the valid number words "one", "four", "three".
// These map to digits 1, 4, 3. Thus, the final result is "143".

// Example 2:
// Input: s = "ninexsix"
// Output: "96"
// Explanation:
// The substring "nine" is a valid number word and maps to 9.
// The character "x" does not match any valid number word prefix and is skipped.
// Then, the substring "six" is a valid number word and maps to 6, so the final result is "96".

// Example 3:
// Input: s = "zeero"
// Output: ""
// Explanation:
// No substring forms a valid number word during left-to-right parsing.
// All characters are skipped and incomplete fragments are ignored, so the result is an empty string.

// Example 4:
// Input: s = "tw"
// Output: ""
// Explanation:
// No substring forms a valid number word during left-to-right parsing.
// All characters are skipped and incomplete fragments are ignored, so the result is an empty string.

// Constraints:
//     1 <= s.length <= 10^5
//     s contains only lowercase English letters.

import "fmt"

func convertNumber(s string) string {
    words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    res := []byte{}
    for i := 0; i < len(s); i++ {
        for j, w := range words {
            r := i + len(w)
            if r <= len(s) && s[i:r] == w {
                res = append(res, '0' + byte(j))
                i = r - 1
                break
            }
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "onefourthree"
    // Output: "143"
    // Explanation:
    // Parsing from left to right, extract the valid number words "one", "four", "three".
    // These map to digits 1, 4, 3. Thus, the final result is "143".
    fmt.Println(convertNumber("onefourthree")) // "143"
    // Example 2:
    // Input: s = "ninexsix"
    // Output: "96"
    // Explanation:
    // The substring "nine" is a valid number word and maps to 9.
    // The character "x" does not match any valid number word prefix and is skipped.
    // Then, the substring "six" is a valid number word and maps to 6, so the final result is "96".
    fmt.Println(convertNumber("ninexsix")) // "96"
    // Example 3:
    // Input: s = "zeero"
    // Output: ""
    // Explanation:
    // No substring forms a valid number word during left-to-right parsing.
    // All characters are skipped and incomplete fragments are ignored, so the result is an empty string.
    fmt.Println(convertNumber("zeero")) // ""
    // Example 4:
    // Input: s = "tw"
    // Output: ""
    // Explanation:
    // No substring forms a valid number word during left-to-right parsing.
    // All characters are skipped and incomplete fragments are ignored, so the result is an empty string.
    fmt.Println(convertNumber("tw")) // ""

    fmt.Println(convertNumber("bleufrog")) // ""
    fmt.Println(convertNumber("leetcode")) // ""
    fmt.Println(convertNumber("zeroonetwothreefourfivesixseveneightnine")) // "0123456789"
    fmt.Println(convertNumber("nineeightsevensixfivefourthreetwoonezero")) // "9876543210"
}
