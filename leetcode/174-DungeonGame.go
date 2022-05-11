package main

/**
174. Dungeon Game
The demons had captured the princess and imprisoned her in the bottom-right corner of a dungeon.
The dungeon consists of m x n rooms laid out in a 2D grid.
Our valiant knight was initially positioned in the top-left room and must fight his way through dungeon to rescue the princess.
The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons (represented by negative integers),
so the knight loses health upon entering these rooms; other rooms are either empty (represented as 0)
or contain magic orbs that increase the knight's health (represented by positive integers).
To reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.
Return the knight's minimum initial health so that he can rescue the princess.
Note that any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

Example 1:

	Input: dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
	Output: 7   -2 -> -3 -> 3 -> 1 -> -5
	Explanation: The initial health of the knight must be at least 7 if he follows the optimal path: RIGHT-> RIGHT -> DOWN -> DOWN.

Example 2:

	Input: dungeon = [[0]]
	Output: 1

Constraints:

	m == dungeon.length
	n == dungeon[i].length
	1 <= m, n <= 200
	-1000 <= dungeon[i][j] <= 1000

 */

import (
	"fmt"
	"math"
)

// 解法一 动态规划
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = max(1-dungeon[m-1][n-1], 1)
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max(1, dp[m-1][i+1]-dungeon[m-1][i])
	}
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(1, dp[i+1][n-1]-dungeon[i][n-1])
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = min(max(1, dp[i][j+1]-dungeon[i][j]), max(1, dp[i+1][j]-dungeon[i][j]))
			fmt.Printf("%v %v %v\n",i,j,dp)
		}
	}
	return dp[0][0]
}

// 解法二 二分搜索
func calculateMinimumHP1(dungeon [][]int) int {
	low, high := 1, math.MaxInt64
	for low < high {
		mid := low + (high-low)>>1
		if canCross(dungeon, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func canCross(dungeon [][]int, start int) bool {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = start + dungeon[0][0]
			} else {
				a, b := math.MinInt64, math.MinInt64
				if i > 0 && dp[i-1][j] > 0 {
					a = dp[i-1][j] + dungeon[i][j]
				}
				if j > 0 && dp[i][j-1] > 0 {
					b = dp[i][j-1] + dungeon[i][j]
				}
				dp[i][j] = max(a, b)
			}
		}
	}
	return dp[m-1][n-1] > 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("calculateMinimumHP1([][]int{{ -2,-3,3}, {-5,-10,1}, {10,30,-5}}) = %v\n",calculateMinimumHP1([][]int{{ -2,-3,3}, {-5,-10,1}, {10,30,-5}})) // 7
	fmt.Printf("calculateMinimumHP1([][]int{ { 0 } }) = %v\n",calculateMinimumHP1([][]int{ { 0 } })) // 1

	fmt.Printf("calculateMinimumHP([][]int{{ -2,-3,3}, {-5,-10,1}, {10,30,-5}}) = %v\n",calculateMinimumHP([][]int{{ -2,-3,3}, {-5,-10,1}, {10,30,-5}})) // 7
	fmt.Printf("calculateMinimumHP([][]int{ { 0 } }) = %v\n",calculateMinimumHP([][]int{ { 0 } })) // 1
}
