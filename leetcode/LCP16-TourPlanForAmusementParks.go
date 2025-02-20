package main

// LCP 16. 游乐园的游览计划
// 又到了一年一度的春游时间，小吴计划去游乐场游玩 1 天，游乐场总共有 N 个游乐项目，编号从 0 到 N-1。
// 小吴给每个游乐项目定义了一个非负整数值 value[i] 表示自己的喜爱值。
// 两个游乐项目之间会有双向路径相连，整个游乐场总共有 M 条双向路径，保存在二维数组 edges中。 
// 小吴计划选择一个游乐项目 A 作为这一天游玩的重点项目。
// 上午小吴准备游玩重点项目 A 以及与项目 A 相邻的两个项目 B、C （项目A、B与C要求是不同的项目，且项目B与项目C要求相邻），并返回 A ，即存在一条 A-B-C-A 的路径。 
// 下午，小吴决定再游玩重点项目 A以及与A相邻的两个项目 B'、C'，（项目A、B'与C'要求是不同的项目，且项目B'与项目C'要求相邻），并返回 A ，即存在一条 A-B'-C'-A 的路径。
// 下午游玩项目 B'、C' 可与上午游玩项目B、C存在重复项目。 
// 小吴希望提前安排好游玩路径，使得喜爱值之和最大。
// 请你返回满足游玩路径选取条件的最大喜爱值之和，如果没有这样的路径，返回 0。 
// 注意：一天中重复游玩同一个项目并不能重复增加喜爱值了。
// 例如：上下午游玩路径分别是 A-B-C-A与A-C-D-A 那么只能获得 value[A] + value[B] + value[C] + value[D] 的总和。

// 示例 1：
// 输入：edges = [[0,1],[1,2],[0,2]], value = [1,2,3]
// 输出：6
// 解释：喜爱值之和最高的方案之一是 0->1->2->0 与 0->2->1->0 。重复游玩同一点不重复计入喜爱值，返回1+2+3=6

// 示例 2：
// 输入：edges = [[0,2],[2,1]], value = [1,2,5]
// 输出：0
// 解释：无满足要求的游玩路径，返回 0

// 示例 3：
// 输入：edges = [[0,1],[0,2],[0,3],[0,4],[0,5],[1,3],[2,4],[2,5],[3,4],[3,5],[4,5]], value = [7,8,6,8,9,7]
// 输出：39
// 解释：喜爱值之和最高的方案之一是 3->0->1->3 与 3->4->5->3 。喜爱值最高为 7+8+8+9+7=39

// 限制：
//     3 <= value.length <= 10000
//     1 <= edges.length <= 10000
//     0 <= edges[i][0],edges[i][1] < value.length
//     0 <= value[i] <= 10000
//     edges中没有重复的边
//     edges[i][0] != edges[i][1]

import "fmt"
import "sort"

func maxWeight(edges [][]int, value []int) int {
    n := len(value) // 点数量
    m := len(edges) // 边数量
    // 保存每个点的度
    cnt := make([]int, n)
    for _, edge := range edges {
        cnt[edge[0]]++
        cnt[edge[1]]++
    }
    // 对边按2个端点的权重和从大到小排序
    sort.Slice(edges, func(i, j int) bool {
        return (value[edges[i][0]] + value[edges[i][1]]) > (value[edges[j][0]] + value[edges[j][1]])
    })
    // 将无向图构建为有向图, 方向为度小的点指向度大的点，度一样的时候，序号小的指向序号大的
    graph := make([][][2]int, n)
    for i := 0; i < n; i++ {
        graph[i] = make([][2]int, 0)
    }
    for i := 0; i < m; i++ {
        node1 := edges[i][0]
        node2 := edges[i][1]
        if cnt[node1] < cnt[node2] || (cnt[node1] == cnt[node2] && node1 < node2) {
            graph[node1] = append(graph[node1], [2]int{node2, i})
        } else {
            graph[node2] = append(graph[node2], [2]int{node1, i})
        }
    }
    // 获取所有的三元环
    edgeAndPoint := make([][]int, m)
    for i := 0; i < m; i++ {
        edgeAndPoint[i] = make([]int, 0)
    }
    tmpEdges1 := make([]int, n)
    for i := range tmpEdges1 {
        tmpEdges1[i] = 10000000
    }
    tmpEdges2 := make([]int, n)
    for i := 0; i < m; i++ {
        node1 := edges[i][0]
        node2 := edges[i][1]
        for _, neighbor := range graph[node1] {
            tmpEdges1[neighbor[0]] = i
            tmpEdges2[neighbor[0]] = neighbor[1]
        }
        for _, neighbor := range graph[node2] {
            if tmpEdges1[neighbor[0]] == i {
                edgeAndPoint[i] = append(edgeAndPoint[i], neighbor[0])
                edgeAndPoint[neighbor[1]] = append(edgeAndPoint[neighbor[1]], node1)
                edgeAndPoint[tmpEdges2[neighbor[0]]] = append(edgeAndPoint[tmpEdges2[neighbor[0]]], node2)
            }
        }
    }
    // 将从边到点的三元环转化为从点到边的三元环
    pointAndEdge := make([][]int, n)
    for i := 0; i < n; i++ {
        pointAndEdge[i] = make([]int, 0)
    }
    for i := 0; i < m; i++ {
        for _, point := range edgeAndPoint[i] {
            pointAndEdge[point] = append(pointAndEdge[point], i)
        }
    }
    res := 0
    for i := 0; i < n; i++ {
        lastIndex := len(pointAndEdge[i]) - 1
        if lastIndex < 0 { continue }
        for j := 0; j < min(3, len(pointAndEdge[i])) && lastIndex >= j; j++ {
            edgeIndex := pointAndEdge[i][j]
            node1 := edges[edgeIndex][0]
            node2 := edges[edgeIndex][1]
            for k := j; k <= lastIndex; k++ {
                edgeIndex2 := pointAndEdge[i][k]
                node3 := edges[edgeIndex2][0]
                node4 := edges[edgeIndex2][1]
                curRes := value[i] + value[node1] + value[node2]
                counter := 0
                if node3 != node1 && node3 != node2 {
                    curRes += value[node3]
                    counter++
                }
                if node4 != node1 && node4 != node2 {
                    curRes += value[node4]
                    counter++
                }
                if curRes > res {
                    res = curRes
                }
                if counter == 2 {
                    lastIndex = k - 1
                    break
                }
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：edges = [[0,1],[1,2],[0,2]], value = [1,2,3]
    // 输出：6
    // 解释：喜爱值之和最高的方案之一是 0->1->2->0 与 0->2->1->0 。重复游玩同一点不重复计入喜爱值，返回1+2+3=6
    fmt.Println(maxWeight([][]int{{0,1},{1,2},{0,2}}, []int{1,2,3})) // 6
    // 示例 2：
    // 输入：edges = [[0,2],[2,1]], value = [1,2,5]
    // 输出：0
    // 解释：无满足要求的游玩路径，返回 0
    fmt.Println(maxWeight([][]int{{0,2},{2,1}}, []int{1,2,5})) // 0
    // 示例 3：
    // 输入：edges = [[0,1],[0,2],[0,3],[0,4],[0,5],[1,3],[2,4],[2,5],[3,4],[3,5],[4,5]], value = [7,8,6,8,9,7]
    // 输出：39
    // 解释：喜爱值之和最高的方案之一是 3->0->1->3 与 3->4->5->3 。喜爱值最高为 7+8+8+9+7=39
    fmt.Println(maxWeight([][]int{{0,1},{0,2},{0,3},{0,4},{0,5},{1,3},{2,4},{2,5},{3,4},{3,5},{4,5}}, []int{7,8,6,8,9,7})) // 39

    fmt.Println(maxWeight([][]int{{0,1},{1,2},{2,0},{1,3},{3,4}}, []int{1,2,3,4,5})) // 6
}