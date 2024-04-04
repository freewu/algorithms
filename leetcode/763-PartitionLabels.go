package main

// 763. Partition Labels
// You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.
// Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.
// Return a list of integers representing the size of these parts.

// Example 1:
// Input: s = "ababcbacadefegdehijhklij"
// Output: [9,7,8]
// Explanation:
// The partition is "ababcbaca", "defegde", "hijhklij".
// This is a partition so that each letter appears in at most one part.
// A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.

// Example 2:
// Input: s = "eccbbbbdec"
// Output: [10]
 
// Constraints:
//     1 <= s.length <= 500
//     s consists of lowercase English letters.

import "fmt"

func partitionLabels(s string) []int {
    // 得到每个字符最后出一的位置
    m := map[rune]int{}
    for i, b := range s {
        m[b] = i
    }
    res, count,lastIndex := []int{}, 0, 0
    for i, b := range s {
        count++
        if val, ok := m[b]; ok && val > lastIndex {
            lastIndex = val
        }
        if lastIndex == i {
            res = append(res, count)
            count = 0
        }
    }
    return res
}

func partitionLabels1(s string) []int {
    maxIdx := make([]int, 26)
    for i := range s {
        maxIdx[s[i]-'a'] = i
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := make([]int, 0, len(s))
    left, right := 0, 0
    for i:= 0; i < len(s); i ++ {
        right = max(right, maxIdx[s[i]-'a'])
        if i == right {
            res = append(res, right-left+1)
            left = right + 1
        }
    }
    return res
}

func main() {
    // The partition is "ababcbaca", "defegde", "hijhklij".
    // This is a partition so that each letter appears in at most one part.
    // A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.
    fmt.Println(partitionLabels("ababcbacadefegdehijhklij")) // [9,7,8]
    fmt.Println(partitionLabels("eccbbbbdec")) // [10]

    fmt.Println(partitionLabels1("ababcbacadefegdehijhklij")) // [9,7,8]
    fmt.Println(partitionLabels1("eccbbbbdec")) // [10]
}