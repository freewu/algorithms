package main

// 424. Longest Repeating Character Replacement
// You are given a string s and an integer k. 
// You can choose any character of the string and change it to any other uppercase English character. 
// You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.

// Example 2:
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achieve this answer too.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of only uppercase English letters.
//     0 <= k <= s.length

import "fmt"

func characterReplacement(s string, k int) int {
    res, left, counter, freq := 0, 0, 0, make([]int, 26)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for right := 0; right < len(s); right++ {
        // 边滑动的过程中边统计出现频次最多的字母，因为最后求得的最长长度的解，一定是在出现频次最多的字母上，
        freq[s[right]-'A']++
        counter = max(counter, freq[s[right]-'A'])
        // 再改变其他字母得到的最长连续长度。窗口滑动的过程中，用窗口的长度减去窗口中出现频次最大的长度，
        // 如果差值比 K 大，就代表需要缩小左窗口了直到差值等于 K
        for right-left+1-counter > k {
            freq[s[left]-'A']--
            left++
        }
        // 不断的取出窗口的长度的最大值就可以了
        res = max(res, right - left + 1)
    }
    return res
}


func main() {
    // Example 1:
    // Input: s = "ABAB", k = 2
    // Output: 4
    // Explanation: Replace the two 'A's with two 'B's or vice versa.
    fmt.Println(characterReplacement("ABAB", 2)) // 4
    // Example 2:
    // Input: s = "AABABBA", k = 1
    // Output: 4
    // Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
    // The substring "BBBB" has the longest repeating letters, which is 4.
    // There may exists other ways to achieve this answer too.
    fmt.Println(characterReplacement("AABABBA", 1)) // 4
}