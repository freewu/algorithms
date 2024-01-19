package main

import (
	"fmt"
	"strings"
)

// 125. Valid Palindrome
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, 
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.


// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.
 

// Constraints:

// 		1 <= s.length <= 2 * 105
// 		s consists only of printable ASCII characters.

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	// 新建一个只有 0-9 a-z的数组
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
		// 从 i -> <- l 判断是否相符
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

func isPalindrome1(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	// 如果 -> <- 相交了说明判断完，说明输入的是回文
	for i < j {
		// i -> 遇到特殊字符跳过
		for i < j && !isChar(s[i]) {
			i++
		}
		// <- j 遇到特殊字符跳过
		for i < j && !isChar(s[j]) {
			j--
		}
		// 判断收尾是否是相同字符，不是返回false 
		if s[i] != s[j] {
			return false
		}
		// 相符则继续向内收 i-> <-j
		i++
		j--
	}
	return true
}

// 判断 c 是否是字符或者数字
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
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


	fmt.Println(isPalindrome1("ab"))   // false
	fmt.Println(isPalindrome1(" Abc")) // false

	fmt.Println(isPalindrome1(""))                               // true
	fmt.Println(isPalindrome1("a"))                              // true
	fmt.Println(isPalindrome1(" Aba"))                           // true
	fmt.Println(isPalindrome1("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome1("12321"))                          // true
	fmt.Println(isPalindrome1("123321"))                         // true
	fmt.Println(isPalindrome1(".,"))                             // true
}
