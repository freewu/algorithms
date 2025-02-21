package main

// LCP 21. 追逐游戏
// 秋游中的小力和小扣设计了一个追逐游戏。
// 他们选了秋日市集景区中的 N 个景点，景点编号为 1~N。
// 此外，他们还选择了 N 条小路，满足任意两个景点之间都可以通过小路互相到达，且不存在两条连接景点相同的小路。
// 整个游戏场景可视作一个无向连通图，记作二维数组edges，数组中以[a,b]形式表示景点 a 与景点 b 之间有一条小路连通。

// 小力和小扣只能沿景点间的小路移动。
// 小力的目标是在最快时间内追到小扣，小扣的目标是尽可能延后被小力追到的时间。
// 游戏开始前，两人分别站在两个不同的景点startA和startB。
// 每一回合，小力先行动，小扣观察到小力的行动后再行动。
// 小力和小扣在每回合可选择以下行动之一：
//     1. 移动至相邻景点
//     2. 留在原地

// 如果小力追到小扣（即两人于某一时刻出现在同一位置），则游戏结束。
// 若小力可以追到小扣，请返回最少需要多少回合；若小力无法追到小扣，请返回 -1。

// 注意：小力和小扣一定会采取最优移动策略。

// 示例 1：
// 输入：edges = [[1,2],[2,3],[3,4],[4,1],[2,5],[5,6]], startA = 3, startB = 5
// 输出：3
// 解释：
// <img src="https://pic.leetcode-cn.com/1597991318-goeHHr-image.png" />
// 第一回合，小力移动至 2 号点，小扣观察到小力的行动后移动至 6 号点； 
// 第二回合，小力移动至 5 号点，小扣无法移动，留在原地； 
// 第三回合，小力移动至 6 号点，小力追到小扣。返回 3。

// 示例 2：
// 输入：edges = [[1,2],[2,3],[3,4],[4,1]], startA = 1, startB = 3
// 输出：-1
// 解释：
// <img src="https://pic.leetcode-cn.com/1597991157-QfeakF-image.png" />
// 小力如果不动，则小扣也不动；否则小扣移动到小力的对角线位置。这样小力无法追到小扣。

// 提示：
//     edges的长度等于图中节点个数
//     3 <= edges.length <= 10^5
//     1 <= edges[i][0], edges[i][1] <= edges.length 且 edges[i][0] != edges[i][1]
//     1 <= startA, startB <= edges.length 且 startA != startB

import "fmt"
import "container/list"

func topsort(graph [][]int, deg []int) map[int]bool {
    ans := make(map[int]bool)
    q := list.New()
    for i, v := range deg {
        if v == 1 {
            q.PushBack(i)
        }
    }
    for q.Len() > 0 {
        now := q.Front().Value.(int)
        q.Remove(q.Front())
        ans[now] = true
        for _, nxt := range graph[now] {
            if ans[nxt] {
                continue
            }
            deg[nxt]--
            if deg[nxt] == 1 {
                q.PushBack(nxt)
            }
        }
    }
    return ans
}

func bfs(start, n int, graph [][]int) []int {
    dis := make([]int, n)
    for i := range dis {
        dis[i] = -1
    }
    step := 0
    q := list.New()
    q.PushBack(start)
    for q.Len() > 0 {
        size := q.Len()
        for i := 0; i < size; i++ {
            now := q.Front().Value.(int)
            q.Remove(q.Front())
            if dis[now] != -1 {
                continue
            }
            dis[now] = step
            for _, nxt := range graph[now] {
                if dis[nxt] == -1 {
                    q.PushBack(nxt)
                }
            }
        }
        step++
    }
    return dis
}

func chaseGame(edges [][]int, startA int, startB int) int {
    n := len(edges)
    startA--
    startB--
    // 建立图
    graph := make([][]int, n)
    deg := make([]int, n)
    for _, edge := range edges {
        u, v := edge[0]-1, edge[1]-1
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
        deg[u]++
        deg[v]++
    }
    // topsort 找 circle
    notCircle := topsort(graph, deg)
    circle := make(map[int]bool)
    for i := 0; i < n; i++ {
        if !notCircle[i] {
            circle[i] = true
        }
    }
    // 特判相邻情况
    for _, neighbor := range graph[startA] {
        if neighbor == startB {
            return 1
        }
    }
    // 特判B在环中，而且环满足 >3，则永远抓不到
    if circle[startB] && len(circle) > 3 {
        return -1
    }
    // 下面的情况就是：不相邻 | not(B在环中+len(circle)>3)
    disA := bfs(startA, n, graph)
    disB := bfs(startB, n, graph)
    // 找 B 距环的入口
    entry := -1
    minDisB := n + 1
    for i := range circle {
        if disB[i] < minDisB {
            minDisB = disB[i]
            entry = i
        }
    }
    // 不可能抓到了，B躲进去环里面去了
    if len(circle) > 3 && disA[entry] > disB[entry] {
        return -1
    }
    // 下面是一定会抓到的情况；disA[i] > disB[i] + 1 关键判定
    res := 0
    for i := 0; i < n; i++ {
        if disA[i] > disB[i]+1 {
            if disA[i] > res {
                res = disA[i]
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：edges = [[1,2],[2,3],[3,4],[4,1],[2,5],[5,6]], startA = 3, startB = 5
    // 输出：3
    // 解释：
    // <img src="https://pic.leetcode-cn.com/1597991318-goeHHr-image.png" />
    // 第一回合，小力移动至 2 号点，小扣观察到小力的行动后移动至 6 号点； 
    // 第二回合，小力移动至 5 号点，小扣无法移动，留在原地； 
    // 第三回合，小力移动至 6 号点，小力追到小扣。返回 3。
    fmt.Println(chaseGame([][]int{{1,2},{2,3},{3,4},{4,1},{2,5},{5,6}}, 3, 5)) // 3
    // 示例 2：
    // 输入：edges = [[1,2],[2,3],[3,4],[4,1]], startA = 1, startB = 3
    // 输出：-1
    // 解释：
    // <img src="https://pic.leetcode-cn.com/1597991157-QfeakF-image.png" />
    // 小力如果不动，则小扣也不动；否则小扣移动到小力的对角线位置。这样小力无法追到小扣。
    fmt.Println(chaseGame([][]int{{1,2},{2,3},{3,4},{4,1}}, 1, 3)) // -1
}