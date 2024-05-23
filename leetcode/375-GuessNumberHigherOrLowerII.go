package main

// 375. Guess Number Higher or Lower II
// We are playing the Guessing Game. The game will work as follows:
//     I pick a number between 1 and n.
//     You guess a number.
//     If you guess the right number, you win the game.
//     If you guess the wrong number, then I will tell you whether the number I picked is higher or lower, and you will continue guessing.
//     Every time you guess a wrong number x, you will pay x dollars. If you run out of money, you lose the game.

// Given a particular n, return the minimum amount of money you need to guarantee a win regardless of what number I pick.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/10/graph.png" />
// Input: n = 10
// Output: 16
// Explanation: The winning strategy is as follows:
// - The range is [1,10]. Guess 7.
//     - If this is my number, your total is $0. Otherwise, you pay $7.
//     - If my number is higher, the range is [8,10]. Guess 9.
//         - If this is my number, your total is $7. Otherwise, you pay $9.
//         - If my number is higher, it must be 10. Guess 10. Your total is $7 + $9 = $16.
//         - If my number is lower, it must be 8. Guess 8. Your total is $7 + $9 = $16.
//     - If my number is lower, the range is [1,6]. Guess 3.
//         - If this is my number, your total is $7. Otherwise, you pay $3.
//         - If my number is higher, the range is [4,6]. Guess 5.
//             - If this is my number, your total is $7 + $3 = $10. Otherwise, you pay $5.
//             - If my number is higher, it must be 6. Guess 6. Your total is $7 + $3 + $5 = $15.
//             - If my number is lower, it must be 4. Guess 4. Your total is $7 + $3 + $5 = $15.
//         - If my number is lower, the range is [1,2]. Guess 1.
//             - If this is my number, your total is $7 + $3 = $10. Otherwise, you pay $1.
//             - If my number is higher, it must be 2. Guess 2. Your total is $7 + $3 + $1 = $11.
// The worst case in all these scenarios is that you pay $16. Hence, you only need $16 to guarantee a win.

// Example 2:
// Input: n = 1
// Output: 0
// Explanation: There is only one possible number, so you can guess 1 and not have to pay anything.

// Example 3:
// Input: n = 2
// Output: 1
// Explanation: There are two possible numbers, 1 and 2.
// - Guess 1.
//     - If this is my number, your total is $0. Otherwise, you pay $1.
//     - If my number is higher, it must be 2. Guess 2. Your total is $1.
// The worst case is that you pay $1.
 
// Constraints:
//     1 <= n <= 200

import "fmt"

func getMoneyAmount(n int) int {
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    bottomUp := func(dp [][]int, begin, end int) int {
        for distance := 1; distance <= end-begin; distance++ {
            for i := begin; i+distance <= end; i++ {
                dp[i][i+distance] = i + dp[i+1][i+distance]
                for j := i + 1; j <= i+distance-1; j++ {
                    dp[i][i+distance] = min(dp[i][i+distance], j + max(dp[i][j-1], dp[j+1][i+distance]))
                }
            }
        }
        return dp[begin][end]
    }
    // var topDown func (dp [][]int, begin, end int) int
    // topDown = func (dp [][]int, begin, end int) int {
    //     if begin >= end {
    //         return 0
    //     }
    //     if dp[begin][end] != 0 {
    //         return dp[begin][end]
    //     }
    //     dp[begin][end] = math.MaxInt
    //     for i := begin; i <= end; i++ {
    //         dp[begin][end] = min(dp[begin][end], i + max(topDown(dp, begin, i-1), topDown(dp, i+1, end)))
    //     }
    //     return dp[begin][end]
    // }
    //return topDown(dp, 1, n)
    return bottomUp(dp, 1, n)
}

func getMoneyAmount1(n int) int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 1; i-- {
        for j := i + 1; j <= n; j++ {
            dp[i][j] = j + dp[i][j-1]
            for k := i; k < j; k++ {
                cost := k + max(dp[i][k-1], dp[k+1][j])
                if cost < dp[i][j] {
                    dp[i][j] = cost
                }
            }
        }
    }
    return dp[1][n]
}

func getMoneyAmount2(n int) int {
    dp, inf := make([][]int, n+1),1 << 63 - 1
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for l := 2; l <= n; l++ {
        for start := 1; start < n+2-l; start++ {
            mn := inf
            for piv := start+(l-1)/2; piv < start+l-1; piv++ {
                res := piv + max(dp[start][piv-1], dp[piv+1][start+l-1])
                mn = min(res, mn)
            }
            dp[start][start+l-1] = mn
        }       
    }
    return dp[1][n]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/10/graph.png" />
    // Input: n = 10
    // Output: 16
    // Explanation: The winning strategy is as follows:
    // - The range is [1,10]. Guess 7.
    //     - If this is my number, your total is $0. Otherwise, you pay $7.
    //     - If my number is higher, the range is [8,10]. Guess 9.
    //         - If this is my number, your total is $7. Otherwise, you pay $9.
    //         - If my number is higher, it must be 10. Guess 10. Your total is $7 + $9 = $16.
    //         - If my number is lower, it must be 8. Guess 8. Your total is $7 + $9 = $16.
    //     - If my number is lower, the range is [1,6]. Guess 3.
    //         - If this is my number, your total is $7. Otherwise, you pay $3.
    //         - If my number is higher, the range is [4,6]. Guess 5.
    //             - If this is my number, your total is $7 + $3 = $10. Otherwise, you pay $5.
    //             - If my number is higher, it must be 6. Guess 6. Your total is $7 + $3 + $5 = $15.
    //             - If my number is lower, it must be 4. Guess 4. Your total is $7 + $3 + $5 = $15.
    //         - If my number is lower, the range is [1,2]. Guess 1.
    //             - If this is my number, your total is $7 + $3 = $10. Otherwise, you pay $1.
    //             - If my number is higher, it must be 2. Guess 2. Your total is $7 + $3 + $1 = $11.
    // The worst case in all these scenarios is that you pay $16. Hence, you only need $16 to guarantee a win.
    fmt.Println(getMoneyAmount(10)) // 16
    // Example 2:
    // Input: n = 1
    // Output: 0
    // Explanation: There is only one possible number, so you can guess 1 and not have to pay anything.
    fmt.Println(getMoneyAmount(1)) // 0
    // Example 3:
    // Input: n = 2
    // Output: 1
    // Explanation: There are two possible numbers, 1 and 2.
    // - Guess 1.
    //     - If this is my number, your total is $0. Otherwise, you pay $1.
    //     - If my number is higher, it must be 2. Guess 2. Your total is $1.
    // The worst case is that you pay $1.
    fmt.Println(getMoneyAmount(2)) // 1

    fmt.Println(getMoneyAmount1(10)) // 16
    fmt.Println(getMoneyAmount1(1)) // 0
    fmt.Println(getMoneyAmount1(2)) // 1
    
    fmt.Println(getMoneyAmount2(10)) // 16
    fmt.Println(getMoneyAmount2(1)) // 0
    fmt.Println(getMoneyAmount2(2)) // 1
}