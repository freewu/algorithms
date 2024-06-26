package main

// 417. Pacific Atlantic Water Flow
// There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. 
// The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

// The island is partitioned into a grid of square cells. 
// You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

// The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west 
// if the neighboring cell's height is less than or equal to the current cell's height. 
// Water can flow from any cell adjacent to an ocean into the ocean.

// Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes 
// that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/08/waterflow-grid.jpg" />
// Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
// Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
// Explanation: The following cells can flow to the Pacific and Atlantic oceans, as shown below:
// [0,4]: [0,4] -> Pacific Ocean 
//        [0,4] -> Atlantic Ocean
// [1,3]: [1,3] -> [0,3] -> Pacific Ocean 
//        [1,3] -> [1,4] -> Atlantic Ocean
// [1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean 
//        [1,4] -> Atlantic Ocean
// [2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean 
//        [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
// [3,0]: [3,0] -> Pacific Ocean 
//        [3,0] -> [4,0] -> Atlantic Ocean
// [3,1]: [3,1] -> [3,0] -> Pacific Ocean 
//        [3,1] -> [4,1] -> Atlantic Ocean
// [4,0]: [4,0] -> Pacific Ocean 
//        [4,0] -> Atlantic Ocean
// Note that there are other possible paths for these cells to flow to the Pacific and Atlantic oceans.

// Example 2:
// Input: heights = [[1]]
// Output: [[0,0]]
// Explanation: The water can flow from the only cell to the Pacific and Atlantic oceans.
 
// Constraints:
//     m == heights.length
//     n == heights[r].length
//     1 <= m, n <= 200
//     0 <= heights[r][c] <= 10^5

import "fmt"

func pacificAtlantic(matrix [][]int) [][]int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return nil
    }
    row, col, res, inf := len(matrix), len(matrix[0]), make([][]int, 0), -1 << 32
    pacific, atlantic := make([][]bool, row), make([][]bool, row)

    var dfs func(matrix [][]int, row, col int, visited *[][]bool, height int) 
    dfs = func(matrix [][]int, row, col int, visited *[][]bool, height int) {
        if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
            return
        }
        if (*visited)[row][col] || matrix[row][col] < height {
            return
        }
        (*visited)[row][col] = true
        dfs(matrix, row+1, col, visited, matrix[row][col])
        dfs(matrix, row-1, col, visited, matrix[row][col])
        dfs(matrix, row, col+1, visited, matrix[row][col])
        dfs(matrix, row, col-1, visited, matrix[row][col])
    }
    for i := 0; i < row; i++ {
        pacific[i] = make([]bool, col)
        atlantic[i] = make([]bool, col)
    }
    // 利用 DFS 把二维数据按照行优先搜索一遍，分别标记出太平洋和大西洋水流能到达的位置
    for i := 0; i < row; i++ {
        dfs(matrix, i, 0, &pacific, inf)
        dfs(matrix, i, col-1, &atlantic, inf)
    }
    for j := 0; j < col; j++ {
        dfs(matrix, 0, j, &pacific, inf)
        dfs(matrix, row-1, j, &atlantic, inf)
    }
    // 按照列优先搜索一遍，标记出太平洋和大西洋水流能到达的位置。最后两者都能到达的坐标即为所求
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if atlantic[i][j] && pacific[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/08/waterflow-grid.jpg" />
    // Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
    // Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
    // Explanation: The following cells can flow to the Pacific and Atlantic oceans, as shown below:
    // [0,4]: [0,4] -> Pacific Ocean 
    //        [0,4] -> Atlantic Ocean
    // [1,3]: [1,3] -> [0,3] -> Pacific Ocean 
    //        [1,3] -> [1,4] -> Atlantic Ocean
    // [1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean 
    //        [1,4] -> Atlantic Ocean
    // [2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean 
    //        [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
    // [3,0]: [3,0] -> Pacific Ocean 
    //        [3,0] -> [4,0] -> Atlantic Ocean
    // [3,1]: [3,1] -> [3,0] -> Pacific Ocean 
    //        [3,1] -> [4,1] -> Atlantic Ocean
    // [4,0]: [4,0] -> Pacific Ocean 
    //        [4,0] -> Atlantic Ocean
    // Note that there are other possible paths for these cells to flow to the Pacific and Atlantic oceans.
    ocean1 := [][]int{
        {1,2,2,3,5},
        {3,2,3,4,4},
        {2,4,5,3,1},
        {6,7,1,4,5},
        {5,1,1,2,4},
    }
    fmt.Println(pacificAtlantic(ocean1)) // [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
    // Example 2:
    // Input: heights = [[1]]
    // Output: [[0,0]]
    // Explanation: The water can flow from the only cell to the Pacific and Atlantic oceans.
    fmt.Println(pacificAtlantic([][]int{{1}})) // [[0,0]]
}