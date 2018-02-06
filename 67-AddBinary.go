package main

import (
	"fmt"
)

/*
Given two binary strings, return their sum (also a binary string).

For example,
a = "11"
b = "1"
Return "100".
*/

func addBinary(a string, b string) string {
	var al = len(a) - 1
	var bl = len(b) - 1

	var s = ""
	var flag = 0 // 进位标识

	for {
		// 短的结束了
		if al < 0 || bl < 0 {
			break
		}
		if '1' == a[al] && '1' == b[bl] { // 1/1
			if 0 == flag {
				s = "0" + s
			} else {
				s = "1" + s
			}
			flag = 1
		} else if '0' == a[al] && '0' == b[bl] { // 0/0
			if 0 == flag {
				s = "0" + s
			} else {
				s = "1" + s
			}
			flag = 0
		} else { // 0/1 1/0的情况
			if 0 == flag {
				s = "1" + s
				flag = 0
			} else {
				s = "0" + s
				flag = 1
			}
		}
		al--
		bl--
	}
	// 连接其它的部分
	for {
		if al < 0 {
			break
		}
		if a[al] == '1' && 1 == flag {
			flag = 1
			s = "0" + s
		} else if a[al] == '0' && 0 == flag {
			flag = 0
			s = string(a[al]) + s
		} else {
			flag = 0
			s = "1" + s
		}
		al--
	}
	for {
		if bl < 0 {
			break
		}
		if b[bl] == '1' && 1 == flag {
			flag = 1
			s = "0" + s
		} else if b[bl] == '0' && 0 == flag {
			flag = 0
			s = string(b[bl]) + s
		} else {
			flag = 0
			s = "1" + s
		}
		bl--
	}

	if flag == 1 {
		s = "1" + s
	}

	return s
}

func main() {
	fmt.Println(addBinary("11", "1"))  // "100"
	fmt.Println(addBinary("11", "11")) // "110"
}
