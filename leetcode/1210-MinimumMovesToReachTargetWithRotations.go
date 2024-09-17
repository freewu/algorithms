package main

// 1210. Minimum Moves to Reach Target with Rotations
// In an n*n grid, there is a snake that spans 2 cells and starts moving from the top left corner at (0, 0) and (0, 1). 
// The grid has empty cells represented by zeros and blocked cells represented by ones. 
// The snake wants to reach the lower right corner at (n-1, n-2) and (n-1, n-1).

// In one move the snake can:
//     1. Move one cell to the right if there are no blocked cells there. 
//        This move keeps the horizontal/vertical position of the snake as it is.
//     2. Move down one cell if there are no blocked cells there. 
//        This move keeps the horizontal/vertical position of the snake as it is.
//     3. Rotate clockwise if it's in a horizontal position and the two cells under it are both empty. 
//        In that case the snake moves from (r, c) and (r, c+1) to (r, c) and (r+1, c).
//        <img src="https://assets.leetcode.com/uploads/2019/09/24/image-2.png" />
//     4. Rotate counterclockwise if it's in a vertical position and the two cells to its right are both empty. 
//        In that case the snake moves from (r, c) and (r+1, c) to (r, c) and (r, c+1).
//        <img src="https://assets.leetcode.com/uploads/2019/09/24/image-1.png" />

// Return the minimum number of moves to reach the target.

// If there is no way to reach the target, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/09/24/image.png" />
// Input: grid = [[0,0,0,0,0,1],
//                [1,1,0,0,1,0],
//                [0,0,0,0,1,1],
//                [0,0,1,0,1,0],
//                [0,1,1,0,0,0],
//                [0,1,1,0,0,0]]
// Output: 11
// Explanation:
// One possible solution is [right, right, rotate clockwise, right, down, down, down, down, rotate counterclockwise, right, down].

// Example 2:
// Input: grid = [[0,0,1,1,1,1],
//                [0,0,0,0,1,1],
//                [1,1,0,0,0,1],
//                [1,1,1,0,0,1],
//                [1,1,1,0,0,1],
//                [1,1,1,0,0,0]]
// Output: 9

// Constraints:
//     2 <= n <= 100
//     0 <= grid[i][j] <= 1
//     It is guaranteed that the snake starts at empty cells.

import "fmt"

func minimumMoves(grid [][]int) int {
    n, inf := len(grid), 1 << 31
    dp := make([][][2]int, n)
    for i := range dp {
        dp[i] = make([][2]int, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k < 2; k++ {
                if i==0 && j==0 &&k==0 {
                    dp[i][j][k] = 0
                } else if grid[i][j] == 1 { // wall
                    dp[i][j][k] = inf
                } else if k == 0 && ( j + 1 >= n || grid[i][j+1] == 1) {
                    dp[i][j][k] = inf
                } else if k == 1 && ( i + 1 >= n || grid[i+1][j] == 1) {
                    dp[i][j][k] = inf
                } else{
                    choice1, choice2 := inf, inf // go down, go right
                    if i - 1 >= 0 { choice1 = dp[i-1][j][k] }
                    if j - 1 >= 0 { choice2 = dp[i][j-1][k] }
                    tempMinChoice := min(choice1,choice2)
                    if tempMinChoice == inf {
                        dp[i][j][k] = tempMinChoice
                    } else {
                        dp[i][j][k] = tempMinChoice + 1
                    }
                }
            }
            // rotate
            if i+1 < n && j + 1 < n && grid[i][j+1] != 1 && grid[i+1][j+1] != 1 && 
               dp[i][j][1] != inf && dp[i][j][0] > dp[i][j][1] + 1 {
                dp[i][j][0] = dp[i][j][1] + 1
            }
            if i + 1 < n && j + 1 < n && grid[i+1][j] != 1 && grid[i+1][j+1] != 1 &&
               dp[i][j][0] != inf && dp[i][j][1] > dp[i][j][0] + 1 {
                dp[i][j][1] = dp[i][j][0]+1
            }
        }
    }
    if dp[n-1][n-2][0]== inf { return -1 }
    return dp[n-1][n-2][0]
}

func minimumMoves1(grid [][]int) int {
    // 网格图-BFS + 化简(异或) O(n^2)
    // 蛇开始为水平状态,结束也为水平状态
    // 化简一!! 如何表示一个蛇?
    // 蛇有两个状态,占两个格子,如何用表示一个蛇, 使用 "尾部所在的格子{x,y}" + 蛇的状态"s" 即可表示
    // 化简二!! 如何表示三种移动?
    // - 水平{x+1,y,s} => 注意:竖直的蛇可以直接水平移动的(不需要先调方向)
    // - 竖直{x,y+1,s}
    // - 旋转,可以发现"根"不边, 只是s改变, 而s可以用0/1表示两种状态, 使用异或的(0/1)做出状态切换
    // 化简三:  如何保证可以行动?
    // - 蛇头在前,所以需要判断蛇头的位置, 蛇头和蛇尾正好根据状态,确定额外的1在那个上,x+s,y+(s^1)恰好是蛇头位置
    // - 水平,竖直,都是保证 新尾部 和新头部 都不能是障碍(竖着方向的蛇可以直接水平平移,头尾都会换)
    // - 旋转,可以看到需要保证四个格子,而蛇尾,初始和终点的"头部"分别占去了三个, 只需判断 x+1,y+1的位置是否被占用即可
    type tuple struct{ x, y, s int }
    dirs := []tuple{{0, 1, 0}, {1, 0, 0}, {0, 0, 1}} //  水平移动,竖直移动,旋转
    m, n := len(grid), len(grid[0])

    vis := make([][][2]bool, m)
    for i := range vis {
        vis[i] = make([][2]bool, n)
    }
    q := []tuple{{0, 0, 0}} // 存储蛇的尾巴位置和状态
    vis[0][0][0] = true
    for step := 1; len(q) > 0; step++ { // step代表到达"入队列"的元素所需要的步数
        tmp := q
        q = nil
        for _, cur := range tmp {
            for _, d := range dirs {
                nx, ny, ns := cur.x+d.x, cur.y+d.y, cur.s^d.s // 尾巴的位置和状态: 如果是旋转,尾部不变,只是状态切换(利用^1实现0/1切换)
                nx2, ny2 := nx+ns, ny+(ns^1)                  // 头部的位置
                if nx2 < m && ny2 < n && !vis[nx][ny][ns] &&  // 1.头部永远在前,只需保证头部不越界即可(只会像右/下,走,不会从上面越界)
                    grid[nx][ny] == 0 && grid[nx2][ny2] == 0 && //  2.新头部/尾部可能都切换
                    (d.s == 0 || grid[nx+1][ny+1] == 0) { // 3.如果是旋转,则额外需要验证其多扫过的一个格子(原头部和尾部在上一轮已经验证,新头部上一行验证,只需多验证一个即可)

                    if nx == m-1 && ny == n-2 { // 尾部到达这个点只有一种合法的状态=>无需验证头部
                        return step
                    }
                    vis[nx][ny][ns] = true
                    q = append(q, tuple{nx, ny, ns})
                }
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/09/24/image.png" />
    // Input: grid = [[0,0,0,0,0,1],
    //                [1,1,0,0,1,0],
    //                [0,0,0,0,1,1],
    //                [0,0,1,0,1,0],
    //                [0,1,1,0,0,0],
    //                [0,1,1,0,0,0]]
    // Output: 11
    // Explanation:
    // One possible solution is [right, right, rotate clockwise, right, down, down, down, down, rotate counterclockwise, right, down].
    grid1 := [][]int{
        {0,0,0,0,0,1},
        {1,1,0,0,1,0},
        {0,0,0,0,1,1},
        {0,0,1,0,1,0},
        {0,1,1,0,0,0},
        {0,1,1,0,0,0},
    }
    fmt.Println(minimumMoves(grid1)) // 11
    // Example 2:
    // Input: grid = [[0,0,1,1,1,1],
    //                [0,0,0,0,1,1],
    //                [1,1,0,0,0,1],
    //                [1,1,1,0,0,1],
    //                [1,1,1,0,0,1],
    //                [1,1,1,0,0,0]]
    // Output: 9
    grid2 := [][]int{
        {0,0,1,1,1,1},
        {0,0,0,0,1,1},
        {1,1,0,0,0,1},
        {1,1,1,0,0,1},
        {1,1,1,0,0,1},
        {1,1,1,0,0,0},
    }
    fmt.Println(minimumMoves(grid2)) // 9

    fmt.Println(minimumMoves1(grid1)) // 11
    fmt.Println(minimumMoves1(grid2)) // 9
}