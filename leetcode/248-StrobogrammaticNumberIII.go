package main

// 248. Strobogrammatic Number III
// Given two strings low and high that represent two integers low and high where low <= high, return the number of strobogrammatic numbers in the range [low, high].
// A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

// Example 1:
// Input: low = "50", high = "100"
// Output: 3

// Example 2:
// Input: low = "0", high = "0"
// Output: 1

// Constraints:
//     1 <= low.length, high.length <= 15
//     low and high consist of only digits.
//     low <= high
//     low and high do not contain any leading zeros except for zero itself.

import "fmt"

func strobogrammaticInRange(low string, high string) int {
    res, queue := 0, []string{"", "0", "1", "8"}
    gte := func (num1, num2 string) bool { // 判断num1 是否 大于等于 num2
        if len(num1) != len(num2) {
            return len(num1) > len(num2)
        }
        return num1 >= num2
    }
    for len(queue) > 0 {
        num := queue[len(queue)-1]
        queue = queue[:len(queue)-1]
        if !(len(num) > 1 && num[0] == '0') && gte(num, low) && gte(high, num) { // 排除多位数以 0 开头
            res++
        }
        if len(num) + 2 <= len(high) {
            queue = append(queue, "0" + num + "0")
            queue = append(queue, "1" + num + "1")
            queue = append(queue, "6" + num + "9")
            queue = append(queue, "8" + num + "8")
            queue = append(queue, "9" + num + "6")
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: low = "50", high = "100"
    // Output: 3
    fmt.Println(strobogrammaticInRange("50","100")) // 3  [69,88,96]
    // Example 2:
    // Input: low = "0", high = "0"
    // Output: 1
    fmt.Println(strobogrammaticInRange("0","0")) // 1  [0]
}