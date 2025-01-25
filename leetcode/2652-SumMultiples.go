package main

// 2652. Sum Multiples
// Given a positive integer n, find the sum of all integers in the range [1, n] inclusive that are divisible by 3, 5, or 7.

// Return an integer denoting the sum of all numbers in the given range satisfying the constraint.

// Example 1:
// Input: n = 7
// Output: 21
// Explanation: Numbers in the range [1, 7] that are divisible by 3, 5, or 7 are 3, 5, 6, 7. 
// The sum of these numbers is 21.

// Example 2:
// Input: n = 10
// Output: 40
// Explanation: Numbers in the range [1, 10] that are divisible by 3, 5, or 7 are 3, 5, 6, 7, 9, 10. 
// The sum of these numbers is 40.

// Example 3:
// Input: n = 9
// Output: 30
// Explanation: Numbers in the range [1, 9] that are divisible by 3, 5, or 7 are 3, 5, 6, 7, 9. 
// The sum of these numbers is 30.

// Constraints:
//     1 <= n <= 10^3

import "fmt"

func sumOfMultiples(n int) int {
    res := 0
    for i := 1; i <= n; i++ {
        if i % 3 == 0 || i % 5 == 0 || i % 7 == 0 {
            res += i
        }
    }
    return res
}

func sumOfMultiples1(n int) int {
    calc := func (n, m int) int { return n / m * (n / m + 1) / 2 * m }
    return calc(n, 3) + calc(n, 5) + calc(n, 7) - calc(n, 15) - calc(n, 21) - calc(n, 35) + calc(n, 105)
}

func main() {
    // Example 1:
    // Input: n = 7
    // Output: 21
    // Explanation: Numbers in the range [1, 7] that are divisible by 3, 5, or 7 are 3, 5, 6, 7. 
    // The sum of these numbers is 21.
    fmt.Println(sumOfMultiples(7)) // 21
    // Example 2:
    // Input: n = 10
    // Output: 40
    // Explanation: Numbers in the range [1, 10] that are divisible by 3, 5, or 7 are 3, 5, 6, 7, 9, 10. 
    // The sum of these numbers is 40.
    fmt.Println(sumOfMultiples(10)) // 40
    // Example 3:
    // Input: n = 9
    // Output: 30
    // Explanation: Numbers in the range [1, 9] that are divisible by 3, 5, or 7 are 3, 5, 6, 7, 9. 
    // The sum of these numbers is 30.
    fmt.Println(sumOfMultiples(9)) // 30

    fmt.Println(sumOfMultiples(1)) // 0
    fmt.Println(sumOfMultiples(2)) // 0
    fmt.Println(sumOfMultiples(3)) // 3
    fmt.Println(sumOfMultiples(8)) // 21
    fmt.Println(sumOfMultiples(64)) // 1087
    fmt.Println(sumOfMultiples(99)) // 2738
    fmt.Println(sumOfMultiples(100)) // 2838
    fmt.Println(sumOfMultiples(999)) // 271066
    fmt.Println(sumOfMultiples(1000)) // 272066

    fmt.Println(sumOfMultiples1(7)) // 21
    fmt.Println(sumOfMultiples1(10)) // 40
    fmt.Println(sumOfMultiples1(9)) // 30
    fmt.Println(sumOfMultiples1(1)) // 0
    fmt.Println(sumOfMultiples1(2)) // 0
    fmt.Println(sumOfMultiples1(3)) // 3
    fmt.Println(sumOfMultiples1(8)) // 21
    fmt.Println(sumOfMultiples1(64)) // 1087
    fmt.Println(sumOfMultiples1(99)) // 2738
    fmt.Println(sumOfMultiples1(100)) // 2838
    fmt.Println(sumOfMultiples1(999)) // 271066
    fmt.Println(sumOfMultiples1(1000)) // 272066
}