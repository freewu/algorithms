package main

// 3990. Create Grid With Exactly K Paths II
// You are given an integer k.

// Construct any grid consisting only of the characters '.' and '#', where:
//     1. '.' represents a free cell.
//     2. '#' represents an obstacle cell.

// The grid must contain at most 25 rows and at most 25 columns.

// A valid path is a sequence of free cells that:
//     1. Starts at the top-left cell (0, 0).
//     2. Ends at the bottom-right cell (m - 1, n - 1), where m and n are the dimensions of your constructed grid.
//     3. Moves only:
//         3.1. Move Right, from (i, j) to (i, j + 1), or
//         3.2. Move Down, from (i, j) to (i + 1, j).

// Return any grid such that there are exactly k valid paths from the top-left cell to the bottom-right cell. 
// If no such grid exists, return an empty array.

// Example 1:
// Input: k = 2
// Output: ["..#","#..","#.."]
// Explanation:
// ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/31/screenshot-2026-05-31-at-82224pm.png" />
// The grid contains exactly 2 valid paths from (0, 0) to (2, 2):
// (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2)
// (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)

// Example 2:
// Input: k = 3
// Output: ["...","#..","#.."]
// Explanation:
// ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/31/screenshot-2026-05-31-at-82251pm.png" />
// The grid contains exactly 3 valid paths from (0, 0) to (2, 2):
// (0, 0) → (0, 1) → (0, 2) → (1, 2) → (2, 2)
// (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2)
// (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)

// Constraints:​​​​​​​
//     1 <= k <= 1000

import "fmt"
import "math/bits"

func createGrid(k int) []string {
    width := bits.Len(uint(k))
    m, n := width * 2, width + 3
    arr := make([][]rune, m)
    for i := range arr {
        row := make([]rune, n)
        for j := 0; j < n-1; j++ {
            row[j] = '#'
        }
        row[n-1] = '.'
        arr[i] = row
    }
    for j := 0; j < width; j++ {
        i := j * 2
        arr[i][j] = '.'
        arr[i][j+1] = '.'
        arr[i+1][j] = '.'
        arr[i+1][j+1] = '.'
    }
    for i := 0; i < width; i++ {
        if (k >> i) & 1 == 1 {
            row := i * 2
            for j := i + 2; j < n; j++ {
                arr[row][j] = '.'
            }
        }
    }
    res := make([]string, m)
    for i, r := range arr {
        res[i] = string(r)
    }
    return res
}

func main() {
    // Example 1:
    // Input: k = 2
    // Output: ["..#","#..","#.."]
    // Explanation:
    // ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/31/screenshot-2026-05-31-at-82224pm.png" />
    // The grid contains exactly 2 valid paths from (0, 0) to (2, 2):
    // (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2)
    // (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)
    fmt.Println(createGrid(2)) // ["..#","#..","#.."]
    // Example 2:
    // Input: k = 3
    // Output: ["...","#..","#.."]
    // Explanation:
    // ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/31/screenshot-2026-05-31-at-82251pm.png" />
    // The grid contains exactly 3 valid paths from (0, 0) to (2, 2):
    // (0, 0) → (0, 1) → (0, 2) → (1, 2) → (2, 2)
    // (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2)
    // (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)
    fmt.Println(createGrid(3)) // ["...", "#..", "#.."]

    fmt.Println(createGrid(1)) // [.... ..#.]
    fmt.Println(createGrid(4)) // [..###. ..###. #..##. #..##. ##.... ##..#.]
    fmt.Println(createGrid(8)) // [..####. ..####. #..###. #..###. ##..##. ##..##. ###.... ###..#.]
    fmt.Println(createGrid(99)) // [.......... ..#######. #......... #..######. ##..#####. ##..#####. ###..####. ###..####. ####..###. ####..###. #####..... #####..##. ######.... ######..#.]
    fmt.Println(createGrid(100)) // [..#######. ..#######. #..######. #..######. ##........ ##..#####. ###..####. ###..####. ####..###. ####..###. #####..... #####..##. ######.... ######..#.]
    fmt.Println(createGrid(101)) // [.......... ..#######. #..######. #..######. ##........ ##..#####. ###..####. ###..####. ####..###. ####..###. #####..... #####..##. ######.... ######..#.]  
    fmt.Println(createGrid(999)) // [............. ..##########. #............ #..#########. ##........... ##..########. ###..#######. ###..#######. ####..######. ####..######. #####........ #####..#####. ######....... ######..####. #######...... #######..###. ########..... ########..##. #########.... #########..#.]
    fmt.Println(createGrid(1000)) // [..##########. ..##########. #..#########. #..#########. ##..########. ##..########. ###.......... ###..#######. ####..######. ####..######. #####........ #####..#####. ######....... ######..####. #######...... #######..###. ########..... ########..##. #########.... #########..#.]
}