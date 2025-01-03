package main

// 2384. Largest Palindromic Number
// You are given a string num consisting of digits only.

// Return the largest palindromic integer (in the form of a string) that can be formed using digits taken from num. 
// It should not contain leading zeroes.

// Notes:
//     You do not need to use all the digits of num, but you must use at least one digit.
//     The digits can be reordered.

// Example 1:
// Input: num = "444947137"
// Output: "7449447"
// Explanation: 
// Use the digits "4449477" from "444947137" to form the palindromic integer "7449447".
// It can be shown that "7449447" is the largest palindromic integer that can be formed.

// Example 2:
// Input: num = "00009"
// Output: "9"
// Explanation: 
// It can be shown that "9" is the largest palindromic integer that can be formed.
// Note that the integer returned should not contain leading zeroes.

// Constraints:
//     1 <= num.length <= 10^5
//     num consists of digits.

import "fmt"
import "bytes"

func largestPalindromic(num string) string {
    mp := map[rune]int{}
    for _, v := range num {
        mp[v]++
    }
    left, right, center := "", "", ""
    for _, k := range "9876543210" {
        v := mp[k]
        for i := 0; i < v / 2; i++ {
            left += string(k)
            right = string(k) + right
        }
        if v % 2 == 1 {
            if center == "" {
                center = string(k)
            }
        }
    }
    if len(left) > 0 && left[0] == '0' {
        if center == "" {
            return "0"
        }
        return center
    }
    return left + center + right
}

func largestPalindromic1(num string) string {
    res, arr :=  []byte{}, make([]int, 10)
    for _, v := range num {
        arr[int(v - '0')]++
    }
    for i := 9; i >= 0; i-- {
        if len(res) == 0 && i == 0 { continue }
        if arr[i] >= 2 {
            res = append(res, bytes.Repeat([]byte{byte(i + '0')}, arr[i] / 2)...)
            arr[i] %= 2
        }
    }
    if len(res) > 0 && arr[0] >= 2 {
        res = append(res, bytes.Repeat([]byte{'0'}, arr[0]/2)...)
        arr[0] %= 2
    }
    mid := len(res) - 1
    for i := 9; i >= 0; i-- {
        if arr[i] > 0 {
            res = append(res, byte(i+'0'))
            break
        }
    }
    for i := mid; i >= 0; i-- {
        res = append(res, res[i])
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: num = "444947137"
    // Output: "7449447"
    // Explanation: 
    // Use the digits "4449477" from "444947137" to form the palindromic integer "7449447".
    // It can be shown that "7449447" is the largest palindromic integer that can be formed.
    fmt.Println(largestPalindromic("444947137")) // "7449447"
    // Example 2:
    // Input: num = "00009"
    // Output: "9"
    // Explanation: 
    // It can be shown that "9" is the largest palindromic integer that can be formed.
    // Note that the integer returned should not contain leading zeroes.
    fmt.Println(largestPalindromic("00009")) // "9"

    fmt.Println(largestPalindromic1("444947137")) // "7449447"
    fmt.Println(largestPalindromic1("00009")) // "9"
}