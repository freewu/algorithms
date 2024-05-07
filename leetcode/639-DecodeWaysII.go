package main

// 639. Decode Ways II
// A message containing letters from A-Z can be encoded into numbers using the following mapping:
//     'A' -> "1"
//     'B' -> "2"
//     ...
//     'Z' -> "26"
// To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:
//     "AAJF" with the grouping (1 1 10 6)
//     "KJF" with the grouping (11 10 6)
// Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
// In addition to the mapping above, an encoded message may contain the '*' character, which can represent any digit from '1' to '9' ('0' is excluded). For example, the encoded message "1*" may represent any of the encoded messages "11", "12", "13", "14", "15", "16", "17", "18", or "19". Decoding "1*" is equivalent to decoding any of the encoded messages it can represent.
// Given a string s consisting of digits and '*' characters, return the number of ways to decode it.
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "*"
// Output: 9
// Explanation: The encoded message can represent any of the encoded messages "1", "2", "3", "4", "5", "6", "7", "8", or "9".
// Each of these can be decoded to the strings "A", "B", "C", "D", "E", "F", "G", "H", and "I" respectively.
// Hence, there are a total of 9 ways to decode "*".

// Example 2:
// Input: s = "1*"
// Output: 18
// Explanation: The encoded message can represent any of the encoded messages "11", "12", "13", "14", "15", "16", "17", "18", or "19".
// Each of these encoded messages have 2 ways to be decoded (e.g. "11" can be decoded to "AA" or "K").
// Hence, there are a total of 9 * 2 = 18 ways to decode "1*".

// Example 3:
// Input: s = "2*"
// Output: 15
// Explanation: The encoded message can represent any of the encoded messages "21", "22", "23", "24", "25", "26", "27", "28", or "29".
// "21", "22", "23", "24", "25", and "26" have 2 ways of being decoded, but "27", "28", and "29" only have 1 way.
// Hence, there are a total of (6 * 2) + (3 * 1) = 12 + 3 = 15 ways to decode "2*".
 
// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is a digit or '*'.

import "fmt"

func numDecodings(s string) int {
    M, n := 1000000007, len(s)
    dp := make([]int, n+1)
    dp[0] = 1
    if string(s[0]) == "*" {
        dp[1] = 9
    } else {
        if string(s[0]) == "0" {
            dp[1] = 0
        } else {
            dp[1] = 1
        }
    }
    for i := 1; i < n; i++ {
        num := string(s[i])
        if num == "*" {
            // Using star as an individual digit
            dp[i+1] = 9 * dp[i] % M
            // Using * in combination with previous digit
            if string(s[i-1]) == "1" {
                // If preceeding digit is 1, there are 9 additional
                // ways to to form 2 digit decodings (11-19)
                dp[i+1] = (dp[i+1] + 9*dp[i-1]) % M
            } else if string(s[i-1]) == "2" {
                // If preceeding digit is 1, there are 6 additional
                // ways to to form 2 digit decodings (21-16)
                dp[i+1] = (dp[i+1] + 6*dp[i-1]) % M
            } else if string(s[i-1]) == "*" {
                // If preceeding digit is 1, there are 9 + 6 additional
                // ways to to perform last 2 digit decodings
                dp[i+1] = (dp[i+1] + 15*dp[i-1]) % M
            }
        } else {
            // Using the individual digit, there are i-1
            // ways to perform single digit decoding
            if num != "0" {
                dp[i+1] = dp[i]
            }
            if string(s[i-1]) == "1" || string(s[i-1]) == "2" && num <= "6" {
                // If preceeding digit is 1, we have an additional
                // i-2 ways of decoding (last 2 digits as one letter)
                // Along with the existing i-1 ways of decoding with
                // last digit as one letter
                dp[i+1] = (dp[i+1] + dp[i-1]) % M
            } else if string(s[i-1]) == "*" {
                // If the preceeding digit is  * and the current digit is
                // less than 6, additonal of 2 * i-2 ways are possible:
                // (1x, 2x) where x is current digit
                // If x is > 6, we can have an additional of i-2 ways (1x)
                if num <= "6" {
                    dp[i+1] = (dp[i+1] + 2*dp[i-1]) % M
                } else {
                    dp[i+1] = (dp[i+1] + dp[i-1]) % M
                }
            }
        }
    }
    return dp[n]
}


func numDecodings1(s string) int {
    const mod int = 1e9 + 7
    a, b, c := 0, 1, 0

    check1digit := func(ch byte) int {
        if ch == '*' { return 9; }
        if ch == '0' { return 0; }
        return 1
    }
    check2digits := func(c0, c1 byte) int {
        if c0 == '*' && c1 == '*' {  return 15; }
        if c0 == '*' { 
            if c1 <= '6' { return 2; }
            return 1
        }
        if c1 == '*' {
            if c0 == '1' { return 9; }
            if c0 == '2' { return 6; }
            return 0
        }
        if c0 != '0' && (c0-'0')*10+(c1-'0') <= 26 {
            return 1
        }
        return 0
    }

    for i := range s {
        c = b * check1digit(s[i]) % mod
        if i > 0 {
            c = (c + a * check2digits(s[i-1], s[i])) % mod
        }
        a, b = b, c
    }
    return c
}

func main() {
    // Example 1:
    // Input: s = "*"
    // Output: 9
    // Explanation: The encoded message can represent any of the encoded messages "1", "2", "3", "4", "5", "6", "7", "8", or "9".
    // Each of these can be decoded to the strings "A", "B", "C", "D", "E", "F", "G", "H", and "I" respectively.
    // Hence, there are a total of 9 ways to decode "*".
    fmt.Println(numDecodings("*")) // 9
    // Example 2:
    // Input: s = "1*"
    // Output: 18
    // Explanation: The encoded message can represent any of the encoded messages "11", "12", "13", "14", "15", "16", "17", "18", or "19".
    // Each of these encoded messages have 2 ways to be decoded (e.g. "11" can be decoded to "AA" or "K").
    // Hence, there are a total of 9 * 2 = 18 ways to decode "1*".
    fmt.Println(numDecodings("1*")) // 18
    // Example 3:
    // Input: s = "2*"
    // Output: 15
    // Explanation: The encoded message can represent any of the encoded messages "21", "22", "23", "24", "25", "26", "27", "28", or "29".
    // "21", "22", "23", "24", "25", and "26" have 2 ways of being decoded, but "27", "28", and "29" only have 1 way.
    // Hence, there are a total of (6 * 2) + (3 * 1) = 12 + 3 = 15 ways to decode "2*".
    fmt.Println(numDecodings("2*")) // 15

    fmt.Println(numDecodings1("*")) // 9
    fmt.Println(numDecodings1("1*")) // 18
    fmt.Println(numDecodings1("2*")) // 15
}