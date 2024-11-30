package main

// 2119. A Number After a Double Reversal
// Reversing an integer means to reverse all its digits.
//     For example, reversing 2021 gives 1202. Reversing 12300 gives 321 as the leading zeros are not retained.

// Given an integer num, reverse num to get reversed1, then reverse reversed1 to get reversed2. 
// Return true if reversed2 equals num. Otherwise return false.

// Example 1:
// Input: num = 526
// Output: true
// Explanation: Reverse num to get 625, then reverse 625 to get 526, which equals num.

// Example 2:
// Input: num = 1800
// Output: false
// Explanation: Reverse num to get 81, then reverse 81 to get 18, which does not equal num.

// Example 3:
// Input: num = 0
// Output: true
// Explanation: Reverse num to get 0, then reverse 0 to get 0, which equals num.

// Constraints:
//     0 <= num <= 10^6

import "fmt"

func isSameAfterReversals(num int) bool {
    if num % 10 == 0 && num != 0 {
        return false
    }
    return true
}

func main() {
    // Example 1:
    // Input: num = 526
    // Output: true
    // Explanation: Reverse num to get 625, then reverse 625 to get 526, which equals num.
    fmt.Println(isSameAfterReversals(526)) // true
    // Example 2:
    // Input: num = 1800
    // Output: false
    // Explanation: Reverse num to get 81, then reverse 81 to get 18, which does not equal num.
    fmt.Println(isSameAfterReversals(1800)) // false
    // Example 3:
    // Input: num = 0
    // Output: true
    // Explanation: Reverse num to get 0, then reverse 0 to get 0, which equals num.
    fmt.Println(isSameAfterReversals(0)) // true

    fmt.Println(isSameAfterReversals(1)) // true
    fmt.Println(isSameAfterReversals(2)) // true
    fmt.Println(isSameAfterReversals(8)) // true
    fmt.Println(isSameAfterReversals(1024)) // true
    fmt.Println(isSameAfterReversals(999_999)) // true
    fmt.Println(isSameAfterReversals(10)) // false
    fmt.Println(isSameAfterReversals(50)) // false
    fmt.Println(isSameAfterReversals(100)) // false
    fmt.Println(isSameAfterReversals(1_000)) // false
    fmt.Println(isSameAfterReversals(1_000_000)) // false
}