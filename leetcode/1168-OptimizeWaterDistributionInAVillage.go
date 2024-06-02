package main

// 1168. Optimize Water Distribution in a Village
// There are n houses in a village. We want to supply water for all the houses by building wells and laying pipes.

// For each house i, we can either build a well inside it directly with cost wells[i - 1] (note the -1 due to 0-indexing),
// or pipe in water from another well to it. 
// The costs to lay pipes between houses are given by the array pipes where each pipes[j] = [house1j, house2j, costj] represents the cost to connect house1j and house2j together using a pipe. 
// Connections are bidirectional, and there could be multiple valid connections between the same two houses with different costs.

// Return the minimum total cost to supply water to all houses.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/05/22/1359_ex1.png" />
// Input: n = 3, wells = [1,2,2], pipes = [[1,2,1],[2,3,1]]
// Output: 3
// Explanation: The image shows the costs of connecting houses using pipes.
// The best strategy is to build a well in the first house with cost 1 and connect the other houses to it with cost 2 so the total cost is 3.

// Example 2:
// Input: n = 2, wells = [1,1], pipes = [[1,2,1],[1,2,2]]
// Output: 2
// Explanation: We can supply water with cost two using one of the three options:
// Option 1:
//   - Build a well inside house 1 with cost 1.
//   - Build a well inside house 2 with cost 1.
// The total cost will be 2.
// Option 2:
//   - Build a well inside house 1 with cost 1.
//   - Connect house 2 with house 1 with cost 1.
// The total cost will be 2.
// Option 3:
//   - Build a well inside house 2 with cost 1.
//   - Connect house 1 with house 2 with cost 1.
// The total cost will be 2.
// Note that we can connect houses 1 and 2 with cost 1 or with cost 2 but we will always choose the cheapest option. 
 
// Constraints:
//     2 <= n <= 10^4
//     wells.length == n
//     0 <= wells[i] <= 10^5
//     1 <= pipes.length <= 10^4
//     pipes[j].length == 3
//     1 <= house1j, house2j <= n
//     0 <= costj <= 10^5
//     house1j != house2j

import "fmt"
import "container/heap"
import "slices"

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
    // 不能用 kruskal，因为至少要有一个 house 是用 well 的
    // 所以不能一开始就将所有 pipe 加入优先队列
    graph := make([][][]int, n) // 将邻接矩阵变为邻接链表
    for i := 0; i < n; i++ {
        graph[i] = make([][]int, 0)
    }
    for _, pipe := range pipes {
        from, to, cost := pipe[0] - 1, pipe[1] - 1, pipe[2] // 将所有 house idx 减 1
        graph[from] = append(graph[from], []int{to, cost})
        graph[to] = append(graph[to], []int{from, cost})
    }
    pq := Heap(make([][]int, 0)) // 初始化 pq，只包括井
    for house, cost := range wells {
        heap.Push(&pq, []int{house, cost})
    }
    visited, numVisited, totalCost := make([]bool, n), 0, 0
    for numVisited < n {
        edge := heap.Pop(&pq).([]int)
        house, cost := edge[0], edge[1]
        // 因为刚开始所有的井都被加入了
        // 而且通往一个 house 的边可能被重复加入
        // 所以 pop 出来的 house 仍有可能重复
        if visited[house] {
            continue
        }
        visited[house] = true
        numVisited++
        totalCost += cost
        for _, edge := range graph[house] {
            if visited[edge[0]] {
                continue
            }
            heap.Push(&pq, edge)
        }
    }
    return totalCost
}

// 优先队列的每个元素是 长度为 2 的[]int
// 0 元素是尚未通水的 house number，1 元素是 cost
type Heap [][]int
func (h Heap) Len() int {return len(h)}
func (h Heap) Less(i, j int) bool {return h[i][1] < h[j][1]}     // compare cost
func (h Heap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.([]int))
}
func (h *Heap) Pop() interface{} {
    old := *h
    ret := old[len(old) - 1]
    *h = old[:len(old) - 1]
    return ret
}


// 水资源分配优化 并查集
// 村里面一共有 n 栋房子。我们希望通过建造水井和铺设管道来为所有房子供水。
// 对于每个房子 i，我们有两种可选的供水方案：一种是直接在房子内建造水井
// 成本为 wells[i - 1] （注意 -1 ，因为 索引从0开始 ）
// 另一种是从另一口井铺设管道引水，数组 pipes 给出了在房子间铺设管道的成本，
// 其中每个 pipes[j] = [house1j, house2j, costj]
// 代表用管道将 house1j 和 house2j连接在一起的成本。连接是双向的。
// 请返回 为所有房子都供水的最低总成本
func minCostToSupplyWater1(n int, wells []int, pipes [][]int) int {
    // 最小生成树
    // 一个巧妙的转换: 给房子建造水井转换为了铺设管道的另一种特例,相当于有一个水源点给所有房子拉了一条水管
    // 房子编号 1->n,  边的数量和顶点的数量差不多,可以使用Kruskal方法
    father := make([]int, n+1)
    build := func() {
        for i := 1; i <= n; i++ {
            father[i] = i
        }
    }
    var find func(x int) int
    find = func(x int) int {
        if x != father[x] {
            father[x] = find(father[x])
        }
        return father[x]
    }
    union := func(x, y int) bool {
        fx := find(x)
        fy := find(y)
        if fx != fy {
            father[fx] = fy
            return true
        }
        return false
    }
    edges := make([][]int, 0, n+len(pipes))
    for i, cost := range wells { // 水井认为是0号顶点
        edges = append(edges, []int{0, i + 1, cost})
    }
    edges = append(edges, pipes...)
    slices.SortFunc(edges, func(a, b []int) int {
        return a[2] - b[2]
    })
    // kruskal 构建 mst
    build()
    res := 0
    for _, edge := range edges {
        if union(edge[0], edge[1]) {
            res += edge[2]
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/05/22/1359_ex1.png" />
    // Input: n = 3, wells = [1,2,2], pipes = [[1,2,1],[2,3,1]]
    // Output: 3
    // Explanation: The image shows the costs of connecting houses using pipes.
    // The best strategy is to build a well in the first house with cost 1 and connect the other houses to it with cost 2 so the total cost is 3.
    fmt.Println(minCostToSupplyWater(3,[]int{1,2,2},[][]int{{1,2,1},{2,3,1}})) // 3
    // Example 2:
    // Input: n = 2, wells = [1,1], pipes = [[1,2,1],[1,2,2]]
    // Output: 2
    // Explanation: We can supply water with cost two using one of the three options:
    // Option 1:
    //   - Build a well inside house 1 with cost 1.
    //   - Build a well inside house 2 with cost 1.
    // The total cost will be 2.
    // Option 2:
    //   - Build a well inside house 1 with cost 1.
    //   - Connect house 2 with house 1 with cost 1.
    // The total cost will be 2.
    // Option 3:
    //   - Build a well inside house 2 with cost 1.
    //   - Connect house 1 with house 2 with cost 1.
    // The total cost will be 2.
    // Note that we can connect houses 1 and 2 with cost 1 or with cost 2 but we will always choose the cheapest option. 
    fmt.Println(minCostToSupplyWater(2,[]int{1,1},[][]int{{1,2,1},{1,2,2}})) // 2

    fmt.Println(minCostToSupplyWater1(3,[]int{1,2,2},[][]int{{1,2,1},{2,3,1}})) // 3
    fmt.Println(minCostToSupplyWater1(2,[]int{1,1},[][]int{{1,2,1},{1,2,2}})) // 2
}