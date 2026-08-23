package main

// 4030. Check ASCII Palindromic
// You are given a string s consisting of lowercase English letters.

// Construct a binary string by replacing each character in s with the 8-bit binary representation of its ASCII value, 
// including leading zeros, while preserving the original order of the characters.

// Return true if the resulting binary string is a palindrome. Otherwise, return false.

// A binary string is a string which contains only the characters '0' and '1'.

// A palindrome is a string that reads the same forward and backward.

// Example 1:
// Input: s = "ff"
// Output: true
// Explanation:
// The ASCII value of f is 102, whose 8-bit binary representation is 01100110.
// Thus, the binary string is 0110011001100110.
// Since this binary string is a palindrome, the output is true.

// Example 2:
// Input: s = "leet"
// Output: false
// Explanation:
// The ASCII values of l, e, e, and t are 108, 101, 101, and 116, respectively.
// Their 8-bit binary representations are 01101100, 01100101, 01100101, and 01110100.
// Thus, the binary string is 01101100011001010110010101110100.
// Since this binary string is not a palindrome, the output is false.

// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters.

import "fmt"

func isPalindromic(s string) bool {
    bin := ""
    // 1. 遍历字符串s，将每个字符的ASCII值转换为8位二进制表示
    // 2. 将所有二进制表示拼接起来，得到一个新的字符串
    // 3. 检查新字符串是否为回文
    for _, c := range s {
        bin += fmt.Sprintf("%08b", c)
    }
    // 4. 检查新字符串是否为回文
    for i := 0; i < len(bin)/2; i++ {
        if bin[i] != bin[len(bin) - 1 - i] {
            return false
        }
    }
    return true
}

func isPalindromic1(s string) bool {
    sum := len(s) * 8
    for i := 0; i < sum/2; i++ {
        j := sum - 1 - i
        ch := s[i/8]
        bi := (ch >> (7 - (i % 8))) & 1
        c := s[j/8]
        b := (c >> (7 - (j % 8))) & 1
        if b != bi {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "ff"
    // Output: true
    // Explanation:
    // The ASCII value of f is 102, whose 8-bit binary representation is 01100110.
    // Thus, the binary string is 0110011001100110.
    // Since this binary string is a palindrome, the output is true.
    fmt.Println(isPalindromic("ff")) // true
    // Example 2:
    // Input: s = "leet"
    // Output: false
    // Explanation:
    // The ASCII values of l, e, e, and t are 108, 101, 101, and 116, respectively.
    // Their 8-bit binary representations are 01101100, 01100101, 01100101, and 01110100.
    // Thus, the binary string is 01101100011001010110010101110100.
    // Since this binary string is not a palindrome, the output is false.  
    fmt.Println(isPalindromic("leet")) // false

    fmt.Println(isPalindromic("leetcode")) // false
    fmt.Println(isPalindromic("bluefrog")) // false
    fmt.Println(isPalindromic("freewu")) // false

    fmt.Println(isPalindromic1("ff")) // true 
    fmt.Println(isPalindromic1("leet")) // false
    fmt.Println(isPalindromic1("leetcode")) // false
    fmt.Println(isPalindromic1("bluefrog")) // false
    fmt.Println(isPalindromic1("freewu")) // false
}