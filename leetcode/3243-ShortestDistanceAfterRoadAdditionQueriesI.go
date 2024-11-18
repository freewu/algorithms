package main

// 3243. Shortest Distance After Road Addition Queries I
// You are given an integer n and a 2D integer array queries.

// There are n cities numbered from 0 to n - 1. 
// Initially, there is a unidirectional road from city i to city i + 1 for all 0 <= i < n - 1.

// queries[i] = [ui, vi] represents the addition of a new unidirectional road from city ui to city vi. 
// After each query, you need to find the length of the shortest path from city 0 to city n - 1.

// Return an array answer where for each i in the range [0, queries.length - 1], 
// answer[i] is the length of the shortest path from city 0 to city n - 1 after processing the first i + 1 queries.

// Example 1:
// Input: n = 5, queries = [[2,4],[0,2],[0,4]]
// Output: [3,2,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image8.jpg" />
// After the addition of the road from 2 to 4, the length of the shortest path from 0 to 4 is 3.
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image9.jpg" />
// After the addition of the road from 0 to 2, the length of the shortest path from 0 to 4 is 2.
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image10.jpg" />
// After the addition of the road from 0 to 4, the length of the shortest path from 0 to 4 is 1.

// Example 2:
// Input: n = 4, queries = [[0,3],[0,2]]
// Output: [1,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image11.jpg" />
// After the addition of the road from 0 to 3, the length of the shortest path from 0 to 3 is 1.
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image12.jpg" />
// After the addition of the road from 0 to 2, the length of the shortest path remains 1.

// Constraints:
//     3 <= n <= 500
//     1 <= queries.length <= 500
//     queries[i].length == 2
//     0 <= queries[i][0] < queries[i][1] < n
//     1 < queries[i][1] - queries[i][0]
//     There are no repeated roads among the queries.

import "fmt"
import "container/heap"

type Item struct {
    City         int
    ShortestPath int
    Index        int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].ShortestPath < pq[j].ShortestPath}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].Index, pq[j].Index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    item.Index = len(*pq)
    (*pq) = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    item := (*pq)[len(*pq)-1]
    (*pq) = (*pq)[:len(*pq)-1]
    return item
}

func (pq *PriorityQueue) Update(item *Item, shortestPath int) {
    item.ShortestPath = shortestPath
    heap.Fix(pq, item.Index)
}


func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    adjList := make([][]int, n)
    for i := 0; i < n-1; i++ {
        adjList[i] = append(adjList[i], i+1)
    }
    shortestPaths := make([]int, 0)
    dijkstra := func(n int, adjList [][]int) int {
        pq := make(PriorityQueue, 0)
        heap.Init(&pq)
        shortestPaths := make([]int, n)
        for idx := range shortestPaths {
            shortestPaths[idx] = 1 << 31
        }
        shortestPaths[0] = 0
        heap.Push(&pq, &Item{City: 0, ShortestPath: 0})
        for len(pq) > 0 {
            item := heap.Pop(&pq).(*Item)
            city := item.City
            shortestPath := item.ShortestPath
            for _, canVisit := range adjList[city] {
                if shortestPath+1 < shortestPaths[canVisit] {
                    shortestPaths[canVisit] = shortestPath + 1
                    heap.Push(&pq, &Item{City: canVisit, ShortestPath: shortestPaths[canVisit]})
                }
            }
        }
        return shortestPaths[n-1]
    }
    for _, query := range queries {
        adjList[query[0]] = append(adjList[query[0]], query[1])
        shortestPaths = append(shortestPaths, dijkstra(n, adjList))
    }
    return shortestPaths
}

// bfs
func shortestDistanceAfterQueries1(n int, queries [][]int) []int {
    nexts := make([][]int, n)
    for i := range nexts[1:] {
        nexts[i] = []int{ i + 1}
    }
    bfs := func() int {
        queue, visited, steps := []int{ 0 }, make([]bool, n), 0
        visited[0] = true
        for len(queue) != 0 {
            size := len(queue)
            steps++
            for _, cur := range queue {
                for _, next := range nexts[cur] {
                    if visited[next] { continue }
                    if next == n - 1 { return steps }
                    queue = append(queue, next) // push
                    visited[next] = true
                }
            }
            queue = queue[size:] // pop
        }
        return -1
    }
    res := make([]int, len(queries))
    for i := range queries {
        u, v := queries[i][0], queries[i][1]
        nexts[u] = append(nexts[u], v)
        res[i] = bfs()
    }
    return res
}

type PriorityQueueItem struct {
    node, distance int
}

type PriorityQueue1 struct {
    items []PriorityQueueItem
}

func (pq *PriorityQueue1) Push(item PriorityQueueItem) {
    pq.items = append(pq.items, item)
    pq.bubbleUp(len(pq.items) - 1)
}

func (pq *PriorityQueue1) bubbleUp(index int) {
    for index > 0 {
        parent := (index - 1) / 2
        if pq.items[index].distance >= pq.items[parent].distance { break }
        pq.items[index], pq.items[parent] = pq.items[parent], pq.items[index]
        index = parent
    }
}

func (pq *PriorityQueue1) bubbleDown(index int) {
    lastIndex := len(pq.items) - 1
    for {
        left, right, smallest := 2 * index + 1, 2 * index + 2, index
        if left <= lastIndex && pq.items[left].distance < pq.items[smallest].distance { smallest = left }
        if right <= lastIndex && pq.items[right].distance < pq.items[smallest].distance { smallest = right }
        if smallest == index { break }
        pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
        index = smallest
    }
}

func shortestDistanceAfterQueries2(n int, queries [][]int) []int {
    res, extraEdges, dist :=  make([]int, len(queries)), make([][]int, n), make([]int, n)
    pq := &PriorityQueue1{}
    for i := 0; i < n; i++ {
        dist[i] = i
    }
    for i, q := range queries {
        u, v := q[0], q[1]
        extraEdges[u] = append(extraEdges[u], v)
        if len(pq.items) > 0 {
            pq.items = pq.items[:0]
        }
        pq.Push(PriorityQueueItem{u, dist[u]})
        for len(pq.items) > 0 {
            current := pq.items[0]
            last := pq.items[len(pq.items)-1]
            pq.items = pq.items[:len(pq.items)-1]
            if len(pq.items) > 0 {
                pq.items[0] = last
                pq.bubbleDown(0)
            }
            if current.distance > dist[current.node] { continue }
            if current.node == n-1 {
                dist[n-1] = current.distance
                break
            }
            if current.node < n-1 {
                neighbor := current.node + 1
                newDist := current.distance + 1
                if newDist < dist[neighbor] {
                    dist[neighbor] = newDist
                    pq.items = append(pq.items, PriorityQueueItem{neighbor, newDist})
                    pq.bubbleUp(len(pq.items) - 1)
                }
            }
            for _, neighbor := range extraEdges[current.node] {
                newDist := current.distance + 1
                if newDist < dist[neighbor] {
                    dist[neighbor] = newDist
                    pq.items = append(pq.items, PriorityQueueItem{neighbor, newDist})
                    pq.bubbleUp(len(pq.items) - 1)
                }
            }
        }
        res[i] = dist[n-1]
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5, queries = [[2,4],[0,2],[0,4]]
    // Output: [3,2,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image8.jpg" />
    // After the addition of the road from 2 to 4, the length of the shortest path from 0 to 4 is 3.
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image9.jpg" />
    // After the addition of the road from 0 to 2, the length of the shortest path from 0 to 4 is 2.
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image10.jpg" />
    // After the addition of the road from 0 to 4, the length of the shortest path from 0 to 4 is 1.
    fmt.Println(shortestDistanceAfterQueries(5,[][]int{{2,4},{0,2},{0,4}})) // [3,2,1]
    // Example 2:
    // Input: n = 4, queries = [[0,3],[0,2]]
    // Output: [1,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image11.jpg" />
    // After the addition of the road from 0 to 3, the length of the shortest path from 0 to 3 is 1.
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image12.jpg" />
    // After the addition of the road from 0 to 2, the length of the shortest path remains 1.
    fmt.Println(shortestDistanceAfterQueries(4,[][]int{{0,3},{0,2}})) // [1,1]

    fmt.Println(shortestDistanceAfterQueries1(5,[][]int{{2,4},{0,2},{0,4}})) // [3,2,1]
    fmt.Println(shortestDistanceAfterQueries1(4,[][]int{{0,3},{0,2}})) // [1,1]

    fmt.Println(shortestDistanceAfterQueries2(5,[][]int{{2,4},{0,2},{0,4}})) // [3,2,1]
    fmt.Println(shortestDistanceAfterQueries2(4,[][]int{{0,3},{0,2}})) // [1,1]
}