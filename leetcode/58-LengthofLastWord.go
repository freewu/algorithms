package main

import (
	"fmt"
	"strings"
)

/*
58. Length of Last Word
Given a string s consisting of some words separated by some number of spaces, return the length of the last word in the string.
A word is a maximal substring consisting of non-space characters only.

Constraints:

	1 <= s.length <= 10^4
	s consists of only English letters and spaces ' '.
	There will be at least one word in s.


Example 1:

	Input: s = "Hello World"
	Output: 5
	Explanation: The last word is "World" with length 5.

Example 2:

	Input: s = "   fly me   to   the moon  "
	Output: 4
	Explanation: The last word is "moon" with length 4.

Example 3:

	Input: s = "luffy is still joyboy"
	Output: 6
	Explanation: The last word is "joyboy" with length 6.

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

func lengthOfLastWord2(s string) int {
	last := len(s) - 1
	for last >= 0 && s[last] == ' ' {
		last--
	}
	if last < 0 {
		return 0
	}
	first := last
	for first >= 0 && s[first] != ' ' {
		first--
	}
	return last - first
}

// best solution
func lengthOfLastWordBest(s string) int {
	s = strings.TrimSpace(s)
	if "" == s {
		return 0
	}
	parts := strings.Fields(s)
	return len(parts[len(parts) - 1])
}

func main() {
	fmt.Printf("lengthOfLastWord(\"Hello World\") = %v\n",lengthOfLastWord("Hello World"))  // 5
	fmt.Printf("lengthOfLastWord(\"Hello World \") = %v\n",lengthOfLastWord("Hello World "))  // 5
	fmt.Printf("lengthOfLastWord(\" \") = %v\n",lengthOfLastWord(" "))  // 0
	fmt.Printf("lengthOfLastWord(\"a\") = %v\n",lengthOfLastWord("a"))  // 1
	fmt.Printf("lengthOfLastWord(\"a33\") = %v\n",lengthOfLastWord("a33"))  // 3

	fmt.Printf("lengthOfLastWord1(\"Hello World\") = %v\n",lengthOfLastWord1("Hello World"))  // 5
	fmt.Printf("lengthOfLastWord1(\"Hello World \") = %v\n",lengthOfLastWord1("Hello World "))  // 5
	fmt.Printf("lengthOfLastWord1(\" \") = %v\n",lengthOfLastWord1(" "))  // 0
	fmt.Printf("lengthOfLastWord1(\"a\") = %v\n",lengthOfLastWord1("a"))  // 1
	fmt.Printf("lengthOfLastWord1(\"a33\") = %v\n",lengthOfLastWord1("a33"))  // 3

	fmt.Printf("lengthOfLastWord2(\"Hello World\") = %v\n",lengthOfLastWord2("Hello World"))  // 5
	fmt.Printf("lengthOfLastWord2(\"Hello World \") = %v\n",lengthOfLastWord2("Hello World "))  // 5
	fmt.Printf("lengthOfLastWord2(\" \") = %v\n",lengthOfLastWord2(" "))  // 0
	fmt.Printf("lengthOfLastWord2(\"a\") = %v\n",lengthOfLastWord2("a"))  // 1
	fmt.Printf("lengthOfLastWord2(\"a33\") = %v\n",lengthOfLastWord2("a33"))  // 3

	fmt.Printf("lengthOfLastWordBest(\"Hello World\") = %v\n",lengthOfLastWordBest("Hello World"))  // 5
	fmt.Printf("lengthOfLastWordBest(\"Hello World \") = %v\n",lengthOfLastWordBest("Hello World "))  // 5
	fmt.Printf("lengthOfLastWordBest(\" \") = %v\n",lengthOfLastWordBest(" "))  // 0
	fmt.Printf("lengthOfLastWordBest(\"a\") = %v\n",lengthOfLastWordBest("a"))  // 1
	fmt.Printf("lengthOfLastWordBest(\"a33\") = %v\n",lengthOfLastWordBest("a33"))  // 3
}
