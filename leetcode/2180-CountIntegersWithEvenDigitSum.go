package main

// 2180. Count Integers With Even Digit Sum
// Given a positive integer num, return the number of positive integers less than or equal to num whose digit sums are even.

// The digit sum of a positive integer is the sum of all its digits.

// Example 1:
// Input: num = 4
// Output: 2
// Explanation:
// The only integers less than or equal to 4 whose digit sums are even are 2 and 4.

// Example 2:
// Input: num = 30
// Output: 14
// Explanation:
// The 14 integers less than or equal to 30 whose digit sums are even are
// 2, 4, 6, 8, 11, 13, 15, 17, 19, 20, 22, 24, 26, and 28.

// Constraints:
//     1 <= num <= 1000

import "fmt"

func countEven(num int) int {
    digitSumIsEven := func(n int) bool {
        sum := 0
        for n != 0 {
            sum += n % 10
            n /= 10
        }
        if sum % 2 == 0 { return true }
        return false
    }
    if digitSumIsEven(num) {
        return num / 2
    }
    return (num - 1) / 2
}

func countEven1(num int) int {
    if num == 1 { return 0 }
    res := 0
    for i := 2; i <= num; i++ {
        tmp, tmp2 := i, 0
        for {
            a := tmp % 10
            tmp = tmp / 10
            tmp2 += a
            if tmp == 0 { break }
        }
        if tmp2 % 2 == 0 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = 4
    // Output: 2
    // Explanation:
    // The only integers less than or equal to 4 whose digit sums are even are 2 and 4.
    fmt.Println(countEven(4)) // 2
    // Example 2:
    // Input: num = 30
    // Output: 14
    // Explanation:
    // The 14 integers less than or equal to 30 whose digit sums are even are
    // 2, 4, 6, 8, 11, 13, 15, 17, 19, 20, 22, 24, 26, and 28.
    fmt.Println(countEven(30)) // 14

    fmt.Println(countEven(1)) // 0
    fmt.Println(countEven(8)) // 4
    fmt.Println(countEven(64)) // 32
    fmt.Println(countEven(128)) // 63
    fmt.Println(countEven(999)) // 499
    fmt.Println(countEven(1000)) // 499

    fmt.Println(countEven1(4)) // 2
    fmt.Println(countEven1(30)) // 14
    fmt.Println(countEven1(1)) // 0
    fmt.Println(countEven1(8)) // 4
    fmt.Println(countEven1(64)) // 32
    fmt.Println(countEven1(128)) // 63
    fmt.Println(countEven1(999)) // 499
    fmt.Println(countEven1(1000)) // 499
}