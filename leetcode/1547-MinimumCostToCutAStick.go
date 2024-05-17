package main

// 1547. Minimum Cost to Cut a Stick
// Given a wooden stick of length n units. The stick is labelled from 0 to n. 
// For example, a stick of length 6 is labelled as follows:
// <img src="https://assets.leetcode.com/uploads/2020/07/21/statement.jpg" />

// Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.
// You should perform the cuts in order, you can change the order of the cuts as you wish.
// The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts. 
// When you cut a stick, it will be split into two smaller sticks (i.e. the sum of their lengths is the length of the stick before the cut). 
// Please refer to the first example for a better explanation.

// Return the minimum total cost of the cuts.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/23/e1.jpg" />
// Input: n = 7, cuts = [1,3,4,5]
// Output: 16
// Explanation: Using cuts order = [1, 3, 4, 5] as in the input leads to the following scenario:
// <img src="https://assets.leetcode.com/uploads/2020/07/21/e11.jpg" />
// The first cut is done to a rod of length 7 so the cost is 7. The second cut is done to a rod of length 6 (i.e. the second part of the first cut), the third is done to a rod of length 4 and the last cut is to a rod of length 3. The total cost is 7 + 6 + 4 + 3 = 20.
// Rearranging the cuts to be [3, 5, 1, 4] for example will lead to a scenario with total cost = 16 (as shown in the example photo 7 + 4 + 3 + 2 = 16).

// Example 2:
// Input: n = 9, cuts = [5,6,1,4,2]
// Output: 22
// Explanation: If you try the given cuts ordering the cost will be 25.
// There are much ordering with total cost <= 25, for example, the order [4, 6, 5, 2, 1] has total cost = 22 which is the minimum possible.
 
// Constraints:
//     2 <= n <= 10^6
//     1 <= cuts.length <= min(n - 1, 100)
//     1 <= cuts[i] <= n - 1
//     All the integers in cuts array are distinct.

import "fmt"
import "sort"

// Recursive  dfs
func minCost(size int, cuts []int) int {
    sort.Ints(cuts)
    cuts = append([]int{0}, cuts...)
    cuts = append(cuts, size)

    n := len(cuts) - 1
    inf, dp := size * (n + 1), make([][]int, n)
    for start := range dp {
        dp[start] = make([]int, n + 1)
        for end := start; end <= n; end++ {
            dp[start][end] = inf
        }
        dp[start][start + 1] = 0
    }
    var calculate func(int, int) int
    calculate = func(start, end int) int {
        if dp[start][end] < inf {
            return dp[start][end]
        }
        length := cuts[end] - cuts[start]
        for mid := start + 1; mid < end; mid++ {
            cost := calculate(start, mid) + calculate(mid, end) + length
            if dp[start][end] > cost {
                dp[start][end] = cost
            }
        }
        return dp[start][end]
    }
    return calculate(0, n)
}

// Iterative  bfs
func minCost1(size int, cuts []int) int {
    sort.Ints(cuts)
    cuts = append([]int{0}, cuts...)
    cuts = append(cuts, size)
    n := len(cuts) - 1

    inf, dp := size * (n + 1), make([][]int, n)
    for start := range dp {
        dp[start] = make([]int, n + 1)
        for end := start; end <= n; end++ {
            dp[start][end] = inf
        }
        dp[start][start + 1] = 0
    }
    for shift := 2; shift <= n; shift++ {
        for end := shift; end <= n; end++ {
            start := end - shift
            length := cuts[end] - cuts[start]
            for mid := start + 1; mid < end; mid++ {
                cost := dp[start][mid] + dp[mid][end] + length
                if dp[start][end] > cost {
                    dp[start][end] = cost
                }
            }
        }
    }
    return dp[0][n]
}

func minCost2(n int, cuts []int) int {
    sort.Ints(cuts)
    m := len(cuts) + 2
    dp, newCuts := make([][]int, m), make([]int, m)
    for i := 1; i < m-1; i++ {
        newCuts[i] = cuts[i-1]
    }
    newCuts[0], newCuts[m-1] = 0, n
    for i := range dp {
        dp[i] = make([]int, m)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for l := m - 2; l >= 1; l-- {
        for r := l; r <= m-2; r++ {
            if l == r {
                dp[l][r] = newCuts[r+1] - newCuts[l-1]
            } else {
                res := 1 << 63 - 1
                for i := l; i <= r; i++ {
                    res = min(res, dp[l][i-1] + dp[i+1][r])
                }
                dp[l][r] = res + newCuts[r+1]- newCuts[l-1]
            }
        }
    }
    return dp[1][m-2]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/23/e1.jpg" />
    // Input: n = 7, cuts = [1,3,4,5]
    // Output: 16
    // Explanation: Using cuts order = [1, 3, 4, 5] as in the input leads to the following scenario:
    // <img src="https://assets.leetcode.com/uploads/2020/07/21/e11.jpg" />
    // The first cut is done to a rod of length 7 so the cost is 7. The second cut is done to a rod of length 6 (i.e. the second part of the first cut), the third is done to a rod of length 4 and the last cut is to a rod of length 3. The total cost is 7 + 6 + 4 + 3 = 20.
    // Rearranging the cuts to be [3, 5, 1, 4] for example will lead to a scenario with total cost = 16 (as shown in the example photo 7 + 4 + 3 + 2 = 16).
    fmt.Println(minCost(7,[]int{1,3,4,5})) // 16
    // Example 2:
    // Input: n = 9, cuts = [5,6,1,4,2]
    // Output: 22
    // Explanation: If you try the given cuts ordering the cost will be 25.
    // There are much ordering with total cost <= 25, for example, the order [4, 6, 5, 2, 1] has total cost = 22 which is the minimum possible.
    fmt.Println(minCost(9,[]int{5,6,1,4,2})) // 22

    fmt.Println(minCost1(7,[]int{1,3,4,5})) // 16
    fmt.Println(minCost1(9,[]int{5,6,1,4,2})) // 22

    fmt.Println(minCost2(7,[]int{1,3,4,5})) // 16
    fmt.Println(minCost2(9,[]int{5,6,1,4,2})) // 22
}