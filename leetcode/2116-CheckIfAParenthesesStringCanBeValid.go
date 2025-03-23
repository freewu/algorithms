package main

// 2116. Check if a Parentheses String Can Be Valid
// A parentheses string is a non-empty string consisting only of '(' and ')'. 
// It is valid if any of the following conditions is true:
//     1. It is ().
//     2. It can be written as AB (A concatenated with B), where A and B are valid parentheses strings.
//     3. It can be written as (A), where A is a valid parentheses string.

// You are given a parentheses string s and a string locked, both of length n. 
// locked is a binary string consisting only of '0's and '1's. For each index i of locked,
//     1. If locked[i] is '1', you cannot change s[i].
//     2. But if locked[i] is '0', you can change s[i] to either '(' or ')'.

// Return true if you can make s a valid parentheses string. Otherwise, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/06/eg1.png" />
// Input: s = "))()))", locked = "010100"
// Output: true
// Explanation: locked[1] == '1' and locked[3] == '1', so we cannot change s[1] or s[3].
// We change s[0] and s[4] to '(' while leaving s[2] and s[5] unchanged to make s valid.

// Example 2:
// Input: s = "()()", locked = "0000"
// Output: true
// Explanation: We do not need to make any changes because s is already valid.

// Example 3:
// Input: s = ")", locked = "0"
// Output: false
// Explanation: locked permits us to change s[0]. 
// Changing s[0] to either '(' or ')' will not make s valid.

// Constraints:
//     n == s.length == locked.length
//     1 <= n <= 10^5
//     s[i] is either '(' or ')'.
//     locked[i] is either '0' or '1'.

import "fmt"

func canBeValid(s string, locked string) bool {
    if len(s) % 2 == 1 { return false }
    balance, soft := 0, 0
    for i := 0; i < len(s); i++ {
        if locked[i] == '0' {
            // Initially treat each unlocked char as a left paren ('(').
            if 2 * soft < balance { soft++ } // Count as soft only if bal is high enough.
            balance++
        } else if s[i] == '(' {
            balance++
        } else {
            if balance == 0 { return false } // It's a locked ')'.
            balance--
            if 2*soft > balance { soft-- }
        }
    }
    if balance > 0 && 2*soft < balance { 
        return false
    }
    return true
}

func canBeValid1(s string, locked string) bool {
    if len(s) % 2 == 1 { return false }
    x := 0
    for i, ch := range s {
        if ch == '(' || locked[i] == '0' {
            x++
        } else if x > 0 {
            x--
        } else {
            return false
        }
    }
    x = 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ')' || locked[i] == '0' {
            x++
        } else if x > 0 {
            x--
        } else {
            return false
        }
    }
    return true
}

func canBeValid2(s string, locked string) bool {
    n := len(s)
    if n & 1 == 1 { return false }
    left, right := 0, 0
    for j := 0; j < n; j++ {
        if s[j] == '(' || locked[j] == '0' {
            left++
        } else {
            left--
        }
        if left < 0 { return false }
    }
    if left == 0 { return true }
    for j := n - 1; j >= 0; j-- {
        if s[j] == ')' || locked[j] == '0' {
            right++
        } else {
            right--
        }
        if right < 0 { return false }
    }
    return true
}

func canBeValid3(s string, locked string) bool {
    low, high := 0, 0
    for i, c := range []byte(s) {
        if locked[i] == '0' {
            low--
            high++
            if low < 0 {
                low = 1
            }
        } else {
            if c == ')' {
                low--
                high--
                if low < 0 {
                    low = 1
                }
                if high < 0 {
                    return false
                }
            } else {
                low++
                high++
            }
        }
    }
    return low <= 0 && len(s) % 2 == 0
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/06/eg1.png" />
    // Input: s = "))()))", locked = "010100"
    // Output: true
    // Explanation: locked[1] == '1' and locked[3] == '1', so we cannot change s[1] or s[3].
    // We change s[0] and s[4] to '(' while leaving s[2] and s[5] unchanged to make s valid.
    fmt.Println(canBeValid("))()))", "010100")) // true
    // Example 2:
    // Input: s = "()()", locked = "0000"
    // Output: true
    // Explanation: We do not need to make any changes because s is already valid.
    fmt.Println(canBeValid("()()", "0000")) // true
    // Example 3:
    // Input: s = ")", locked = "0"
    // Output: false
    // Explanation: locked permits us to change s[0]. 
    // Changing s[0] to either '(' or ')' will not make s valid.
    fmt.Println(canBeValid(")", "0")) // false

    fmt.Println(canBeValid1("))()))", "010100")) // true
    fmt.Println(canBeValid1("()()", "0000")) // true
    fmt.Println(canBeValid1(")", "0")) // false

    
    fmt.Println(canBeValid2("))()))", "010100")) // true
    fmt.Println(canBeValid2("()()", "0000")) // true
    fmt.Println(canBeValid2(")", "0")) // false

    fmt.Println(canBeValid3("))()))", "010100")) // true
    fmt.Println(canBeValid3("()()", "0000")) // true
    fmt.Println(canBeValid3(")", "0")) // false
}