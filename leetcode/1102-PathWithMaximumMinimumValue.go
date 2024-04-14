package main

// 1102. Path With Maximum Minimum Value
// Given an m x n integer matrix grid, 
// return the maximum score of a path starting at (0, 0) and ending at (m - 1, n - 1) moving in the 4 cardinal directions.
//     The score of a path is the minimum value in that path.

// For example, the score of the path 8 → 4 → 5 → 9 is 4.
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/05/maxgrid1.jpg" />
// Input: grid = [[5,4,5],[1,2,6],[7,4,6]]
// Output: 4
// Explanation: The path with the maximum score is highlighted in yellow. 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/08/05/maxgrid2.jpg" />
// Input: grid = [[2,2,1,2,2,2],[1,2,2,2,1,2]]
// Output: 2

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/08/05/maxgrid3.jpg" />
// Input: grid = [[3,4,6,3,4],[0,2,1,1,7],[8,8,3,2,7],[3,2,4,9,8],[4,1,2,0,0],[4,6,5,4,3]]
// Output: 3

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 100
//     0 <= grid[i][j] <= 10^9

import "fmt"
import "container/heap"

type Cell struct {
    id    int
    score int
}

type CellPQ []*Cell
func (p CellPQ) Len() int { return len(p) }
func (p CellPQ) Less(i, j int) bool {
    return p[i].score > p[j].score
}
func (p CellPQ) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p *CellPQ) Push(x interface{}) {
    (*p) = append(*p, x.(*Cell))
}
func (p *CellPQ) Pop() interface{} {
    n := p.Len()
    item := (*p)[n-1]
    *p = (*p)[:n-1]
    return item
}

func maximumMinimumPath(grid [][]int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    q := &CellPQ{} 
    heap.Init(q)
    tr, tc := len(grid), len(grid[0])
    visited := make(map[int]bool)
    heap.Push(q, &Cell{0, grid[0][0]})
    visited[0] = true
    res := min(grid[0][0], grid[tr - 1][tc - 1])
    for q.Len() > 0 {
        cc := heap.Pop(q).(*Cell)
        res = min(res, cc.score)
        crow, ccol := cc.id/tc, cc.id%tc
        // 4 directions explore, adding not visited nodes
        if crow > 0 {
            nid := (crow-1)*tc + ccol
            if _, ok := visited[nid]; !ok {
                heap.Push(q, &Cell{nid, grid[crow-1][ccol]})
                visited[nid] = true
            }
        }
        if ccol > 0 {
            nid := crow*tc + ccol - 1
            if _, ok := visited[nid]; !ok {
                heap.Push(q, &Cell{nid, grid[crow][ccol-1]})
                visited[nid] = true
            }
        }
        if crow < len(grid)-1 {
            nid := (crow+1)*tc + ccol
            if _, ok := visited[nid]; !ok {
                if nid == tr*tc-1 { // 到终点了
                    break
                }
                heap.Push(q, &Cell{nid, grid[crow+1][ccol]})
                visited[nid] = true
            }
        }
        if ccol < tc-1 {
            nid := crow*tc + ccol + 1
            if _, ok := visited[nid]; !ok {
                if nid == tr*tc-1 { // 到终点了
                    break
                }
                heap.Push(q, &Cell{nid, grid[crow][ccol+1]})
                visited[nid] = true
            }
        }
    }
    return res
}

func maximumMinimumPath1(grid [][]int) int {
    var visited [101][101]bool
    dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} // 4个方向
    row, col := len(grid), len(grid[0])
    min := func (x, y int) int { if x < y { return x; }; return y; }
    left, right := 0, min(grid[0][0],grid[row - 1][col - 1])

    var pathExists func(val, curRow, curCol int) bool
    pathExists = func(val, curRow, curCol int) bool {
        if curRow == row - 1 && curCol == col - 1 {
            return true
        }
        visited[curRow][curCol] = true
        for _, dir := range dirs {
            newRow, newCol := curRow + dir[0], curCol + dir[1]
            if newRow >= 0 && newRow < row && newCol >= 0 && newCol < col &&
                !visited[newRow][newCol] && grid[newRow][newCol] >= val {
                if pathExists(val, newRow, newCol) {
                    return true
                }
            }
        }
        return false
    }
    for left < right {
        middle := (left + right + 1) >> 1
        for i := 0; i < row; i++ {
            for j := 0; j < col; j++ {
                visited[i][j] = false
            }
        }
        if pathExists(middle, 0, 0) {
            left = middle
        } else {
            right = middle - 1
        }
    }
    return left
}

func main() {
    fmt.Println(maximumMinimumPath([][]int{{5,4,5},{1,2,6},{7,4,6}})) // 4
    fmt.Println(maximumMinimumPath([][]int{{2,2,1,2,2,2},{1,2,2,2,1,2}})) // 2
    fmt.Println(maximumMinimumPath([][]int{
        {3,4,6,3,4},
        {0,2,1,1,7},
        {8,8,3,2,7},
        {3,2,4,9,8},
        {4,1,2,0,0},
        {4,6,5,4,3},
    })) // 3

    fmt.Println(maximumMinimumPath1([][]int{{5,4,5},{1,2,6},{7,4,6}})) // 4
    fmt.Println(maximumMinimumPath1([][]int{{2,2,1,2,2,2},{1,2,2,2,1,2}})) // 2
    fmt.Println(maximumMinimumPath1([][]int{
        {3,4,6,3,4},
        {0,2,1,1,7},
        {8,8,3,2,7},
        {3,2,4,9,8},
        {4,1,2,0,0},
        {4,6,5,4,3},
    })) // 3
}