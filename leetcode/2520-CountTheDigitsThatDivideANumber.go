package main

// 2520. Count the Digits That Divide a Number
// Given an integer num, return the number of digits in num that divide num.

// An integer val divides nums if nums % val == 0.

// Example 1:
// Input: num = 7
// Output: 1
// Explanation: 7 divides itself, hence the answer is 1.

// Example 2:
// Input: num = 121
// Output: 2
// Explanation: 121 is divisible by 1, but not 2. Since 1 occurs twice as a digit, we return 2.

// Example 3:
// Input: num = 1248
// Output: 4
// Explanation: 1248 is divisible by all of its digits, hence the answer is 4.

// Constraints:
//     1 <= num <= 10^9
//     num does not contain 0 as one of its digits.

import "fmt"
import "strconv"

func countDigits1(num int) (r int) {
    res, s := 0, strconv.Itoa(num)
    for _, v := range s {
        n, _ := strconv.Atoi(string(v))
        if num % n == 0 {
            res++
        }
    }
    return res
}

func countDigits(num int) int {
    res, tmp := 0, num
    for num > 0 {
        d := num % 10
        num /= 10
        if tmp % d == 0 { res++ }
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = 7
    // Output: 1
    // Explanation: 7 divides itself, hence the answer is 1.
    fmt.Println(countDigits(7)) // 1
    // Example 2:
    // Input: num = 121
    // Output: 2
    // Explanation: 121 is divisible by 1, but not 2. Since 1 occurs twice as a digit, we return 2.
    fmt.Println(countDigits(121)) // 2
    // Example 3:
    // Input: num = 1248
    // Output: 4
    // Explanation: 1248 is divisible by all of its digits, hence the answer is 4.
    fmt.Println(countDigits(1248)) // 4

    fmt.Println(countDigits(1)) // 1
    fmt.Println(countDigits(2)) // 1
    //fmt.Println(countDigits(1024)) // 
    fmt.Println(countDigits(9999)) // 4
    fmt.Println(countDigits(999_999_999)) // 9
    //fmt.Println(countDigits(1_000_000_000)) //

    fmt.Println(countDigits1(7)) // 1
    fmt.Println(countDigits1(121)) // 2
    fmt.Println(countDigits1(1248)) // 4
    fmt.Println(countDigits1(1)) // 1
    fmt.Println(countDigits1(2)) // 1
    //fmt.Println(countDigits1(1024)) // 
    fmt.Println(countDigits1(9999)) // 4
    fmt.Println(countDigits1(999_999_999)) // 9
    //fmt.Println(countDigits1(1_000_000_000)) //
}