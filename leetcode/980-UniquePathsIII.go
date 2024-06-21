package main

// 980. Unique Paths III
// You are given an m x n integer array grid where grid[i][j] could be:
//     1 representing the starting square. There is exactly one starting square.
//     2 representing the ending square. There is exactly one ending square.
//     0 representing empty squares we can walk over.
//     -1 representing obstacles that we cannot walk over.

// Return the number of 4-directional walks from the starting square to the ending square, 
// that walk over every non-obstacle square exactly once.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/02/lc-unique1.jpg" />
// Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
// Output: 2
// Explanation: We have the following two paths: 
// 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
// 2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/08/02/lc-unique2.jpg" />
// Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
// Output: 4
// Explanation: We have the following four paths: 
// 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
// 2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
// 3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
// 4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/08/02/lc-unique3-.jpg" />
// Input: grid = [[0,1],[2,0]]
// Output: 0
// Explanation: There is no path that walks over every empty square exactly once.
// Note that the starting and ending square can be anywhere in the grid.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 20
//     1 <= m * n <= 20
//     -1 <= grid[i][j] <= 2
//     There is exactly one starting cell and one ending cell.

import "fmt"

// backtracking
func uniquePathsIII(grid [][]int) int {
    dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
    res, cellCnt := 0, 0
    dst, src, seen := [2]int{}, [2]int{}, map[[2]int]bool{}
    for i := range grid {
        for j := range grid[0] {
             if grid[i][j] == 0 { cellCnt++ }
             if grid[i][j] == 2 { dst = [2]int{i,j} }
             if grid[i][j] == 1 { src = [2]int{i,j} }
        }
    }
    cellCnt++ // include starting cell in # of cells that should be seen
    var dfs func(src, dst [2]int, g [][]int, uPath *int, cellCnt int, seen map[[2]int]bool)
    dfs = func (src, dst [2]int, g [][]int, uPath *int, cellCnt int, seen map[[2]int]bool) {
        if len(seen) == cellCnt && src == dst {
            *uPath++
            return
        }
        m, n := len(g), len(g[0])
        seen[src] = true // set current cell to true so its not revisited
        for _, d := range dirs {
            r, c:= src[0]+d[0], src[1] + d[1] // next cells are out of bound
            next := [2]int{r,c}
            if r < 0 || r > m-1 || c < 0 || c > n-1 { continue } // border check
            if g[r][c] == -1 || seen[next] { continue } // next cell is a obstacle, or its been traveled before
            dfs(next, dst, g, uPath, cellCnt, seen)  // search with next cell
        }    
        delete(seen, src) // remove cell from seen 
    }
    dfs(src, dst, grid, &res, cellCnt, seen)
    return res
}

func uniquePathsIII1(grid [][]int) int {
    gGrid := grid
    gNext := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    gVisit := make([][]bool, len(grid))
    startI, startJ, visitNum, maxNum := 0, 0, 0, 0
    for i := 0; i < len(gVisit); i++ {
        gVisit[i] = make([]bool, len(grid[i]))
    }
    var getRes func(i, j int) int
    getRes = func(i, j int) int {
        if i < 0 || i >= len(gGrid) || j < 0 || j >= len(gGrid[i]) {
            return 0
        }
        if gGrid[i][j] == -1 || gVisit[i][j] {
            return 0
        }
        if gGrid[i][j] == 2 {
            if visitNum == maxNum {
                return 1
            } else {
                return 0
            }
        }
        gVisit[i][j] = true
        visitNum++
        res := 0
        for _, next := range gNext {
            res += getRes(i+next[0], j+next[1])
        }
        visitNum--
        gVisit[i][j] = false
        return res
    }
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                startI, startJ = i, j
                maxNum++
            }
            if gGrid[i][j] == 0 {
                maxNum++
            }
        }
    }
    return getRes(startI, startJ)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/02/lc-unique1.jpg" />
    // Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
    // Output: 2
    // Explanation: We have the following two paths: 
    // 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
    // 2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
    grid1 := [][]int{
        {1,0,0,0},
        {0,0,0,0},
        {0,0,2,-1},
    }
    fmt.Println(uniquePathsIII(grid1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/08/02/lc-unique2.jpg" />
    // Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
    // Output: 4
    // Explanation: We have the following four paths: 
    // 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
    // 2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
    // 3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
    // 4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
    grid2 := [][]int{
        {1,0,0,0},
        {0,0,0,0},
        {00,0,0,2},
    }
    fmt.Println(uniquePathsIII(grid2)) // 4
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/08/02/lc-unique3-.jpg" />
    // Input: grid = [[0,1],[2,0]]
    // Output: 0
    // Explanation: There is no path that walks over every empty square exactly once.
    // Note that the starting and ending square can be anywhere in the grid.
    grid3 := [][]int{
        {0,1},
        {2,0},
    }
    fmt.Println(uniquePathsIII(grid3)) // 0

    fmt.Println(uniquePathsIII1(grid1)) // 2
    fmt.Println(uniquePathsIII1(grid2)) // 4
    fmt.Println(uniquePathsIII1(grid3)) // 0
}