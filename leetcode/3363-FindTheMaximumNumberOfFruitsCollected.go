package main

// 3363. Find the Maximum Number of Fruits Collected
// There is a game dungeon comprised of n x n rooms arranged in a grid.

// You are given a 2D array fruits of size n x n, where fruits[i][j] represents the number of fruits in the room (i, j). 
// Three children will play in the game dungeon, with initial positions at the corner rooms (0, 0), (0, n - 1), and (n - 1, 0).

// The children will make exactly n - 1 moves according to the following rules to reach the room (n - 1, n - 1):
//     1. The child starting from (0, 0) must move from their current room (i, j) to one of the rooms (i + 1, j + 1), 
//        (i + 1, j), and (i, j + 1) if the target room exists.
//     2. The child starting from (0, n - 1) must move from their current room (i, j) to one of the rooms (i + 1, j - 1), 
//        (i + 1, j), and (i + 1, j + 1) if the target room exists.
//     3. The child starting from (n - 1, 0) must move from their current room (i, j) to one of the rooms (i - 1, j + 1), 
//        (i, j + 1), and (i + 1, j + 1) if the target room exists.

// When a child enters a room, they will collect all the fruits there. 
// If two or more children enter the same room, only one child will collect the fruits, and the room will be emptied after they leave.

// Return the maximum number of fruits the children can collect from the dungeon.

// Example 1:
// Input: fruits = [[1,2,3,4],[5,6,8,7],[9,10,11,12],[13,14,15,16]]
// Output: 100
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/10/15/example_1.gif" />
// In this example:
// The 1st child (green) moves on the path (0,0) -> (1,1) -> (2,2) -> (3, 3).
// The 2nd child (red) moves on the path (0,3) -> (1,2) -> (2,3) -> (3, 3).
// The 3rd child (blue) moves on the path (3,0) -> (3,1) -> (3,2) -> (3, 3).
// In total they collect 1 + 6 + 11 + 16 + 4 + 8 + 12 + 13 + 14 + 15 = 100 fruits.

// Example 2:
// Input: fruits = [[1,1],[1,1]]
// Output: 4
// Explanation:
// In this example:
// The 1st child moves on the path (0,0) -> (1,1).
// The 2nd child moves on the path (0,1) -> (1,1).
// The 3rd child moves on the path (1,0) -> (1,1).
// In total they collect 1 + 1 + 1 + 1 = 4 fruits.

// Constraints:
//     2 <= n == fruits.length == fruits[i].length <= 1000
//     0 <= fruits[i][j] <= 1000

import "fmt"

func maxCollectedFruits(fruits [][]int) int {
    green, red, blue, n := 0, 0, 0, len(fruits)
    for i := 0; i < n; i++ {
        green += fruits[i][i]
        fruits[i][i] = 0
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if j < n - i - 1 {
                fruits[i][j] = 0
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // red
    for i := 1; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            l, m, r := 0, 0, 0
            if j - 1 >= 0 {
                l = fruits[i - 1][j - 1]
            }
            m = fruits[i - 1][j]
            if j+1 < n {
                r = fruits[i - 1][j + 1]
            }
            fruits[i][j] += max(max(l, m), r)
        }
    }
    // blue
    for j := 1; j < n - 1; j++ {
        for i := n - 1; i >= n - i - 1; i-- {
            u, m, d := 0, 0, 0
            if i-1 >= 0 {
                u = fruits[i - 1][j - 1]
            }
            m = fruits[i][j - 1]
            if i + 1 < n {
                d = fruits[i + 1][j - 1]
            }
            fruits[i][j] += max(max(u, m), d)
        }
    }
    red, blue = fruits[n - 2][n - 1], fruits[n - 1][n - 2]
    return green + red + blue
}

func maxCollectedFruits1(fruits [][]int) int {
    res, n := 0, len(fruits)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp[0][n-1] = fruits[0][n-1]
    for i := 1; i < n; i++ {
        for j := max(i + 1, n - i - 1); j < n; j++ {
            dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - 1])
            if j + 1 < n {
                dp[i][j] = max(dp[i][j], dp[i - 1][j + 1])
            }
            dp[i][j] += fruits[i][j]
        }
    }
    for i := 0; i < n; i++ {
        res += fruits[i][i]
    }
    res += dp[n - 2][n - 1]
    dp[n - 1][0] = fruits[n-1][0]
    for j := 1; j < n; j++ {
        for i := max(j + 1, n - j - 1); i < n; i++ {
            dp[i][j] = max(dp[i][j - 1], dp[i - 1][j - 1])
            if i + 1 < n {
                dp[i][j] = max(dp[i][j], dp[i + 1][j - 1])
            }
            dp[i][j] += fruits[i][j]
        }
    }
    res += dp[n - 1][n - 2]
    return res
}

func maxCollectedFruits2(fruits [][]int) int {
    res, n := 0, len(fruits)
    // 对角线
    for i := 0; i < n; i++ {
        res += fruits[i][i]
        fruits[i][i] = 0
    }
    // 不可达区域
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i < n-1-j && i != j {
                fruits[i][j] = 0
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 对角线两侧分别动态规划
    for i := 2; i < n; i++ {
        for j := 0; j < i && j < n-i; j++ {
            // 右侧
            toDiag := 0
            if j >= 1 {
                toDiag = fruits[i-2][n-j]
            }
            fruits[i-1][n-1-j] += max(max(fruits[i-2][n-1-j], fruits[i-2][n-2-j]), toDiag)
            // 左侧
            toDiag = 0
            if j >= 1 {
                toDiag = fruits[n-j][i-2]
            }
            fruits[n-1-j][i-1] += max(max(fruits[n-1-j][i-2], fruits[n-2-j][i-2]), toDiag)
        }
    }
    res += fruits[n-1][n-2]
    res += fruits[n-2][n-1]
    return res
}

func main() {
    // Example 1:
    // Input: fruits = [[1,2,3,4],[5,6,8,7],[9,10,11,12],[13,14,15,16]]
    // Output: 100
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/10/15/example_1.gif" />
    // In this example:
    // The 1st child (green) moves on the path (0,0) -> (1,1) -> (2,2) -> (3, 3).
    // The 2nd child (red) moves on the path (0,3) -> (1,2) -> (2,3) -> (3, 3).
    // The 3rd child (blue) moves on the path (3,0) -> (3,1) -> (3,2) -> (3, 3).
    // In total they collect 1 + 6 + 11 + 16 + 4 + 8 + 12 + 13 + 14 + 15 = 100 fruits.
    fmt.Println(maxCollectedFruits([][]int{{1,2,3,4},{5,6,8,7},{9,10,11,12},{13,14,15,16}})) // 100
    // Example 2:
    // Input: fruits = [[1,1],[1,1]]
    // Output: 4
    // Explanation:
    // In this example:
    // The 1st child moves on the path (0,0) -> (1,1).
    // The 2nd child moves on the path (0,1) -> (1,1).
    // The 3rd child moves on the path (1,0) -> (1,1).
    // In total they collect 1 + 1 + 1 + 1 = 4 fruits.
    fmt.Println(maxCollectedFruits([][]int{{1,1},{1,1}})) // 4

    fmt.Println(maxCollectedFruits1([][]int{{1,2,3,4},{5,6,8,7},{9,10,11,12},{13,14,15,16}})) // 100
    fmt.Println(maxCollectedFruits1([][]int{{1,1},{1,1}})) // 4

    fmt.Println(maxCollectedFruits2([][]int{{1,2,3,4},{5,6,8,7},{9,10,11,12},{13,14,15,16}})) // 100
    fmt.Println(maxCollectedFruits2([][]int{{1,1},{1,1}})) // 4
}