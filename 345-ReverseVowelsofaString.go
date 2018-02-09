package main

import (
	"fmt"
)

/*
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:
Given s = "hello", return "holle".

Example 2:
Given s = "leetcode", return "leotcede".

Note:
The vowels does not include the letter "y".
*/

func reverseVowels(s string) string {
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
		// find the vowel from begin
		if 'a' == s[i] || 'e' == s[i] || 'i' == s[i] || 'o' == s[i] || 'u' == s[i] ||
			'A' == s[i] || 'E' == s[i] || 'I' == s[i] || 'O' == s[i] || 'U' == s[i] {
		} else {
			t[i] = s[i]
			i++
			continue
		}
		// find the vowel from end
		if 'a' == s[l] || 'e' == s[l] || 'i' == s[l] || 'o' == s[l] || 'u' == s[l] ||
			'A' == s[l] || 'E' == s[l] || 'I' == s[l] || 'O' == s[l] || 'U' == s[l] {
		} else {
			t[l] = s[l]
			l--
			continue
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
	fmt.Println(reverseVowels("aA"))    // Aa
	fmt.Println(reverseVowels("abc"))   // abc
	fmt.Println(reverseVowels("ooee"))  // ee
	fmt.Println(reverseVowels("hello")) // holle
}
