package main

// 3271. Hash Divided String
// You are given a string s of length n and an integer k, where n is a multiple of k. 
// Your task is to hash the string s into a new string called result, which has a length of n / k.

// First, divide s into n / k substrings, each with a length of k. Then, initialize result as an empty string.

// For each substring in order from the beginning:
//     1. The hash value of a character is the index of that character in the English alphabet (e.g., 'a' → 0, 'b' → 1, ..., 'z' → 25).
//     2. Calculate the sum of all the hash values of the characters in the substring.
//     3. Find the remainder of this sum when divided by 26, which is called hashedChar.
//     4. Identify the character in the English lowercase alphabet that corresponds to hashedChar.
//     5. Append that character to the end of result.

// Return result.

// Example 1:
// Input: s = "abcd", k = 2
// Output: "bf"
// Explanation:
// First substring: "ab", 0 + 1 = 1, 1 % 26 = 1, result[0] = 'b'.
// Second substring: "cd", 2 + 3 = 5, 5 % 26 = 5, result[1] = 'f'.

// Example 2:
// Input: s = "mxz", k = 3
// Output: "i"
// Explanation:
// The only substring: "mxz", 12 + 23 + 25 = 60, 60 % 26 = 8, result[0] = 'i'.

// Constraints:
//     1 <= k <= 100
//     k <= s.length <= 1000
//     s.length is divisible by k.
//     s consists only of lowercase English letters.

import "fmt"

func stringHash(s string, k int) string {
    res, n := "", len(s)
    calc := func(s string) rune {
        var res rune
        for _, v := range s {
            res += (v - 'a')
        }
        return res % 26
    }
    for i := 0; i < n; i += k {
        end := i + k
        if end > n {
            end = n
        }
        res += string(calc(s[i:end]) + 'a')
    }
    return res
}

func stringHash1(s string, k int) string {
    res, sum := []byte{}, 0
    for i := range s {
        sum += int(s[i] - 'a')
        if i % k == k - 1 {
            sum, res = 0, append(res, byte('a' + sum % 26))
        }
    }
    return string(res)
}


func main() {
    // Example 1:
    // Input: s = "abcd", k = 2
    // Output: "bf"
    // Explanation:
    // First substring: "ab", 0 + 1 = 1, 1 % 26 = 1, result[0] = 'b'.
    // Second substring: "cd", 2 + 3 = 5, 5 % 26 = 5, result[1] = 'f'.
    fmt.Println(stringHash("abcd", 2)) // "bf"
    // Example 2:
    // Input: s = "mxz", k = 3
    // Output: "i"
    // Explanation:
    // The only substring: "mxz", 12 + 23 + 25 = 60, 60 % 26 = 8, result[0] = 'i'.
    fmt.Println(stringHash("mxz", 3)) // "i"

    fmt.Println(stringHash("bluefrog", 3)) // "gau"
    fmt.Println(stringHash("leetcode", 3)) // "tjh"

    fmt.Println(stringHash1("abcd", 2)) // "bf"
    fmt.Println(stringHash1("mxz", 3)) // "i"
    fmt.Println(stringHash1("bluefrog", 3)) // "ga"
    fmt.Println(stringHash1("leetcode", 3)) // "tj"
}