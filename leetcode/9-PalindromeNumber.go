package main

/*
Determine whether an integer is a palindrome. Do this without extra space.

click to show spoilers.

Some hints:
Could negative integers be palindromes? (ie, -1)

If you are thinking of converting the integer to string, note the restriction of using extra space.

You could also try reversing an integer. However, if you have solved the problem "Reverse Integer", you know that the reversed integer might overflow. How would you handle such case?

There is a more generic way of solving this problem.
*/

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// 负数和个位数不可能是回文
	if x < 10 {
		return false
	}

	var t = x
	var s int64 = 0
	for {
		s = s * 10 + int64(x % 10)
		x /= 10;
		if x == 0 {
			break
		}
	}
	return int(s) == t
}

func main() {
	fmt.Println(isPalindrome(-12321))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(123211))
}