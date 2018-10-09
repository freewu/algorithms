package main

/*
Write a function to find the longest common prefix string amongst an array of strings.
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

func main () {
	var d = []string {
		"xxxxxx",
		"xxxxccc",
		"xxabva",
	}
	fmt.Println(longestCommonPrefix(d))
}