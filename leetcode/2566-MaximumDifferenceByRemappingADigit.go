package main

// 2566. Maximum Difference by Remapping a Digit
// You are given an integer num. 
// You know that Bob will sneakily remap one of the 10 possible digits (0 to 9) to another digit.

// Return the difference between the maximum and minimum values Bob can make by remapping exactly one digit in num.

// Notes:
//     1. When Bob remaps a digit d1 to another digit d2, Bob replaces all occurrences of d1 in num with d2.
//     2. Bob can remap a digit to itself, in which case num does not change.
//     3. Bob can remap different digits for obtaining minimum and maximum values respectively.
//     4. The resulting number after remapping can contain leading zeroes.

// Example 1:
// Input: num = 11891
// Output: 99009
// Explanation: 
// To achieve the maximum value, Bob can remap the digit 1 to the digit 9 to yield 99899.
// To achieve the minimum value, Bob can remap the digit 1 to the digit 0, yielding 890.
// The difference between these two numbers is 99009.

// Example 2:
// Input: num = 90
// Output: 99
// Explanation:
// The maximum value that can be returned by the function is 99 (if 0 is replaced by 9) and the minimum value that can be returned by the function is 0 (if 9 is replaced by 0).
// Thus, we return 99.

// Constraints:
//     1 <= num <= 10^8

import "fmt"
import "strings"
import "strconv"

func minMaxDifference(num int) int {
    str := strconv.Itoa(num)
    tmp := str
    for _, v := range str {
        if v != '9' {
            tmp = strings.ReplaceAll(str, string(v), "9")
            break
        }
    }
    res, _ := strconv.Atoi(tmp)
    tmp = str
    for _, v := range str {
        if v != '0' {
            tmp = strings.ReplaceAll(str, string(v), "0")
            break
        }
    }
    t, _ := strconv.Atoi(tmp)
    return res - t
}

func minMaxDifference1(num int) int {
    str := strconv.Itoa(num)
    mx, mn := num, num
    for _, v := range str {
        if v != '9' {
            mx, _ = strconv.Atoi(strings.ReplaceAll(str, string(v), "9"))
            break
        }
    }
    mn, _ = strconv.Atoi(strings.ReplaceAll(str, str[:1], "0"))
    return mx - mn
}

func main() {
    // Example 1:
    // Input: num = 11891
    // Output: 99009
    // Explanation: 
    // To achieve the maximum value, Bob can remap the digit 1 to the digit 9 to yield 99899.
    // To achieve the minimum value, Bob can remap the digit 1 to the digit 0, yielding 890.
    // The difference between these two numbers is 99009.
    fmt.Println(minMaxDifference(11891)) // 99009
    // Example 2:
    // Input: num = 90
    // Output: 99
    // Explanation:
    // The maximum value that can be returned by the function is 99 (if 0 is replaced by 9) and the minimum value that can be returned by the function is 0 (if 9 is replaced by 0).
    // Thus, we return 99.
    fmt.Println(minMaxDifference(90)) // 99

    fmt.Println(minMaxDifference(1)) // 9
    fmt.Println(minMaxDifference(2)) // 9
    fmt.Println(minMaxDifference(1024)) // 9000
    fmt.Println(minMaxDifference(99_999_999)) // 99999999
    fmt.Println(minMaxDifference(100_000_000)) // 900000000

    fmt.Println(minMaxDifference1(11891)) // 99009
    fmt.Println(minMaxDifference1(90)) // 99
    fmt.Println(minMaxDifference1(1)) // 9
    fmt.Println(minMaxDifference1(2)) // 9
    fmt.Println(minMaxDifference1(1024)) // 9000
    fmt.Println(minMaxDifference1(99_999_999)) // 99999999
    fmt.Println(minMaxDifference1(100_000_000)) // 900000000
}