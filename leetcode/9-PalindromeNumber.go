package main

// 9. Palindrome Number
// Given an integer x, return true if x is a palindrome, and false otherwise.

// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 
// Constraints:    
//     -2^31 <= x <= 2^31 - 1
 
import "fmt"

func isPalindrome(x int) bool {
    if x < 0 { // 负数不可能是回文
        return false
    }
    t, s := x, 0
    for {
        s = s * 10 + (x % 10)
        x /= 10;
        if x == 0 {
            break
        }
    }
    return s == t
}

func main() {
    // Explanation: 121 reads as 121 from left to right and from right to left.
    fmt.Println(isPalindrome(121)) // true
    // Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
    fmt.Println(isPalindrome(-121)) // false
    // Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
    fmt.Println(isPalindrome(10)) // false
 
    fmt.Println(isPalindrome(-12321)) // false
    fmt.Println(isPalindrome(1))
    fmt.Println(isPalindrome(12321))
    fmt.Println(isPalindrome(123211))
}