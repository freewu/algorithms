package main

// 420. Strong Password Checker
// A password is considered strong if the below conditions are all met:
//     It has at least 6 characters and at most 20 characters.
//     It contains at least one lowercase letter, at least one uppercase letter, and at least one digit.
//     It does not contain three repeating characters in a row (i.e., "Baaabb0" is weak, but "Baaba0" is strong).

// Given a string password, return the minimum number of steps required to make password strong. 
// if password is already strong, return 0.

// In one step, you can:
//     Insert one character to password,
//     Delete one character from password, or
//     Replace one character of password with another character.
    
// Example 1:
// Input: password = "a"
// Output: 5

// Example 2:
// Input: password = "aA1"
// Output: 3

// Example 3:
// Input: password = "1337C0d3"
// Output: 0
 
// Constraints:
//     1 <= password.length <= 50
//     password consists of letters, digits, dot '.' or exclamation mark '!'.

import "fmt"

func strongPasswordChecker(password string) int {
    n := len(password)
    missingLC, missingUC, missingD := 1, 1, 1
    mustChanged, deleteOne, deleteTwo := 0, 0, 0
    for idx := 0; idx < n; idx++ {
        char := password[idx]
        if 97 <= char && char <= 122 { // 'a' - 'z'
            missingLC = 0
        }
        if 65 <= char && char <= 90 { // 'A' - 'Z'
            missingUC = 0
        }
        if 48 <= char && char <= 57 { // '0' - '9'
            missingD = 0
        }
        if idx > 1 && password[idx] == password[idx-1] && password[idx] == password[idx-2] {
            seqLen := 3
            for idx < n-1 && password[idx+1] == password[idx] {
                seqLen++
                idx++
            }
            mustChanged += seqLen / 3
            if seqLen % 3 == 0 { // "aaa" => "aa"
                deleteOne++
            } else if seqLen%3 == 1 { // "aaaa" => "aa"
                deleteTwo++
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    missingTotal := missingLC + missingUC + missingD
    if n < 6 {
        return max(missingTotal, 6 - n)
    } else if n <= 20 {
        return max(missingTotal, mustChanged)
    } else {
        mustDeleted := n - 20
        mustChanged -= min(mustDeleted, deleteOne)
        mustChanged -= min(max(mustDeleted - deleteOne, 0), deleteTwo*2) / 2
        mustChanged -= max(mustDeleted-deleteOne-2*deleteTwo, 0) / 3

        return mustDeleted + max(missingTotal, mustChanged)
    }
}

func main() {
    // Example 1:
    // Input: password = "a"
    // Output: 5
    fmt.Println(strongPasswordChecker("a")) // 5
    // Example 2:
    // Input: password = "aA1"
    // Output: 3
    fmt.Println(strongPasswordChecker("aA1")) // 3
    // Example 3:
    // Input: password = "1337C0d3"
    // Output: 0
    fmt.Println(strongPasswordChecker("1337C0d3")) // 0
    fmt.Println(strongPasswordChecker("123456")) // 2
}