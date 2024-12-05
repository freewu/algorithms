package main

// 2310. Sum of Numbers With Units Digit K
// Given two integers num and k, consider a set of positive integers with the following properties:
//     The units digit of each integer is k.
//     The sum of the integers is num.

// Return the minimum possible size of such a set, or -1 if no such set exists.

// Note:
//     The set can contain multiple instances of the same integer, and the sum of an empty set is considered 0.
//     The units digit of a number is the rightmost digit of the number.

// Example 1:
// Input: num = 58, k = 9
// Output: 2
// Explanation:
// One valid set is [9,49], as the sum is 58 and each integer has a units digit of 9.
// Another valid set is [19,39].
// It can be shown that 2 is the minimum possible size of a valid set.

// Example 2:
// Input: num = 37, k = 2
// Output: -1
// Explanation: It is not possible to obtain a sum of 37 using only integers that have a units digit of 2.

// Example 3:
// Input: num = 0, k = 7
// Output: 0
// Explanation: The sum of an empty set is considered 0.

// Constraints:
//     0 <= num <= 3000
//     0 <= k <= 9

import "fmt"

func minimumNumbers(num int, k int) int {
    if num == 0 { return 0 } // base case:
    getLastDigit := func(num int) int {
        if num < 10 { return num }
        return num % 10
    }
    if getLastDigit(num) == k { return 1 }
    if num < k { return -1 }
    count := 0
    for num > k {
        if getLastDigit(num) == k {
            count++
            break
        }
        num -= k
        count ++
        if count > 9 { return -1 }
    }
    if num < 10 && num != k { return -1 }
    if num < 10 && num == k { return count + 1 }
    return count
}

func minimumNumbers1(num int, k int) int {
    if num == 0 { return 0 }
    for i := 1; i <= 10; i++ {
        if k * i <= num && (num - k * i) % 10 == 0 {
            return i
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: num = 58, k = 9
    // Output: 2
    // Explanation:
    // One valid set is [9,49], as the sum is 58 and each integer has a units digit of 9.
    // Another valid set is [19,39].
    // It can be shown that 2 is the minimum possible size of a valid set.
    fmt.Println(minimumNumbers(58, 9)) // 2
    // Example 2:
    // Input: num = 37, k = 2
    // Output: -1
    // Explanation: It is not possible to obtain a sum of 37 using only integers that have a units digit of 2.
    fmt.Println(minimumNumbers(37, 2)) // -1
    // Example 3:
    // Input: num = 0, k = 7
    // Output: 0
    // Explanation: The sum of an empty set is considered 0.
    fmt.Println(minimumNumbers(0, 7)) // 0

    fmt.Println(minimumNumbers(0, 0)) // 0
    fmt.Println(minimumNumbers(3000, 9)) // 10
    fmt.Println(minimumNumbers(0, 9)) // 0
    fmt.Println(minimumNumbers(3000, 0)) // 1

    fmt.Println(minimumNumbers1(58, 9)) // 2
    fmt.Println(minimumNumbers1(37, 2)) // -1
    fmt.Println(minimumNumbers1(0, 7)) // 0
    fmt.Println(minimumNumbers1(0, 0)) // 0
    fmt.Println(minimumNumbers1(3000, 9)) // 10
    fmt.Println(minimumNumbers1(0, 9)) // 0
    fmt.Println(minimumNumbers1(3000, 0)) // 1
}