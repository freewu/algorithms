package main

// 2405. Optimal Partition of String
// Given a string s, partition the string into one or more substrings such that the characters in each substring are unique.
// That is, no letter appears in a single substring more than once.

// Return the minimum number of substrings in such a partition.

// Note that each character should belong to exactly one substring in a partition.

// Example 1:
// Input: s = "abacaba"
// Output: 4
// Explanation:
// Two possible partitions are ("a","ba","cab","a") and ("ab","a","ca","ba").
// It can be shown that 4 is the minimum number of substrings needed.

// Example 2:
// Input: s = "ssssss"
// Output: 6
// Explanation:
// The only valid partition is ("s","s","s","s","s","s").

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of only English lowercase letters.

import "fmt"

func partitionString(s string) int {
    res, c := 1, uint32(0)
    runeBit := func(r rune) uint32 { return 1 << (r - 'a') }
    for _, r := range s {
        bit := runeBit(r)
        if 0 < c & bit {
            c = 0
            res++
        }
        c |= bit
    }
    return res
}

func partitionString1(s string) int {
    res, mask := 1, 0
    for i := 0; i < len(s); i++ {
        v := int(s[i] - 'a')
        if mask >> v & 1 == 1 {
            mask = 0
            res += 1
        }
        mask |= 1 << v
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abacaba"
    // Output: 4
    // Explanation:
    // Two possible partitions are ("a","ba","cab","a") and ("ab","a","ca","ba").
    // It can be shown that 4 is the minimum number of substrings needed.
    fmt.Println(partitionString("abacaba")) // 4
    // Example 2:
    // Input: s = "ssssss"
    // Output: 6
    // Explanation:
    // The only valid partition is ("s","s","s","s","s","s").
    fmt.Println(partitionString("ssssss")) // 6

    fmt.Println(partitionString("bluefrog")) // 1
    fmt.Println(partitionString("leetcode")) // 3

    fmt.Println(partitionString1("abacaba")) // 4
    fmt.Println(partitionString1("ssssss")) // 6
    fmt.Println(partitionString1("bluefrog")) // 1
    fmt.Println(partitionString1("leetcode")) // 3
}