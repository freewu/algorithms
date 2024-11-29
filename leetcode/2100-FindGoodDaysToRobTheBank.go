package main

// 2100. Find Good Days to Rob the Bank
// You and a gang of thieves are planning on robbing a bank. 
// You are given a 0-indexed integer array security, where security[i] is the number of guards on duty on the ith day. 
// he days are numbered starting from 0. You are also given an integer time.

// The ith day is a good day to rob the bank if:
//     There are at least time days before and after the ith day,
//     The number of guards at the bank for the time days before i are non-increasing, and
//     The number of guards at the bank for the time days after i are non-decreasing.

// More formally, this means day i is a good day to rob the bank if 
// and only if security[i - time] >= security[i - time + 1] >= ... >= security[i] <= ... <= security[i + time - 1] <= security[i + time].

// Return a list of all days (0-indexed) that are good days to rob the bank. 
// The order that the days are returned in does not matter.

// Example 1:
// Input: security = [5,3,3,3,5,6,2], time = 2
// Output: [2,3]
// Explanation:
// On day 2, we have security[0] >= security[1] >= security[2] <= security[3] <= security[4].
// On day 3, we have security[1] >= security[2] >= security[3] <= security[4] <= security[5].
// No other days satisfy this condition, so days 2 and 3 are the only good days to rob the bank.

// Example 2:
// Input: security = [1,1,1,1,1], time = 0
// Output: [0,1,2,3,4]
// Explanation:
// Since time equals 0, every day is a good day to rob the bank, so return every day.

// Example 3:
// Input: security = [1,2,3,4,5,6], time = 2
// Output: []
// Explanation:
// No day has 2 days before it that have a non-increasing number of guards.
// Thus, no day is a good day to rob the bank, so return an empty list.

// Constraints:
//     1 <= security.length <= 10^5
//     0 <= security[i], time <= 10^5

import "fmt"

func goodDaysToRobBank(security []int, time int) []int {
    res, n := []int{}, len(security)
    if n < 2 * time + 1{ return res }
    dp1, dp2 := make([]int, n), make([]int, n)
    for i := 1; i < n - time; i++ {
        if security[i-1] >= security[i] {
            dp1[i] = dp1[i-1] + 1
        }else{
            dp1[i] = 0
        }
    }
    for j := n - 2; j >= time; j-- {
        if security[j + 1] >= security[j] {
            dp2[j] = dp2[j+1] + 1 
        }else{
            dp2[j] = 0
        }
    }
    for i := time; i < n - time; i++ {
        if dp1[i] >= time && dp2[i] >= time {
            res = append(res,i)
        }   
    }
    return res
}

func goodDaysToRobBank1(security []int, time int) []int {
    res, n := []int{}, len(security)
    left, right := make([]int, n), make([]int, n)
    for i := 1; i < n; i++ {
        if security[i] <= security[i-1] {
            left[i] = left[i-1] + 1
        }
        if security[n - i - 1] <= security[n - i] {
            right[n - i - 1] = right[n - i] + 1
        }
    }
    for i := time; i < n - time; i++ {
        if left[i] >= time && right[i] >= time {
            res = append(res, i)
        }
    }
    return res
}

func goodDaysToRobBank2(security []int, time int) []int {
    res, left, right := []int{}, 0, 0
    for i := 0; i < len(security) - time; i++ {
        if i > 0 && security[i] <= security[i - 1] {
            left++
        } else {
            left = 0
        }
        if i + time - 1 >= 0 && security[i + time - 1] <= security[i + time] {
            right++
        } else {
            right = 0
        }
        if left >= time && right >= time {
            res = append(res, i)
        }
    }

    return res
}

func main() {
    // Example 1:
    // Input: security = [5,3,3,3,5,6,2], time = 2
    // Output: [2,3]
    // Explanation:
    // On day 2, we have security[0] >= security[1] >= security[2] <= security[3] <= security[4].
    // On day 3, we have security[1] >= security[2] >= security[3] <= security[4] <= security[5].
    // No other days satisfy this condition, so days 2 and 3 are the only good days to rob the bank.
    fmt.Println(goodDaysToRobBank([]int{5,3,3,3,5,6,2}, 2)) // [2,3]
    // Example 2:
    // Input: security = [1,1,1,1,1], time = 0
    // Output: [0,1,2,3,4]
    // Explanation:
    // Since time equals 0, every day is a good day to rob the bank, so return every day.
    fmt.Println(goodDaysToRobBank([]int{1,1,1,1,1}, 0)) // [0,1,2,3,4]
    // Example 3:
    // Input: security = [1,2,3,4,5,6], time = 2
    // Output: []
    // Explanation:
    // No day has 2 days before it that have a non-increasing number of guards.
    // Thus, no day is a good day to rob the bank, so return an empty list.
    fmt.Println(goodDaysToRobBank([]int{1,2,3,4,5,6}, 2)) // []

    fmt.Println(goodDaysToRobBank1([]int{5,3,3,3,5,6,2}, 2)) // [2,3]
    fmt.Println(goodDaysToRobBank1([]int{1,1,1,1,1}, 0)) // [0,1,2,3,4]
    fmt.Println(goodDaysToRobBank1([]int{1,2,3,4,5,6}, 2)) // []

    fmt.Println(goodDaysToRobBank2([]int{5,3,3,3,5,6,2}, 2)) // [2,3]
    fmt.Println(goodDaysToRobBank2([]int{1,1,1,1,1}, 0)) // [0,1,2,3,4]
    fmt.Println(goodDaysToRobBank2([]int{1,2,3,4,5,6}, 2)) // []
}