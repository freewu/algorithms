package main

// 1140. Stone Game II
// Alice and Bob continue their games with piles of stones.  
// There are a number of piles arranged in a row, and each pile has a positive integer number of stones piles[i].  
// The objective of the game is to end with the most stones. 

// Alice and Bob take turns, with Alice starting first. Initially, M = 1.

// On each player's turn, that player can take all the stones in the first X remaining piles, where 1 <= X <= 2M.  Then, we set M = max(M, X).
// The game continues until all the stones have been taken.
// Assuming Alice and Bob play optimally, return the maximum number of stones Alice can get.

// Example 1:
// Input: piles = [2,7,9,4,4]
// Output: 10
// Explanation:  If Alice takes one pile at the beginning, Bob takes two piles, then Alice takes 2 piles again. Alice can get 2 + 4 + 4 = 10 piles in total. If Alice takes two piles at the beginning, then Bob can take all three piles left. In this case, Alice get 2 + 7 = 9 piles in total. So we return 10 since it's larger. 

// Example 2:
// Input: piles = [1,2,3,4,5,100]
// Output: 104
 
// Constraints:
//     1 <= piles.length <= 100
//     1 <= piles[i] <= 10^4

import "fmt"

func stoneGameII(piles []int) int {
    l := len(piles)
    sum, dp := make([]int, l + 1), make([][]int, l + 1)
    for i := 0; i < l; i++ {
        sum[i+1] = sum[i] + piles[i]
    }
    for i := 0; i < l+1; i++ {
        dp[i] = make([]int, l + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := l - 1; i >= 0; i-- {
        for j := l; j >= 1; j-- {
            for k := 1; k <= 2 * j; k++ {
                if i + k >= l {
                    dp[i][j] = max(dp[i][j], sum[l] - sum[i])
                    break
                } else {
                    dp[i][j] = max(dp[i][j], sum[l] - sum[i] - dp[i+k][max(k, j)])
                }
            }
        }
    }
    return dp[0][1]
}

func stoneGameII1(piles []int) int {
    n, inf := len(piles), 1 << 32 - 1
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, s := n - 1, 0; i >= 0; i-- {
        s += piles[i]
        for m := 1; m <= i/2+1; m++ {
            if i + m * 2 >= n { // 全拿
                dp[i][m] = s
            } else {
                mn := inf
                for x := 1; x <= m*2; x++ {
                    mn = min(mn, dp[i+x][max(m, x)])
                }
                dp[i][m] = s - mn
            }
        }
    }
    return dp[0][1]
}

func main() {
    // Example 1:
    // Input: piles = [2,7,9,4,4]
    // Output: 10
    // Explanation:  
    // If Alice takes one pile at the beginning, Bob takes two piles, 
    // then Alice takes 2 piles again. Alice can get 2 + 4 + 4 = 10 piles in total. 
    // If Alice takes two piles at the beginning, then Bob can take all three piles left. 
    // In this case, Alice get 2 + 7 = 9 piles in total. So we return 10 since it's larger. 
    fmt.Println(stoneGameII([]int{ 2,7,9,4,4 })) // 10
    // Example 2:
    // Input: piles = [1,2,3,4,5,100]
    // Output: 104
    fmt.Println(stoneGameII([]int{ 1,2,3,4,5,100 })) // 104

    fmt.Println(stoneGameII1([]int{ 2,7,9,4,4 })) // 10
    fmt.Println(stoneGameII1([]int{ 1,2,3,4,5,100 })) // 104
}