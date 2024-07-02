package main

// 3099. Harshad Number
// An integer divisible by the sum of its digits is said to be a Harshad number. 
// You are given an integer x. Return the sum of the digits of x if x is a Harshad number, otherwise, return -1.

// Example 1:
// Input: x = 18
// Output: 9
// Explanation:
// The sum of digits of x is 9. 18 is divisible by 9. So 18 is a Harshad number and the answer is 9.

// Example 2:
// Input: x = 23
// Output: -1
// Explanation:
// The sum of digits of x is 5. 23 is not divisible by 5. So 23 is not a Harshad number and the answer is -1.

// Constraints:
//     1 <= x <= 100

import "fmt"

func sumOfTheDigitsOfHarshadNumber(x int) int {
    sumOfDigits := func(x int) int {
        sum := 0
        for x > 0 { // till x is greater than 0
            sum += x % 10
            x /=  10
        }
        return sum
    }
    sum := sumOfDigits(x)
    if x % sum == 0 {
        return sum
    }
    return -1
}

func main() {
    // Example 1:
    // Input: x = 18
    // Output: 9
    // Explanation:
    // The sum of digits of x is 9. 18 is divisible by 9. So 18 is a Harshad number and the answer is 9.
    fmt.Println(sumOfTheDigitsOfHarshadNumber(18)) // 9
    // Example 2:
    // Input: x = 23
    // Output: -1
    // Explanation:
    // The sum of digits of x is 5. 23 is not divisible by 5. So 23 is not a Harshad number and the answer is -1.
    fmt.Println(sumOfTheDigitsOfHarshadNumber(23)) // -1
}