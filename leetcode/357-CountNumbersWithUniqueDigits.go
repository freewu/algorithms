package main

// 357. Count Numbers with Unique Digits
// Given an integer n, return the count of all numbers with unique digits, x, where 0 <= x < 10n.

// Example 1:
// Input: n = 2
// Output: 91
// Explanation: The answer should be the total numbers in the range of 0 ≤ x < 100, excluding 11,22,33,44,55,66,77,88,99

// Example 2:
// Input: n = 0
// Output: 1
 
// Constraints:
//     0 <= n <= 8

import "fmt"

// 暴力打表法
func countNumbersWithUniqueDigits1(n int) int {
    res := []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}
    if n >= 10 {
        return res[10]
    }
    return res[n]
}

// 如果只是一位数，不存在重复的数字，结果是 10 。
// 如果是二位数，第一位一定不能取 0，那么第一位有 1-9，9 种取法，第二位为了和第一位不重复，只能有 0-9，10 种取法中减去第一位取的数字，那么也是 9 种取法。
// 如果是三位数，第三位是 8 种取法；
// 四位数，第四位是 7 种取法；
// 五位数，第五位是 6 种取法；
// 六位数，第六位是 5 种取法；
// 七位数，第七位是 4 种取法；
// 八位数，第八位是 3 种取法；
// 九位数，第九位是 2 种取法；
// 十位数，第十位是 1 种取法；
// 十一位数，第十一位是 0 种取法；十二位数，第十二位是 0 种取法
func countNumbersWithUniqueDigits(n int) int {
    if n == 0 {
        return 1
    }
    res, uniqueDigits, availableNumber := 10, 9, 9
    for n > 1 && availableNumber > 0 {
        uniqueDigits = uniqueDigits * availableNumber
        res += uniqueDigits
        availableNumber--
        n--
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 91
    // Explanation: The answer should be the total numbers in the range of 0 ≤ x < 100, excluding 11,22,33,44,55,66,77,88,99
    fmt.Println(countNumbersWithUniqueDigits(2)) // 91
    // Example 2:
    // Input: n = 0
    // Output: 1
    fmt.Println(countNumbersWithUniqueDigits(0)) // 1
    fmt.Println(countNumbersWithUniqueDigits(1)) // 10
    fmt.Println(countNumbersWithUniqueDigits(3)) // 739
    fmt.Println(countNumbersWithUniqueDigits(4)) // 5275
    fmt.Println(countNumbersWithUniqueDigits(5)) // 32491
    fmt.Println(countNumbersWithUniqueDigits(6)) // 168571
    fmt.Println(countNumbersWithUniqueDigits(7)) // 712891
    fmt.Println(countNumbersWithUniqueDigits(8)) // 2345851

    fmt.Println(countNumbersWithUniqueDigits1(0)) // 1
    fmt.Println(countNumbersWithUniqueDigits1(1)) // 10
    fmt.Println(countNumbersWithUniqueDigits1(2)) // 91
    fmt.Println(countNumbersWithUniqueDigits1(3)) // 739
    fmt.Println(countNumbersWithUniqueDigits1(4)) // 5275
    fmt.Println(countNumbersWithUniqueDigits1(5)) // 32491
    fmt.Println(countNumbersWithUniqueDigits1(6)) // 168571
    fmt.Println(countNumbersWithUniqueDigits1(7)) // 712891
    fmt.Println(countNumbersWithUniqueDigits1(8)) // 2345851
}