package main

// 2506. Count Pairs Of Similar Strings
// You are given a 0-indexed string array words.

// Two strings are similar if they consist of the same characters.
//     1. For example, "abca" and "cba" are similar since both consist of characters 'a', 'b', and 'c'.
//     2. However, "abacba" and "bcfd" are not similar since they do not consist of the same characters.

// Return the number of pairs (i, j) such that 0 <= i < j <= word.length - 1 and the two strings words[i] and words[j] are similar.

// Example 1:
// Input: words = ["aba","aabb","abcd","bac","aabc"]
// Output: 2
// Explanation: There are 2 pairs that satisfy the conditions:
// - i = 0 and j = 1 : both words[0] and words[1] only consist of characters 'a' and 'b'. 
// - i = 3 and j = 4 : both words[3] and words[4] only consist of characters 'a', 'b', and 'c'. 

// Example 2:
// Input: words = ["aabb","ab","ba"]
// Output: 3
// Explanation: There are 3 pairs that satisfy the conditions:
// - i = 0 and j = 1 : both words[0] and words[1] only consist of characters 'a' and 'b'. 
// - i = 0 and j = 2 : both words[0] and words[2] only consist of characters 'a' and 'b'.
// - i = 1 and j = 2 : both words[1] and words[2] only consist of characters 'a' and 'b'.

// Example 3:
// Input: words = ["nba","cba","dba"]
// Output: 0
// Explanation: Since there does not exist any pair that satisfies the conditions, we return 0.

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 100
//     words[i] consist of only lowercase English letters.

import "fmt"

func similarPairs(words []string) int {
    res, mp := 0, make(map[int]int)
    mask := func(s string) int {
        res := 0
        for i := 0; i < len(s); i++ {
            res |= 1 << (s[i] - 'a')
        }
        return res
    }
    for i := 0; i < len(words); i++ {
        mp[mask(words[i])]++
    }
    for _, v := range mp {
        res += (v - 1) * v / 2
    }
    return res
}

func similarPairs1(words []string) int {
    res, mp := 0, make(map[int]int) // k: mask, v: count
    for _, word := range words {
        mask := 0
        for _, c := range word {
            mask |= 1 << (c - 'a')
        }
        res += mp[mask]
        mp[mask]++
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["aba","aabb","abcd","bac","aabc"]
    // Output: 2
    // Explanation: There are 2 pairs that satisfy the conditions:
    // - i = 0 and j = 1 : both words[0] and words[1] only consist of characters 'a' and 'b'. 
    // - i = 3 and j = 4 : both words[3] and words[4] only consist of characters 'a', 'b', and 'c'. 
    fmt.Println(similarPairs([]string{"aba","aabb","abcd","bac","aabc"})) // 2
    // Example 2:
    // Input: words = ["aabb","ab","ba"]
    // Output: 3
    // Explanation: There are 3 pairs that satisfy the conditions:
    // - i = 0 and j = 1 : both words[0] and words[1] only consist of characters 'a' and 'b'. 
    // - i = 0 and j = 2 : both words[0] and words[2] only consist of characters 'a' and 'b'.
    // - i = 1 and j = 2 : both words[1] and words[2] only consist of characters 'a' and 'b'.
    fmt.Println(similarPairs([]string{"aabb","ab","ba"})) // 3
    // Example 3:
    // Input: words = ["nba","cba","dba"]
    // Output: 0
    // Explanation: Since there does not exist any pair that satisfies the conditions, we return 0.
    fmt.Println(similarPairs([]string{"nba","cba","dba"})) // 0

    fmt.Println(similarPairs([]string{"bluefrog","leetcode"})) // 0

    fmt.Println(similarPairs1([]string{"aba","aabb","abcd","bac","aabc"})) // 2
    fmt.Println(similarPairs1([]string{"aabb","ab","ba"})) // 3
    fmt.Println(similarPairs1([]string{"nba","cba","dba"})) // 0
    fmt.Println(similarPairs1([]string{"bluefrog","leetcode"})) // 0
}