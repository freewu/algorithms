package main

// 3376. Minimum Time to Break Locks I
// Bob is stuck in a dungeon and must break n locks, each requiring some amount of energy to break. 
// The required energy for each lock is stored in an array called strength where strength[i] indicates the energy needed to break the ith lock.

// To break a lock, Bob uses a sword with the following characteristics:
//     1. The initial energy of the sword is 0.
//     2. The initial factor X by which the energy of the sword increases is 1.
//     3. Every minute, the energy of the sword increases by the current factor X.
//     4. To break the ith lock, the energy of the sword must reach at least strength[i].
//     5. After breaking a lock, the energy of the sword resets to 0, and the factor X increases by a given value K.

// Your task is to determine the minimum time in minutes required for Bob to break all n locks and escape the dungeon.

// Return the minimum time required for Bob to break all n locks.

// Example 1:
// Input: strength = [3,4,1], K = 1
// Output: 4
// Explanation:
// Time    Energy  X   Action          Updated X
// 0       0       1   Nothing             1
// 1       1       1   Break 3rd Lock      2
// 2       2       2   Nothing             2
// 3       4       2   Break 2nd Lock      3
// 4       3       3   Break 1st Lock      3
// The locks cannot be broken in less than 4 minutes; thus, the answer is 4.

// Example 2:
// Input: strength = [2,5,4], K = 2
// Output: 5
// Explanation:
// Time    Energy  X   Action      Updated X
// 0       0       1   Nothing	       1
// 1       1       1   Nothing         1
// 2       2       1   Break 1st Lock  3
// 3       3       3   Nothing         3
// 4       6       3   Break 2nd Lock  5
// 5       5       5   Break 3rd Lock  7
// The locks cannot be broken in less than 5 minutes; thus, the answer is 5.

// Constraints:
//     n == strength.length
//     1 <= n <= 8
//     1 <= K <= 10
//     1 <= strength[i] <= 10^6

import "fmt"
import "math/bits"

// 排列型回溯+最优性剪枝
func findMinimumTime(strength []int, k int) int {
    res, n := 1 << 31, len(strength)
    done := make([]bool, n)
    var dfs func(i int, time int)
    dfs = func(i, time int) {
        if time >= res { return } // 最优性剪枝：答案不可能变小
        if i == n {
            res = time
            return
        }
        x := 1 + k * i
        for j, v := range done {
            if !v {
                done[j] = true // 已开锁
                dfs(i + 1, time + (strength[j] - 1) / x + 1)
                done[j] = false // 恢复现场
            }
        }
    }
    dfs(0, 0)
    return res
}

// 状压 DP
func findMinimumTime1(strength []int, k int) int {
    n := len(strength)
    memo := make([]int, 1 << n)
    for i := range memo {
        memo[i] = -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(i int) int
    dfs = func(i int) int {
        if i == 0 { return 0 }
        if memo[i] != -1 { return memo[i] }
        res, x := 1 << 31, 1 + k * (n - bits.OnesCount(uint(i)))
        for j, s := range strength {
            if i >> j & 1 > 0 {
                res = min(res, dfs(i ^ 1 << j) + (s - 1) / x + 1)
            }
        }
        memo[i] = res // 记忆化
        return res
    }
    return dfs(1 << n - 1)
}

func findMinimumTime2(strength []int, k int) int {
    n := len(strength)
    m := 1 << n
    facts := make([]int, m)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < m; i++ {
        x := 1 + k * (n - bits.OnesCount(uint(i)))
        facts[i] = 1 << 31
        for j, s := range strength {
            if i >> j & 1 > 0 {
                facts[i] = min(facts[i], facts[i ^ 1 << j] + (s - 1) / x + 1)
            }
        }
    }
    return facts[m - 1]
}

// 二分图
func findMinimumTime3(strength []int, k int) int {
    n := len(strength)
    S := n * 2
    T := S + 1
    // rid 为反向边在邻接表中的下标
    type neighbor struct{ to, rid, cap, cost int }
    g := make([][]neighbor, T+1)
    addEdge := func(from, to, cap, cost int) {
        g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
        g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
    }
    for i, s := range strength {
        // 枚举这个锁是第几次开的
        for j := 0; j < n; j++ {
            x := 1 + k*j
            addEdge(i, n+j, 1, (s-1)/x+1)
        }
        addEdge(S, i, 1, 0)
    }
    for i := n; i < n*2; i++ {
        addEdge(i, T, 1, 0)
    }
    // 下面是最小费用最大流模板
    dis := make([]int, len(g))
    type vi struct{ v, i int }
    fa := make([]vi, len(g))
    inQ := make([]bool, len(g))
    spfa := func() bool {
        for i := range dis {
            dis[i] = 1 << 31
        }
        dis[S] = 0
        inQ[S] = true
        q := []int{S}
        for len(q) > 0 {
            v := q[0]
            q = q[1:]
            inQ[v] = false
            for i, e := range g[v] {
                if e.cap == 0 {
                    continue
                }
                w := e.to
                newD := dis[v] + e.cost
                if newD < dis[w] {
                    dis[w] = newD
                    fa[w] = vi{v, i}
                    if !inQ[w] {
                        inQ[w] = true
                        q = append(q, w)
                    }
                }
            }
        }
        // 循环结束后所有 inQ[v] 都为 false，无需重置
        return dis[T] < 1 << 31
    }
    minCost := 0
    for spfa() {
        // 沿 st-end 的最短路尽量增广
        // 特别地，如果建图时所有边的容量都设为 1，那么 minF 必然为 1，下面第一个 for 循环可以省略
        minF := 1 << 31
        for v := T; v != S; {
            p := fa[v]
            minF = min(minF, g[p.v][p.i].cap)
            v = p.v
        }
        for v := T; v != S; {
            p := fa[v]
            e := &g[p.v][p.i]
            e.cap -= minF
            g[v][e.rid].cap += minF
            v = p.v
        }
        minCost += dis[T] * minF
    }
    return minCost
}

func main() {
    // Example 1:
    // Input: strength = [3,4,1], K = 1
    // Output: 4
    // Explanation:
    // Time    Energy  X   Action          Updated X
    // 0       0       1   Nothing             1
    // 1       1       1   Break 3rd Lock      2
    // 2       2       2   Nothing             2
    // 3       4       2   Break 2nd Lock      3
    // 4       3       3   Break 1st Lock      3
    // The locks cannot be broken in less than 4 minutes; thus, the answer is 4.
    fmt.Println(findMinimumTime([]int{3,4,1}, 1)) // 4
    // Example 2:
    // Input: strength = [2,5,4], K = 2
    // Output: 5
    // Explanation:
    // Time    Energy  X   Action      Updated X
    // 0       0       1   Nothing	        1
    // 1       1       1   Nothing         1
    // 2       2       1   Break 1st Lock  3
    // 3       3       3   Nothing         3
    // 4       6       3   Break 2nd Lock  5
    // 5       5       5   Break 3rd Lock  7
    // The locks cannot be broken in less than 5 minutes; thus, the answer is 5.
    fmt.Println(findMinimumTime([]int{2,5,4}, 2)) // 5

    fmt.Println(findMinimumTime1([]int{3,4,1}, 1)) // 4
    fmt.Println(findMinimumTime1([]int{2,5,4}, 2)) // 5

    fmt.Println(findMinimumTime2([]int{3,4,1}, 1)) // 4
    fmt.Println(findMinimumTime2([]int{2,5,4}, 2)) // 5

    fmt.Println(findMinimumTime3([]int{3,4,1}, 1)) // 4
    fmt.Println(findMinimumTime3([]int{2,5,4}, 2)) // 5
}