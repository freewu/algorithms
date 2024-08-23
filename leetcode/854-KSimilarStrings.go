package main

// 854. K-Similar Strings
// Strings s1 and s2 are k-similar (for some non-negative integer k) if we can swap the positions of two letters in s1 exactly k times so that the resulting string equals s2.

// Given two anagrams s1 and s2, return the smallest k for which s1 and s2 are k-similar.

// Example 1:
// Input: s1 = "ab", s2 = "ba"
// Output: 1
// Explanation: The two string are 1-similar because we can use one swap to change s1 to s2: "ab" --> "ba".

// Example 2:
// Input: s1 = "abc", s2 = "bca"
// Output: 2
// Explanation: The two strings are 2-similar because we can use two swaps to change s1 to s2: "abc" --> "bac" --> "bca".

// Constraints:
//     1 <= s1.length <= 20
//     s2.length == s1.length
//     s1 and s2 contain only lowercase letters from the set {'a', 'b', 'c', 'd', 'e', 'f'}.
//     s2 is an anagram of s1.

import "fmt"

func kSimilarity(a string, b string) int {
    res, achars, bchars := 0, []byte(a), []byte(b)
    checkAllOptions := func(achars []byte, bchars []byte, i int, b string) int {
        res := 1 << 31 - 1
        for j := i + 1; j < len(achars); j++ {
            if achars[j] == bchars[i] && achars[j] != bchars[j] {
                achars[i], achars[j] = achars[j], achars[i]
                newStr := string(achars)
                res = min(res, 1 + kSimilarity(newStr, b))
                achars[i], achars[j] = achars[j], achars[i] // backtrack
            }
        }
        return res
    }
    getAllPerfectMatches := func(achars []byte, bchars []byte) int {
        res := 0
        for i := 0; i < len(achars); i++ {
            if achars[i] == bchars[i] {
                continue
            }
            for j := i + 1; j < len(achars); j++ {
                if achars[j] == bchars[i] && bchars[j] == achars[i] {
                    achars[i], achars[j] = achars[j], achars[i]
                    res++
                    break
                }
            }
        }
        return res
    }
    res += getAllPerfectMatches(achars, bchars)
    for i := 0; i < len(achars); i++ {
        if achars[i] == bchars[i] {
            continue
        }
        return res + checkAllOptions(achars, bchars, i, b)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s1 = "ab", s2 = "ba"
    // Output: 1
    // Explanation: The two string are 1-similar because we can use one swap to change s1 to s2: "ab" --> "ba".
    fmt.Println(kSimilarity("ab", "ba")) // 1
    // Example 2:
    // Input: s1 = "abc", s2 = "bca"
    // Output: 2
    // Explanation: The two strings are 2-similar because we can use two swaps to change s1 to s2: "abc" --> "bac" --> "bca".
    fmt.Println(kSimilarity("abc", "bca")) // 2
}