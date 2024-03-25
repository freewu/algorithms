package main

// 2642. Design Graph With Shortest Path Calculator
// There is a directed weighted graph that consists of n nodes numbered from 0 to n - 1. 
// The edges of the graph are initially represented by the given array edges where edges[i] = [fromi, toi, edgeCosti] meaning that there is an edge from fromi to toi with the cost edgeCosti.

// Implement the Graph class:
//     Graph(int n, int[][] edges) initializes the object with n nodes and the given edges.
//     addEdge(int[] edge) adds an edge to the list of edges where edge = [from, to, edgeCost]. It is guaranteed that there is no edge between the two nodes before adding this one.
//     int shortestPath(int node1, int node2) returns the minimum cost of a path from node1 to node2. If no path exists, return -1. The cost of a path is the sum of the costs of the edges in the path.
    
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/01/11/graph3drawio-2.png" /> 
// Input
// ["Graph", "shortestPath", "shortestPath", "addEdge", "shortestPath"]
// [[4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]], [3, 2], [0, 3], [[1, 3, 4]], [0, 3]]
// Output
// [null, 6, -1, null, 6]
// Explanation
// Graph g = new Graph(4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]);
// g.shortestPath(3, 2); // return 6. The shortest path from 3 to 2 in the first diagram above is 3 -> 0 -> 1 -> 2 with a total cost of 3 + 2 + 1 = 6.
// g.shortestPath(0, 3); // return -1. There is no path from 0 to 3.
// g.addEdge([1, 3, 4]); // We add an edge from node 1 to node 3, and we get the second diagram above.
// g.shortestPath(0, 3); // return 6. The shortest path from 0 to 3 now is 0 -> 1 -> 3 with a total cost of 2 + 4 = 6.
 
// Constraints:
//     1 <= n <= 100
//     0 <= edges.length <= n * (n - 1)
//     edges[i].length == edge.length == 3
//     0 <= fromi, toi, from, to, node1, node2 <= n - 1
//     1 <= edgeCosti, edgeCost <= 10^6
//     There are no repeated edges and no self-loops in the graph at any point.
//     At most 100 calls will be made for addEdge.
//     At most 100 calls will be made for shortestPath.

import "fmt"
import "container/heap"
import "math"

type Edge struct {
    head int
    cost int
}

type Graph struct {
    adj [][]Edge
}

func Constructor(n int, edges [][]int) Graph {
    adj := make([][]Edge, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
    }
    return Graph{adj}
}

func (this *Graph) AddEdge(edge []int)  {
    this.adj[edge[0]] = append(this.adj[edge[0]], Edge{edge[1], edge[2]})
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
    n := len(this.adj)
    items := make([]*Item, n)
    for i := 0; i < n; i++ {
        items[i] = &Item{i, math.MaxInt, i}
    }
    items[node1].priority = 0
    queue := make(PriorityQueue, n)
    copy(queue, items)
    heap.Init(&queue)
    for len(queue) > 0 {
        i := heap.Pop(&queue).(*Item)
        node, priority := i.value, i.priority
        if priority == math.MaxInt || node == node2 {
            break
        }
        for _, e := range this.adj[node] {
            j := items[e.head]
            if h := priority + e.cost; h < j.priority {
                j.priority = h
                heap.Fix(&queue, j.index)
            }
        } 
    }
    if items[node2].priority == math.MaxInt {
        return -1
    }
    return items[node2].priority
}

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

const inf = math.MaxInt / 2 // 防止更新最短路时加法溢出

type Graph1 [][]int

func Constructor1(n int, edges [][]int) Graph1 {
    g := make([][]int, n) // 邻接矩阵
    for i := range g {
        g[i] = make([]int, n)
        for j := range g[i] {
            g[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
        }
    }
    for _, e := range edges {
        g[e[0]][e[1]] = e[2] // 添加一条边（题目保证没有重边）
    }
    return g
}

func (g Graph1) AddEdge(e []int) {
    g[e[0]][e[1]] = e[2] // 添加一条边（题目保证这条边之前不存在）
}

func (g Graph1) ShortestPath(start, end int) int {
    n := len(g)
    dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
    for i := range dis {
        dis[i] = inf
    }
    dis[start] = 0
    vis := make([]bool, n)
    for {
        x := -1
        for i, b := range vis {
            if !b && (x < 0 || dis[i] < dis[x]) {
                x = i
            }
        }
        if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
            return -1
        }
        if x == end { // 找到终点，提前退出
            return dis[x]
        }
        vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
        for y, w := range g[x] {
            dis[y] = min(dis[y], dis[x]+w) // 更新最短路长度
        }
    }
}

func main() {
    // Graph g = new Graph(4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]);
    obj := Constructor(
        4,
        [][]int{
            []int{0, 2, 5},
            []int{0, 1, 2},
            []int{1, 2, 1},
            []int{3, 0, 3},
        },
    )
    // g.shortestPath(3, 2); // return 6. The shortest path from 3 to 2 in the first diagram above is 3 -> 0 -> 1 -> 2 with a total cost of 3 + 2 + 1 = 6.
    fmt.Println(obj.ShortestPath(3,2)) // 6
    // g.shortestPath(0, 3); // return -1. There is no path from 0 to 3.
    fmt.Println(obj.ShortestPath(0,3)) // -1
    // g.addEdge([1, 3, 4]); // We add an edge from node 1 to node 3, and we get the second diagram above.
    obj.AddEdge([]int{1, 3, 4})
    // g.shortestPath(0, 3); // return 6. The shortest path from 0 to 3 now is 0 -> 1 -> 3 with a total cost of 2 + 4 = 6.
    fmt.Println(obj.ShortestPath(0,3)) // 6


    // Graph g = new Graph(4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]);
    obj1 := Constructor1(
        4,
        [][]int{
            []int{0, 2, 5},
            []int{0, 1, 2},
            []int{1, 2, 1},
            []int{3, 0, 3},
        },
    )
    // g.shortestPath(3, 2); // return 6. The shortest path from 3 to 2 in the first diagram above is 3 -> 0 -> 1 -> 2 with a total cost of 3 + 2 + 1 = 6.
    fmt.Println(obj1.ShortestPath(3,2)) // 6
    // g.shortestPath(0, 3); // return -1. There is no path from 0 to 3.
    fmt.Println(obj1.ShortestPath(0,3)) // -1
    // g.addEdge([1, 3, 4]); // We add an edge from node 1 to node 3, and we get the second diagram above.
    obj1.AddEdge([]int{1, 3, 4})
    // g.shortestPath(0, 3); // return 6. The shortest path from 0 to 3 now is 0 -> 1 -> 3 with a total cost of 2 + 4 = 6.
    fmt.Println(obj1.ShortestPath(0,3)) // 6
}