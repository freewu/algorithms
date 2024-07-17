package main

// 1584. Min Cost to Connect All Points
// You are given an array points representing integer coordinates of some points on a 2D-plane, where points[i] = [xi, yi].

// The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|, 
// where |val| denotes the absolute value of val.

// Return the minimum cost to make all points connected. 
// All points are connected if there is exactly one simple path between any two points.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/26/d.png" />
// Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
// Output: 20
// Explanation: 
// <img src="https://assets.leetcode.com/uploads/2020/08/26/c.png" />
// We can connect the points as shown above to get the minimum cost of 20.
// Notice that there is a unique path between every pair of points.

// Example 2:
// Input: points = [[3,12],[-2,5],[-4,1]]
// Output: 18

// Constraints:
//     1 <= points.length <= 1000
//     -10^6 <= xi, yi <= 10^6
//     All pairs (xi, yi) are distinct.

import "fmt"
import "container/heap"

type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func minCostConnectPoints(points [][]int) int {
    n := len(points)
    adj := make([][][]int, n)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n; i++ {
        x1, y1 := points[i][0], points[i][1]
        for j := i + 1; j < n; j++ {
            x2, y2 := points[j][0], points[j][1]
            dist := abs(x1 - x2) + abs(y1 - y2)
            adj[i] = append(adj[i], []int{ dist, j })
            adj[j] = append(adj[j], []int{ dist, i })
        }
    }
    res, visited := 0, make(map[int]bool)
    h := &IntHeap{ [2]int{ 0, 0 } }
    for len(visited) < n {
        tmp := heap.Pop(h).([2]int)
        cost, idx := tmp[0], tmp[1]
        if visited[idx] { continue }
        res += cost
        visited[idx] = true
        for _, v := range adj[idx] {
            nCost, nIdx := v[0], v[1]
            if !visited[nIdx] { 
                heap.Push(h, [2]int{ nCost, nIdx }) 
            }
        }
    }
    return res
}

func minCostConnectPoints1(points [][]int) int {
    res, inf, minDistance := 0, 1 << 32  - 1, make([]int, len(points))
    for i := range minDistance {
        minDistance[i] = inf
    }
    minDistanceTree := []int{0}
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    calculateDistance := func (x1, y1, x2, y2 int) int {
        return abs(x1 - x2) + abs(y1 - y2)
    }
    visited := make([]bool, len(points))
    visited[0] = true
    for len(minDistanceTree) < len(points) {
        for i := 0; i < len(points); i++ {
            if visited[i] {
                continue
            }
            currentIndex := minDistanceTree[len(minDistanceTree)-1]
            if i == currentIndex { continue }
            minDistance[i] = min(minDistance[i], calculateDistance(points[i][0], points[i][1], points[currentIndex][0], points[currentIndex][1]))
        }
        currentMinIndex, currentMinDistance := 0, inf
        for j := 0; j < len(points); j++ {
            if visited[j] { continue }
            if j == minDistanceTree[len(minDistanceTree)-1] { continue }
            if currentMinDistance > minDistance[j] {
                currentMinDistance = minDistance[j]
                currentMinIndex = j
            }
        }
        visited[currentMinIndex] = true
        minDistanceTree = append(minDistanceTree, currentMinIndex)
    }
    for i := 1; i < len(minDistance); i++ {
        res += minDistance[i]
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/08/26/d.png" />
    // Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
    // Output: 20
    // Explanation: 
    // <img src="https://assets.leetcode.com/uploads/2020/08/26/c.png" />
    // We can connect the points as shown above to get the minimum cost of 20.
    // Notice that there is a unique path between every pair of points.
    fmt.Println(minCostConnectPoints([][]int{{0,0},{2,2},{3,10},{5,2},{7,0}})) // 20
    // Example 2:
    // Input: points = [[3,12],[-2,5],[-4,1]]
    // Output: 18
    fmt.Println(minCostConnectPoints([][]int{{3,12},{-2,5},{-4,1}})) // 18

    fmt.Println(minCostConnectPoints1([][]int{{0,0},{2,2},{3,10},{5,2},{7,0}})) // 20
    fmt.Println(minCostConnectPoints1([][]int{{3,12},{-2,5},{-4,1}})) // 18
}