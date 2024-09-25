package main

// 1317. Convert Integer to the Sum of Two No-Zero Integers
// No-Zero integer is a positive integer that does not contain any 0 in its decimal representation.

// Given an integer n, return a list of two integers [a, b] where:
//     a and b are No-Zero integers.
//     a + b = n

// The test cases are generated so that there is at least one valid solution. 
// If there are many valid solutions, you can return any of them.

// Example 1:
// Input: n = 2
// Output: [1,1]
// Explanation: Let a = 1 and b = 1.
// Both a and b are no-zero integers, and a + b = 2 = n.

// Example 2:
// Input: n = 11
// Output: [2,9]
// Explanation: Let a = 2 and b = 9.
// Both a and b are no-zero integers, and a + b = 9 = n.
// Note that there are other valid answers as [8, 3] that can be accepted.

// Constraints:
//     2 <= n <= 10^4

import "fmt"

func getNoZeroIntegers(n int) []int {
    a, b := 1, n - 1
    notContainZero := func(n int) bool {
        for n > 10 {
            if n % 10 == 0 { return false }
            n = n / 10
        }
        return n != 10
    }
    for {
        if notContainZero(a) && notContainZero(b) { 
            return []int{a, b} 
        }
        a++
        b--
    }
    return nil
}

func getNoZeroIntegers1(n int) []int {
    notContainZero := func(n int) bool {
        for n > 10 {
            if n % 10 == 0 { return false }
            n = n / 10
        }
        return n != 10
    }
    for i := 1; i <= n; i++ {
        if notContainZero(i) && notContainZero(n-i) {
            return []int{i, n - i}
        }
    }
    return []int{}
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: [1,1]
    // Explanation: Let a = 1 and b = 1.
    // Both a and b are no-zero integers, and a + b = 2 = n.
    fmt.Println(getNoZeroIntegers(2)) // [1,1]
    // Example 2:
    // Input: n = 11
    // Output: [2,9]
    // Explanation: Let a = 2 and b = 9.
    // Both a and b are no-zero integers, and a + b = 9 = n.
    // Note that there are other valid answers as [8, 3] that can be accepted.
    fmt.Println(getNoZeroIntegers(11)) // [2,9]

    fmt.Println(getNoZeroIntegers(1024)) // [25 999]
    fmt.Println(getNoZeroIntegers(4000)) // [1 3999]

    fmt.Println(getNoZeroIntegers1(2)) // [1,1]
    fmt.Println(getNoZeroIntegers1(11)) // [2,9]
    fmt.Println(getNoZeroIntegers1(1024)) // [25 999]
    fmt.Println(getNoZeroIntegers1(4000)) // [1 3999]
}