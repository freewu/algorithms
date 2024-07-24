package main

// 1001. Grid Illumination
// There is a 2D grid of size n x n where each cell of this grid has a lamp that is initially turned off.

// You are given a 2D array of lamp positions lamps, where lamps[i] = [rowi, coli] indicates 
// that the lamp at grid[rowi][coli] is turned on. Even if the same lamp is listed more than once, it is turned on.

// When a lamp is turned on, it illuminates its cell and all other cells in the same row, column, or diagonal.

// You are also given another 2D array queries, where queries[j] = [rowj, colj]. 
// For the jth query, determine whether grid[rowj][colj] is illuminated or not. 
// After answering the jth query, turn off the lamp at grid[rowj][colj] and its 8 adjacent lamps if they exist. 
// A lamp is adjacent if its cell shares either a side or corner with grid[rowj][colj].

// Return an array of integers ans, 
// where ans[j] should be 1 if the cell in the jth query was illuminated, or 0 if the lamp was not.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/19/illu_1.jpg" />
// Input: n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
// Output: [1,0]
// Explanation: We have the initial grid with all lamps turned off. In the above picture we see the grid after turning on the lamp at grid[0][0] then turning on the lamp at grid[4][4].
// The 0th query asks if the lamp at grid[1][1] is illuminated or not (the blue square). 
// It is illuminated, so set ans[0] = 1. Then, we turn off all lamps in the red square.
// <img src="https://assets.leetcode.com/uploads/2020/08/19/illu_step1.jpg" />
// The 1st query asks if the lamp at grid[1][0] is illuminated or not (the blue square). 
// It is not illuminated, so set ans[1] = 0. Then, we turn off all lamps in the red rectangle.
// <img src="https://assets.leetcode.com/uploads/2020/08/19/illu_step2.jpg" />

// Example 2:
// Input: n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,1]]
// Output: [1,1]

// Example 3:
// Input: n = 5, lamps = [[0,0],[0,4]], queries = [[0,4],[0,1],[1,4]]
// Output: [1,1,0]

// Constraints:
//     1 <= n <= 10^9
//     0 <= lamps.length <= 20000
//     0 <= queries.length <= 20000
//     lamps[i].length == 2
//     0 <= rowi, coli < n
//     queries[j].length == 2
//     0 <= rowj, colj < n

import "fmt"

func gridIllumination(n int, lamps, queries [][]int) []int {
    type pair struct{ x, y int }
    points := map[pair]bool{}
    row, col, diagonal, antiDiagonal := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
    for _, lamp := range lamps {
        r, c := lamp[0], lamp[1]
        p := pair{r, c}
        if points[p] { // 重复的就不需要继续徐打开
            continue
        }
        points[p] = true
        row[r]++            // 照亮行
        col[c]++            // 照亮列
        diagonal[r-c]++     // 照亮左上角
        antiDiagonal[r+c]++ // 照亮右下对角
    }
    res := make([]int, len(queries))
    for i, query := range queries {
        r, c := query[0], query[1]
        if row[r] > 0 || col[c] > 0 || diagonal[r-c] > 0 || antiDiagonal[r+c] > 0 { // 如果是被照亮的则结果为1
            res[i] = 1
        }
        for x := r - 1; x <= r+1; x++ { // 关闭周围的8个
            for y := c - 1; y <= c+1; y++ {
                if x < 0 || y < 0 || x >= n || y >= n || !points[pair{x, y}] {
                    continue
                }
                delete(points, pair{x, y})
                row[x]--
                col[y]--
                diagonal[x-y]--
                antiDiagonal[x+y]--
            }
        }
    }
    return res
}

func gridIllumination1(n int, lamps [][]int, queries [][]int) []int {
    r, c:= make(map[int]int), make(map[int]int) // 使用 map，降低空间复杂度
    pie, na := make(map[int]int), make(map[int]int) // 使用 map，防止 int 溢出：1 <= n <= 1e9
    memo := make(map[[2]int]struct{}) // 去重
    res := make([]int, len(queries))
    for _, l := range lamps {
        x, y := l[0], l[1]
        if _, ok := memo[[2]int{x, y}]; !ok {
            memo[[2]int{x, y}] = struct{}{}
            r[x]++
            c[y]++
            pie[x-n+y]++ // \ 的算法 x-n+y：防止溢出（或直接使用 x+y）
            na[x-y]++    // / 的算法 x-y
        }
    }
    for idx, q := range queries {
        x, y := q[0], q[1]
        if r[x] > 0 || c[y] > 0 || pie[x-n+y] > 0 || na[x-y] > 0 {
            res[idx] = 1
        }
        for i := x - 1; i <= x+1; i++ {
            for j := y - 1; j <= y+1; j++ {
                if 0 <= i && i < n && 0 <= j && j < n {
                    if _, ok := memo[[2]int{i, j}]; ok {
                        r[i]--
                        c[j]--
                        pie[i-n+j]--
                        na[i-j]--
                        delete(memo, [2]int{i, j})
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/08/19/illu_1.jpg" />
    // Input: n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
    // Output: [1,0]
    // Explanation: We have the initial grid with all lamps turned off. In the above picture we see the grid after turning on the lamp at grid[0][0] then turning on the lamp at grid[4][4].
    // The 0th query asks if the lamp at grid[1][1] is illuminated or not (the blue square). 
    // It is illuminated, so set ans[0] = 1. Then, we turn off all lamps in the red square.
    // <img src="https://assets.leetcode.com/uploads/2020/08/19/illu_step1.jpg" />
    // The 1st query asks if the lamp at grid[1][0] is illuminated or not (the blue square). 
    // It is not illuminated, so set ans[1] = 0. Then, we turn off all lamps in the red rectangle.
    // <img src="https://assets.leetcode.com/uploads/2020/08/19/illu_step2.jpg" />
    fmt.Println(gridIllumination(5,[][]int{{0,0},{4,4}},[][]int{{1,1},{1,0}})) // [1,0]
    // Example 2:
    // Input: n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,1]]
    // Output: [1,1]
    fmt.Println(gridIllumination(5,[][]int{{0,0},{4,4}},[][]int{{1,1},{1,1}})) // [1,1]
    // Example 3:
    // Input: n = 5, lamps = [[0,0],[0,4]], queries = [[0,4],[0,1],[1,4]]
    // Output: [1,1,0]
    fmt.Println(gridIllumination(5,[][]int{{0,0},{4,4}},[][]int{{0,4},{0,1},{1,4}})) // [1,1,0]

    fmt.Println(gridIllumination1(5,[][]int{{0,0},{4,4}},[][]int{{1,1},{1,0}})) // [1,0]
    fmt.Println(gridIllumination1(5,[][]int{{0,0},{4,4}},[][]int{{1,1},{1,1}})) // [1,1]
    fmt.Println(gridIllumination1(5,[][]int{{0,0},{4,4}},[][]int{{0,4},{0,1},{1,4}})) // [1,1,0]
}