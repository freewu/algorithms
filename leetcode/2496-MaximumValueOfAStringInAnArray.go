package main

// 2496. Maximum Value of a String in an Array
// The value of an alphanumeric string can be defined as:
//     The numeric representation of the string in base 10, if it comprises of digits only.
//     The length of the string, otherwise.

// Given an array strs of alphanumeric strings, return the maximum value of any string in strs.

// Example 1:
// Input: strs = ["alic3","bob","3","4","00000"]
// Output: 5
// Explanation: 
// - "alic3" consists of both letters and digits, so its value is its length, i.e. 5.
// - "bob" consists only of letters, so its value is also its length, i.e. 3.
// - "3" consists only of digits, so its value is its numeric equivalent, i.e. 3.
// - "4" also consists only of digits, so its value is 4.
// - "00000" consists only of digits, so its value is 0.
// Hence, the maximum value is 5, of "alic3".

// Example 2:
// Input: strs = ["1","01","001","0001"]
// Output: 1
// Explanation: 
// Each string in the array has value 1. Hence, we return 1.

// Constraints:
//     1 <= strs.length <= 100
//     1 <= strs[i].length <= 9
//     strs[i] consists of only lowercase English letters and digits.

import "fmt"
import "strconv"

func maximumValue(strs []string) int {
    res := 0
    for _, str := range strs {
        for i, s := range str {
            if s >= 'a' {
                if len(str) > res {
                    res = len(str)
                }
                break
            } else if i == len(str) - 1 {
                v, _ := strconv.Atoi(str)
                if v > res {
                    res = v
                }
            }
        }
    }
    return res
}

func maximumValue1(strs []string) int {
    helper := func(s string) int {
        res := 0
        for _, v := range s {
            if v >= '0' && v <= '9' {
                res *= 10
                res += int(v - '0')
            } else {
                return len(s)
            }
        }
        return res
    }
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, str := range strs {
        res = max(res, helper(str))
    }
    return res
}

func main() {
    // Example 1:
    // Input: strs = ["alic3","bob","3","4","00000"]
    // Output: 5
    // Explanation: 
    // - "alic3" consists of both letters and digits, so its value is its length, i.e. 5.
    // - "bob" consists only of letters, so its value is also its length, i.e. 3.
    // - "3" consists only of digits, so its value is its numeric equivalent, i.e. 3.
    // - "4" also consists only of digits, so its value is 4.
    // - "00000" consists only of digits, so its value is 0.
    // Hence, the maximum value is 5, of "alic3".
    fmt.Println(maximumValue([]string{"alic3","bob","3","4","00000"})) // 5
    // Example 2:
    // Input: strs = ["1","01","001","0001"]
    // Output: 1
    // Explanation: 
    // Each string in the array has value 1. Hence, we return 1.
    fmt.Println(maximumValue([]string{"1","01","001","0001"})) // 1

    fmt.Println(maximumValue1([]string{"alic3","bob","3","4","00000"})) // 5
    fmt.Println(maximumValue1([]string{"1","01","001","0001"})) // 1
}