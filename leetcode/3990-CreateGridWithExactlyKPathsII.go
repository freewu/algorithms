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

func createGrid(k int) []string {
	if k == 0 {
		return []string{}
	}
	// 存储每一位需要的2的幂次
	var powers []int
	tmp := k
	pow := 0
	for tmp > 0 {
		if tmp%2 == 1 {
			powers = append(powers, pow)
		}
		tmp /= 2
		pow++
	}
	// 每一个power对应一个通路块，块之间用1列#分隔
	var row0, row1 string
	for idx, p := range powers {
		length := (1 << p) + 1 // 2^p +1
		// 通路块第一行：全.
		block0 := make([]byte, length)
		for i := range block0 {
			block0[i] = '.'
		}
		// 通路块第二行：首尾.，中间#
		block1 := make([]byte, length)
		block1[0] = '.'
		block1[length-1] = '.'
		for i := 1; i < length-1; i++ {
			block1[i] = '#'
		}
		row0 += string(block0)
		row1 += string(block1)
		// 块之间加分隔墙（1列#），最后一块不加
		if idx != len(powers)-1 {
			row0 += "#"
			row1 += "#"
		}
	}
	return []string{row0, row1}
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
    fmt.Println(createGrid(3)) // ["...","#..","#.."]

    fmt.Println(createGrid(1)) // ["...","#..","#.."]
    fmt.Println(createGrid(4)) // []
    fmt.Println(createGrid(8)) // ["...","#..","#.."]
    fmt.Println(createGrid(99)) // ["...","#..","#.."]
    fmt.Println(createGrid(100)) // ["...","#..","#.."]
    fmt.Println(createGrid(101)) // ["...","#..","#.."]
    fmt.Println(createGrid(999)) // ["...","#..","#.."]
    fmt.Println(createGrid(1000)) // ["...","#..","#.."]
}