package main

// 2554. Maximum Number of Integers to Choose From a Range I
// You are given an integer array banned and two integers n and maxSum. 
// You are choosing some number of integers following the below rules:
//     The chosen integers have to be in the range [1, n].
//     Each integer can be chosen at most once.
//     The chosen integers should not be in the array banned.
//     The sum of the chosen integers should not exceed maxSum.

// Return the maximum number of integers you can choose following the mentioned rules.

// Example 1:
// Input: banned = [1,6,5], n = 5, maxSum = 6
// Output: 2
// Explanation: You can choose the integers 2 and 4.
// 2 and 4 are from the range [1, 5], both did not appear in banned, and their sum is 6, which did not exceed maxSum.

// Example 2:
// Input: banned = [1,2,3,4,5,6,7], n = 8, maxSum = 1
// Output: 0
// Explanation: You cannot choose any integer while following the mentioned conditions.

// Example 3:
// Input: banned = [11], n = 7, maxSum = 50
// Output: 7
// Explanation: You can choose the integers 1, 2, 3, 4, 5, 6, and 7.
// They are from the range [1, 7], all did not appear in banned, and their sum is 28, which did not exceed maxSum.

// Constraints:
//     1 <= banned.length <= 10^4
//     1 <= banned[i], n <= 10^4
//     1 <= maxSum <= 10^9

import "fmt"
import "sort"

func maxCount(banned []int, n int, maxSum int) int {
    visited := map[int]bool{}
    for _, v := range banned {
        visited[v] = true
    }
    res, sum := 0, 0
    for i := 1; i <= n && sum + i <= maxSum; i++ {
        if !visited[i] {
            res++
            sum += i
        }
    }
    return res
}

func maxCount1(banned []int, n int, maxSum int) int {
    res, visited := 0, make(map[int]bool)
    for _, v := range banned {
        visited[v] = true
    }
    for i := 1; i <= n && maxSum >= i; i++ {
        if visited[i] { continue }
        res++
        maxSum -= i
    }
    return res 
}

func maxCount2(banned []int, n int, maxSum int) int {
    sort.Ints(banned)
    res, sum := 0, 0
    for i := 1; i <= n; i++ {
        if len(banned) > 0 && i == banned[0] {
            for len(banned) > 0 && i == banned[0] {
                banned = banned[1: ]
            }
        } else {
            sum += i
            res++
        }
        if sum > maxSum {
            return res - 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: banned = [1,6,5], n = 5, maxSum = 6
    // Output: 2
    // Explanation: You can choose the integers 2 and 4.
    // 2 and 4 are from the range [1, 5], both did not appear in banned, and their sum is 6, which did not exceed maxSum.
    fmt.Println(maxCount([]int{1,6,5}, 5, 6)) // 2
    // Example 2:
    // Input: banned = [1,2,3,4,5,6,7], n = 8, maxSum = 1
    // Output: 0
    // Explanation: You cannot choose any integer while following the mentioned conditions.
    fmt.Println(maxCount([]int{1,2,3,4,5,6,7}, 8, 1)) // 0
    // Example 3:
    // Input: banned = [11], n = 7, maxSum = 50
    // Output: 7
    // Explanation: You can choose the integers 1, 2, 3, 4, 5, 6, and 7.
    // They are from the range [1, 7], all did not appear in banned, and their sum is 28, which did not exceed maxSum.
    fmt.Println(maxCount([]int{11}, 7, 50)) // 7

    fmt.Println(maxCount1([]int{1,6,5}, 5, 6)) // 2
    fmt.Println(maxCount1([]int{1,2,3,4,5,6,7}, 8, 1)) // 0
    fmt.Println(maxCount1([]int{11}, 7, 50)) // 7

    fmt.Println(maxCount2([]int{1,6,5}, 5, 6)) // 2
    fmt.Println(maxCount2([]int{1,2,3,4,5,6,7}, 8, 1)) // 0
    fmt.Println(maxCount2([]int{11}, 7, 50)) // 7
}