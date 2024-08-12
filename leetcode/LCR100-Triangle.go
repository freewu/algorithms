package main

// LCR 100. 三角形最小路径和
// 给定一个三角形 triangle ，找出自顶向下的最小路径和。

// 每一步只能移动到下一行中相邻的结点上。
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
// 也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。


// 示例 1：
// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
//    (2)
//   (3) 4
//  6 (5) 7
// 4 (1) 8 3
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

// 示例 2：
// 输入：triangle = [[-10]]
// 输出：-10

// 提示：
//     1 <= triangle.length <= 200
//     triangle[0].length == 1
//     triangle[i].length == triangle[i - 1].length + 1
//     -10^4 <= triangle[i][j] <= 10^4

// 进阶：
//     你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
 
// # 解题思路
//     求出从三角形顶端到底端的最小和。要求最好用 O(n) 的时间复杂度。
//     直接从下层往上层推。普通解法是用二维数组 DP，稍微优化的解法是一维数组 DP

import "fmt"

// 倒序 DP，无辅助空间
func minimumTotal(triangle [][]int) int {
    if triangle == nil {
        return 0
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for row := len(triangle) - 2; row >= 0; row-- { // 从倒数第二层开始向上计算最小值
        for col := 0; col < len(triangle[row]); col++ {
            triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
            // fmt.Printf("row = %v, col = %v, triangle = %v\n",row,col,triangle)
        }
    }
    return triangle[0][0]
}

// 正常 DP，空间复杂度 O(n)
func minimumTotal1(triangle [][]int) int {
    if len(triangle) == 0 {
        return 0
    }
    dp, res, index := make([]int, len(triangle[len(triangle)-1])), 1 << 32, 0
    for ; index < len(triangle[0]); index++ {
        dp[index] = triangle[0][index]
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < len(triangle); i++ {
        for j := len(triangle[i]) - 1; j >= 0; j-- {
            if j == 0 { // 最左边
                dp[j] += triangle[i][0]
            } else if j == len(triangle[i]) - 1 { // 最右边
                dp[j] += dp[j-1] + triangle[i][j]
            } else { // 中间
                dp[j] = min(dp[j-1] + triangle[i][j], dp[j] + triangle[i][j])
            }
        }
    }
    for i := 0; i < len(dp); i++ {
        if dp[i] < res {
            res = dp[i]
        }
    }
    return res
}

func minimumTotal2(triangle [][]int) int {
    n := len(triangle)
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, len(triangle[i]))
        copy(dp[i], triangle[i])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n-2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            dp[i][j] = dp[i][j] + min(dp[i+1][j], dp[i+1][j+1])
        }
    }
    return dp[0][0]
}

func main() {
    // Explanation: The triangle looks like:
    //    2
    //   3 4
    //  6 5 7
    // 4 1 8 3
    // The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
    fmt.Printf("minimumTotal([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}}) = %v\n",minimumTotal([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}})) // 11
    fmt.Printf("minimumTotal([][]int{{-10}}) = %v\n",minimumTotal([][]int{{-10}})) // -10

    fmt.Printf("minimumTotal1([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}}) = %v\n",minimumTotal1([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}})) // 11
    fmt.Printf("minimumTotal1([][]int{{-10}}) = %v\n",minimumTotal1([][]int{{-10}})) // -10

    fmt.Printf("minimumTotal2([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}}) = %v\n",minimumTotal2([][]int{{2}, {3,4},{6,5,7},{4,1,8,3}})) // 11
    fmt.Printf("minimumTotal2([][]int{{-10}}) = %v\n",minimumTotal2([][]int{{-10}})) // -10
}
