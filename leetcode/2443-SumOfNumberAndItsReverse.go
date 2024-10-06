package main

// 2443. Sum of Number and Its Reverse
// Given a non-negative integer num, 
// return true if num can be expressed as the sum of any non-negative integer and its reverse, or false otherwise.

// Example 1:
// Input: num = 443
// Output: true
// Explanation: 172 + 271 = 443 so we return true.

// Example 2:
// Input: num = 63
// Output: false
// Explanation: 63 cannot be expressed as the sum of a non-negative integer and its reverse so we return false.

// Example 3:
// Input: num = 181
// Output: true
// Explanation: 140 + 041 = 181 so we return true. Note that when a number is reversed, there may be leading zeros.

// Constraints:
//     0 <= num <= 10^5

import "fmt"

func sumOfNumberAndReverse(num int) bool {
    reverse := func (i int) int {
        res := 0
        for i > 0 {
            res *= 10
            res += i % 10
            i /= 10
        }
        return res
    }
    for i := 0; i <= num; i++ {
        if i + reverse(i) == num {
            return true
        }
    }
    return false
}

func sumOfNumberAndReverse1(num int) bool {
    reverse := func(num int) int {
        res := 0
        for num > 0 {
            digit := num % 10
            num = num / 10
            res = res * 10 + digit
        }
        return res
    }
    for i := 0; i <= num / 2; i++ {
        diff := num - i
        if reverse(diff) == i {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: num = 443
    // Output: true
    // Explanation: 172 + 271 = 443 so we return true.
    fmt.Println(sumOfNumberAndReverse(443)) // true
    // Example 2:
    // Input: num = 63
    // Output: false
    // Explanation: 63 cannot be expressed as the sum of a non-negative integer and its reverse so we return false.
    fmt.Println(sumOfNumberAndReverse(63)) // false
    // Example 3:
    // Input: num = 181
    // Output: true
    // Explanation: 140 + 041 = 181 so we return true. Note that when a number is reversed, there may be leading zeros.
    fmt.Println(sumOfNumberAndReverse(181)) // true

    fmt.Println(sumOfNumberAndReverse(0)) // true
    fmt.Println(sumOfNumberAndReverse(1)) // false
    fmt.Println(sumOfNumberAndReverse(99999)) // false
    fmt.Println(sumOfNumberAndReverse(100000)) // false

    fmt.Println(sumOfNumberAndReverse1(443)) // true
    fmt.Println(sumOfNumberAndReverse1(63)) // false
    fmt.Println(sumOfNumberAndReverse1(181)) // true
    fmt.Println(sumOfNumberAndReverse1(0)) // true
    fmt.Println(sumOfNumberAndReverse1(1)) // false
    fmt.Println(sumOfNumberAndReverse1(99999)) // false
    fmt.Println(sumOfNumberAndReverse1(100000)) // false
}