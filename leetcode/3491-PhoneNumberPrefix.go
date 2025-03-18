package main

// 3491. Phone Number Prefix
// You are given a string array numbers that represents phone numbers. 
// Return true if no phone number is a prefix of any other phone number; otherwise, return false.

// Example 1:
// Input: numbers = ["1","2","4","3"]
// Output: true
// Explanation:
// No number is a prefix of another number, so the output is true.

// Example 2:
// Input: numbers = ["001","007","15","00153"]
// Output: false
// Explanation:
// The string "001" is a prefix of the string "00153". Thus, the output is false.

// Constraints:
//     2 <= numbers.length <= 50
//     1 <= numbers[i].length <= 50
//     All numbers contain only digits '0' to '9'.

import "fmt"
import "strings"

func phonePrefix(numbers []string) bool {
    for i := range numbers {
        for j := range numbers {
            if i == j { continue } // 一样则跳过
            if strings.HasPrefix(numbers[i], numbers[j]) {
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: numbers = ["1","2","4","3"]
    // Output: true
    // Explanation:
    // No number is a prefix of another number, so the output is true.
    fmt.Println(phonePrefix([]string{"1","2","4","3"})) // true
    // Example 2:
    // Input: numbers = ["001","007","15","00153"]
    // Output: false
    // Explanation:
    // The string "001" is a prefix of the string "00153". Thus, the output is false.
    fmt.Println(phonePrefix([]string{"001","007","15","00153"})) // false
}