package main

import (
	"fmt"
)

/**
  Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

  Example 1:

  Input: "babad"
  Output: "bab"
  Note: "aba" is also a valid answer.
  Example 2:

  Input: "cbbd"
  Output: "bb"
*/

func longestPalindrome(s string) string {
	var len = len(s)
	if len <= 1 {
		return s
	}
	var maxlen = 1
	var begin = 0

	// 创建一个len * len的数组
	//var arr [len][len]int
	//var arr = make([][]int, len, len)
	var arr [][]int
	for i := 0; i < len; i++ {
		tmp := make([]int, len)
		arr = append(arr, tmp)
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if i == j {
				arr[i][j] = 1
			} else {
				arr[i][j] = 0
			}
		}
	}

	fmt.Println(arr)
	fmt.Println(maxlen)
	fmt.Println(begin)

	return s
}

func main() {
	fmt.Println(longestPalindrome("1234"))
}
