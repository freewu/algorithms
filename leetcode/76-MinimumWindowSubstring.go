package main

// 76. Minimum Window Substring
// Given two strings s and t of lengths m and n respectively, 
// return the minimum window substring of s such that every character in t (including duplicates) is included in the window. 
// If there is no such substring, return the empty string "".

// The testcases will be generated such that the answer is unique.

// Example 1:
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

// Example 2:
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.

// Example 3:
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.

// Constraints:
//     m == s.length
//     n == t.length
//     1 <= m, n <= 10^5
//     s and t consist of uppercase and lowercase English letters.

// Follow up: Could you find an algorithm that runs in O(m + n) time?

// 解题思路:
//     滑动窗口
//     在窗口滑动的过程中不断的包含字符串 T，直到完全包含字符串 T 的字符以后，记下左右窗口的位置和窗口大小。
//     每次都不断更新这个符合条件的窗口和窗口大小的最小值。最后输出结果即可。

import "fmt"

func minWindow(s string, t string) string {
    if s == "" || t == "" {
        return ""
    }
    tFreq, sFreq := [256]int{}, [256]int{}
    result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s) + 1, 0
    for i := 0; i < len(t); i++ {
        tFreq[t[i]-'a']++
    }
    for left < len(s) {
        if right+1 < len(s) && count < len(t) {
            sFreq[s[right+1]-'a']++
            if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
                count++
            }
            right++
        } else {
            if right-left+1 < minW && count == len(t) {
                minW = right - left + 1
                finalLeft = left
                finalRight = right
            }
            if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
                count--
            }
            sFreq[s[left]-'a']--
            left++
        }
    }
    if finalLeft != -1 {
        result = string(s[finalLeft : finalRight+1])
    }
    return result
}

// best solution
func minWindow1(s string, t string) string {
    dict, inf := [128]int{}, 1 << 32 - 1
    for _, ch := range t {
        dict[ch]++
    }
    counter, begin, end, head, d:= len(t), 0, 0, 0, inf
    for end < len(s) {
        c1 := s[end]

        if dict[c1] > 0 {
            counter--
        }
        dict[c1]--
        end++
        for counter == 0 {
            if d > end-begin {
                d = end - begin
                head = begin
            }
            c2 := s[begin]
            dict[c2]++
            if dict[c2] > 0 {
                counter++
            }
            begin++
        }
    }
    if d == inf {
        return ""
    }
    return s[head : head+d]
}

func main() {
    // Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
    fmt.Printf("minWindow(\"ADOBECODEBANC\",\"ABC\") = %v\n",minWindow("ADOBECODEBANC","ABC")) // "BANC"
    // Explanation: The entire string s is the minimum window.
    fmt.Printf("minWindow(\"a\",\"a\") = %v\n",minWindow("a","a")) // "a"
    // Explanation: Both 'a's from t must be included in the window.
    // Since the largest window of s only has one 'a', return empty string.
    fmt.Printf("minWindow(\"a\",\"aa\") = %v\n",minWindow("a","aa")) // ""

    fmt.Printf("minWindow1(\"ADOBECODEBANC\",\"ABC\") = %v\n",minWindow1("ADOBECODEBANC","ABC")) // "BANC"
    fmt.Printf("minWindow1(\"a\",\"a\") = %v\n",minWindow1("a","a")) // "a"
    fmt.Printf("minWindow1(\"a\",\"aa\") = %v\n",minWindow1("a","aa")) // ""
}
