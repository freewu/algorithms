package main

// 1616. Split Two Strings to Make Palindrome
// You are given two strings a and b of the same length. 
// Choose an index and split both strings at the same index, 
// splitting a into two strings: aprefix and asuffix where a = aprefix + asuffix, 
// and splitting b into two strings: bprefix and bsuffix where b = bprefix + bsuffix. 
// Check if aprefix + bsuffix or bprefix + asuffix forms a palindrome.

// When you split a string s into sprefix and ssuffix, either ssuffix or sprefix is allowed to be empty. 
// For example, if s = "abc", then "" + "abc", "a" + "bc", "ab" + "c" , and "abc" + "" are valid splits.

// Return true if it is possible to form a palindrome string, otherwise return false.

// Notice that x + y denotes the concatenation of strings x and y.

// Example 1:
// Input: a = "x", b = "y"
// Output: true
// Explaination: If either a or b are palindromes the answer is true since you can split in the following way:
// aprefix = "", asuffix = "x"
// bprefix = "", bsuffix = "y"
// Then, aprefix + bsuffix = "" + "y" = "y", which is a palindrome.

// Example 2:
// Input: a = "xbdef", b = "xecab"
// Output: false

// Example 3:
// Input: a = "ulacfd", b = "jizalu"
// Output: true
// Explaination: Split them at index 3:
// aprefix = "ula", asuffix = "cfd"
// bprefix = "jiz", bsuffix = "alu"
// Then, aprefix + bsuffix = "ula" + "alu" = "ulaalu", which is a palindrome.

// Constraints:
//     1 <= a.length, b.length <= 10^5
//     a.length == b.length
//     a and b consist of lowercase English letters

import "fmt"

func checkPalindromeFormation(a string, b string) bool {
    checkPalindrome := func(s string) bool {
        start, end := 0, len(s) - 1 
        for start < end {
            if s[start] != s[end] {
                return false
            }
            start++
            end--
        }
        return true
    }
    check := func(a string, b string) bool {
        n := len(a)
        start, end := 0, n - 1
        for start < end && a[start] == b[end] {
            start++
            end--
        }
        // The prefix and suffix characters have been matched till start in a and end in b. 
        // Now, if the remainder of the string from start to end is palindrome in either a or b, the concat will be a palindrome, splitting at either start or end 
        return checkPalindrome(a[start:end + 1]) || checkPalindrome(b[start:end + 1])
    }
    return check(a, b) || check(b, a)
}

func main() {
    // Example 1:
    // Input: a = "x", b = "y"
    // Output: true
    // Explaination: If either a or b are palindromes the answer is true since you can split in the following way:
    // aprefix = "", asuffix = "x"
    // bprefix = "", bsuffix = "y"
    // Then, aprefix + bsuffix = "" + "y" = "y", which is a palindrome.
    fmt.Println(checkPalindromeFormation("x", "y")) // true
    // Example 2:
    // Input: a = "xbdef", b = "xecab"
    // Output: false
    fmt.Println(checkPalindromeFormation("xbdef", "xecab")) // false
    // Example 3:
    // Input: a = "ulacfd", b = "jizalu"
    // Output: true
    // Explaination: Split them at index 3:
    // aprefix = "ula", asuffix = "cfd"
    // bprefix = "jiz", bsuffix = "alu"
    // Then, aprefix + bsuffix = "ula" + "alu" = "ulaalu", which is a palindrome.
    fmt.Println(checkPalindromeFormation("ulacfd", "jizalu")) // true
}