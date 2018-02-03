package main

import (
	"fmt"
	"strings"
)

/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/

func lengthOfLastWord1(s string) int {
	var w = 0
	s = strings.TrimSpace(s)
	for i := 0; i < len(s); i++ {
		if ' ' == s[i] {
			w = 0
			continue // 不执行到下面了
		}
		w++
	}
	return w
}

// without trim
func lengthOfLastWord(s string) int {
	var w = 0
	var m = 0
	for i := 0; i < len(s); i++ {
		if ' ' == s[i] {
			w = 0
		} else {
			w++
			m = w
		}
	}
	return m
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))  // 5
	fmt.Println(lengthOfLastWord("Hello World ")) // 5
	fmt.Println(lengthOfLastWord(" "))            // 0
	fmt.Println(lengthOfLastWord("a"))            // 1
	fmt.Println(lengthOfLastWord("a33"))          // 3
}
