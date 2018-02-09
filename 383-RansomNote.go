package main

import (
	"fmt"
)

/*
Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
*/

func canConstruct(ransomNote string, magazine string) bool {
	return false
}

func main() {
	fmt.Println(canConstruct("abc", "abcc")) // true
	fmt.Println(canConstruct("", "abcc"))    // true
	fmt.Println(canConstruct("abc", ""))     // false
	fmt.Println(canConstruct("a", "b"))      // false
	fmt.Println(canConstruct("ab", "ba"))    // true
	fmt.Println(canConstruct("ab", "bac"))   // true
	fmt.Println(canConstruct("cab", "ba"))   // false
}
