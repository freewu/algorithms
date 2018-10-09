package main

import (
	"fmt"
	"math"
	"strings"
)

/*
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.
*/

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	var l = len(str)
	if l == 0 {
		return 0
	}
	var flag = 1
	var s = 0

	switch true {
	case '-' == str[0]:
		flag = -1
	case '+' == str[0]:
		s = 0
	case (str[0] >= 48 && str[0] <= 57):
		s = int(str[0]) - 48
	default:
		return 0
	}

	for i := 1; i < l; i++ {
		if str[i] > 57 || str[i] < 48 {
			break
		}
		s = s*10 + (int(str[i]) - 48)

		if s > math.MaxInt32 {
			if -1 == flag {
				s = math.MinInt32
				flag = 1
			} else {
				s = math.MaxInt32
			}
			break
		}
	}
	return flag * s
}

func main() {
	//fmt.Println('0')
	//fmt.Println('9')
	//fmt.Println('-')
	//fmt.Println('+')

	fmt.Println(myAtoi("aaaa"))
	fmt.Println(myAtoi("1a1aa"))
	fmt.Println(myAtoi("12345"))

	fmt.Println(myAtoi("-12345"))
	fmt.Println(myAtoi("+12345"))

	fmt.Println(myAtoi("   +12345"))

	fmt.Println(myAtoi("2147483648"))          // 2147483647
	fmt.Println(myAtoi("9223372036854775809")) // 2147483647
}
