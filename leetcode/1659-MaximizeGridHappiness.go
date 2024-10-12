package main

// 1659. Maximize Grid Happiness
// You are given four integers, m, n, introvertsCount, and extrovertsCount. 
// You have an m x n grid, and there are two types of people: introverts and extroverts. 
// There are introvertsCount introverts and extrovertsCount extroverts.

// You should decide how many people you want to live in the grid and assign each of them one grid cell. 
// Note that you do not have to have all the people living in the grid.

// The happiness of each person is calculated as follows:
//     Introverts start with 120 happiness and lose 30 happiness for each neighbor (introvert or extrovert).
//     Extroverts start with 40 happiness and gain 20 happiness for each neighbor (introvert or extrovert).

// Neighbors live in the directly adjacent cells north, east, south, and west of a person's cell.

// The grid happiness is the sum of each person's happiness. 
// Return the maximum possible grid happiness.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/grid_happiness.png" />
// Input: m = 2, n = 3, introvertsCount = 1, extrovertsCount = 2
// Output: 240
// Explanation: Assume the grid is 1-indexed with coordinates (row, column).
// We can put the introvert in cell (1,1) and put the extroverts in cells (1,3) and (2,3).
// - Introvert at (1,1) happiness: 120 (starting happiness) - (0 * 30) (0 neighbors) = 120
// - Extrovert at (1,3) happiness: 40 (starting happiness) + (1 * 20) (1 neighbor) = 60
// - Extrovert at (2,3) happiness: 40 (starting happiness) + (1 * 20) (1 neighbor) = 60
// The grid happiness is 120 + 60 + 60 = 240.
// The above figure shows the grid in this example with each person's happiness. The introvert stays in the light green cell while the extroverts live on the light purple cells.

// Example 2:
// Input: m = 3, n = 1, introvertsCount = 2, extrovertsCount = 1
// Output: 260
// Explanation: Place the two introverts in (1,1) and (3,1) and the extrovert at (2,1).
// - Introvert at (1,1) happiness: 120 (starting happiness) - (1 * 30) (1 neighbor) = 90
// - Extrovert at (2,1) happiness: 40 (starting happiness) + (2 * 20) (2 neighbors) = 80
// - Introvert at (3,1) happiness: 120 (starting happiness) - (1 * 30) (1 neighbor) = 90
// The grid happiness is 90 + 80 + 90 = 260.

// Example 3:
// Input: m = 2, n = 2, introvertsCount = 4, extrovertsCount = 0
// Output: 240

// Constraints:
//     1 <= m, n <= 5
//     0 <= introvertsCount, extrovertsCount <= min(m * n, 6)

import "fmt"

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
    memo := make([][][7][7]int, m)
    for i := 0; i < len(memo); i++ {
        memo[i] = make([][7][7]int, 1 << (n*2))
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dp func(m, n int, i, j int, curr, prev int, intro, extro int, memo [][][7][7]int) int
    dp = func(m, n int, i, j int, curr, prev int, intro, extro int, memo [][][7][7]int) int {
        if i == m { return 0 } // // We've reached the end of the grid
        if j == n { return dp(m, n, i+1, 0, 0, curr, intro, extro, memo) } // We've reached the end of the current row, go to the next row
        if j == 0 && memo[i][prev][intro][extro] != 0 { return memo[i][prev][intro][extro] }
        // We have three options; We can skip this cell, put an introvert (if we have one)
        // or an extrovert (if we have one)
        res, up, left := dp(m, n, i, j+1, curr, prev, intro, extro, memo), 0, 0
        if j > 0 {
            left = (curr >> ((j-1)*2)) & 3 // left comes from curr
        }
        // intro = 1
        // extro = 2
        up = (prev >> (j*2)) & 3
        if intro > 0 {
            curr := 120 + dp(m, n, i, j+1, curr | (1 << (j*2)), prev, intro-1, extro, memo)
            if up == 1 {
                curr -= 30 // up introvert loses happiness
                curr -= 30 // curr introvert loses happiness
            } else if up == 2 {
                curr += 20 // up extrovert gains happiness
                curr -= 30 // curr introvert loses happiness
            }
            if left == 1 {
                curr -= 30 // left introvert loses happiness
                curr -= 30 // curr introvert loses happiness
            } else if left == 2 {
                curr += 20 // left extrovert gains happiness
                curr -= 30 // curr introvert loses happiness
            }
            res = max(res, curr)
        }
        if extro > 0 {
            curr := 40 + dp(m, n, i, j+1, curr | (2 << (j*2)), prev, intro, extro-1, memo)
            if up == 1 {
                curr -= 30 // up introvert loses happiness
                curr += 20 // curr extrovert gains happiness
            } else if up == 2 {
                curr += 20 // up extrovert gains happiness
                curr += 20 // curr extrovert gains happiness
            }
            if left == 1 {
                curr -= 30 // left introvert loses happiness
                curr += 20 // curr extrovert gains happiness
            } else if left == 2 {
                curr += 20 // left extrovert gains happiness
                curr += 20 // curr extrovert gains happiness
            }
            res = max(res, curr)
        }
        if j == 0 {
            memo[i][prev][intro][extro] = res
        }
        return res
    }
    return dp(m, n, 0, 0, 0, 0, introvertsCount, extrovertsCount, memo)  
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/05/grid_happiness.png" />
    // Input: m = 2, n = 3, introvertsCount = 1, extrovertsCount = 2
    // Output: 240
    // Explanation: Assume the grid is 1-indexed with coordinates (row, column).
    // We can put the introvert in cell (1,1) and put the extroverts in cells (1,3) and (2,3).
    // - Introvert at (1,1) happiness: 120 (starting happiness) - (0 * 30) (0 neighbors) = 120
    // - Extrovert at (1,3) happiness: 40 (starting happiness) + (1 * 20) (1 neighbor) = 60
    // - Extrovert at (2,3) happiness: 40 (starting happiness) + (1 * 20) (1 neighbor) = 60
    // The grid happiness is 120 + 60 + 60 = 240.
    // The above figure shows the grid in this example with each person's happiness. The introvert stays in the light green cell while the extroverts live on the light purple cells.
    fmt.Println(getMaxGridHappiness(2,3,1,2)) // 240
    // Example 2:
    // Input: m = 3, n = 1, introvertsCount = 2, extrovertsCount = 1
    // Output: 260
    // Explanation: Place the two introverts in (1,1) and (3,1) and the extrovert at (2,1).
    // - Introvert at (1,1) happiness: 120 (starting happiness) - (1 * 30) (1 neighbor) = 90
    // - Extrovert at (2,1) happiness: 40 (starting happiness) + (2 * 20) (2 neighbors) = 80
    // - Introvert at (3,1) happiness: 120 (starting happiness) - (1 * 30) (1 neighbor) = 90
    // The grid happiness is 90 + 80 + 90 = 260.
    fmt.Println(getMaxGridHappiness(3,1,2,1)) // 260
    // Example 3:
    // Input: m = 2, n = 2, introvertsCount = 4, extrovertsCount = 0
    // Output: 240
    fmt.Println(getMaxGridHappiness(2,2,4,0)) // 240
}