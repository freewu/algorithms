package main

// 504. Base 7
// Given an integer num, return a string of its base 7 representation.

// Example 1:
// Input: num = 100
// Output: "202"

// Example 2:
// Input: num = -7
// Output: "-10"
 
// Constraints:
//     -10^7 <= num <= 10^7

import "fmt"
import "strconv"

func convertToBase7(num int) string {
    if num == 0 {
        return "0"
    }
    negative := false // 处理负数
    if num < 0 {
        negative = true
        num = -num
    }
    res, nums := "", []int{}
    for num != 0 {
        remainder := num % 7 // 按 7 取余
        nums = append(nums, remainder)
        num = num / 7
    }
    if negative {
        res += "-"
    }
    for i := len(nums) - 1; i >= 0; i-- {
        res += strconv.Itoa(nums[i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = 100
    // Output: "202"
    fmt.Println(convertToBase7(100)) // "202"
    // Example 2:
    // Input: num = -7
    // Output: "-10"
    fmt.Println(convertToBase7(7)) // "-10"
}