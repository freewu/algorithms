package main

// 2850. Minimum Moves to Spread Stones Over Grid
// You are given a 0-indexed 2D integer matrix grid of size 3 * 3, representing the number of stones in each cell. 
// The grid contains exactly 9 stones, and there can be multiple stones in a single cell.

// In one move, you can move a single stone from its current cell to any other cell if the two cells share a side.
// Return the minimum number of moves required to place one stone in each cell.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/08/23/example1-3.svg" />
// Input: grid = [[1,1,0],[1,1,1],[1,2,1]]
// Output: 3
// Explanation: One possible sequence of moves to place one stone in each cell is: 
// 1- Move one stone from cell (2,1) to cell (2,2).
// 2- Move one stone from cell (2,2) to cell (1,2).
// 3- Move one stone from cell (1,2) to cell (0,2).
// In total, it takes 3 moves to place one stone in each cell of the grid.
// It can be shown that 3 is the minimum number of moves required to place one stone in each cell.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/08/23/example2-2.svg" />
// Input: grid = [[1,3,0],[1,0,0],[1,0,3]]
// Output: 4
// Explanation: One possible sequence of moves to place one stone in each cell is:
// 1- Move one stone from cell (0,1) to cell (0,2).
// 2- Move one stone from cell (0,1) to cell (1,1).
// 3- Move one stone from cell (2,2) to cell (1,2).
// 4- Move one stone from cell (2,2) to cell (2,1).
// In total, it takes 4 moves to place one stone in each cell of the grid.
// It can be shown that 4 is the minimum number of moves required to place one stone in each cell.
 
// Constraints:
//     grid.length == grid[i].length == 3
//     0 <= grid[i][j] <= 9
//     Sum of grid is equal to 9.

import "fmt"

func minimumMoves(grid [][]int) int {
    emptyCells, stoneCells := [][2]int{}, [][2]int{}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[i][j] == 0 {
                emptyCells = append(emptyCells, [2]int{i, j})
            } else if grid[i][j] > 1 {
                stoneCells = append(stoneCells, [2]int{i, j})
            }
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var backtracking func(grid [][]int, emptyCellIdx int) int
    backtracking = func(grid [][]int, emptyCellIdx int) int {
        if emptyCellIdx == len(emptyCells) { return 0 }
        curEmpty := emptyCells[emptyCellIdx]
        dist, res := 0, 1 << 32 - 1
        for i := 0; i < len(stoneCells); i++ {
            take := stoneCells[i]
            if grid[take[0]][take[1]] == 1 { continue }  // Only take if the grid allows it
            d := abs(curEmpty[0]-take[0]) + abs(curEmpty[1]-take[1]) // Calculate manhattan distance
            grid[take[0]][take[1]]-- // Accout for taking this stone
            dist = d + backtracking(grid, emptyCellIdx+1)
            grid[take[0]][take[1]]++
            res = min(dist, res)
        }
        return res
    }
    return backtracking(grid, 0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/08/23/example1-3.svg" />
    // Input: grid = [[1,1,0],[1,1,1],[1,2,1]]
    // Output: 3
    // Explanation: One possible sequence of moves to place one stone in each cell is: 
    // 1- Move one stone from cell (2,1) to cell (2,2).
    // 2- Move one stone from cell (2,2) to cell (1,2).
    // 3- Move one stone from cell (1,2) to cell (0,2).
    // In total, it takes 3 moves to place one stone in each cell of the grid.
    // It can be shown that 3 is the minimum number of moves required to place one stone in each cell.
    fmt.Println(minimumMoves([][]int{{1,1,0},{1,1,1},{1,2,1}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/08/23/example2-2.svg" />
    // Input: grid = [[1,3,0],[1,0,0],[1,0,3]]
    // Output: 4
    // Explanation: One possible sequence of moves to place one stone in each cell is:
    // 1- Move one stone from cell (0,1) to cell (0,2).
    // 2- Move one stone from cell (0,1) to cell (1,1).
    // 3- Move one stone from cell (2,2) to cell (1,2).
    // 4- Move one stone from cell (2,2) to cell (2,1).
    // In total, it takes 4 moves to place one stone in each cell of the grid.
    // It can be shown that 4 is the minimum number of moves required to place one stone in each cell.
    fmt.Println(minimumMoves([][]int{{1,3,0},{1,0,0},{1,0,3}})) // 4
}