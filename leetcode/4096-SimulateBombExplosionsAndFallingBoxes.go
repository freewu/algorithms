package main

// 4096. Simulate Bomb Explosions and Falling Boxes
// You are given a 2D character array grid of size m x n, where each cell contains one of the following symbols:
//     1. '#': a box.
//     2. '-': an empty cell.
//     3. '*': a bomb.

// Each second, the following events occur in order:
//     1. Explosions: A bomb explodes if it is orthogonally adjacent to at least one box (shares a side with it). 
//        An exploding bomb destroys all boxes in its 8 neighboring cells, including diagonally adjacent cells, turning them into empty cells. 
//        All explosions happen simultaneously, based on the grid at the beginning of that second. 
//        Empty cells and bombs remain unchanged.
//     2. Gravity: After the explosions, all non-static boxes simultaneously fall down by exactly one row. 
//        A box is static if it is in the bottommost row, immediately above a bomb, or immediately above another static box.

// Bombs never move or disappear, and the same bomb may explode multiple times. 
// Boxes above non-static boxes can fall together with them.

// Return the final grid after 10^{100} seconds.

// Example 1:
// Input: grid = [["#","#","#"],["-","-","*"],["-","-","-"]]
// Output: [["-","-","-"],["-","-","*"],["#","-","-"]]
// Explanation:
// Let (r, c) represent the cell in the r^{th} row from the top and c^{th} column from the left. 
// For example, (0, 0) denotes the top-left most cell.
// The 1st second:
// - Since the bomb at (1, 2) is orthogonally adjacent to the box at (0, 2), the boxes at (0, 1) and (0, 2) become empty cells.
// - The box at (0, 0) falls to (1, 0).
// The 2nd second:
// - The bomb at (1, 2) is not orthogonally adjacent to any box. Therefore, nothing happens.
// - The box at (1, 0) falls to (2, 0).
// Nothing happens afterwards.

// Example 2:
// Input: grid = [["-","*"],["#","-"]]
// Output: [["-","*"],["#","-"]]
// Explanation:
//  Nothing happens because the bomb is not orthogonally adjacent to any box.

// Example 3:
// Input: grid = [["-","*","#"],["#","-","-"]]
// Output: [["-","*","-"],["-","-","-"]]
// Explanation:
// The 1st second:
// - The bomb at (0, 1) is orthogonally adjacent to the box at (0, 2). 
//   Therefore, the boxes at (0, 2) and (1, 0) become empty cells.
// Nothing happens afterwards.

// Example 4:
// Input: grid = [["#","*","-","#"],["-","-","*","-"]]
// Output: [["-","*","-","-"],["-","-","*","-"]]
// Explanation:
// The 1st second:
// - Since the bomb at (0, 1) is orthogonally adjacent to the box at (0, 0), the box at (0, 0) becomes an empty cell. 
//   The other bomb at (1, 2) is not orthogonally adjacent to any box.
// - The box at (0, 3) falls to (1, 3).
// The 2nd second:
// - Since the bomb at (1, 2) is orthogonally adjacent to the box at (1, 3), the box at (1, 3) becomes an empty cell. 
//   The other bomb at (0, 1) is not orthogonally adjacent to any box.

// Nothing happens afterwards.
//     - 1 <= m == grid.length <= 100
//     - 1 <= n == grid[i].length <= 100
//     - grid[i][j] is '#', '-', or '*'.

import "fmt"

func finalGrid(grid [][]byte) [][]byte {
    m, n := len(grid), len(grid[0])
    if m == 0 {
        return nil
    }
    // 方向
    ortho := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    dir8 := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
    // 复制初始状态
    curr := make([][]byte, m)
    for i := range curr {
        curr[i] = append([]byte(nil), grid[i]...)
    }
    equal := func(a, b [][]byte) bool {
        for i := range a {
            for j := range a[i] {
                if a[i][j] != b[i][j] {
                    return false
                }
            }
        }
        return true
    }
    for {
        // 1. 收集本轮要爆炸的炸弹
        var boom [][]int
        for r := 0; r < m; r++ {
            for c := 0; c < n; c++ {
                if curr[r][c] != '*' {
                    continue
                }
                ok := false
                for _, d := range ortho {
                    nr, nc := r+d[0], c+d[1]
                    if nr >= 0 && nr < m && nc >= 0 && nc < n && curr[nr][nc] == '#' {
                        ok = true
                        break
                    }
                }
                if ok {
                    boom = append(boom, []int{r, c})
                }
            }
        }
        // 爆炸快照
        afterExp := make([][]byte, m)
        for i := range afterExp {
            afterExp[i] = append([]byte(nil), curr[i]...)
        }
        for _, b := range boom {
            br, bc := b[0], b[1]
            for _, d := range dir8 {
                nr, nc := br+d[0], bc+d[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n {
                    if afterExp[nr][nc] == '#' {
                        afterExp[nr][nc] = '-'
                    }
                }
            }
        }
        // 2. 重力下落，逐列
        afterGrav := make([][]byte, m)
        for i := range afterGrav {
            afterGrav[i] = append([]byte(nil), afterExp[i]...)
        }
        for c := 0; c < n; c++ {
            row := m - 1
            for row >= 0 {
                if afterGrav[row][c] == '*' {
                    row--
                    continue
                }
                cnt := 0
                ptr := row
                for ptr >= 0 && afterGrav[ptr][c] != '*' {
                    if afterGrav[ptr][c] == '#' {
                        cnt++
                    }
                    ptr--
                }
                for k := ptr + 1; k <= row; k++ {
                    if k > row-cnt {
                        afterGrav[k][c] = '#'
                    } else {
                        afterGrav[k][c] = '-'
                    }
                }
                row = ptr
            }
        }
        // 稳态判断：爆炸没变化，重力也没变化 → 停止
        if equal(afterExp, curr) && equal(afterGrav, afterExp) {
            return curr
        }
        curr = afterGrav
    }
}

func main() {
    // Example 1:
    // Input: grid = [["#","#","#"],["-","-","*"],["-","-","-"]]
    // Output: [["-","-","-"],["-","-","*"],["#","-","-"]]
    // Explanation:
    // Let (r, c) represent the cell in the r^{th} row from the top and c^{th} column from the left. 
    // For example, (0, 0) denotes the top-left most cell.
    // The 1st second:
    // - Since the bomb at (1, 2) is orthogonally adjacent to the box at (0, 2), the boxes at (0, 1) and (0, 2) become empty cells.
    // - The box at (0, 0) falls to (1, 0).
    // The 2nd second:
    // - The bomb at (1, 2) is not orthogonally adjacent to any box. Therefore, nothing happens.
    // - The box at (1, 0) falls to (2, 0).
    // Nothing happens afterwards.
    fmt.Println(finalGrid([][]byte{{byte('#'), byte('#'), byte('#')}, {byte('-'), byte('-'), byte('*')}, {byte('-'), byte('-'), byte('-')}})) // [["-","-","-"],["-","-","*"],["#","-","-"]]
    // Example 2:
    // Input: grid = [["-","*"],["#","-"]]
    // Output: [["-","*"],["#","-"]]
    // Explanation:
    //  Nothing happens because the bomb is not orthogonally adjacent to any box.
    fmt.Println(finalGrid([][]byte{{byte('-'), byte('*')}, {byte('#'), byte('-')}})) // [["-","*","#","-"]]
    // Example 3:
    // Input: grid = [["-","*","#"],["#","-","-"]]
    // Output: [["-","*","-"],["-","-","-"]]
    // Explanation:
    // The 1st second:
    // - The bomb at (0, 1) is orthogonally adjacent to the box at (0, 2). 
    //   Therefore, the boxes at (0, 2) and (1, 0) become empty cells.
    // Nothing happens afterwards.
    fmt.Println(finalGrid([][]byte{{byte('-'), byte('*')}, {byte('#'), byte('#')}})) // [["-","*","-","-"],["-","-","*"],["#","-","-"]]
    // Example 4
    // Input: grid = [["#","*","-","#"],["-","-","*","-"]]
    // Output: [["-","*","-","-"],["-","-","*","-"]]
    // Explanation:
    // The 1st second:
    // - Since the bomb at (0, 1) is orthogonally adjacent to the box at (0, 0), the box at (0, 0) becomes an empty cell. 
    //   The other bomb at (1, 2) is not orthogonally adjacent to any box.
    // - The box at (0, 3) falls to (1, 3).
    // The 2nd second:
    // - Since the bomb at (1, 2) is orthogonally adjacent to the box at (1, 3), the box at (1, 3) becomes an empty cell. 
    //   The other bomb at (0, 1) is not orthogonally adjacent to any box.
    fmt.Println(finalGrid([][]byte{{byte('#'), byte('#'), byte('#')}, {byte('-'), byte('-'), byte('*')}, {byte('-'), byte('-'), byte('-')}})) // [["-","-","-"],["-","-","*"],["#","-","-"]]        
}