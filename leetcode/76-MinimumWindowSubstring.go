package main

import (
	"fmt"
	"math"
)

/**
76. Minimum Window Substring
Given two strings s and t of lengths m and n respectively,
return the minimum window substring of s such that every character in t (including duplicates) is included in the window.
If there is no such substring, return the empty string "".
The testcases will be generated such that the answer is unique.
A substring is a contiguous sequence of characters within the string.

Constraints:

	m == s.length
	n == t.length
	1 <= m, n <= 10^5
	s and t consist of uppercase and lowercase English letters.

Follow up: Could you find an algorithm that runs in O(m + n) time?

Example 1:

	Input: s = "ADOBECODEBANC", t = "ABC"
	Output: "BANC"
	Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:

	Input: s = "a", t = "a"
	Output: "a"
	Explanation: The entire string s is the minimum window.

Example 3:

	Input: s = "a", t = "aa"
	Output: ""
	Explanation: Both 'a's from t must be included in the window.
	Since the largest window of s only has one 'a', return empty string.

解题思路:
	滑动窗口
	在窗口滑动的过程中不断的包含字符串 T，直到完全包含字符串 T 的字符以后，记下左右窗口的位置和窗口大小。
	每次都不断更新这个符合条件的窗口和窗口大小的最小值。最后输出结果即可。
*/

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var tFreq, sFreq [256]int
	result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s)+1, 0
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
func minWindowBest(s string, t string) string {
	dict := [128]int{}
	for _, ch := range t {
		dict[ch]++
	}
	counter := len(t)
	begin := 0
	end := 0
	head := 0
	d := math.MaxInt32
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
	if d == math.MaxInt32 {
		return ""
	}
	return s[head : head+d]
}

func main() {
	fmt.Printf("minWindow(\"ADOBECODEBANC\",\"ABC\") = %v\n",minWindow("ADOBECODEBANC","ABC")) // "BANC"
	fmt.Printf("minWindow(\"a\",\"a\") = %v\n",minWindow("a","a")) // "a"
	fmt.Printf("minWindow(\"a\",\"aa\") = %v\n",minWindow("a","aa")) // ""

	fmt.Printf("minWindowBest(\"ADOBECODEBANC\",\"ABC\") = %v\n",minWindowBest("ADOBECODEBANC","ABC")) // "BANC"
	fmt.Printf("minWindowBest(\"a\",\"a\") = %v\n",minWindowBest("a","a")) // "a"
	fmt.Printf("minWindowBest(\"a\",\"aa\") = %v\n",minWindowBest("a","aa")) // ""
}
