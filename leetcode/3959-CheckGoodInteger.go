package main

// 3959. Check Good Integer
// You are given a positive integer n.

// Let digitSum be the sum of the digits of n, and let squareSum be the sum of the squares of the digits of n.

// An integer is called good if squareSum - digitSum >= 50.

// Return true if n is good. Otherwise, return false.

// Example 1:
// Input: n = 1000
// Output: false
// Explanation:
// The digits of 1000 are 1, 0, 0, and 0.
// The digitSum is 1 + 0 + 0 + 0 = 1.
// The squareSum is 12 + 02 + 02 + 02 = 1.
// The squareSum - digitSum is 1 - 1 = 0. As 0 is not greater than or equal to 50, the output is false.

// Example 2:
// Input: n = 19
// Output: true
// Explanation:
// The digits of 19 are 1 and 9.
// The digitSum is 1 + 9 = 10.
// The squareSum is 12 + 92 = 1 + 81 = 82.
// The squareSum - digitSum is 82 - 10 = 72. As 72 is greater than or equal to 50, the output is true.

// Constraints:
//     1 <= n <= 10^9

import "fmt"

func checkGoodInteger(n int) bool {
    sum, val := 0, 0
    for n > 0 {
        val = n % 10
        if(val > 7) { // This is Unique isnt it
            return true
         } 
        if(val > 1) {
            sum += val * (val - 1)
        }
        n /= 10
    }
    return sum > 49
}

func checkGoodInteger1(n int) bool {
    sum, squareSum := 0,0
    for n > 0 {
        rem := n % 10
        sum += rem
        squareSum += (rem * rem)
        n /= 10
    }
    if squareSum - sum >= 50 {
        return true
    }
    return false
}

func main() {
    // Example 1:
    // Input: n = 1000
    // Output: false
    // Explanation:
    // The digits of 1000 are 1, 0, 0, and 0.
    // The digitSum is 1 + 0 + 0 + 0 = 1.
    // The squareSum is 12 + 02 + 02 + 02 = 1.
    // The squareSum - digitSum is 1 - 1 = 0. As 0 is not greater than or equal to 50, the output is false.
    fmt.Println(checkGoodInteger(1000)) // false
    // Example 2:
    // Input: n = 19
    // Output: true
    // Explanation:
    // The digits of 19 are 1 and 9.
    // The digitSum is 1 + 9 = 10.
    // The squareSum is 12 + 92 = 1 + 81 = 82.
    // The squareSum - digitSum is 82 - 10 = 72. As 72 is greater than or equal to 50, the output is true. 
    fmt.Println(checkGoodInteger(19)) // true

    fmt.Println(checkGoodInteger(1)) // false
    fmt.Println(checkGoodInteger(2)) // false
    fmt.Println(checkGoodInteger(3)) // false
    fmt.Println(checkGoodInteger(8)) // true
    fmt.Println(checkGoodInteger(64)) // false
    fmt.Println(checkGoodInteger(99)) // true
    fmt.Println(checkGoodInteger(100)) // false
    fmt.Println(checkGoodInteger(101)) // false
    fmt.Println(checkGoodInteger(999)) // true
    fmt.Println(checkGoodInteger(1024)) // false
    fmt.Println(checkGoodInteger(999_999_999)) // true
    fmt.Println(checkGoodInteger(1_000_000_000)) // false

    fmt.Println(checkGoodInteger1(1000)) // false
    fmt.Println(checkGoodInteger1(19)) // true
    fmt.Println(checkGoodInteger1(1)) // false
    fmt.Println(checkGoodInteger1(2)) // false
    fmt.Println(checkGoodInteger1(3)) // false
    fmt.Println(checkGoodInteger1(8)) // true
    fmt.Println(checkGoodInteger1(64)) // false
    fmt.Println(checkGoodInteger1(99)) // true
    fmt.Println(checkGoodInteger1(100)) // false
    fmt.Println(checkGoodInteger1(101)) // false
    fmt.Println(checkGoodInteger1(999)) // true
    fmt.Println(checkGoodInteger1(1024)) // false
    fmt.Println(checkGoodInteger1(999_999_999)) // true
    fmt.Println(checkGoodInteger1(1_000_000_000)) // false
}