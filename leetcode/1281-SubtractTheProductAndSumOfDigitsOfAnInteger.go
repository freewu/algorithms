package main

// 1281. Subtract the Product and Sum of Digits of an Integer
// Given an integer number n, return the difference between the product of its digits and the sum of its digits.
 
// Example 1:
// Input: n = 234
// Output: 15 
// Explanation: 
// Product of digits = 2 * 3 * 4 = 24 
// Sum of digits = 2 + 3 + 4 = 9 
// Result = 24 - 9 = 15

// Example 2:
// Input: n = 4421
// Output: 21
// Explanation: 
// Product of digits = 4 * 4 * 2 * 1 = 32 
// Sum of digits = 4 + 4 + 2 + 1 = 11 
// Result = 32 - 11 = 21

// Constraints:
//     1 <= n <= 10^5

import "fmt"

// 迭代
func subtractProductAndSum(n int) int {
    sum := 0
    product := 1
    for n != 0 {
        t := n % 10 // 取余
        sum += t
        product *= t
        n = n / 10
    }
    return product - sum
}

// 递归
func subtractProductAndSum1(n int) int {
    return mul(n) - sum(n)
}

func mul(n int) int {
    if n < 10 {
        return n
    }
    return n % 10 * mul(n / 10)
}

func sum(n int) int {
    if n < 10 {
        return n
    }
    return n % 10 + sum(n / 10)
}

func main() {
    // Product of digits = 2 * 3 * 4 = 24 
    // Sum of digits = 2 + 3 + 4 = 9 
    // Result = 24 - 9 = 15
    fmt.Println(subtractProductAndSum(234)) // 15
    // Product of digits = 4 * 4 * 2 * 1 = 32 
    // Sum of digits = 4 + 4 + 2 + 1 = 11 
    // Result = 32 - 11 = 21
    fmt.Println(subtractProductAndSum(4421)) // 21

    fmt.Println(subtractProductAndSum1(234)) // 15
    fmt.Println(subtractProductAndSum1(4421)) // 21
}