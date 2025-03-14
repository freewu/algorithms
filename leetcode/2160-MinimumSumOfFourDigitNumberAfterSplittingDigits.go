package main

// 2160. Minimum Sum of Four Digit Number After Splitting Digits
// You are given a positive integer num consisting of exactly four digits. 
// Split num into two new integers new1 and new2 by using the digits found in num. 
// Leading zeros are allowed in new1 and new2, and all the digits found in num must be used.

// For example, given num = 2932, you have the following digits: 
// two 2's, one 9 and one 3. Some of the possible pairs [new1, new2] are [22, 93], [23, 92], [223, 9] and [2, 329].

// Return the minimum possible sum of new1 and new2.

// Example 1:
// Input: num = 2932
// Output: 52
// Explanation: Some possible pairs [new1, new2] are [29, 23], [223, 9], etc.
// The minimum sum can be obtained by the pair [29, 23]: 29 + 23 = 52.

// Example 2:
// Input: num = 4009
// Output: 13
// Explanation: Some possible pairs [new1, new2] are [0, 49], [490, 0], etc. 
// The minimum sum can be obtained by the pair [4, 9]: 4 + 9 = 13.

// Constraints:
//     1000 <= num <= 9999

import "fmt"
import "sort"

func minimumSum(num int) int {
    digits := make([]int, 4)
    for i := 0; num > 0; i, num, digits[i] = i + 1, num / 10, num % 10 {}
    sort.Ints(digits)
    return digits[3] + digits[2] + digits[1] * 10 + digits[0] * 10
}

func main() {
    // Example 1:
    // Input: num = 2932
    // Output: 52
    // Explanation: Some possible pairs [new1, new2] are [29, 23], [223, 9], etc.
    // The minimum sum can be obtained by the pair [29, 23]: 29 + 23 = 52.
    fmt.Println(minimumSum(2932)) // 52
    // Example 2:
    // Input: num = 4009
    // Output: 13
    // Explanation: Some possible pairs [new1, new2] are [0, 49], [490, 0], etc. 
    // The minimum sum can be obtained by the pair [4, 9]: 4 + 9 = 13.
    fmt.Println(minimumSum(4009)) // 13

    fmt.Println(minimumSum(1000)) // 1    1000 => 0 + 0 + 0 + 1
    fmt.Println(minimumSum(1024)) // 16   1024 => 10 + 2 + 4
    fmt.Println(minimumSum(9999)) // 198  9999 => 99 + 99
}