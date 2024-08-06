package main

// 2165. Smallest Value of the Rearranged Number
// You are given an integer num. Rearrange the digits of num such that its value is minimized and it does not contain any leading zeros.
// Return the rearranged number with minimal value.
// Note that the sign of the number does not change after rearranging the digits.

// Example 1:
// Input: num = 310
// Output: 103
// Explanation: The possible arrangements for the digits of 310 are 013, 031, 103, 130, 301, 310. 
// The arrangement with the smallest value that does not contain any leading zeros is 103.

// Example 2:
// Input: num = -7605
// Output: -7650
// Explanation: Some possible arrangements for the digits of -7605 are -7650, -6705, -5076, -0567.
// The arrangement with the smallest value that does not contain any leading zeros is -7650.

// Constraints:
//     -10^15 <= num <= 10^15

import "fmt"
import "sort"

func smallestNumber(num int64) int64 {
    if num == 0 {
        return 0
    }
    res, digits := int64(0), []int64{}
    for {
        if num == 0 {
            break
        }
        // get the ones place
        digit := num % 10
        digits = append(digits, digit)
        num /= 10
    }
    // It doesn't matter whether the number is negative
    sort.Slice(digits, func(i, j int) bool {
        return digits[i] < digits[j]
    })
    if digits[0] == 0 { // 如果开头是 0 需要移动到里面
        for i, d := range digits {
            if d != 0 {
                digits[0], digits[i] = digits[i], digits[0]
                break
            }
        }
    }
    for _, digit := range digits {
        res = res * 10 + digit
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = 310
    // Output: 103
    // Explanation: The possible arrangements for the digits of 310 are 013, 031, 103, 130, 301, 310. 
    // The arrangement with the smallest value that does not contain any leading zeros is 103.
    fmt.Println(smallestNumber(310)) // 103
    // Example 2:
    // Input: num = -7605
    // Output: -7650
    // Explanation: Some possible arrangements for the digits of -7605 are -7650, -6705, -5076, -0567.
    // The arrangement with the smallest value that does not contain any leading zeros is -7650.
    fmt.Println(smallestNumber(-7605)) // -7650
}