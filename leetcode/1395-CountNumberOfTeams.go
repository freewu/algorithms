package main

// 1395. Count Number of Teams
// There are n soldiers standing in a line. Each soldier is assigned a unique rating value.
// You have to form a team of 3 soldiers amongst them under the following rules:
//     Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).
//     A team is valid if: (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).

// Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).

// Example 1:
// Input: rating = [2,5,3,4,1]
// Output: 3
// Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1). 

// Example 2:
// Input: rating = [2,1,3]
// Output: 0
// Explanation: We can't form any team given the conditions.

// Example 3:
// Input: rating = [1,2,3,4]
// Output: 4

// Constraints:
//     n == rating.length
//     3 <= n <= 1000
//     1 <= rating[i] <= 10^5
//     All the integers in rating are unique.

import "fmt"

// 中间位 从 1 到 n-1，再分别找左右两边大于小于该数的数目
func numTeams(rating []int) int {
    res, n := 0, len(rating)
    for i := 1; i < n - 1; i++ {
        iLess, iMore, jLess, jMore := 0,0,0,0
        for j := 0; j < i ;j ++ {
            if rating[j] > rating[i] {
                iMore++
            } else {
                iLess++
            }
        }
        for j := i + 1; j < n; j++ {
            if rating[j] > rating[i] {
                jMore++
            } else {
                jLess++
            }
        }
        res += iMore * jLess + iLess * jMore
    }
    return res
}

// dp 
// 遍历数组，用dp[i][0]存储大于/小于该数的数目，即顺序的两元组的个数；
// 用dp[i][1]存储顺序的三元组的个数，最后只需要求dp[i][1]的和即可
func numTeams1(rating []int) int {
    res, n := 0, len(rating)
    dp1, dp2 := make([][]int,n), make([][]int,n)
    for i := 0; i < n; i++ {
        dp1[i], dp2[i] = []int{0,0}, []int{0,0}
        for j := 0; j < i; j++ {
            if rating[i] > rating[j] {
                dp1[i][0]++
                dp1[i][1] += dp1[j][0]
            } else {
                dp2[i][0]++
                dp2[i][1] += dp2[j][0]
            }
        }
    }
    for k := range dp1 {
        res += dp1[k][1]
        res += dp2[k][1]
    }
    return res
}

func main() {
    // Example 1:
    // Input: rating = [2,5,3,4,1]
    // Output: 3
    // Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1). 
    fmt.Println(numTeams([]int{2,5,3,4,1})) // 3
    // Example 2:
    // Input: rating = [2,1,3]
    // Output: 0
    // Explanation: We can't form any team given the conditions.
    fmt.Println(numTeams([]int{2,1,3})) // 0
    // Example 3:
    // Input: rating = [1,2,3,4]
    // Output: 4
    fmt.Println(numTeams([]int{1,2,3,4})) // 4

    fmt.Println(numTeams1([]int{2,5,3,4,1})) // 3
    fmt.Println(numTeams1([]int{2,1,3})) // 0
    fmt.Println(numTeams1([]int{1,2,3,4})) // 4
}