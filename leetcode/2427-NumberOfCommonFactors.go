package main

// 2427. Number of Common Factors
// Given two positive integers a and b, return the number of common factors of a and b.

// An integer x is a common factor of a and b if x divides both a and b.

// Example 1:
// Input: a = 12, b = 6
// Output: 4
// Explanation: The common factors of 12 and 6 are 1, 2, 3, 6.

// Example 2:
// Input: a = 25, b = 30
// Output: 2
// Explanation: The common factors of 25 and 30 are 1, 5.

// Constraints:
//     1 <= a, b <= 1000

import "fmt"

func commonFactors(a int, b int) int {
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    res, factor := 0, gcd(a, b)
    for i := 1 ; i <= factor ; i++{
        if factor % i == 0 {
            res++
        }
    }
    return res
}

func commonFactors1(a int, b int) int {
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    res, factor := 0, gcd(a, b)
    for i := 1; i <= factor; i++ {
        if i * i > factor { break }
        if factor % i != 0 { continue }
        res++
        if factor / i != i {
            res++
        }
    }
    return res
}

func commonFactors2(a int, b int) int {
    res, mn := 0, b
    if a < b {
        mn = a
    }
    for i := 1; i <= mn; i++ {
        if a % i == 0 && b % i == 0 {
            res++
        }
    } 
    return res
}

func main() {
    // Example 1:
    // Input: a = 12, b = 6
    // Output: 4
    // Explanation: The common factors of 12 and 6 are 1, 2, 3, 6.
    fmt.Println(commonFactors(12, 6)) // 4
    // Example 2:
    // Input: a = 25, b = 30
    // Output: 2
    // Explanation: The common factors of 25 and 30 are 1, 5.
    fmt.Println(commonFactors(25, 30)) // 2

    fmt.Println(commonFactors(1, 1)) // 1
    fmt.Println(commonFactors(1, 1000)) // 1
    fmt.Println(commonFactors(1000, 1)) // 1
    fmt.Println(commonFactors(1000, 1000)) // 16

    fmt.Println(commonFactors1(12, 6)) // 4
    fmt.Println(commonFactors1(25, 30)) // 2
    fmt.Println(commonFactors1(1, 1)) // 1
    fmt.Println(commonFactors1(1, 1000)) // 1
    fmt.Println(commonFactors1(1000, 1)) // 1
    fmt.Println(commonFactors1(1000, 1000)) // 16

    fmt.Println(commonFactors2(12, 6)) // 4
    fmt.Println(commonFactors2(25, 30)) // 2
    fmt.Println(commonFactors2(1, 1)) // 1
    fmt.Println(commonFactors2(1, 1000)) // 1
    fmt.Println(commonFactors2(1000, 1)) // 1
    fmt.Println(commonFactors2(1000, 1000)) // 16
}