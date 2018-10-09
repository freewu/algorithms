package main

import (
	"fmt"
	"strings"
)

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

For example,
"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.

Note:
Have you consider that the string might be empty? This is a good question to ask during an interview.

For the purpose of this problem, we define empty string as valid palindrome.
*/
func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	var m = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 97 && s[i] <= 122) {
			m = append(m, s[i])
		}
	}
	s = string(m)
	if len(s) <= 1 {
		return true
	}
	var i = 0
	var l = len(s) - 1
	// 判断是否是回文件
	for {
		if i >= l {
			break
		}
		if s[i] != s[l] {
			return false
		}
		i++
		l--
	}

	// fmt.Println('0') // 48
	// fmt.Println('9') // 57
	// fmt.Println('a') // 97
	// fmt.Println('z') // 122

	return true
}

func main() {
	fmt.Println(isPalindrome("ab"))   // false
	fmt.Println(isPalindrome(" Abc")) // false

	fmt.Println(isPalindrome(""))                               // true
	fmt.Println(isPalindrome("a"))                              // true
	fmt.Println(isPalindrome(" Aba"))                           // true
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("12321"))                          // true
	fmt.Println(isPalindrome("123321"))                         // true
	fmt.Println(isPalindrome(".,"))                             // true
}
