package main

import (
	"fmt"
)

/*
Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".
*/

func reverseString(s string) string {
	var l = len(s)
	if l <= 1 {
		return s
	}
	var i = 0
	var t = make([]byte, l)
	l = l - 1
	for {
		if i > l {
			break
		}
		// syntax sugar
		t[l], t[i] = s[i], s[l]
		//t[l] = s[i]
		//t[i] = s[l]
		i++
		l--
	}
	return string(t)
}

func main() {
	fmt.Println(reverseString("ab"))    // ba
	fmt.Println(reverseString("123"))   // 321
	fmt.Println(reverseString("hello")) // olleh
	fmt.Println(reverseString(""))      //
}
