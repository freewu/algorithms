package main

// 面试题 16.19. Pond Sizes LCCI
// You have an integer matrix representing a plot of land, where the value at that loca­tion represents the height above sea level. 
// A value of zero indicates water. 
// A pond is a region of water connected vertically, horizontally, or diagonally. 
// The size of the pond is the total number of connected water cells. 
// Write a method to compute the sizes of all ponds in the matrix, the return values need to be sorted in ascending order.

// Example:
// Input: 
// [
//   [0,2,1,0],
//   [0,1,0,1],
//   [1,1,0,1],
//   [0,1,0,1]
// ]
// Output: [1,2,4]

// Note:
//     0 < len(land) <= 1000
//     0 < len(land[i]) <= 1000

import "fmt"
import "sort"

// dfs
func pondSizes(land [][]int) []int {
    m, n := len(land), len(land[0])
    var dfs func(int, int) int
    dfs = func(x, y int) int {
        if x < 0 || x >= m || y < 0 || y >= n || land[x][y] != 0 { return 0 }
        land[x][y] = -1
        res := 1
        for dx := -1; dx <= 1; dx++ {
            for dy := -1; dy <= 1; dy++ {
                if dx == 0 && dy == 0 { continue }
                res += dfs(x + dx, y + dy)
            }
        }
        return res
    }
    res := []int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if land[i][j] == 0 {
                res = append(res, dfs(i, j))
            }
        }
    }
    sort.Ints(res)
    return res
}

// bfs
func pondSizes1(land [][]int) []int {
    m, n := len(land), len(land[0])
    bfs := func(x, y int) int {
        queue, res := [][]int{}, 0
        queue, land[x][y] = append(queue, []int{x, y}), -1
        for len(queue) > 0 {
            x, y, queue = queue[0][0], queue[0][1], queue[1:]
            res++
            for dx := -1; dx <= 1; dx++ {
                for dy := -1; dy <= 1; dy++ {
                    if dx == 0 && dy == 0 { continue }
                    if x + dx < 0 || x + dx >= m || y + dy < 0 || y + dy >= n || land[x + dx][y + dy] != 0 { continue }
                    land[x + dx][y + dy] = -1
                    queue = append(queue, []int{ x + dx, y + dy })
                }
            }
        }
        return res
    }
    res := []int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if land[i][j] == 0 {
                res = append(res, bfs(i, j))
            }
        }
    }
    sort.Ints(res)
    return res
}

func pondSizes2(land [][]int) []int {
    // 最大岛屿面积变种
    row, col := len(land), len(land[0])
    var dfs func(x, y int) int
    dfs =  func (x, y int) int {
        if (x < 0 || x >= row || y < 0 || y >= col || land[x][y] != 0 ) { return 0 }
        land[x][y] = -1
        res := 1
        for i := x -1; i <= x + 1; i++ {
            for j := y - 1; j <= y + 1; j++ {
                res += dfs(i, j)
            }
        }
        return res
    }
    res := []int{}
    for x := 0; x < row; x++ {
        for y := 0; y < col; y++ {
            if (land[x][y] == 0) {
                res = append(res, dfs(x, y))
            }
        }
    }
    sort.Ints(res)
    return res
}

func main() {
    // Example:
    // Input: 
    // [
    //   [0,2,1,0],
    //   [0,1,0,1],
    //   [1,1,0,1],
    //   [0,1,0,1]
    // ]
    // Output: [1,2,4]
    grid1 := [][]int{
        {0,2,1,0},
        {0,1,0,1},
        {1,1,0,1},
        {0,1,0,1},
    }
    fmt.Println(pondSizes(grid1)) // [1,2,4]

    grid11 := [][]int{
        {0,2,1,0},
        {0,1,0,1},
        {1,1,0,1},
        {0,1,0,1},
    }
    fmt.Println(pondSizes1(grid11)) // [1,2,4]

    grid21 := [][]int{
        {0,2,1,0},
        {0,1,0,1},
        {1,1,0,1},
        {0,1,0,1},
    }
    fmt.Println(pondSizes1(grid21)) // [1,2,4]
}