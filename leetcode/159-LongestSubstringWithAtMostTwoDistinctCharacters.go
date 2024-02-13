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
// 	1 <= s.length <= 105
// 	s consists of English letters.

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	m := make(map[byte]int,58)
	start, end := 0, 0
	counter, l := 0, 0

	max := func (a, b int) int {
		if a <= b {
			return b
		}
		return a
	}

	for end < len(s) {
		if _, ok := m[s[end]]; ok {
			m[s[end]] += 1
		} else {
			m[s[end]] = 1
		}
		if m[s[end]] == 1 {
			counter += 1 // 出现新的字符了
		}
		for counter > 2  {
			m[s[start]] -= 1
			if m[s[start]] == 0 {
				counter -= 1
			}
			start += 1
		}
		l = max(end - start + 1, l)
		end += 1
	} 
	return l
}

const CHARSET int = 58
// best solution
func lengthOfLongestSubstringTwoDistinct1(s string) int {
	n := len(s)
	cnts := make([]int, CHARSET)
	diffCnt := 0
	i, j := 0, 0
	ans := 0
	for ; j < n; j++ {
		if cnts[s[j]-'A'] == 0 {
			diffCnt++
		}
		cnts[s[j]-'A']++
		for diffCnt > 2 {
			if cnts[s[i]-'A'] == 1 {
				diffCnt--
			}
			cnts[s[i]-'A']--
			i++
		}
		ans = max(ans, j-i+1)
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba")) // 3
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb")) // 5

	fmt.Println(lengthOfLongestSubstringTwoDistinct1("eceba")) // 3
	fmt.Println(lengthOfLongestSubstringTwoDistinct1("ccaabbb")) // 5
}