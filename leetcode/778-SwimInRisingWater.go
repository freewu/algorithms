package main

// 778. Swim in Rising Water
// You are given an n x n integer matrix grid where each value grid[i][j] represents the elevation at that point (i, j).

// The rain starts to fall. At time t, the depth of the water everywhere is t. 
// You can swim from a square to another 4-directionally adjacent square if and only if the elevation of both squares individually are at most t. 
// You can swim infinite distances in zero time. Of course, you must stay within the boundaries of the grid during your swim.

// Return the least time until you can reach the bottom right square (n - 1, n - 1) if you start at the top left square (0, 0).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/swim1-grid.jpg" />
// Input: grid = [[0,2],[1,3]]
// Output: 3
// Explanation:
// At time 0, you are in grid location (0, 0).
// You cannot go anywhere else because 4-directionally adjacent neighbors have a higher elevation than t = 0.
// You cannot reach point (1, 1) until time 3.
// When the depth of water is 3, we can swim anywhere inside the grid.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/swim2-grid-1.jpg" />
// Input: grid = [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
// Output: 16
// Explanation: The final route is shown.
// We need to wait until time 16 so that (0, 0) and (4, 4) are connected.

// Constraints:
//     n == grid.length
//     n == grid[i].length
//     1 <= n <= 50
//     0 <= grid[i][j] < n2
//     Each value grid[i][j] is unique.

import "fmt"
import "container/heap"

// Heap
type MyHeap [][3]int
func (h MyHeap) Len() int           { return len(h) }
func (h MyHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MyHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.([3]int))
}

func (h *MyHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func swimInWater(grid [][]int) int {
    r, c := len(grid), len(grid[0])
    h := &MyHeap{[3]int{grid[0][0], 0, 0}}
    directions := [4][2]int{ {1, 0}, {-1, 0}, {0, 1}, {0, -1}, }
    visited := map[[2]int]bool{}
    res := 2600

    max := func (x, y int) int { if x > y { return x; }; return y; }
    // Dijkstra's algorithm
    // O(R*C*log(R*C))
    for h.Len() > 0 {
        topHeap := heap.Pop(h).([3]int)
        height, row, col := topHeap[0], topHeap[1], topHeap[2]
        if row == r - 1 && col == c - 1 {
            return height
        }
        for _, dir := range directions {
            neiRow, neiCol := row + dir[0], col + dir[1]
            if neiRow < 0 || neiRow == r || neiCol < 0 || neiCol == c || visited[[2]int{neiRow, neiCol}] {
                continue
            }
            neiHeight := max(grid[neiRow][neiCol], height)
            visited[[2]int{neiRow, neiCol}] = true
            heap.Push(h, [3]int{neiHeight, neiRow, neiCol})
        }
    }
    return res
}

func swimInWater1(grid [][]int) int {
    row, col := len(grid), len(grid[0])
    visit := make([][]bool, row)
    for i := 0; i < row; i++ {
        visit[i] = make([]bool, col)
    }
    directions := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
    var dfs func(t int, i, j int) bool
    dfs = func(t int, i, j int) bool {
        visit[i][j] = true
        if grid[i][j] > t {
            return false
        }
        if i == row -1 && j  == col -1 {
            return true
        }        
        for _, direction := range directions {
            i0, j0 := i + direction[0], j+ direction[1]
            if  i0 >= 0 && i0 < row && j0 >= 0 && j0 < col && !visit[i0][j0] && grid[i0][j0] <= t { // 边界检测
                if dfs(t, i0, j0) {
                    return true
                }
            }
        }
        return false
    }
    mn, mx := grid[row-1][col-1], grid[row-1][col-1]
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if grid[i][j] > mx {
                mx = grid[i][j]
            }
        }
    }
    l, r := mn, mx
    for l <= r {
        mid := (l + r ) / 2
        for i := 0; i < row; i++ {
            for j := 0; j < row; j++ { 
                visit[i][j] =  false
            }            
        }
        if dfs(mid, 0, 0) {
            r = mid -1 
        } else {
            l = mid + 1
        }
    }
    return l 
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/swim1-grid.jpg" />
    // Input: grid = [[0,2],[1,3]]
    // Output: 3
    // Explanation:
    // At time 0, you are in grid location (0, 0).
    // You cannot go anywhere else because 4-directionally adjacent neighbors have a higher elevation than t = 0.
    // You cannot reach point (1, 1) until time 3.
    // When the depth of water is 3, we can swim anywhere inside the grid.
    fmt.Println(swimInWater([][]int{{0,2},{1,3}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/swim2-grid-1.jpg" />
    // Input: grid = [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
    // Output: 16
    // Explanation: The final route is shown.
    // We need to wait until time 16 so that (0, 0) and (4, 4) are connected.
    fmt.Println(swimInWater([][]int{{0,1,2,3,4},{24,23,22,21,5},{12,13,14,15,16},{11,17,18,19,20},{10,9,8,7,6}})) // 16
    
    fmt.Println(swimInWater1([][]int{{0,2},{1,3}})) // 3
    fmt.Println(swimInWater1([][]int{{0,1,2,3,4},{24,23,22,21,5},{12,13,14,15,16},{11,17,18,19,20},{10,9,8,7,6}})) // 16
}