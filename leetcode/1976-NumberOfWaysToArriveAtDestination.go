package main

// 1976. Number of Ways to Arrive at Destination
// You are in a city that consists of n intersections numbered from 0 to n - 1 with bi-directional roads between some intersections. 
// The inputs are generated such that you can reach any intersection from any other intersection and that there is at most one road between any two intersections.
// You are given an integer n and a 2D integer array roads where roads[i] = [ui, vi, timei] means that there is a road between intersections ui and vi that takes timei minutes to travel. 
// You want to know in how many ways you can travel from intersection 0 to intersection n - 1 in the shortest amount of time.
// Return the number of ways you can arrive at your destination in the shortest amount of time. 
// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/17/graph2.png" />
// Input: n = 7, roads = [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]]
// Output: 4
// Explanation: The shortest amount of time it takes to go from intersection 0 to intersection 6 is 7 minutes.
// The four ways to get there in 7 minutes are:
// - 0 ➝ 6
// - 0 ➝ 4 ➝ 6
// - 0 ➝ 1 ➝ 2 ➝ 5 ➝ 6
// - 0 ➝ 1 ➝ 3 ➝ 5 ➝ 6

// Example 2:
// Input: n = 2, roads = [[1,0,10]]
// Output: 1
// Explanation: There is only one way to go from intersection 0 to intersection 1, and it takes 10 minutes.
 
// Constraints:
//         1 <= n <= 200
//         n - 1 <= roads.length <= n * (n - 1) / 2
//         roads[i].length == 3
//         0 <= ui, vi <= n - 1
//         1 <= timei <= 10^9
//         ui != vi
//         There is at most one road connecting any two intersections.
//         You can reach any intersection from any other intersection.

import "fmt"
import "container/heap"

type Item struct {
    City  int
    Time  int64
    Index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Time < pq[j].Time }
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].Index, pq[j].Index = i, j
}
func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    item.Index = len(*pq)
    *pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
    item := (*pq)[len(*pq)-1]
    *pq = (*pq)[:len(*pq)-1]
    return item
}
func (pq *PriorityQueue) Update(item *Item, Time int64) {
    item.Time = Time
    heap.Fix(pq, item.Index)
}

type Pair struct {
    City int
    Time int64
}

func dijkstra(n int, adjList [][]Pair) int {
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    numberOfShortestWays := make([]int, n)
    numberOfShortestWays[0] = 1
    shortestTimes := make([]int64, n)
    for idx := range shortestTimes {
        shortestTimes[idx] = 1 << 61
    }
    shortestTimes[0] = 0
    heap.Push(&pq, &Item{City: 0, Time: 0})
    for len(pq) > 0 {
        item := heap.Pop(&pq).(*Item)
        City, Time := item.City, item.Time
        for _, toVisit := range adjList[City] {
            c1, t1 := toVisit.City, toVisit.Time
            if t1+Time < shortestTimes[c1] {
                numberOfShortestWays[c1] = numberOfShortestWays[City]
                shortestTimes[c1] = t1 + Time
                heap.Push(&pq, &Item{City: c1, Time: shortestTimes[c1]})
            } else if t1+Time == shortestTimes[c1] {
                numberOfShortestWays[c1] = (numberOfShortestWays[c1] + numberOfShortestWays[City]) % 1_000_000_007
            }
        }
    }
    return numberOfShortestWays[n-1]
}

func countPaths(n int, roads [][]int) int {
    adjList := make([][]Pair, n)
    for _, road := range roads {
        adjList[road[0]] = append(adjList[road[0]], Pair{City: road[1], Time: int64(road[2])})
        adjList[road[1]] = append(adjList[road[1]], Pair{City: road[0], Time: int64(road[2])})
    }
    return dijkstra(n, adjList)
}

// best solution
func countPaths1(n int, roads [][]int) (ans int) {
    g := make([][]int, n)
    for i := range g {
        g[i] = make([]int, n)
        for j := range g[i] {
            g[i][j] = 1e18
        }
    }
    for _, r := range roads {
        v, w, wt := r[0], r[1], r[2]
        g[v][w] = wt
        g[w][v] = wt
    }
    // 求 0 到其余点的最短路
    d := make([]int, n)
    for i := range d {
        d[i] = 1e18
    }
    d[0] = 0
    used := make([]bool, n)
    for {
        v := -1
        for w, u := range used {
            if !u && (v < 0 || d[w] < d[v]) {
                v = w
            }
        }
        if v < 0 {
            break
        }
        used[v] = true
        for w, wt := range g[v] {
            if newD := d[v] + wt; newD < d[w] {
                d[w] = newD
            }
        }
    }
    // 最短路构成了一个 DAG，这里不需要建一个新图，直接根据距离来判断每条边是否在 DAG 上
    // 计算 DAG 的入度数组
    deg := make([]int, n)
    for v, r := range g {
        for w, wt := range r {
            if d[v]+wt == d[w] { // 这条边在 DAG 上
                deg[w]++
            }
        }
    }
    // 在 DAG 上跑拓扑排序
    dp := make([]int, n) // dp[i] 表示 0 到 i 的最短路个数
    dp[0] = 1
    q := []int{0}
    for len(q) > 0 {
        v := q[0]
        q = q[1:]
        for w, wt := range g[v] {
            if d[v]+wt == d[w] { // 这条边在 DAG 上
                dp[w] = (dp[w] + dp[v]) % (1e9 + 7)
                if deg[w]--; deg[w] == 0 {
                    q = append(q, w)
                }
            }
        }
    }
    return dp[n-1]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/17/graph2.png" />
    // Input: n = 7, roads = [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]]
    // Output: 4
    // Explanation: The shortest amount of time it takes to go from intersection 0 to intersection 6 is 7 minutes.
    // The four ways to get there in 7 minutes are:
    // - 0 ➝ 6
    // - 0 ➝ 4 ➝ 6
    // - 0 ➝ 1 ➝ 2 ➝ 5 ➝ 6
    // - 0 ➝ 1 ➝ 3 ➝ 5 ➝ 6
    fmt.Println(countPaths(
        7,
        [][]int{[]int{0,6,7},[]int{0,1,2},[]int{1,2,3},[]int{1,3,3},[]int{6,3,3},[]int{3,5,1},[]int{6,5,1},[]int{0,4,5},[]int{2,5,1},[]int{4,6,2}},
    )) // 4
    // Example 2:
    // Input: n = 2, roads = [[1,0,10]]
    // Output: 1
    // Explanation: There is only one way to go from intersection 0 to intersection 1, and it takes 10 minutes.
    fmt.Println(countPaths(
        2,
        [][]int{[]int{1,0,10}},
    )) // 1

    fmt.Println(countPaths1(
        7,
        [][]int{[]int{0,6,7},[]int{0,1,2},[]int{1,2,3},[]int{1,3,3},[]int{6,3,3},[]int{3,5,1},[]int{6,5,1},[]int{0,4,5},[]int{2,5,1},[]int{4,6,2}},
    )) // 4
    fmt.Println(countPaths1(
        2,
        [][]int{[]int{1,0,10}},
    )) // 1
}