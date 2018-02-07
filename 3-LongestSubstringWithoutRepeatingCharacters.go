package main

import (
	"fmt"
	"strings"
)

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	if 0 == len(s) {
		return 0
	}
	var l = 0
	var sl = 0
	var m = ""
	for i := 0; i < len(s); i++ {
		if strings.Index(m, string(s[i])) == -1 {
			l++
			m += string(s[i])
		} else {
			if sl < l {
				sl = l
			}
			// 返回到s[i]之后开始的字符串
			var t = s[i]
			for {
				i--
				if t == s[i] {
					m = string(s[i+1])
					i++
					break
				}
			}
			l = 1
		}
	}
	// 如果最后一段字符是最长的
	if sl < l {
		sl = l
	}
	return sl
}

// wrong solution
func lengthOfLongestSubstring1(s string) int {
	var l = len(s)

	if 0 == l {
		return 0
	}

	var m = 0
	var t = 0

	for i := 0; i < l; i++ {
		t = 1
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				t = 1
			} else {
				t++
			}
			if t > m {
				m = t
			}
		}
	}
	if t > m {
		m = t
	}
	return m
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
	fmt.Println(lengthOfLongestSubstring("aac"))      // 2
	fmt.Println(lengthOfLongestSubstring("abc"))      // 3
	fmt.Println(lengthOfLongestSubstring(""))         // 0
	fmt.Println(lengthOfLongestSubstring("a"))        // 1
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}
