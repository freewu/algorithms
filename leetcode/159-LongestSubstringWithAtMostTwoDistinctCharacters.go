package main

// 159. Longest Substring with At Most Two Distinct Characters
// Given a string s, return the length of the longest substring that contains at most two distinct characters.

// Example 1:
// Input: s = "eceba"
// Output: 3
// Explanation: The substring is "ece" which its length is 3.

// Example 2:
// Input: s = "ccaabbb"
// Output: 5
// Explanation: The substring is "aabbb" which its length is 5.
 
// Constraints:
//     1 <= s.length <= 10^5
//     s consists of English letters.

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
    mp := make(map[byte]int,58)
    res, count, start, end := 0, 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for end < len(s) {
        mp[s[end]]++
        if mp[s[end]] == 1 {
            count++ // 出现新的字符了
        }
        for count > 2  {
            mp[s[start]]--
            if mp[s[start]] == 0 {
                count--
            }
            start++
        }
        res = max(res, end - start + 1)
        end++
    } 
    return res
}

// best solution
func lengthOfLongestSubstringTwoDistinct1(s string) int {
    res, diff, i, j, n := 0, 0, 0, 0, len(s)
    count := make([]int, 58)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for ; j < n; j++ {
        if count[s[j] - 'A'] == 0 {
            diff++
        }
        count[s[j] - 'A']++
        for diff > 2 {
            if count[s[i] - 'A'] == 1 {
                diff--
            }
            count[s[i] - 'A']--
            i++
        }
        res = max(res, j - i + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "eceba"
    // Output: 3
    // Explanation: The substring is "ece" which its length is 3.
    fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba")) // 3
    // Example 2:
    // Input: s = "ccaabbb"
    // Output: 5
    // Explanation: The substring is "aabbb" which its length is 5.
    fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb")) // 5

    fmt.Println(lengthOfLongestSubstringTwoDistinct("bluefrog")) // 2
    fmt.Println(lengthOfLongestSubstringTwoDistinct("leetcode")) // 3

    fmt.Println(lengthOfLongestSubstringTwoDistinct1("eceba")) // 3
    fmt.Println(lengthOfLongestSubstringTwoDistinct1("ccaabbb")) // 5
    fmt.Println(lengthOfLongestSubstringTwoDistinct1("bluefrog")) // 2
    fmt.Println(lengthOfLongestSubstringTwoDistinct1("leetcode")) // 3
}