package main

// 1824. Minimum Sideway Jumps
// There is a 3 lane road of length n that consists of n + 1 points labeled from 0 to n. 
// A frog starts at point 0 in the second lane and wants to jump to point n. 
// However, there could be obstacles along the way.

// You are given an array obstacles of length n + 1 where each obstacles[i] (ranging from 0 to 3) describes an obstacle on the lane obstacles[i] at point i. 
// If obstacles[i] == 0, there are no obstacles at point i. 
// There will be at most one obstacle in the 3 lanes at each point.
//     For example, if obstacles[2] == 1, then there is an obstacle on lane 1 at point 2.

// The frog can only travel from point i to point i + 1 on the same lane if there is not an obstacle on the lane at point i + 1. 
// To avoid obstacles, the frog can also perform a side jump to jump to another lane (even if they are not adjacent) at the same point if there is no obstacle on the new lane.
//     For example, the frog can jump from lane 3 at point 3 to lane 1 at point 3.

// Return the minimum number of side jumps the frog needs to reach any lane at point n starting from lane 2 at point 0.

// Note: There will be no obstacles on points 0 and n.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/25/ic234-q3-ex1.png" />
// Input: obstacles = [0,1,2,3,0]
// Output: 2 
// Explanation: The optimal solution is shown by the arrows above. There are 2 side jumps (red arrows).
// Note that the frog can jump over obstacles only when making side jumps (as shown at point 2).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/25/ic234-q3-ex2.png" />
// Input: obstacles = [0,1,1,3,3,0]
// Output: 0
// Explanation: There are no obstacles on lane 2. No side jumps are required.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/03/25/ic234-q3-ex3.png" />
// Input: obstacles = [0,2,1,0,3,0]
// Output: 2
// Explanation: The optimal solution is shown by the arrows above. There are 2 side jumps.

// Constraints:
//     obstacles.length == n + 1
//     1 <= n <= 5 * 10^5
//     0 <= obstacles[i] <= 3
//     obstacles[0] == obstacles[n] == 0

import "fmt"

func minSideJumps(obstacles []int) int {
    l1, l2, l3, inf := 1, 0, 1, 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(obstacles)-1; i++ {
        switch obstacles[i+1] {
            case 1:
                if obstacles[i] != 2 {
                    l2 = min(l2, 1 + l1)
                }
                if obstacles[i] != 3 {
                    l3 = min(l3, 1 + l1)
                }
                l1 = inf
            case 2:
                if obstacles[i] != 1 {
                    l1 = min(l1, 1 + l2)
                }
                if obstacles[i] != 3 {
                    l3 = min(l3, 1 + l2)
                }
                l2 = inf
            case 3:
                if obstacles[i] != 1 {
                    l1 = min(l1, 1 + l3)
                }
                if obstacles[i] != 2 {
                    l2 = min(l2, 1 + l3)
                }
                l3 = inf
        }
    }
    return min(l1, min(l2, l3))
}

func minSideJumps1(obstacles []int) int {
    dp, n, inf := make([][]int, 3), len(obstacles), 1 << 31
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i := 0; i < 3; i++ {
        for j := 0; j < len(obstacles); j++ {
            dp[i][j] = inf
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // 表示青蛙到达第0个坐标下的2号跑道上所需的最少侧跳次数
    dp[0][0], dp[1][0], dp[2][0] = 1, 0, 1
    for j := 1; j < n; j++ {
        for i := 0; i < 3; i++ {
            if obstacles[j] != i+1 {
                dp[i][j] = dp[i][j-1]
            }
        }
        for i := 0; i < 3; i++ {
            if obstacles[j] != i+1 {
                dp[i][j] = min(dp[i][j], min(dp[0][j], min(dp[1][j], dp[2][j])) + 1)
            }
        }
    }
    return min(dp[0][n-1], min(dp[1][n-1], dp[2][n-1]))
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/25/ic234-q3-ex1.png" />
    // Input: obstacles = [0,1,2,3,0]
    // Output: 2 
    // Explanation: The optimal solution is shown by the arrows above. There are 2 side jumps (red arrows).
    // Note that the frog can jump over obstacles only when making side jumps (as shown at point 2).
    fmt.Println(minSideJumps([]int{0,1,2,3,0})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/25/ic234-q3-ex2.png" />
    // Input: obstacles = [0,1,1,3,3,0]
    // Output: 0
    // Explanation: There are no obstacles on lane 2. No side jumps are required.
    fmt.Println(minSideJumps([]int{0,1,1,3,3,0})) // 0
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/03/25/ic234-q3-ex3.png" />
    // Input: obstacles = [0,2,1,0,3,0]
    // Output: 2
    // Explanation: The optimal solution is shown by the arrows above. There are 2 side jumps.
    fmt.Println(minSideJumps([]int{0,2,1,0,3,0})) // 2

    fmt.Println(minSideJumps1([]int{0,1,2,3,0})) // 2
    fmt.Println(minSideJumps1([]int{0,1,1,3,3,0})) // 0
    fmt.Println(minSideJumps1([]int{0,2,1,0,3,0})) // 2
}