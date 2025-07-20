package main

// 3622. Check Divisibility by Digit Sum and Product
// You are given a positive integer n. 
// Determine whether n is divisible by the sum of the following two values:
//     1. The digit sum of n (the sum of its digits).
//     2. The digit product of n (the product of its digits).

// Return true if n is divisible by this sum; otherwise, return false.

// Example 1:
// Input: n = 99
// Output: true
// Explanation:
// Since 99 is divisible by the sum (9 + 9 = 18) plus product (9 * 9 = 81) of its digits (total 99), the output is true.

// Example 2:
// Input: n = 23
// Output: false
// Explanation:
// Since 23 is not divisible by the sum (2 + 3 = 5) plus product (2 * 3 = 6) of its digits (total 11), the output is false.

// Constraints:
//     1 <= n <= 10^6

import "fmt"

func checkDivisibility(n int) bool {
    if n == 0 { return true }
    sum, prod, m, arr := 0, 1, n, []int{}
    for n != 0 {
        rem := n % 10
        n = n / 10
        arr = append(arr, rem)
    }
    for i := 0; i < len(arr); i++ {
        sum += arr[i];
        prod = prod * arr[i];
    }
    if (m % (sum + prod)) == 0 { return true }
    return false
}

func main() {
    // Example 1:
    // Input: n = 99
    // Output: true
    // Explanation:
    // Since 99 is divisible by the sum (9 + 9 = 18) plus product (9 * 9 = 81) of its digits (total 99), the output is true.
    fmt.Println(checkDivisibility(90)) // true
    // Example 2:
    // Input: n = 23
    // Output: false
    // Explanation:
    // Since 23 is not divisible by the sum (2 + 3 = 5) plus product (2 * 3 = 6) of its digits (total 11), the output is false.
    fmt.Println(checkDivisibility(23)) // false

    fmt.Println(checkDivisibility(1)) // false
    fmt.Println(checkDivisibility(1024)) // false
    fmt.Println(checkDivisibility(999_999)) // false
    fmt.Println(checkDivisibility(1_000_000)) // true
}