package main

// LCP 54. 夺回据点
// 欢迎各位勇者来到力扣城，本次试炼主题为「夺回据点」。
// 魔物了占领若干据点，这些据点被若干条道路相连接，roads[i] = [x, y] 表示编号 x、y 的两个据点通过一条道路连接。
// 现在勇者要将按照以下原则将这些据点逐一夺回：
//     在开始的时候，勇者可以花费资源先夺回一些据点，初始夺回第 j 个据点所需消耗的资源数量为 cost[j]
//     接下来，勇者在不消耗资源情况下，每次可以夺回一个和「已夺回据点」相连接的魔物据点，并对其进行夺回

// 注：为了防止魔物暴动，勇者在每一次夺回据点后（包括花费资源夺回据点后），需要保证剩余的所有魔物据点之间是相连通的（不经过「已夺回据点」）。
// 请返回勇者夺回所有据点需要消耗的最少资源数量。

// 注意：
//     输入保证初始所有据点都是连通的，且不存在重边和自环

// 示例 1：
// 输入： cost = [1,2,3,4,5,6] roads = [[0,1],[0,2],[1,3],[2,3],[1,2],[2,4],[2,5]]
// 输出：6
// 解释： 勇者消耗资源 6 夺回据点 0 和 4，魔物据点 1、2、3、5 相连通； 
// 第一次夺回据点 1，魔物据点 2、3、5 相连通； 
// 第二次夺回据点 3，魔物据点 2、5 相连通； 
// 第三次夺回据点 2，剩余魔物据点 5； 
// 第四次夺回据点 5，无剩余魔物据点； 
// 因此最少需要消耗资源为 6，可占领所有据点。
// <img src="https://pic.leetcode-cn.com/1648706944-KJstUN-image.png" />

// 示例 2：
// 输入： cost = [3,2,1,4] roads = [[0,2],[2,3],[3,1]]
// 输出：2
// 解释： 勇者消耗资源 2 夺回据点 1，魔物据点 0、2、3 相连通； 
// 第一次夺回据点 3，魔物据点 2、0 相连通； 
// 第二次夺回据点 2，剩余魔物据点 0； 
// 第三次夺回据点 0，无剩余魔物据点； 
// 因此最少需要消耗资源为 2，可占领所有据点。
// <img src="https://pic.leetcode-cn.com/1648707186-LJRwzU-image.png" />

// 提示：
//     1 <= roads.length, cost.length <= 10^5
//     0 <= roads[i][0], roads[i][1] < cost.length
//     1 <= cost[i] <= 10^9

import "fmt"
import "container/list"
import "math"
import "sort"

func minimumCost(cost []int, roads [][]int) int64 {
    n := len(cost)
    if n == 1 {
        return int64(cost[0])
    }
    adjvex := make(map[int][]int)
    for _, road := range roads {
        x, y := road[0], road[1]
        adjvex[x] = append(adjvex[x], y)
        adjvex[y] = append(adjvex[y], x)
    }
    root := 0
    dfn := make([]int, n)
    low := make([]int, n)
    isCut := make([]bool, n)
    stk := list.New()
    var dcc [][]int
    dfsClock := 1
    var tarjan func(int)
    tarjan = func(x int) {
        low[x] = dfsClock
        dfn[x] = dfsClock
        dfsClock++
        stk.PushBack(x)
        flag := 0
        for _, y := range adjvex[x] {
            if dfn[y] == 0 {
                tarjan(y)
                low[x] = min(low[x], low[y])
                if low[y] >= dfn[x] {
                    flag++
                    if x != root || flag > 1 {
                        isCut[x] = true
                    }
                    t := stk.Back().Value.(int)
                    stk.Remove(stk.Back())
                    dcc = append(dcc, []int{t})
                    for t != y {
                        t = stk.Back().Value.(int)
                        stk.Remove(stk.Back())
                        dcc[len(dcc)-1] = append(dcc[len(dcc)-1], t)
                    }
                    dcc[len(dcc)-1] = append(dcc[len(dcc)-1], x)
                }
            } else {
                low[x] = min(low[x], dfn[y])
            }
        }
    }
    tarjan(root)
    if len(dcc) == 1 {
        res := math.MaxInt64
        for _, co := range cost {
            res = min(res, co)
        }
        return int64(res)
    }
    var leaf []int
    for _, dc := range dcc {
        cutCnt := 0
        minCost := math.MaxInt64
        for _, x := range dc {
            if isCut[x] {
                cutCnt++
            } else {
                minCost = min(minCost, cost[x])
            }
        }
        if cutCnt == 1 {
            leaf = append(leaf, minCost)
        }
    }
    sort.Ints(leaf)
    res := 0
    for i := 0; i < len(leaf)-1; i++ {
        res += leaf[i]
    }
    return int64(res)
}


func main() {
    // 示例 1：
    // 输入： cost = [1,2,3,4,5,6] roads = [[0,1],[0,2],[1,3],[2,3],[1,2],[2,4],[2,5]]
    // 输出：6
    // 解释： 勇者消耗资源 6 夺回据点 0 和 4，魔物据点 1、2、3、5 相连通； 
    // 第一次夺回据点 1，魔物据点 2、3、5 相连通； 
    // 第二次夺回据点 3，魔物据点 2、5 相连通； 
    // 第三次夺回据点 2，剩余魔物据点 5； 
    // 第四次夺回据点 5，无剩余魔物据点； 
    // 因此最少需要消耗资源为 6，可占领所有据点。
    // <img src="https://pic.leetcode-cn.com/1648706944-KJstUN-image.png" />
    fmt.Println(minimumCost([]int{1,2,3,4,5,6},[][]int{{0,1},{0,2},{1,3},{2,3},{1,2},{2,4},{2,5}})) // 6
    // 示例 2：
    // 输入： cost = [3,2,1,4] roads = [[0,2],[2,3],[3,1]]
    // 输出：2
    // 解释： 勇者消耗资源 2 夺回据点 1，魔物据点 0、2、3 相连通； 
    // 第一次夺回据点 3，魔物据点 2、0 相连通； 
    // 第二次夺回据点 2，剩余魔物据点 0； 
    // 第三次夺回据点 0，无剩余魔物据点； 
    // 因此最少需要消耗资源为 2，可占领所有据点。
    // <img src="https://pic.leetcode-cn.com/1648707186-LJRwzU-image.png" />
    fmt.Println(minimumCost([]int{3,2,1,4},[][]int{{0,2},{2,3},{3,1}})) // 2
}