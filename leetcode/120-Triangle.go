package main

/**
120. Triangle
Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below. More formally,
if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Constraints:

	1 <= triangle.length <= 200
	triangle[0].length == 1
	triangle[i].length == triangle[i - 1].length + 1
	-10^4 <= triangle[i][j] <= 10^4

Example 1:

	Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
	Output: 11
	Explanation: The triangle looks like:
	   2
	  3 4
	 6 5 7
	4 1 8 3
	The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

Example 2:

	Input: triangle = [[-10]]
	Output: -10

# 解题思路
	DP
	求出从三角形顶端到底端的最小和。要求最好用 O(n) 的时间复杂度。
	直接从下层往上层推。普通解法是用二维数组 DP，稍微优化的解法是一维数组 DP
 */

import (
	"fmt"
	"math"
)

// 解法一 倒序 DP，无辅助空间
func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	for row := len(triangle) - 2; row >= 0; row-- { // 从倒数第二层开始向上计算最小值
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
			fmt.Printf("row = %v, col = %v, triangle = %v\n",row,col,triangle)
		}
	}
	return triangle[0][0]
}

func min(a, b int) int{
	if a >b {
		return b
	}
	return a
}

// 解法二 正常 DP，空间复杂度 O(n)
func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp, minNum, index := make([]int, len(triangle[len(triangle)-1])), math.MaxInt64, 0
	for ; index < len(triangle[0]); index++ {
		dp[index] = triangle[0][index]
	}
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {
				// 最左边
				dp[j] += triangle[i][0]
			} else if j == len(triangle[i])-1 {
				// 最右边
				dp[j] += dp[j-1] + triangle[i][j]
			} else {
				// 中间
				dp[j] = min(dp[j-1]+triangle[i][j], dp[j]+triangle[i][j])
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		if dp[i] < minNum {
			minNum = dp[i]
		}
	}
	return minNum
}

func main() {
	fmt.Printf("minimumTotal([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}}) = %v\n",minimumTotal([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}})) // 11
	fmt.Printf("minimumTotal([][]int{{-10}}) = %v\n",minimumTotal([][]int{{-10}})) // -10

	fmt.Printf("minimumTotal1([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}}) = %v\n",minimumTotal1([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}})) // 11
	fmt.Printf("minimumTotal1([][]int{{-10}}) = %v\n",minimumTotal1([][]int{{-10}})) // -10
}
