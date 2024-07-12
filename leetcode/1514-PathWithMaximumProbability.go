package main

// 1514. Path with Maximum Probability
// You are given an undirected weighted graph of n nodes (0-indexed), 
// represented by an edge list where edges[i] = [a, b] is an undirected edge connecting the nodes a 
// and b with a probability of success of traversing that edge succProb[i].

// Given two nodes start and end, find the path with the maximum probability of success to go from start to end 
// and return its success probability.

// If there is no path from start to end, return 0. 
// Your answer will be accepted if it differs from the correct answer by at most 1e-5.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/09/20/1558_ex1.png" />
// Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
// Output: 0.25000
// Explanation: There are two paths from start to end, one having a probability of success = 0.2 and the other has 0.5 * 0.5 = 0.25.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/09/20/1558_ex2.png" />
// Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2
// Output: 0.30000

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/09/20/1558_ex3.png" />
// Input: n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
// Output: 0.00000
// Explanation: There is no path between 0 and 2.

// Constraints:
//     2 <= n <= 10^4
//     0 <= start, end < n
//     start != end
//     0 <= a, b < n
//     a != b
//     0 <= succProb.length == edges.length <= 2*10^4
//     0 <= succProb[i] <= 1
//     There is at most one edge between every two nodes.

import "fmt"
import "container/heap"
import "math"

type Edge struct {
    to     int
    weight float64
}

type Node struct {
    node   int
    weight float64
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].weight > pq[j].weight }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[0 : n-1]
    return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    graph := make([][]Edge, n)
    for i, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], Edge{edge[1], math.Log(succProb[i])})
        graph[edge[1]] = append(graph[edge[1]], Edge{edge[0], math.Log(succProb[i])})
    }
    dist := make([]float64, n)
    for i := range dist {
        dist[i] = math.Inf(-1)
    }
    dist[start] = 0
    pq := PriorityQueue{}
    heap.Push(&pq, &Node{start, 0})
    for pq.Len() > 0 {
        curr := heap.Pop(&pq).(*Node)
        if curr.node == end {
            return math.Exp(dist[end])
        }
        if curr.weight < dist[curr.node] {
            continue
        }
        for _, edge := range graph[curr.node] {
            if nextDist := dist[curr.node] + edge.weight; nextDist > dist[edge.to] {
                dist[edge.to] = nextDist
                heap.Push(&pq, &Node{edge.to, nextDist})
            }
        }
    }
    return 0
}

func maxProbability1(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
    dp := make([]float64, n)
    dp[start_node] = 1.0
    for {
        k := false
        for j := 0; j < len(edges); j++ {
            if dp[edges[j][0]] * succProb[j] > dp[edges[j][1]] {
                dp[edges[j][1]] = dp[edges[j][0]] * succProb[j]
                k = true
            }
            if dp[edges[j][1]] * succProb[j] > dp[edges[j][0]] {
                dp[edges[j][0]] = dp[edges[j][1]] * succProb[j]
                k = true
            }
        }
        if !k {
            break
        }
    }
    return dp[end_node]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/09/20/1558_ex1.png" />
    // Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
    // Output: 0.25000
    // Explanation: There are two paths from start to end, one having a probability of success = 0.2 and the other has 0.5 * 0.5 = 0.25.
    fmt.Println(maxProbability(3,[][]int{{0,1},{1,2},{0,2}}, []float64{0.5,0.5,0.2}, 0, 2)) // 0.25000
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/09/20/1558_ex2.png" />
    // Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2
    // Output: 0.30000
    fmt.Println(maxProbability(3,[][]int{{0,1},{1,2},{0,2}}, []float64{0.5,0.5,0.3}, 0, 2)) // 0.30000
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/09/20/1558_ex3.png" />
    // Input: n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
    // Output: 0.00000
    // Explanation: There is no path between 0 and 2.
    fmt.Println(maxProbability(3,[][]int{{0,1}}, []float64{0.5}, 0, 2)) // 0.00000

    fmt.Println(maxProbability1(3,[][]int{{0,1},{1,2},{0,2}}, []float64{0.5,0.5,0.2}, 0, 2)) // 0.25000
    fmt.Println(maxProbability1(3,[][]int{{0,1},{1,2},{0,2}}, []float64{0.5,0.5,0.3}, 0, 2)) // 0.30000
    fmt.Println(maxProbability1(3,[][]int{{0,1}}, []float64{0.5}, 0, 2)) // 0.00000
}