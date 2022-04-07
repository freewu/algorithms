package main

import (
	"fmt"
	"regexp"
)

/**
65. Valid Number
A valid number can be split up into these components (in order):

	A decimal number or an integer.
	(Optional) An 'e' or 'E', followed by an integer.

A decimal number can be split up into these components (in order):

	(Optional) A sign character (either '+' or '-').
	One of the following formats:
		One or more digits, followed by a dot '.'.
		One or more digits, followed by a dot '.', followed by one or more digits.
		A dot '.', followed by one or more digits.

An integer can be split up into these components (in order):

	(Optional) A sign character (either '+' or '-').
	One or more digits.

For example, all the following are valid numbers:
["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"],
while the following are not valid numbers: ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"].

Given a string s, return true if s is a valid number.


Example 1:

	Input: s = "0"
	Output: true

Example 2:

	Input: s = "e"
	Output: false

Example 3:

	Input: s = "."
	Output: false

Constraints:

	1 <= s.length <= 20
	s consists of only English letters (both uppercase and lowercase), digits (0-9), plus '+', minus '-', or dot '.'.
 */

func isNumber(s string) bool {
	numFlag, dotFlag, eFlag := false, false, false
	for i := 0; i < len(s); i++ {
		// 如果是数字，则标记数字出现过
		if '0' <= s[i] && s[i] <= '9' {
			numFlag = true
		// 如果是 ‘.’, 则需要 ‘.‘没有出现过，并且 ‘e/E’ 没有出现过，才会进行标记
		} else if s[i] == '.' && !dotFlag && !eFlag {
			dotFlag = true
		// 如果是 ‘e/E’, 则需要 ‘e/E’没有出现过，并且前面出现过数字，才会进行标记
		} else if (s[i] == 'e' || s[i] == 'E') && !eFlag && numFlag {
			eFlag = true
			numFlag = false // reJudge integer after 'e' or 'E'
		// 如果是 ‘+/-’, 则需要是第一个字符，或者前一个字符是 ‘e/E’，才会进行标记，并重置数字出现的标识
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			continue
		} else {
			return false
		}
	}
	// 最后返回时，需要字符串中至少出现过数字，避免下列case: s == ‘.’ or ‘e/E’ or ‘+/e’ and etc…
	return numFlag
}


var validNum = regexp.MustCompile(`^(([+-]?[0-9]+)|([+-]?[0-9]+\.[0-9]*)|([+-]?\.[0-9]+))([eE][+-]?[0-9]+)?$`)
// best solution
func isNumberBest(s string) bool {
	return validNum.MatchString(s)
}

func main() {
	fmt.Printf("isNumber(\"0\") = %v\n",isNumber("0")) // true
	fmt.Printf("isNumber(\"e\") = %v\n",isNumber("e")) // false
	fmt.Printf("isNumber(\".\") = %v\n",isNumber(",")) // false
	fmt.Printf("isNumber(\"1.e\") = %v\n",isNumber("1.e")) // false

	fmt.Printf("isNumberBest(\"0\") = %v\n",isNumberBest("0")) // true
	fmt.Printf("isNumberBest(\"e\") = %v\n",isNumberBest("e")) // false
	fmt.Printf("isNumberBest(\".\") = %v\n",isNumberBest(",")) // false
	fmt.Printf("isNumberBest(\"1.e\") = %v\n",isNumberBest("1.e")) // false
}
