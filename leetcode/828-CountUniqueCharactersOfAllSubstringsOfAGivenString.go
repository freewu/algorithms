package main

// 828. Count Unique Characters of All Substrings of a Given String
// Let's define a function countUniqueChars(s) that returns the number of unique characters in s.
//     For example, calling countUniqueChars(s) if s = "LEETCODE" 
//     then "L", "T", "C", "O", "D" are the unique characters since they appear only once in s,
//     therefore countUniqueChars(s) = 5.

// Given a string s, return the sum of countUniqueChars(t) where t is a substring of s. 
// The test cases are generated such that the answer fits in a 32-bit integer.

// Notice that some substrings can be repeated so in this case you have to count the repeated ones too.

// Example 1:
// Input: s = "ABC"
// Output: 10
// Explanation: All possible substrings are: "A","B","C","AB","BC" and "ABC".
// Every substring is composed with only unique letters.
// Sum of lengths of all substring is 1 + 1 + 1 + 2 + 2 + 3 = 10

// Example 2:
// Input: s = "ABA"
// Output: 8
// Explanation: The same as example 1, except countUniqueChars("ABA") = 1.

// Example 3:
// Input: s = "LEETCODE"
// Output: 92

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of uppercase English letters only.

import "fmt"

func uniqueLetterString(s string) int {
    index := [26]int{ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, }
    count, summary, prev := make([]int, 26), 1, 1
    v := s[0] - 'A'
    index[v], count[v] = 0, 1
    for i := 1; i < len(s); i++ {
        v := s[i] - 'A'
        prev = prev + i - index[v] - count[v]
        count[v] = i - index[v]
        index[v] = i
        summary += prev
    }
    return summary
}

func uniqueLetterString1(s string) int {
    res, total, last0, last1 := 0, 0, [26]int{}, [26]int{}
    for i := range last0 {
        last0[i], last1[i] = -1, -1
    }
    for i, c := range s {
        c -= 'A'
        total += i - 2 * last0[c] + last1[c]
        res += total
        last1[c] = last0[c]
        last0[c] = i
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "ABC"
    // Output: 10
    // Explanation: All possible substrings are: "A","B","C","AB","BC" and "ABC".
    // Every substring is composed with only unique letters.
    // Sum of lengths of all substring is 1 + 1 + 1 + 2 + 2 + 3 = 10
    fmt.Println(uniqueLetterString("ABC")) // 10
    // Example 2:
    // Input: s = "ABA"
    // Output: 8
    // Explanation: The same as example 1, except countUniqueChars("ABA") = 1.
    fmt.Println(uniqueLetterString("ABA")) // 8
    // Example 3:
    // Input: s = "LEETCODE"
    // Output: 92
    fmt.Println(uniqueLetterString("LEETCODE")) // 92

    fmt.Println(uniqueLetterString1("ABC")) // 10
    fmt.Println(uniqueLetterString1("ABA")) // 8
    fmt.Println(uniqueLetterString1("LEETCODE")) // 92
}