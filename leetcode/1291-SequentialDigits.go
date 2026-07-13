package main

// 1291. Sequential Digits
// An integer has sequential digits if and only if each digit in the number is one more than the previous digit.
// Return a sorted list of all the integers in the range [low, high] inclusive that have sequential digits.

// Example 1:
// Input: low = 100, high = 300
// Output: [123,234]

// Example 2:
// Input: low = 1000, high = 13000
// Output: [1234,2345,3456,4567,5678,6789,12345]
 
// Constraints:
//     10 <= low <= high <= 10^9

// 范围最大为 9 位数，所以直接枚举即可。采用的方法是滑动窗口；
// 定义最长串  s := "123456789";
// 然后以最小长度到最大长长度在最长串上截取

import "fmt"
import "strconv"

func sequentialDigits(low int, high int) []int {
    res := make([]int, 0)
    s := "123456789";
    mn, max := len(strconv.Itoa(low)), len(strconv.Itoa(high))
    for mn <= max {
        temp := 0;
        for i := 0; i <= 9 - mn; i = i + 1 {
            temp, _ = strconv.Atoi(s[i : i + mn])
            if temp < low {
                continue
            } else if temp > high {
                break
            } else {
                res = append(res, temp)
            }
        }
        mn = mn + 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: low = 100, high = 300
    // Output: [123,234]
    fmt.Println(sequentialDigits(100,300)) // [123,234]
    // Example 2:
    // Input: low = 1000, high = 13000
    // Output: [1234,2345,3456,4567,5678,6789,12345]
    fmt.Println(sequentialDigits(1000,13000)) // [1234,2345,3456,4567,5678,6789,12345]

    fmt.Println(sequentialDigits(10,10)) // []
    fmt.Println(sequentialDigits(10,1_000_000_000)) // [12 23 34 45 56 67 78 89 123 234 345 456 567 678 789 1234 2345 3456 4567 5678 6789 12345 23456 34567 45678 56789 123456 234567 345678 456789 1234567 2345678 3456789 12345678 23456789 123456789]
    fmt.Println(sequentialDigits(1_000_000_000,1_000_000_000)) // []
}