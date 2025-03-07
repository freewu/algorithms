package main

// 1780. Check if Number is a Sum of Powers of Three
// Given an integer n, return true if it is possible to represent n as the sum of distinct powers of three. 
// Otherwise, return false.

// An integer y is a power of three if there exists an integer x such that y == 3x.

// Example 1:
// Input: n = 12
// Output: true
// Explanation: 12 = 31 + 32

// Example 2:
// Input: n = 91
// Output: true
// Explanation: 91 = 30 + 32 + 34

// Example 3:
// Input: n = 21
// Output: false

// Constraints:
//     1 <= n <= 10^7

import "fmt"

func checkPowersOfThree(n int) bool {
    quotient, remainder := n / 3, n % 3
    if remainder == 2 {
        return false
    } else if quotient == 0 || quotient == 1 {
        return true
    } else {
        return checkPowersOfThree(quotient)
    }
}

func checkPowersOfThree1(n int) bool {
    // Find the largest power that is smaller or equal to n
    power3 := 1
    for ; power3 <= n; power3 *= 3 {}
    for n > 0 {
        // Subtract current power from n
        if n >= power3 {
            n -= power3
        }
        // We cannot use the same power twice
        if n >= power3 {
            return false
        }
        // Move to the next lower power
        power3 /= 3
    }
    // n has reached 0
    return true
}

func main() {
    // Example 1:
    // Input: n = 12
    // Output: true
    // Explanation: 12 = 31 + 32
    fmt.Println(checkPowersOfThree(12)) // true 12 = 31 + 32
    // Example 2:
    // Input: n = 91
    // Output: true
    // Explanation: 91 = 30 + 32 + 34
    fmt.Println(checkPowersOfThree(91)) // true 91 = 30 + 32 + 34
    // Example 3:
    // Input: n = 21
    // Output: false
    fmt.Println(checkPowersOfThree(21)) // false

    fmt.Println(checkPowersOfThree(1024)) // false
    fmt.Println(checkPowersOfThree(1)) // true
    fmt.Println(checkPowersOfThree(10000000)) // false
    fmt.Println(checkPowersOfThree(999999)) // false

    fmt.Println(checkPowersOfThree1(12)) // true 12 = 31 + 32
    fmt.Println(checkPowersOfThree1(91)) // true 91 = 30 + 32 + 34
    fmt.Println(checkPowersOfThree1(21)) // false
    fmt.Println(checkPowersOfThree1(1024)) // false
    fmt.Println(checkPowersOfThree1(1)) // true
    fmt.Println(checkPowersOfThree1(10000000)) // false
    fmt.Println(checkPowersOfThree1(999999)) // false
}