package main

// 2578. Split With Minimum Sum
// Given a positive integer num, split it into two non-negative integers num1 and num2 such that:
//     1. The concatenation of num1 and num2 is a permutation of num.
//         In other words, the sum of the number of occurrences of each digit in num1 and num2 is equal to the number of occurrences of that digit in num.
//     2. num1 and num2 can contain leading zeros.

// Return the minimum possible sum of num1 and num2.

// Notes:
//     It is guaranteed that num does not contain any leading zeros.
//     The order of occurrence of the digits in num1 and num2 may differ from the order of occurrence of num.

// Example 1:
// Input: num = 4325
// Output: 59
// Explanation: We can split 4325 so that num1 is 24 and num2 is 35, giving a sum of 59. We can prove that 59 is indeed the minimal possible sum.

// Example 2:
// Input: num = 687
// Output: 75
// Explanation: We can split 687 so that num1 is 68 and num2 is 7, which would give an optimal sum of 75.

// Constraints:
//     10 <= num <= 10^9

import "fmt"
import "strconv"
import "sort"

func splitNum(num int) int {
    str := strconv.Itoa(num)
    arr := []int{}
    for _, ch := range str {
        v, _ := strconv.Atoi(string(ch))
        if v != 0 {
            arr = append(arr, v)
        }
    }
    if len(arr) == 2 { return arr[0] + arr[1] }
    sort.Ints(arr)
    num1, num2 := 0, 0
    for i, digit := range arr {
        if i % 2 == 0 {
            num1 = num1 * 10 + digit
        } else {
            num2 = num2 * 10 + digit
        }
    }
    return num1 + num2
}

func main() {
    // Example 1:
    // Input: num = 4325
    // Output: 59
    // Explanation: We can split 4325 so that num1 is 24 and num2 is 35, giving a sum of 59. We can prove that 59 is indeed the minimal possible sum.
    fmt.Println(splitNum(4325)) // 59
    // Example 2:
    // Input: num = 687
    // Output: 75
    // Explanation: We can split 687 so that num1 is 68 and num2 is 7, which would give an optimal sum of 75.
    fmt.Println(splitNum(687)) // 75
}