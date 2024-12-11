package main

// 2177. Find Three Consecutive Integers That Sum to a Given Number
// Given an integer num, return three consecutive integers (as a sorted array) that sum to num. 
// If num cannot be expressed as the sum of three consecutive integers, return an empty array.

// Example 1:
// Input: num = 33
// Output: [10,11,12]
// Explanation: 33 can be expressed as 10 + 11 + 12 = 33.
// 10, 11, 12 are 3 consecutive integers, so we return [10, 11, 12].

// Example 2:
// Input: num = 4
// Output: []
// Explanation: There is no way to express 4 as the sum of 3 consecutive integers.

// Constraints:
//     0 <= num <= 10^15

import "fmt"

func sumOfThree(num int64) []int64 {
    if num % 3 != 0 {  return []int64{} }
    x := num / 3
    return []int64{ x - 1, x, x + 1}
}

func main() {
    // Example 1:
    // Input: num = 33
    // Output: [10,11,12]
    // Explanation: 33 can be expressed as 10 + 11 + 12 = 33.
    // 10, 11, 12 are 3 consecutive integers, so we return [10, 11, 12].
    fmt.Println(sumOfThree(33)) // [10,11,12]
    // Example 2:
    // Input: num = 4
    // Output: []
    // Explanation: There is no way to express 4 as the sum of 3 consecutive integers.
    fmt.Println(sumOfThree(4)) // []

    fmt.Println(sumOfThree(0)) // [-1 0 1]
    fmt.Println(sumOfThree(9)) // [2 3 4]
    fmt.Println(sumOfThree(66)) // [21 22 23]
    fmt.Println(sumOfThree(999)) // [332 333 334]
    fmt.Println(sumOfThree(1024)) // []
    fmt.Println(sumOfThree(1e17)) // []
}