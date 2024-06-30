package main

// 2065. Maximum Path Quality of a Graph
// There is an undirected graph with n nodes numbered from 0 to n - 1 (inclusive). 
// You are given a 0-indexed integer array values where values[i] is the value of the ith node. 
// You are also given a 0-indexed 2D integer array edges, where each edges[j] = [uj, vj, timej] indicates 
// that there is an undirected edge between the nodes uj and vj, and it takes timej seconds to travel between the two nodes. 
// Finally, you are given an integer maxTime.

// A valid path in the graph is any path that starts at node 0, ends at node 0, and takes at most maxTime seconds to complete. 
// You may visit the same node multiple times. 
// The quality of a valid path is the sum of the values of the unique nodes visited in the path (each node's value is added at most once to the sum).

// Return the maximum quality of a valid path.
// Note: There are at most four edges connected to each node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/19/ex1drawio.png" />
// Input: values = [0,32,10,43], edges = [[0,1,10],[1,2,15],[0,3,10]], maxTime = 49
// Output: 75
// Explanation:
// One possible path is 0 -> 1 -> 0 -> 3 -> 0. The total time taken is 10 + 10 + 10 + 10 = 40 <= 49.
// The nodes visited are 0, 1, and 3, giving a maximal path quality of 0 + 32 + 43 = 75.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/19/ex2drawio.png" />
// Input: values = [5,10,15,20], edges = [[0,1,10],[1,2,10],[0,3,10]], maxTime = 30
// Output: 25
// Explanation:
// One possible path is 0 -> 3 -> 0. The total time taken is 10 + 10 = 20 <= 30.
// The nodes visited are 0 and 3, giving a maximal path quality of 5 + 20 = 25.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/10/19/ex31drawio.png" />
// Input: values = [1,2,3,4], edges = [[0,1,10],[1,2,11],[2,3,12],[1,3,13]], maxTime = 50
// Output: 7
// Explanation:
// One possible path is 0 -> 1 -> 3 -> 1 -> 0. The total time taken is 10 + 13 + 13 + 10 = 46 <= 50.
// The nodes visited are 0, 1, and 3, giving a maximal path quality of 1 + 2 + 4 = 7.

// Constraints:
//     n == values.length
//     1 <= n <= 1000
//     0 <= values[i] <= 10^8
//     0 <= edges.length <= 2000
//     edges[j].length == 3
//     0 <= uj < vj <= n - 1
//     10 <= timej, maxTime <= 100
//     All the pairs [uj, vj] are unique.
//     There are at most four edges connected to each node.
//     The graph may not be connected.

import "fmt"
import "container/heap"

// BackTracking
func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
    type Edge struct {
        u, v, time  int
    }
    type State struct {
        timeSpent, quality int
        visited []int
    }
    res, adjustmentMap := 0, make(map[int][]Edge)
    for _, e := range edges {
        adjustmentMap[e[0]] = append(adjustmentMap[e[0]], Edge{e[0], e[1], e[2]})
        adjustmentMap[e[1]] = append(adjustmentMap[e[1]], Edge{e[1], e[0], e[2]})
    }
    var backtracking func(state State, edge int) 
    backtracking = func(state State, edge int) {
        if state.timeSpent > maxTime {
            return
        }
        if edge == 0 && state.quality > res {
            res = state.quality
        }
        neighbours := adjustmentMap[edge]
        for _, nei := range neighbours {
            neiVisited := state.visited[nei.v]
            state.visited[nei.v] = 1
            qualityDelta := 0
            if neiVisited == 0 {
                qualityDelta += values[nei.v]
            }
            state.quality += qualityDelta
            state.timeSpent += nei.time
            backtracking(state, nei.v)	
            state.visited[nei.v] = neiVisited
            state.quality -= qualityDelta
            state.timeSpent -= nei.time
        }
    }
    visited := make([]int, len(values))
    visited[0] = 1
    backtracking(State{0, values[0], visited}, 0)
    return res
}

func maximalPathQuality1(values []int, edges [][]int, maxTime int) int {
    type edge struct{ to, time int }
    res, n, inf := 0, len(values), 1 << 32 - 1
    g := make([][]edge, n)
    for _, e := range edges {
        x, y, t := e[0], e[1], e[2]
        g[x] = append(g[x], edge{y, t})
        g[y] = append(g[y], edge{x, t})
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // Dijkstra 算法
    dis := make([]int, n)
    for i := 1; i < n; i++ {
        dis[i] = inf
    }
    h := hp{{0, 0}}
    for len(h) > 0 {
        p := heap.Pop(&h).(pair)
        dx := p.dis
        x := p.x
        if dx > dis[x] { // x 之前出堆过
            continue
        }
        for _, e := range g[x] {
            y := e.to
            newDis := dx + e.time
            if newDis < dis[y] {
                dis[y] = newDis // 更新 x 的邻居的最短路
                heap.Push(&h, pair{newDis, y})
            }
        }
    }
    vis := make([]bool, n)
    vis[0] = true
    var dfs func(int, int, int)
    dfs = func(x, sumTime, sumValue int) {
        if x == 0 {
            res = max(res, sumValue)
            // 注意这里没有 return，还可以继续走
        }
        for _, e := range g[x] {
            y, t := e.to, e.time
            // 相比方法一，这里多了 dis[y]
            if sumTime+t+dis[y] > maxTime {
                continue
            }
            if vis[y] {
                dfs(y, sumTime+t, sumValue)
            } else {
                vis[y] = true
                // 每个节点的价值至多算入价值总和中一次
                dfs(y, sumTime+t, sumValue+values[y])
                vis[y] = false // 恢复现场
            }
        }
    }
    dfs(0, 0, values[0])
    return res
}

type pair struct{ dis, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/19/ex1drawio.png" />
    // Input: values = [0,32,10,43], edges = [[0,1,10],[1,2,15],[0,3,10]], maxTime = 49
    // Output: 75
    // Explanation:
    // One possible path is 0 -> 1 -> 0 -> 3 -> 0. The total time taken is 10 + 10 + 10 + 10 = 40 <= 49.
    // The nodes visited are 0, 1, and 3, giving a maximal path quality of 0 + 32 + 43 = 75.
    fmt.Println(maximalPathQuality([]int{0,32,10,43},[][]int{{0,1,10},{1,2,15},{0,3,10}}, 49)) // 75
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/19/ex2drawio.png" />
    // Input: values = [5,10,15,20], edges = [[0,1,10],[1,2,10],[0,3,10]], maxTime = 30
    // Output: 25
    // Explanation:
    // One possible path is 0 -> 3 -> 0. The total time taken is 10 + 10 = 20 <= 30.
    // The nodes visited are 0 and 3, giving a maximal path quality of 5 + 20 = 25.
    fmt.Println(maximalPathQuality([]int{5,10,15,20},[][]int{{0,1,10},{1,2,15},{0,3,10}}, 30)) // 25
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/10/19/ex31drawio.png" />
    // Input: values = [1,2,3,4], edges = [[0,1,10],[1,2,11],[2,3,12],[1,3,13]], maxTime = 50
    // Output: 7
    // Explanation:
    // One possible path is 0 -> 1 -> 3 -> 1 -> 0. The total time taken is 10 + 13 + 13 + 10 = 46 <= 50.
    // The nodes visited are 0, 1, and 3, giving a maximal path quality of 1 + 2 + 4 = 7.
    fmt.Println(maximalPathQuality([]int{1,2,3,4},[][]int{{0,1,10},{1,2,11},{2,3,12},{1,3,13}}, 50)) // 7

    fmt.Println(maximalPathQuality1([]int{0,32,10,43},[][]int{{0,1,10},{1,2,15},{0,3,10}}, 49)) // 75
    fmt.Println(maximalPathQuality1([]int{5,10,15,20},[][]int{{0,1,10},{1,2,15},{0,3,10}}, 30)) // 25
    fmt.Println(maximalPathQuality1([]int{1,2,3,4},[][]int{{0,1,10},{1,2,11},{2,3,12},{1,3,13}}, 50)) // 7
}