package main

// 340. Longest Substring with At Most K Distinct Characters
// Given a string s and an integer k, 
// return the length of the longest substring of s that contains at most k distinct characters.

// Example 1:
// Input: s = "eceba", k = 2
// Output: 3
// Explanation: The substring is "ece" with length 3.

// Example 2:
// Input: s = "aa", k = 1
// Output: 2
// Explanation: The substring is "aa" with length 2.
 
// Constraints:
//     1 <= s.length <= 5 * 10^4
//     0 <= k <= 50

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
    // 向右枚举，左滑满足条件 curStr 里包含 > k 的字符
    res, left, right, l, m := 0, 0, 0, len(s), make(map[byte]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for ; right < l; right++ {
        m[s[right]]++
        for len(m) > k {
            lc := s[left]
            m[lc]--
            if m[lc] == 0 {
                delete(m, lc)
            }
            left++
        }
        res = max(res, right - left + 1)
    }
    return res
}

func lengthOfLongestSubstringKDistinct1(s string, k int) int {
    res, inf, positions := -1 << 32 - 1, 1 << 32 - 1, map[byte]int{} // 为每一个字符记录最右位置的索引
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for left, right := 0, 0; right < len(s); right++ {
        positions[s[right]] = right
        if len(positions) == k + 1 { // 长度刚好由  增加到 k+1 时，需要进行左边界调整 窗口大小
            indexToRemove := inf // 获取最小的位置
            for _, pos := range positions {
                indexToRemove = min(indexToRemove, pos)
            }
            delete(positions, s[indexToRemove])
            left = indexToRemove + 1 // 调整左边界
        }
        res = max(res, right - left + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "eceba", k = 2
    // Output: 3
    // Explanation: The substring is "ece" with length 3.
    fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2)) // 3
    // Example 2:
    // Input: s = "aa", k = 1
    // Output: 2
    // Explanation: The substring is "aa" with length 2.
    fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1)) // 2

    fmt.Println(lengthOfLongestSubstringKDistinct1("eceba", 2)) // 3
    fmt.Println(lengthOfLongestSubstringKDistinct1("aa", 1)) // 2
}