package main

// 3908. Valid Digit Number
// You are given an integer n and a digit x.

// A number is considered valid if:
//     1. It contains at least one occurrence of digit x, and
//     2. It does not start with digit x.

// Return true if n is valid, otherwise return false.

// Example 1:
// Input: n = 101, x = 0
// Output: true
// Explanation:
// The number contains digit 0 at index 1. It does not start with 0, so it satisfies both conditions. Thus, the answer is true​​​​​​​.

// Example 2:
// Input: n = 232, x = 2
// Output: false
// Explanation:
// The number starts with 2, which violates the condition. Thus, the answer is false.

// Example 3:
// Input: n = 5, x = 1
// Output: false
// Explanation:
// The number does not contain digit 1. Thus, the answer is false.

// Constraints:
//     0 <= n <= 10^5​​​​​​​
//     0 <= x <= 9

import "fmt"
import "strconv"
import "strings"

func validDigit(n int, x int) bool {
    s := strconv.Itoa(n)
    digit := strconv.Itoa(x)
    // 包含 至少一个 数字 x，并且 不以 数字 x 开头。
    return strings.Contains(s, digit) && s[0] != digit[0]
}

func validDigit1(n int, x int) bool {
    digitArray := func (n int) []int{
        res := make([]int,0)
        for {
            b := n % 10
            res = append(res,b)
            n = int(n / 10)
            if n == 0 {
                break
            }
        }
        return res
    }
    arr := digitArray(n)
    if arr[len(arr)-1] == x { // 不以 数字 x 开头。
        return false
    }
    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] == x { // 包含 至少一个 数字 x。
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: n = 101, x = 0
    // Output: true
    // Explanation:
    // The number contains digit 0 at index 1. It does not start with 0, so it satisfies both conditions. Thus, the answer is true​​​​​​​.
    fmt.Println(validDigit(101, 0)) // true
    // Example 2:
    // Input: n = 232, x = 2
    // Output: false
    // Explanation:
    // The number starts with 2, which violates the condition. Thus, the answer is false.
    fmt.Println(validDigit(232, 2)) // false
    // Example 3:
    // Input: n = 5, x = 1
    // Output: false
    // Explanation:
    // The number does not contain digit 1. Thus, the answer is false.
    fmt.Println(validDigit(5, 1)) // false

    fmt.Println(validDigit(0, 0)) // false
    fmt.Println(validDigit(0, 9)) // false
    fmt.Println(validDigit(100_000, 0)) // true
    fmt.Println(validDigit(100_000, 9)) // false

    fmt.Println(validDigit1(101, 0)) // true
    fmt.Println(validDigit1(232, 2)) // false
    fmt.Println(validDigit1(5, 1)) // false
    fmt.Println(validDigit1(0, 0)) // false
    fmt.Println(validDigit1(0, 9)) // false
    fmt.Println(validDigit1(100_000, 0)) // true
    fmt.Println(validDigit1(100_000, 9)) // false
}
