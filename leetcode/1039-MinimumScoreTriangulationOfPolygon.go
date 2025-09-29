package main

// 1039. Minimum Score Triangulation of Polygon
// You have a convex n-sided polygon where each vertex has an integer value. 
// You are given an integer array values where values[i] is the value of the ith vertex (i.e., clockwise order).

// You will triangulate the polygon into n - 2 triangles. 
// For each triangle, the value of that triangle is the product of the values of its vertices, 
// and the total score of the triangulation is the sum of these values over all n - 2 triangles in the triangulation.

// Return the smallest possible total score that you can achieve with some triangulation of the polygon.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/25/shape1.jpg" />
// Input: values = [1,2,3]
// Output: 6
// Explanation: The polygon is already triangulated, and the score of the only triangle is 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/25/shape2.jpg" />
// Input: values = [3,7,4,5]
// Output: 144
// Explanation: There are two triangulations, with possible scores: 3*7*5 + 4*5*7 = 245, or 3*4*5 + 3*4*7 = 144.
// The minimum score is 144.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/02/25/shape3.jpg" />
// Input: values = [1,3,1,4,1,5]
// Output: 13
// Explanation: The minimum score triangulation has score 1*1*3 + 1*1*4 + 1*1*5 + 1*1*1 = 13.

// Constraints:
//     n == values.length
//     3 <= n <= 50
//     1 <= values[i] <= 100

import "fmt"

func minScoreTriangulation(values []int) int {
    n := len(values)
    memo := make([][]int, n)
    for i := range memo { 
        memo[i] = make([]int, n) 
    }
    var dfs func(i, j int, arr []int, memo [][]int) int
    dfs = func(i, j int, arr []int, memo [][]int) int {
        if memo[i][j] != 0 {
            return memo[i][j]
        }
        res := 0
        for k := i + 1; k < j; k++ {
            left := dfs(i, k, arr, memo)
            mid := arr[i] * arr[k] * arr[j]
            right := dfs(k, j, arr, memo)
            cur := left + mid + right
            if res == 0 || cur < res {
                res = cur
            }
        }
        memo[i][j] = res
        return res
    }
    return dfs(0, n - 1, values, memo)
}

// 区间DP 对一个成环的数组，找到一种切分方法，使得计算分数后的总和最小。
// 这里一个关键点是，对于数组两端的顶点0和顶点n-1，他们组成一条边，必定出现在某个三角形中，所以可以用第三个顶点来划分数组
// 那么第三个顶点取值就是1...n-2之间
// 设dp[i][j]表示values[i:j+1]的顶点，组成的三角形划分的最低分，则dp[0][n-1]就是结果
//      if j-i<2, 无需计算, 返回0
//      if j-i==2, dp[i][j] = values[i]*values[i+1]*values[j]
//      if j-i>2, dp[i][j] = min(values[i]*values[j]*values[x]+dp[i][x]+dp[x][j]), 其中i<x<j， 可以取多个值，需要记忆化搜索
func minScoreTriangulation1(values []int) int {
    n := len(values)
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = -1
        }
    }
    // 计算values[i:j+1]的划分的最小得分
    var dfs func(values []int, i int, j int, dp [][]int) int 
    dfs = func(values []int, i int, j int, dp [][]int) int {
        if dp[i][j] != -1 { // 记忆化搜索
            return dp[i][j]
        }
        if j-i < 2 {
            dp[i][j] = 0
            return dp[i][j]
        }
        if j-i == 2 {
            dp[i][j] = values[i] * values[i+1] * values[i+2]
            return dp[i][j]
        }
        res := 1 << 32 -1
        for x := i + 1; x < j; x++ {
            tmp := values[i]*values[j]*values[x] + dfs(values, i, x, dp) + dfs(values, x, j, dp)
            if res > tmp {
                res = tmp
            }
        }
        dp[i][j] = res
        return dp[i][j]
    }
    return dfs(values, 0, n-1, dp)
}

func minScoreTriangulation2(values []int) int {
    n := len(values)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n) 
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i + 1 == j { return 0 }
        p := memo[i][j]
        if p != -1 { return p }
        res := 1 << 32 - 1
        for k := i + 1; k < j; k++ {
            res = min(res, dfs(i, k) + dfs(k, j) + values[i] * values[k]*values[j])
        }
        memo[i][j] = res
        return res
    }
    return dfs(0, n - 1)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/02/25/shape1.jpg" />
    // Input: values = [1,2,3]
    // Output: 6
    // Explanation: The polygon is already triangulated, and the score of the only triangle is 6.
    fmt.Println(minScoreTriangulation([]int{1,2,3})) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/02/25/shape2.jpg" />
    // Input: values = [3,7,4,5]
    // Output: 144
    // Explanation: There are two triangulations, with possible scores: 3*7*5 + 4*5*7 = 245, or 3*4*5 + 3*4*7 = 144.
    // The minimum score is 144.
    fmt.Println(minScoreTriangulation([]int{3,7,4,5})) // 144
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/02/25/shape3.jpg" />
    // Input: values = [1,3,1,4,1,5]
    // Output: 13
    // Explanation: The minimum score triangulation has score 1*1*3 + 1*1*4 + 1*1*5 + 1*1*1 = 13.
    fmt.Println(minScoreTriangulation([]int{1,3,1,4,1,5})) // 13
    
    fmt.Println(minScoreTriangulation([]int{1,2,3,4,5,6,7,8,9})) // 238
    fmt.Println(minScoreTriangulation([]int{9,8,7,6,5,4,3,2,1})) // 238

    fmt.Println(minScoreTriangulation1([]int{1,2,3})) // 6
    fmt.Println(minScoreTriangulation1([]int{3,7,4,5})) // 144
    fmt.Println(minScoreTriangulation1([]int{1,3,1,4,1,5})) // 13
    fmt.Println(minScoreTriangulation1([]int{1,2,3,4,5,6,7,8,9})) // 238
    fmt.Println(minScoreTriangulation1([]int{9,8,7,6,5,4,3,2,1})) // 238

    fmt.Println(minScoreTriangulation2([]int{1,2,3})) // 6
    fmt.Println(minScoreTriangulation2([]int{3,7,4,5})) // 144
    fmt.Println(minScoreTriangulation2([]int{1,3,1,4,1,5})) // 13
    fmt.Println(minScoreTriangulation2([]int{1,2,3,4,5,6,7,8,9})) // 238
    fmt.Println(minScoreTriangulation2([]int{9,8,7,6,5,4,3,2,1})) // 238
}