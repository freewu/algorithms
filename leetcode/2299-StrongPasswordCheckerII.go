package main

// 2299. Strong Password Checker II
// A password is said to be strong if it satisfies all the following criteria:
//     It has at least 8 characters.
//     It contains at least one lowercase letter.
//     It contains at least one uppercase letter.
//     It contains at least one digit.
//     It contains at least one special character. The special characters are the characters in the following string: "!@#$%^&*()-+".
//     It does not contain 2 of the same character in adjacent positions (i.e., "aab" violates this condition, but "aba" does not).

// Given a string password, return true if it is a strong password. Otherwise, return false.

// Example 1:
// Input: password = "IloveLe3tcode!"
// Output: true
// Explanation: The password meets all the requirements. Therefore, we return true.

// Example 2:
// Input: password = "Me+You--IsMyDream"
// Output: false
// Explanation: The password does not contain a digit and also contains 2 of the same character in adjacent positions. Therefore, we return false.

// Example 3:
// Input: password = "1aB!"
// Output: false
// Explanation: The password does not meet the length requirement. Therefore, we return false. 

// Constraints:
//     1 <= password.length <= 100
//     password consists of letters, digits, and special characters: "!@#$%^&*()-+".

import "fmt"

func strongPasswordCheckerII(password string) bool {
    if len(password) < 8 { return false }
    hasLower, hasUpper, hasDigit, hasSpecial := false, false, false, false
    for i := 0; i < len(password); i++ {
        if i > 0 && password[i] == password[i-1] { return false }
        if password[i] >= 'a' && password[i] <= 'z' { hasLower = true }
        if password[i] >= 'A' && password[i] <= 'Z' { hasUpper = true }
        if password[i] >= '0' && password[i] <= '9' { hasDigit = true }
        if password[i] == '!' || password[i] == '@' || password[i] == '#' ||
           password[i] == '$' || password[i] == '%' || password[i] == '^' || 
           password[i] == '&' || password[i] == '*' || password[i] == '(' || 
           password[i] == ')' || password[i] == '-' || password[i] == '+' {
            hasSpecial = true
        }
    }
    return hasLower && hasUpper && hasDigit && hasSpecial
}

func main() {
    // Example 1:
    // Input: password = "IloveLe3tcode!"
    // Output: true
    // Explanation: The password meets all the requirements. Therefore, we return true.
    fmt.Println(strongPasswordCheckerII("IloveLe3tcode!")) // true
    // Example 2:
    // Input: password = "Me+You--IsMyDream"
    // Output: false
    // Explanation: The password does not contain a digit and also contains 2 of the same character in adjacent positions. Therefore, we return false.
    fmt.Println(strongPasswordCheckerII("Me+You--IsMyDream")) // false
    // Example 3:
    // Input: password = "1aB!"
    // Output: false
    // Explanation: The password does not meet the length requirement. Therefore, we return false. 
    fmt.Println(strongPasswordCheckerII("1aB!")) // false

    fmt.Println(strongPasswordCheckerII("bluefrog")) // false
}