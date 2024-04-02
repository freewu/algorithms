package main

// 441. Arranging Coins
// You have n coins and you want to build a staircase with these coins. 
// The staircase consists of k rows where the ith row has exactly i coins. 
// The last row of the staircase may be incomplete.

// Given the integer n, return the number of complete rows of the staircase you will build.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/arrangecoins1-grid.jpg" />
// Input: n = 5
// Output: 2
// Explanation: Because the 3rd row is incomplete, we return 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/arrangecoins2-grid.jpg" />
// Input: n = 8
// Output: 3
// Explanation: Because the 4th row is incomplete, we return 3.
 
// Constraints:
//     1 <= n <= 2^31 - 1

import "fmt"

func arrangeCoins(n int) int {
    arrange := func (start int, end int, n int) int {
        res := 0
        for start <= end {
            mid := end + (start - end) / 2 
            sum := mid * (mid + 1) / 2
            if sum > n {
                end = mid - 1
            } else {
                res = mid
                start = mid + 1
            }
        }
        return res
    }
    return arrange(0, n, n)
}

func main() {
    // Explanation: Because the 3rd row is incomplete, we return 2.
    fmt.Println(arrangeCoins(5)) // 2
    // Explanation: Because the 4th row is incomplete, we return 3.
    fmt.Println(arrangeCoins(8)) // 3
}