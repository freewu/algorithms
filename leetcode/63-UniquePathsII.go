package main

import "fmt"

/**
63. Unique Paths II
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
Now consider if some obstacles are added to the grids. How many unique paths would there be?
An obstacle and space is marked as 1 and 0 respectively in the grid.

Constraints:

	m == obstacleGrid.length
	n == obstacleGrid[i].length
	1 <= m, n <= 100
	obstacleGrid[i][j] is 0 or 1.

Example 1:

	Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
	Output: 2
	Explanation: There is one obstacle in the middle of the 3x3 grid above.
	There are two ways to reach the bottom-right corner:
	1. Right -> Right -> Down -> Down
	2. Down -> Down -> Right -> Right


Example 2:

	Input: obstacleGrid = [[0,1],[0,0]]
	Output: 1

解题思路:
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径
 */

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 地图长度为0 或 开始位置就是障碍物 直接返回 0
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0]) // 获取地图的 高 & 宽
	dp := make([][]int, m)
	// 先构造一个 m * n 的数组
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// 遍历第一行
	for i := 1; i < n; i++ {
		if dp[0][i-1] != 0 && obstacleGrid[0][i] != 1 {
			dp[0][i] = 1
		}
	}
	// 遍历第一列
	for i := 1; i < m; i++ {
		if dp[i-1][0] != 0 && obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// best solution
func uniquePathsWithObstaclesBest(obstacleGrid [][]int) int {
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 0 {
			obstacleGrid[i][0] = -1
		} else {
			break
		}
	}
	if obstacleGrid[0][0] != 1 {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[0][j] == 0 {
				obstacleGrid[0][j] = -1
			} else {
				break
			}
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 0 {
				if obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1 {
					obstacleGrid[i][j] = 0
				} else if obstacleGrid[i-1][j] == 1 {
					obstacleGrid[i][j] = obstacleGrid[i][j-1]
				} else if obstacleGrid[i][j-1] == 1 {
					obstacleGrid[i][j] = obstacleGrid[i-1][j]
				} else {
					obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
				}
			}
		}
	}
	v := obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
	if v < 0 {
		return v * -1
	}
	return 0
}

func main() {
	fmt.Printf("uniquePathsWithObstacles([][]int{[]int{0,0,0},[]int{ 0,1,0 },[]int{0,0,0}}) = %v\n",uniquePathsWithObstacles([][]int{[]int{0,0,0},[]int{ 0,1,0 },[]int{0,0,0}}))
	fmt.Printf("uniquePathsWithObstacles([][]int{[]int{0,1},[]int{ 0,0 }}) = %v\n",uniquePathsWithObstacles([][]int{[]int{0,1},[]int{ 0,0 }}))

	fmt.Printf("uniquePathsWithObstaclesBest([][]int{[]int{0,0,0},[]int{ 0,1,0 },[]int{0,0,0}}) = %v\n",uniquePathsWithObstaclesBest([][]int{[]int{0,0,0},[]int{ 0,1,0 },[]int{0,0,0}}))
	fmt.Printf("uniquePathsWithObstaclesBest([][]int{[]int{0,1},[]int{ 0,0 }}) = %v\n",uniquePathsWithObstaclesBest([][]int{[]int{0,1},[]int{ 0,0 }}))
}
