package main

// LCP 09. 最小跳跃次数
// 为了给刷题的同学一些奖励，力扣团队引入了一个弹簧游戏机。
// 游戏机由 N 个特殊弹簧排成一排，编号为 0 到 N-1。初始有一个小球在编号 0 的弹簧处。
// 若小球在编号为 i 的弹簧处，通过按动弹簧，可以选择把小球向右弹射 jump[i] 的距离，或者向左弹射到任意左侧弹簧的位置。
// 也就是说，在编号为 i 弹簧处按动弹簧，小球可以弹向 0 到 i-1 中任意弹簧或者 i+jump[i] 的弹簧（若 i+jump[i]>=N ，则表示小球弹出了机器）。
// 小球位于编号 0 处的弹簧时不能再向左弹。

// 为了获得奖励，你需要将小球弹出机器。
// 请求出最少需要按动多少次弹簧，可以将小球从编号 0 弹簧弹出整个机器，即向右越过编号 N-1 的弹簧。

// 示例 1：
// 输入：jump = [2, 5, 1, 1, 1, 1]
// 输出：3
// 解释：小 Z 最少需要按动 3 次弹簧，小球依次到达的顺序为 0 -> 2 -> 1 -> 6，最终小球弹出了机器。

// 限制：
//     1 <= jump.length <= 10^6
//     1 <= jump[i] <= 10000

import "fmt"

func minJump(jump []int) int {
    n := len(jump)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 2 * n
        if i + jump[i] >= n {
            dp[i] = 1
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 2; i >= 0; i-- {
        k := min(i + jump[i], n - 1)
        dp[i]  = min(dp[i], 1 + dp[k] )
        for j := i + 1; j <= k && dp[j] >= dp[i] + 1; j++ {
            dp[j] = dp[i] + 1
        }
    }
    return dp[0]
}

// bfs
func minJump1(jump []int) int {
    n := len(jump)
    visited := make([]bool, n) // 标记是否访问过某个位置
    queue := []int{ 0 } // BFS 队列，存储当前弹簧位置
    steps := 0 // 记录弹簧次数
    maxLeft := 0 // 记录最远左跳范围
    for len(queue) > 0 {
        nextQueue := []int{} // 下一层的队列
        for _, i := range queue {
            if i + jump[i] >= n { return steps + 1 } // 如果可以向右跳出机器
            if !visited[i + jump[i]] { // 向右跳
                visited[i + jump[i]] = true
                nextQueue = append(nextQueue, i+jump[i])
            }
            for j := maxLeft; j < i; j++ { // 向左跳
                if !visited[j] {
                    visited[j] = true
                    nextQueue = append(nextQueue, j)
                }
            }
            maxLeft = i // 更新最远左跳范围
        }
        queue = nextQueue // 更新当前队列为下一层
        steps++
    }
    return -1 // 如果无法跳出机器，返回 -1
}

func main() {
    // 示例 1：
    // 输入：jump = [2, 5, 1, 1, 1, 1]
    // 输出：3
    // 解释：小 Z 最少需要按动 3 次弹簧，小球依次到达的顺序为 0 -> 2 -> 1 -> 6，最终小球弹出了机器。
    fmt.Println(minJump([]int{2,5,1,1,1,1})) // 3

    fmt.Println(minJump([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(minJump([]int{9,8,7,6,5,4,3,2,1})) // 3

    fmt.Println(minJump1([]int{2,5,1,1,1,1})) // 3
    fmt.Println(minJump1([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(minJump1([]int{9,8,7,6,5,4,3,2,1})) // 3
}