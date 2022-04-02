package main

/*
14. Longest Common Prefix
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string “”.

Constraints:

	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] consists of only lower-case English letters.

Example 1:

	Input: strs = ["flower","flow","flight"]
	Output: "fl"

Example 2:

	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.

*/
import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	// 如果为空
	if 0 == len(strs) {
		return ""
	}
	// 如果只是有一个子串
	if 1 == len(strs) {
		return strs[0]
	}
	var sl = len(strs[0])
	// 获得最短的长度
	for i := 1; i < len(strs); i++ {
		var l = len(strs[i]);
		if l < sl {
			sl = l
		}
	}
	var s = ""
	var al = len(strs)

	for i := 0; i < sl; i++ {
		for j := 1; j < al; j++ {
			if strs[j][i] != strs[j - 1][i] {
				return s
			}
		}
		s += string(strs[0][i])
	}
	return s
}

func longestCommonPrefix1(strs []string) string {
	prefix := strs[0] // 取第一个字符串
	for i := 1; i < len(strs); i++ { // 循环数组其它的字符
		for j := 0; j < len(prefix); j++ { // 循环第一了符串
			if len(strs[i]) <= j || strs[i][j] != prefix[j] { // 如果出现不相等
				prefix = prefix[0:j] // 截取保留公共长度 推出循环
				break
			}
		}
	}
	return prefix
}

func main () {
	var d = []string {
		"xxxxxx",
		"xxxxccc",
		"xxabva",
	}
	fmt.Printf("longestCommonPrefix(d) = %v\n",longestCommonPrefix(d)) // xx
	fmt.Printf("longestCommonPrefix1(d) = %v\n",longestCommonPrefix1(d)) // xx

	fmt.Printf("longestCommonPrefix([]string{\"flower\",\"flow\",\"flight\"}) = %v\n",longestCommonPrefix([]string{"flower","flow","flight"})) // fl
	fmt.Printf("longestCommonPrefix1([]string{\"flower\",\"flow\",\"flight\"}) = %v\n",longestCommonPrefix1([]string{"flower","flow","flight"})) // fl

	fmt.Printf("longestCommonPrefix([]string{\"dog\",\"racecar\",\"car\"}) = %v\n",longestCommonPrefix([]string{"dog","racecar","car"})) // ""
	fmt.Printf("longestCommonPrefix1([]string{\"dog\",\"racecar\",\"car\"}) = %v\n",longestCommonPrefix1([]string{"dog","racecar","car"})) // ""
}