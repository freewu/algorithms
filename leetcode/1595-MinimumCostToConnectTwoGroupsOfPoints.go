package main

// 1595. Minimum Cost to Connect Two Groups of Points
// You are given two groups of points where the first group has size1 points, the second group has size2 points, and size1 >= size2.

// The cost of the connection between any two points are given in an size1 x size2 matrix where cost[i][j] is the cost of connecting point i of the first group and point j of the second group. 
// The groups are connected if each point in both groups is connected to one or more points in the opposite group. 
// In other words, each point in the first group must be connected to at least one point in the second group, 
// and each point in the second group must be connected to at least one point in the first group.

// Return the minimum cost it takes to connect the two groups.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/03/ex1.jpg" />
// Input: cost = [[15, 96], [36, 2]]
// Output: 17
// Explanation: The optimal way of connecting the groups is:
// 1--A
// 2--B
// This results in a total cost of 17.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/03/ex2.jpg" />
// Input: cost = [[1, 3, 5], [4, 1, 1], [1, 5, 3]]
// Output: 4
// Explanation: The optimal way of connecting the groups is:
// 1--A
// 2--B
// 2--C
// 3--A
// This results in a total cost of 4.
// Note that there are multiple points connected to point 2 in the first group and point A in the second group. This does not matter as there is no limit to the number of points that can be connected. We only care about the minimum total cost.

// Example 3:
// Input: cost = [[2, 5, 1], [3, 4, 7], [8, 1, 2], [6, 2, 4], [3, 8, 8]]
// Output: 10

// Constraints:
//     size1 == cost.length
//     size2 == cost[i].length
//     1 <= size1, size2 <= 12
//     size1 >= size2
//     0 <= cost[i][j] <= 100

import "fmt"

func connectTwoGroups(cost [][]int) int {
    rbm, n, inf := 0, len(cost), 1 << 31
    for i := range cost[0] {
        rbm |= 1 << i
    }
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, rbm + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dp func(i, bmj int) int
    dp = func(i, bmj int) int {
        if i >= n {
            if bmj == rbm { return 0 }
            return -1
        }
        if memo[i][bmj] != -1 { return memo[i][bmj] }
        res, flag := inf, false
        for j := range cost[i] {
            x := 1 << j
            v1 := dp(i + 1, bmj | x)
            if v1 >= 0 {
                res, flag = min(res, v1 + cost[i][j]), true
            }
            if (bmj & x) != 0 { continue }
            v1 = dp(i, bmj | x)
            if v1 >= 0 {
                res, flag = min(res, v1 + cost[i][j]), true
            }
        }
        memo[i][bmj] = res
        if !flag { return -1 }
        return res
    }
    return dp(0, 0)
}

func connectTwoGroups1(cost [][]int) int {
    // 转为递推,并进行优化
    // 1. 空间压缩 ( dp[i][s] 只依赖于 dp[i-1][sub] 其中sub为s的子集
    // 2. 处理base case (即i<0时候, 需要枚举 s的各种状态,计算 dp[-1][s]的值 (正常i偏移1位为 dp[i+1],所以i可以取到 -1)
    m, n, inf := len(cost), len(cost[0]), 1 << 31
    u := 1 << n - 1
    dp := make([]int, 1 << n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for j := 0; j < n; j++ {
        mn := inf
        for i := 0; i < m; i++ {
            mn = min(mn, cost[i][j])
        }
        for mask, bit := 0, 1 << j; mask < bit; mask++ {
            dp[mask | bit] = dp[mask] + mn
        }
    }
    for _, row := range cost { // dp[i][state]
        for s := u; s >= 0; s-- { // 空间压缩,倒序枚举,需要枚举到空集(但这里是s--进行递减的, 不是枚举子集的方式 sub = (sub-1)&s; 所以判断条件可以放在for的声明上
            res := inf
            for j, c := range row {
                res = min(res, dp[s &^(1 << j)] + c) // 特别注意!! newS=s&^(1<<j)是在s集合中移除j,这里s可能是不包含j的!! 所以存在newS==s,计算过程中不能覆盖 dp[s]!!
            }
            dp[s] = res // 特别注意!! 这里是枚举选哪个,而不是选与不选, dp[s]一定要用新值,不能使用 dp[s] = min(dp[s], res)!! 因为dp[s]还是dp[i-1][s]的值 => 在i位置上必须选一个, dp[i-1][s]是当前i不选!!
        }
    }
    return dp[u] // 看dfs的写法, dfs(n-1,u)是起始阶段,这里也应该用 dp[u] 而不是 dp[0], dp[n-1][u]代表,使用集合a0->n-1的所有元素,匹配集合2的未选过的所有元素的最小花费
}

func connectTwoGroups2(cost [][]int) int {
    m, n, inf := len(cost), len(cost[0]), 1 << 31
    minCosts := make([]int, n) // 预处理,计算组b元素连组a元素的最小花费
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for j := range minCosts {
        mn := inf
        for i := 0; i < m; i++ {
            mn = min(mn, cost[i][j])
        }
        minCosts[j] = mn
    }
    memo := make([][]int, m) // 状压dp + 枚举选哪个
    for i := range memo {
        memo[i] = make([]int, 1 << n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(i int, state int) int // 让组a去的元素[0...i] 取匹配 组b的元素(可以使用之前用过的),并且保证组b的元素最终全部被匹配的 最小成本
    dfs = func(i int, state int) int { // state bit 1代表还未匹配,方便最后处理未匹配的b组元素
        if i < 0 {
            x := 0
            for j := 0; j < n; j++ {
                if state>>j&1 == 1 { // 贪心!! 如果j还未连到组a,它可以寻找一个最低花费的组a元素
                    x += minCosts[j]
                }
            }
            return x
        }
        // 枚举选哪个
        if memo[i][state] != -1 { return memo[i][state] }
        res := inf
        for j, c := range cost[i] {
            res = min(res, c + dfs(i - 1, state& ^ ( 1 << j)))
        }
        memo[i][state] = res
        return res
    }
    return dfs(m - 1, 1 << n - 1) // 倒序遍历,所以i应该从m - 1而不是0开始
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/03/ex1.jpg" />
    // Input: cost = [[15, 96], [36, 2]]
    // Output: 17
    // Explanation: The optimal way of connecting the groups is:
    // 1--A
    // 2--B
    // This results in a total cost of 17.
    fmt.Println(connectTwoGroups([][]int{{15, 96}, {36, 2}})) // 17
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/03/ex2.jpg" />
    // Input: cost = [[1, 3, 5], [4, 1, 1], [1, 5, 3]]
    // Output: 4
    // Explanation: The optimal way of connecting the groups is:
    // 1--A
    // 2--B
    // 2--C
    // 3--A
    // This results in a total cost of 4.
    // Note that there are multiple points connected to point 2 in the first group and point A in the second group. This does not matter as there is no limit to the number of points that can be connected. We only care about the minimum total cost.
    fmt.Println(connectTwoGroups([][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}})) // 4
    // Example 3:
    // Input: cost = [[2, 5, 1], [3, 4, 7], [8, 1, 2], [6, 2, 4], [3, 8, 8]]
    // Output: 10
    fmt.Println(connectTwoGroups([][]int{{2, 5, 1}, {3, 4, 7}, {8, 1, 2}, {6, 2, 4}, {3, 8, 8}})) // 10

    fmt.Println(connectTwoGroups1([][]int{{15, 96}, {36, 2}})) // 17
    fmt.Println(connectTwoGroups1([][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}})) // 4
    fmt.Println(connectTwoGroups1([][]int{{2, 5, 1}, {3, 4, 7}, {8, 1, 2}, {6, 2, 4}, {3, 8, 8}})) // 10

    fmt.Println(connectTwoGroups2([][]int{{15, 96}, {36, 2}})) // 17
    fmt.Println(connectTwoGroups2([][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}})) // 4
    fmt.Println(connectTwoGroups2([][]int{{2, 5, 1}, {3, 4, 7}, {8, 1, 2}, {6, 2, 4}, {3, 8, 8}})) // 10
}