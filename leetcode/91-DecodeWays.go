package main

import (
	"fmt"
	"strconv"
)

/**
91. Decode Ways
A message containing letters from A-Z can be encoded into numbers using the following mapping:
	'A' -> "1"
	'B' -> "2"
	...
	'Z' -> "26"
To decode an encoded message,
all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways).
For example, "11106" can be mapped into:

	"AAJF" with the grouping (1 1 10 6)
	"KJF" with the grouping (11 10 6)

Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
Given a string s containing only digits, return the number of ways to decode it.
The test cases are generated so that the answer fits in a 32-bit integer.

Constraints:

	1 <= s.length <= 100
	s contains only digits and may contain leading zero(s).

Example 1:

	Input: s = "12"
	Output: 2
	Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).

Example 2:

	Input: s = "226"
	Output: 3
	Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

Example 3:

	Input: s = "06"
	Output: 0
	Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").

解题思路:
	给出一个数字字符串，题目要求把数字映射成 26 个字母，映射以后问有多少种可能的翻译方法。
	DP。
 */

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1) // 用一个数组来保存合结果
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' { // 判断当前字符 是否为 不是 0
			dp[i] += dp[i-1]
		}
		// 如果前两个字符是在1-26之间说明还有组合
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
		fmt.Printf("%v = %v\n",i,dp)
	}
	return dp[n]
}

// best solution
func numDecodingsBest(s string) int {
	dp := make(map[int]int)
	return dfs(s, 0, dp)
}

func dfs(s string, start int, dp map[int]int) int {
	if start == len(s) {
		return 1
	}
	if _,ok := dp[start]; ok {
		return dp[start]
	}
	res := 0
	if s[start] == '0' {
		return 0
	}
	res += dfs(s, start+1, dp)
	if start+2 <= len(s) {
		num,err := strconv.Atoi(s[start:start+2])
		if err == nil && num <= 26 {
			res += dfs(s, start+2, dp)
		}
	}
	dp[start] = res
	return res

}

func main() {
	fmt.Printf("numDecodings(\"12\") = %v\n",numDecodings("12")) // 2   "AB" (1 2) or "L" (12)
	fmt.Printf("numDecodings(\"226\") = %v\n",numDecodings("226")) // 3  "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
	fmt.Printf("numDecodings(\"06\") = %v\n",numDecodings("06")) // 0

	fmt.Printf("numDecodingsBest(\"12\") = %v\n",numDecodingsBest("12")) // 2   "AB" (1 2) or "L" (12)
	fmt.Printf("numDecodingsBest(\"226\") = %v\n",numDecodingsBest("226")) // 3  "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
	fmt.Printf("numDecodingsBest(\"06\") = %v\n",numDecodingsBest("06")) // 0
}
