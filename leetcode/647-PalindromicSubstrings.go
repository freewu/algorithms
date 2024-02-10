package main

import "fmt"

// 647. Palindromic Substrings
// Given a string s, return the number of palindromic substrings in it.
// A string is a palindrome when it reads the same backward as forward.
// A substring is a contiguous sequence of characters within the string.

// Example 1:
// Input: s = "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".

// Example 2:
// Input: s = "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

// Constraints:
// 		1 <= s.length <= 1000
// 		s consists of lowercase English letters.

func countSubstrings(s string) int {
	res := 0
	// 从左往右扫一遍字符串，以每个字符做轴，用中心扩散法，依次遍历计数回文子串
	countPalindrome := func (s string, left, right int) int {
		res := 0
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			left--
			right++
			res++
		}
		return res
	}

	for i := 0; i < len(s); i++ {
		res += countPalindrome(s, i, i)
		res += countPalindrome(s, i, i+1)
	}
	return res
}

func main() {
	fmt.Println(countSubstrings("abc")) // 3
	fmt.Println(countSubstrings("aaa")) // 6
}
