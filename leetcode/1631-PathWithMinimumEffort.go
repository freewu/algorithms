package main

// 1631. Path With Minimum Effort
// You are a hiker preparing for an upcoming hike. 
// You are given heights, a 2D array of size rows x columns, where heights[row][col] represents the height of cell (row, col). 
// You are situated in the top-left cell, (0, 0), and you hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed). 
// You can move up, down, left, or right, and you wish to find a route that requires the minimum effort.

// A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.
// Return the minimum effort required to travel from the top-left cell to the bottom-right cell.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/04/ex1.png" />
// Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
// Output: 2
// Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
// This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/04/ex2.png" />
// Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
// Output: 1
// Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in consecutive cells, which is better than route [1,3,5,3,5].

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/10/04/ex3.png" />
// Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
// Output: 0
// Explanation: This route does not require any effort.
 
// Constraints:
//     rows == heights.length
//     columns == heights[i].length
//     1 <= rows, columns <= 100
//     1 <= heights[i][j] <= 10^6

import "fmt"
import "container/heap"
import "math"

type MinHeap [][3]int

func (h MinHeap) Len() int  {
    return len(h)
}

func (h MinHeap) Less(i int, j int) bool {
    return h[i][2] < h[j][2]
}

func (h MinHeap) Swap(i int, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(a interface{}) {
    *h = append(*h, a.([3]int))
}

func (h *MinHeap) Pop() interface{} {
    l := len(*h)
    res := (*h)[l - 1]
    *h = (*h)[:l - 1]
    return res
}

// dijkstra + minheap
func minimumEffortPath(heights [][]int) int {
    m, n := len(heights), len(heights[0])
    offset := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
    set := make(map[[2]int]bool)

    abs := func (x int) int { if x >= 0 { return x; }; return -1 * x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }

    h := &MinHeap{}
    heap.Push(h, [3]int{0, 0, 0})
    res := 0
    for h.Len() > 0{
        p := heap.Pop(h).([3]int)
        if set[[2]int{p[0], p[1]}] {
            continue
        } else {
            set[[2]int{p[0], p[1]}] = true
        }
        res = max(res, p[2])
        if p[0] == m - 1 && p[1] == n - 1 {
            break
        }
        for _, dir := range offset {
            x := p[0] + dir[0]
            y := p[1] + dir[1]
            if x < 0 || y < 0 || x >= m || y >= n {
                continue
            }
            heap.Push(h, [3]int{x, y, abs(heights[x][y] - heights[p[0]][p[1]])})
        }
    }
    return res
}

type State struct {
    x, y int
    costFromStart int
}

type PriorityQueue []*State

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].costFromStart < pq[j].costFromStart
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[:n-1]
    return item
}

func minimumEffortPath1(heights [][]int) int {
    dijsktra := func(heights [][]int) int {
        m, n := len(heights), len(heights[0])
        costTo := make([][]int, m)
        for i := 0; i < m; i++ {
            costTo[i]= make([]int, n)
            for j:=0; j<n; j++ {
                costTo[i][j] = math.MaxInt64
            }
        }
        costTo[0][0] = 0
        var pq PriorityQueue
        heap.Init(&pq)
        heap.Push(&pq, &State{0, 0, 0})
        dirs := [][2]int{{1, 0}, {0,1}, {-1,0}, {0, -1}}
        abs := func (x int) int { if x >= 0 { return x; }; return -1 * x; }
        max := func (x, y int) int { if x > y { return x; }; return y; }
        for len(pq) > 0 {
            curState := heap.Pop(&pq).(*State)
            curX, curY := curState.x, curState.y
            if curX == m-1 && curY == n-1 {
                return costTo[curX][curY]
            }
            for _, dir := range dirs {
                x0, y0 := curX + dir[0], curY + dir[1]
                if x0 < 0 || y0 < 0 || x0 >=m || y0 >= n{
                    continue
                }
                nextCostFromStart := max(costTo[curX][curY], abs(heights[x0][y0] - heights[curX][curY]))
                if nextCostFromStart < costTo[x0][y0] {
                    costTo[x0][y0] = nextCostFromStart
                    heap.Push(&pq, &State{x0, y0, nextCostFromStart})
                }
            }
        }
        return 0
    }
    return dijsktra(heights)
} 

func main() {
    // Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
    // This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.
    fmt.Println(minimumEffortPath([][]int{{1,2,2},{3,8,2},{5,3,5}})) // 2
    // Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in consecutive cells, which is better than route [1,3,5,3,5].
    fmt.Println(minimumEffortPath([][]int{{1,2,3},{3,8,4},{5,3,5}})) // 1
    // Explanation: This route does not require any effort.
    fmt.Println(minimumEffortPath([][]int{{1,2,1,1,1},{1,2,1,2,1},{1,2,1,2,1},{1,2,1,2,1},{1,1,1,2,1}})) // 0

    fmt.Println(minimumEffortPath1([][]int{{1,2,2},{3,8,2},{5,3,5}})) // 2
    fmt.Println(minimumEffortPath1([][]int{{1,2,3},{3,8,4},{5,3,5}})) // 1
    fmt.Println(minimumEffortPath1([][]int{{1,2,1,1,1},{1,2,1,2,1},{1,2,1,2,1},{1,2,1,2,1},{1,1,1,2,1}})) // 0
}