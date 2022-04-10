package main

import "fmt"

/**
97. Interleaving String
Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.
An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:

	s = s1 + s2 + ... + sn
	t = t1 + t2 + ... + tm
	|n - m| <= 1
The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
Note: a + b is the concatenation of strings a and b.

Constraints:

	0 <= s1.length, s2.length <= 100
	0 <= s3.length <= 200
	s1, s2, and s3 consist of lowercase English letters.

Example 1:

	Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
	Output: true

Example 2:

	Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
	Output: false

Example 3:

	Input: s1 = "", s2 = "", s3 = ""
	Output: true

解题思路:
	dfs
	记录 s1 和 s2 串当前比较的位置 p1 和 p2。
	如果 s3[p1+p2] 的位置上等于 s1[p1] 或者 s2[p2] 代表能匹配上，那么继续往后移动 p1 和 p2 相应的位置。
	因为是交错字符串，所以判断匹配的位置是 s3[p1+p2] 的位置。

 */

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	visited := make(map[int]bool)
	return dfs(s1, s2, s3, 0, 0, visited)
}

func dfs(s1, s2, s3 string, p1, p2 int, visited map[int]bool) bool {
	if p1+p2 == len(s3) {
		return true
	}
	if _, ok := visited[(p1*len(s3))+p2]; ok {
		return false
	}
	visited[(p1*len(s3))+p2] = true
	var match1, match2 bool
	if p1 < len(s1) && s3[p1+p2] == s1[p1] {
		match1 = true
	}
	if p2 < len(s2) && s3[p1+p2] == s2[p2] {
		match2 = true
	}
	if match1 && match2 {
		return dfs(s1, s2, s3, p1+1, p2, visited) || dfs(s1, s2, s3, p1, p2+1, visited)
	} else if match1 {
		return dfs(s1, s2, s3, p1+1, p2, visited)
	} else if match2 {
		return dfs(s1, s2, s3, p1, p2+1, visited)
	} else {
		return false
	}
}

// best solution
func isInterleaveBest(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] && s3[j-1] == s2[j-1]
				continue
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j] && s3[i-1] == s1[i-1]
				continue
			}
			dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1] || dp[i][j-1] && s2[j-1] == s3[i+j-1]
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	fmt.Printf("isInterleave(\"aabcc\",\"dbbca\",\"aadbbcbcac\") = %v\n",isInterleave("aabcc","dbbca","aadbbcbcac")) // true
	fmt.Printf("isInterleave(\"aabcc\",\"dbbca\",\"aadbbbaccc\") = %v\n",isInterleave("aabcc","dbbca","aadbbbaccc")) // false
	fmt.Printf("isInterleave(\"\",\"\",\"\") = %v\n",isInterleave("","","")) // true

	fmt.Printf("isInterleaveBest(\"aabcc\",\"dbbca\",\"aadbbcbcac\") = %v\n",isInterleaveBest("aabcc","dbbca","aadbbcbcac")) // true
	fmt.Printf("isInterleaveBest(\"aabcc\",\"dbbca\",\"aadbbbaccc\") = %v\n",isInterleaveBest("aabcc","dbbca","aadbbbaccc")) // false
	fmt.Printf("isInterleaveBest(\"\",\"\",\"\") = %v\n",isInterleaveBest("","","")) // true
}
