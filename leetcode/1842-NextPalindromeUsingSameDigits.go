package main

// 1842. Next Palindrome Using Same Digits
// You are given a numeric string num, representing a very large palindrome.

// Return the smallest palindrome larger than num that can be created by rearranging its digits. 
// If no such palindrome exists, return an empty string "".

// A palindrome is a number that reads the same backward as forward.

// Example 1:
// Input: num = "1221"
// Output: "2112"
// Explanation: The next palindrome larger than "1221" is "2112".

// Example 2:
// Input: num = "32123"
// Output: ""
// Explanation: No palindromes larger than "32123" can be made by rearranging the digits.

// Example 3:
// Input: num = "45544554"
// Output: "54455445"
// Explanation: The next palindrome larger than "45544554" is "54455445".

// Constraints:
//     1 <= num.length <= 10^5
//     num is a palindrome.

import "fmt"

func nextPalindrome(num string) string {
    s := []byte(num)
    nextPermutation := func(arr []byte) bool {
        n := len(arr)
        i := n - 2
        for i >= 0 && arr[i] >= arr[i+1] {
            i--
        }
        if i < 0 { return false }
        j := n - 1
        for j >= 0 && arr[i] >= arr[j] {
            j--
        }
        arr[i], arr[j] = arr[j], arr[i]
        b := arr[i+1:]
        for i, m := 0, len(b); i < m/2; i++ {
            b[i], b[m-1-i] = b[m-1-i], b[i]
        }
        return true
    }
    if !nextPermutation(s[:len(s)/2]) { return "" }
    for i, n := 0, len(s); i < n/2; i++ {
        s[n-1-i] = s[i]
    }
    return string(s)
}

func main() {
    // Example 1:
    // Input: num = "1221"
    // Output: "2112"
    // Explanation: The next palindrome larger than "1221" is "2112".
    fmt.Println(nextPalindrome("1221")) // "2112"
    // Example 2:
    // Input: num = "32123"
    // Output: ""
    // Explanation: No palindromes larger than "32123" can be made by rearranging the digits.
    fmt.Println(nextPalindrome("32123")) // ""
    // Example 3:
    // Input: num = "45544554"
    // Output: "54455445"
    // Explanation: The next palindrome larger than "45544554" is "54455445".
    fmt.Println(nextPalindrome("45544554")) // "54455445"
}