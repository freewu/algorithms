package main

// 2606. Find the Substring With Maximum Cost
// You are given a string s, a string chars of distinct characters and an integer array vals of the same length as chars.

// The cost of the substring is the sum of the values of each character in the substring. 
// The cost of an empty string is considered 0.

// The value of the character is defined in the following way:
//     1. If the character is not in the string chars, then its value is its corresponding position (1-indexed) in the alphabet.
//         For example, the value of 'a' is 1, the value of 'b' is 2, and so on. The value of 'z' is 26.
//     2. Otherwise, assuming i is the index where the character occurs in the string chars, then its value is vals[i].

// Return the maximum cost among all substrings of the string s.

// Example 1:
// Input: s = "adaa", chars = "d", vals = [-1000]
// Output: 2
// Explanation: The value of the characters "a" and "d" is 1 and -1000 respectively.
// The substring with the maximum cost is "aa" and its cost is 1 + 1 = 2.
// It can be proven that 2 is the maximum cost.

// Example 2:
// Input: s = "abc", chars = "abc", vals = [-1,-1,-1]
// Output: 0
// Explanation: The value of the characters "a", "b" and "c" is -1, -1, and -1 respectively.
// The substring with the maximum cost is the empty substring "" and its cost is 0.
// It can be proven that 0 is the maximum cost.

// Constraints:
//     1 <= s.length <= 10^5
//     s consist of lowercase English letters.
//     1 <= chars.length <= 26
//     chars consist of distinct lowercase English letters.
//     vals.length == chars.length
//     -1000 <= vals[i] <= 1000

import "fmt"

func maximumCostSubstring(s string, chars string, vals []int) int {
    res, end, cost, n := 0, 0, 0, len(s)
    costs := make(map[byte]int)
    for i := 0; i < len(chars); i++ {
        costs[chars[i]] = vals[i]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        b := s[i]
        if _, ok := costs[b]; ok {
            cost = costs[b]
        } else {
            cost = int(b) - 96
        }
        end += cost
        res = max(res, end)
        end = max(end, 0)
    }
    return res
}

func maximumCostSubstring1(s, chars string, vals []int) int {
    mp := [26]int{}
    for i := range mp {
        mp[i] = i + 1
    }
    for i, v := range chars {
        mp[v - 'a'] = vals[i]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, mx := 0, 0
    for _, v := range s {
        mx = max(mx, 0) + mp[v - 'a']
        res = max(res, mx)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "adaa", chars = "d", vals = [-1000]
    // Output: 2
    // Explanation: The value of the characters "a" and "d" is 1 and -1000 respectively.
    // The substring with the maximum cost is "aa" and its cost is 1 + 1 = 2.
    // It can be proven that 2 is the maximum cost.
    fmt.Println(maximumCostSubstring("adaa", "d", []int{-1000})) // 2
    // Example 2:
    // Input: s = "abc", chars = "abc", vals = [-1,-1,-1]
    // Output: 0
    // Explanation: The value of the characters "a", "b" and "c" is -1, -1, and -1 respectively.
    // The substring with the maximum cost is the empty substring "" and its cost is 0.
    // It can be proven that 0 is the maximum cost.
    fmt.Println(maximumCostSubstring("abc", "abc", []int{-1,-1,-1})) // 0

    fmt.Println(maximumCostSubstring("bluefrog", "abc", []int{-1,-1,-1})) // 84
    fmt.Println(maximumCostSubstring("leetcode", "abc", []int{-1,-1,-1})) // 65

    fmt.Println(maximumCostSubstring1("adaa", "d", []int{-1000})) // 2
    fmt.Println(maximumCostSubstring1("abc", "abc", []int{-1,-1,-1})) // 0
    fmt.Println(maximumCostSubstring1("bluefrog", "abc", []int{-1,-1,-1})) // 84
    fmt.Println(maximumCostSubstring1("leetcode", "abc", []int{-1,-1,-1})) // 65
}