package main

// 4084. Last Visited Cell in a Directional Grid
// You are given an m x n matrix grid where each cell contains one of the following characters:
//    1. 'L': move one cell left.
//    2. 'R': move one cell right.
//    3. 'U': move one cell up.
//    4. 'D': move one cell down.

// You start at cell (0, 0), which is considered visited. 
// At each step, move to the cell indicated by the character in your current cell.

// The process ends if the indicated move would either leave the grid or enter a cell that has already been visited. 
// In either case, the move is not made.

// Return an integer array [r, c] representing the row and column indices of the last visited cell. 
// The indices are 0-based.

// Example 1:
// Input: grid = [["R","D"],["U","L"]]
// Output: [1,0]
// Explanation:
// - Start at cell (0, 0), which contains 'R', so move right to (0, 1).
// - At cell (0, 1), the character is 'D', so move down to (1, 1).
// - At cell (1, 1), the character is 'L', so move left to (1, 0).
// - At cell (1, 0), the character is 'U', which indicates moving to (0, 0). 
//   However, (0, 0) has already been visited, so this move is not made.
// - Therefore, the process ends at the last visited cell (1, 0), so the answer is [1, 0].

// Example 2:
// Input: grid = [["R","R","R"]]
// Output: [0,2]
// Explanation:
// - Start at cell (0, 0), which contains 'R', so move right to (0, 1).
// - At cell (0, 1), the character is 'R', so move right to (0, 2).
// - At cell (0, 2), the character is 'R', which indicates moving to (0, 3). 
//   However, (0, 3) is outside the grid, so this move is not made.
// - Therefore, the process ends at the last visited cell (0, 2), so the answer is [0, 2].

// Constraints:
//     - 1 <= m == grid.length <= 10^{5}
//     - 1 <= n == grid[i].length <= 10^{5}
//     - m * n <= 10^{5}
//     - grid[i][j] is one of 'L', 'R', 'U', or 'D'.

import "fmt"

func lastVisitedCell(grid [][]byte) []int {
    i, j, n, m := 0, 0, len(grid), len(grid[0])
    if n == 0 {
        return []int{0, 0}
    }
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, m)
    }
    visited[i][j] = true
    for {
        ii, jj := i, j
        c := grid[i][j]
        switch c {
        case 'U':
            ii--
        case 'D':
            ii++
        case 'L':
            jj--
        case 'R':
            jj++
        }
        // 越界 或者 已经访问过，停止
        if ii < 0 || jj < 0 || ii >= n || jj >= m || visited[ii][jj] {
            break
        }
        i, j = ii, jj
        visited[i][j] = true
    }
    return []int{ i, j }
}

func lastVisitedCell1(grid [][]byte) []int {
    x, y, lx, ly := 0, 0, 0, 0
    for ; x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) ; {
        b := grid[x][y]
        if b == 'E' {
            break
        }
        lx, ly = x, y
        if b == 'L' {
            y--
        }
        if b == 'R' {
            y++
        }
        if b == 'U' {
            x--
        }
        if b == 'D' {
            x++
        }
        grid[lx][ly] = 'E'
    }
    return []int{ lx, ly }
}

func main() {
    // Example 1:
    // Input: grid = [["R","D"],["U","L"]]
    // Output: [1,0]
    // Explanation:
    // - Start at cell (0, 0), which contains 'R', so move right to (0, 1).
    // - At cell (0, 1), the character is 'D', so move down to (1, 1).
    // - At cell (1, 1), the character is 'L', so move left to (1, 0).
    // - At cell (1, 0), the character is 'U', which indicates moving to (0, 0). 
    //   However, (0, 0) has already been visited, so this move is not made.
    // - Therefore, the process ends at the last visited cell (1, 0), so the answer is [1, 0].
    fmt.Println(lastVisitedCell([][]byte{{'R','D'},{'U','L'}})) // [1,0] 
    // Example 2:
    // Input: grid = [["R","R","R"]]
    // Output: [0,2]
    // Explanation:
    // - Start at cell (0, 0), which contains 'R', so move right to (0, 1).
    // - At cell (0, 1), the character is 'R', so move right to (0, 2).
    // - At cell (0, 2), the character is 'R', which indicates moving to (0, 3). 
    //   However, (0, 3) is outside the grid, so this move is not made.
    // - Therefore, the process ends at the last visited cell (0, 2), so the answer is [0, 2].
    fmt.Println(lastVisitedCell([][]byte{{'R','R','R'}})) // [0,2] 

    fmt.Println(lastVisitedCell1([][]byte{{'R','D'},{'U','L'}})) // [1,0] 
    fmt.Println(lastVisitedCell1([][]byte{{'R','R','R'}})) // [0,2]
}