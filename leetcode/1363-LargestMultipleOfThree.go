package main

// 1363. Largest Multiple of Three
// Given an array of digits digits, 
// return the largest multiple of three that can be formed by concatenating some of the given digits in any order. 
// If there is no answer return an empty string.

// Since the answer may not fit in an integer data type, return the answer as a string. 
// Note that the returning answer must not contain unnecessary leading zeros.

// Example 1:
// Input: digits = [8,1,9]
// Output: "981"

// Example 2:
// Input: digits = [8,6,7,1,0]
// Output: "8760"

// Example 3:
// Input: digits = [1]
// Output: ""

// Constraints:
//     1 <= digits.length <= 10^4
//     0 <= digits[i] <= 9

import "fmt"

func largestMultipleOfThree(digits []int) string {
    digitCount, mod3Count, sum := [10]int{}, [3]int{}, 0
    for _, d := range digits {
        digitCount[d]++
        mod3Count[d % 3]++
        sum += d
    }
    extra1, extra2 := 0, 0
    switch sum % 3 {
    case 1:
        if mod3Count[1] > 0 {
            extra1 = 1
        } else {
            extra2 = 2
        }
    case 2:
        if mod3Count[2] > 0 {
            extra2 = 1
        } else {
            extra1 = 2
        }
    case 0:
    }
    for digit, count := range digitCount {
        if digit % 3 == 1 && extra1 > 0 && count > 0 {
            if count >= extra1 {
                digitCount[digit] -= extra1
                extra1 = 0
            } else {
                extra1 -= count
                digitCount[digit] = 0
            }
        } else if digit % 3 == 2 && extra2 > 0 && count > 0 {
            if count >= extra2 {
                digitCount[digit] -= extra2
                extra2 = 0
            } else {
                extra2 -= count
                digitCount[digit] = 0
            }
        }
    }
    res := []byte{}
    for digit := 9; digit >= 0; digit-- {
        count := digitCount[digit]
        for i := 0; i < count; i++ {
            res = append(res, byte('0' + digit))
        }
    }
    if len(res) > 0 && res[0] == '0' {
        res = []byte{'0'}
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: digits = [8,1,9]
    // Output: "981"
    fmt.Println(largestMultipleOfThree([]int{8,1,9})) // "981"
    // Example 2:
    // Input: digits = [8,6,7,1,0]
    // Output: "8760"
    fmt.Println(largestMultipleOfThree([]int{8,6,7,1,0})) // "8760"
    // Example 3:
    // Input: digits = [1]
    // Output: ""
    fmt.Println(largestMultipleOfThree([]int{1})) // ""
}