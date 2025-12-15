package main

// 2110. Number of Smooth Descent Periods of a Stock
// You are given an integer array prices representing the daily price history of a stock, 
// where prices[i] is the stock price on the ith day.

// A smooth descent period of a stock consists of one or more contiguous days 
// such that the price on each day is lower than the price on the preceding day by exactly 1. 
// The first day of the period is exempted from this rule.

// Return the number of smooth descent periods.

// Example 1:
// Input: prices = [3,2,1,4]
// Output: 7
// Explanation: There are 7 smooth descent periods:
// [3], [2], [1], [4], [3,2], [2,1], and [3,2,1]
// Note that a period with one day is a smooth descent period by the definition.

// Example 2:
// Input: prices = [8,6,7,7]
// Output: 4
// Explanation: There are 4 smooth descent periods: [8], [6], [7], and [7]
// Note that [8,6] is not a smooth descent period as 8 - 6 ≠ 1.

// Example 3:
// Input: prices = [1]
// Output: 1
// Explanation: There is 1 smooth descent period: [1]

// Constraints:
//     1 <= prices.length <= 10^5
//     1 <= prices[i] <= 10^5

import "fmt"

// brute force 超出时间限制 166 / 168
func getDescentPeriods(prices []int) int64 {
    res, n := 0, len(prices)
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            if j == i {
                res++
            } else {
                if prices[j] + 1 != prices[j - 1] { break }
                res++
            }
        }
    }
    return int64(res)
}

// two pointer
func getDescentPeriods1(prices []int) int64 {
    res, i, j, n := 1, 0, 1, len(prices)
    for j < n {
        if prices[j] + 1 == prices[j - 1] {
            res += (j - i + 1)  // [4,3,2,1] -> {[4,3,2,1], [3,2,1], [2,1], [1]}
        }else{
            i = j
            res++
        }
        j++
    }
    return int64(res)
}

func getDescentPeriods2(prices []int) int64 {
    res, prev := 1, 1
    for i := 1; i < len(prices); i++ {
        if prices[i] == prices[i - 1] - 1 {
            prev++
        } else {
            prev=1
        }
        res += prev
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: prices = [3,2,1,4]
    // Output: 7
    // Explanation: There are 7 smooth descent periods:
    // [3], [2], [1], [4], [3,2], [2,1], and [3,2,1]
    // Note that a period with one day is a smooth descent period by the definition.
    fmt.Println(getDescentPeriods([]int{3,2,1,4})) // 7
    // Example 2:
    // Input: prices = [8,6,7,7]
    // Output: 4
    // Explanation: There are 4 smooth descent periods: [8], [6], [7], and [7]
    // Note that [8,6] is not a smooth descent period as 8 - 6 ≠ 1.
    fmt.Println(getDescentPeriods([]int{8,6,7,7})) // 4
    // Example 3:
    // Input: prices = [1]
    // Output: 1
    // Explanation: There is 1 smooth descent period: [1]
    fmt.Println(getDescentPeriods([]int{1})) // 1

    fmt.Println(getDescentPeriods([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(getDescentPeriods([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(getDescentPeriods1([]int{3,2,1,4})) // 7
    fmt.Println(getDescentPeriods1([]int{8,6,7,7})) // 4
    fmt.Println(getDescentPeriods1([]int{1})) // 1
    fmt.Println(getDescentPeriods1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(getDescentPeriods1([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(getDescentPeriods2([]int{3,2,1,4})) // 7
    fmt.Println(getDescentPeriods2([]int{8,6,7,7})) // 4
    fmt.Println(getDescentPeriods2([]int{1})) // 1
    fmt.Println(getDescentPeriods2([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(getDescentPeriods2([]int{9,8,7,6,5,4,3,2,1})) // 45
}