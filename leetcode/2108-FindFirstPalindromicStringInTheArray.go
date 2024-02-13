package main

import "fmt"

// 2108. Find First Palindromic String in the Array
// Given an array of strings words, return the first palindromic string in the array. 
// If there is no such string, return an empty string "".

// A string is palindromic if it reads the same forward and backward.

// Example 1:
// Input: words = ["abc","car","ada","racecar","cool"]
// Output: "ada"
// Explanation: The first string that is palindromic is "ada".
// Note that "racecar" is also palindromic, but it is not the first.

// Example 2:
// Input: words = ["notapalindrome","racecar"]
// Output: "racecar"
// Explanation: The first and only string that is palindromic is "racecar".

// Example 3:
// Input: words = ["def","ghi"]
// Output: ""
// Explanation: There are no palindromic strings, so the empty string is returned.
 
// Constraints:
// 		1 <= words.length <= 100
// 		1 <= words[i].length <= 100
// 		words[i] consists only of lowercase English letters.

func firstPalindrome(words []string) string {
	// 判断是否为回文
    isPalindrome := func(word string) bool {
		i := 0
		j := len(word) - 1
		for ; i < j ;  {
			if word[i] != word[j] {
				return false
			}
			i = i + 1
			j = j - 1 
		}
		return true
	}
	for i := 0; i < len(words); i = i + 1 {
		//  找到第一个回文就返回
		if(isPalindrome(words[i])) {
			return words[i]
		}
	}
	return ""
}

func main() {
	fmt.Println(firstPalindrome([]string{"abc","car","ada","racecar","cool"})) // ada
	fmt.Println(firstPalindrome([]string{"notapalindrome","racecar"})) // racecar
	fmt.Println(firstPalindrome([]string{"def","ghi"})) // ""
}
