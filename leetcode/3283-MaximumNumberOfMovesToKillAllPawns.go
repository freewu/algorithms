package main

// 3283. Maximum Number of Moves to Kill All Pawns
// There is a 50 x 50 chessboard with one knight and some pawns on it. 
// You are given two integers kx and ky where (kx, ky) denotes the position of the knight, 
// and a 2D array positions where positions[i] = [xi, yi] denotes the position of the pawns on the chessboard.

// Alice and Bob play a turn-based game, where Alice goes first. In each player's turn:
//     1. The player selects a pawn that still exists on the board and captures it with the knight in the fewest possible moves. 
//        Note that the player can select any pawn, it might not be one that can be captured in the least number of moves.
//     2. In the process of capturing the selected pawn, the knight may pass other pawns without capturing them. 
//        Only the selected pawn can be captured in this turn.

// Alice is trying to maximize the sum of the number of moves made by both players until there are no more pawns on the board, 
// whereas Bob tries to minimize them.

// Return the maximum total number of moves made during the game that Alice can achieve, assuming both players play optimally.

// Note that in one move, a chess knight has eight possible positions it can move to, as illustrated below. 
// Each move is two cells in a cardinal direction, then one cell in an orthogonal direction.
// <img src="https://assets.leetcode.com/uploads/2024/08/01/chess_knight.jpg" />

// Example 1:
// Input: kx = 1, ky = 1, positions = [[0,0]]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/16/gif3.gif" />
// The knight takes 4 moves to reach the pawn at (0, 0).

// Example 2:
// Input: kx = 0, ky = 2, positions = [[1,1],[2,2],[3,3]]
// Output: 8
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/16/gif4.gif" />
// Alice picks the pawn at (2, 2) and captures it in two moves: (0, 2) -> (1, 4) -> (2, 2).
// Bob picks the pawn at (3, 3) and captures it in two moves: (2, 2) -> (4, 1) -> (3, 3).
// Alice picks the pawn at (1, 1) and captures it in four moves: (3, 3) -> (4, 1) -> (2, 2) -> (0, 3) -> (1, 1).

// Example 3:
// Input: kx = 0, ky = 0, positions = [[1,2],[2,4]]
// Output: 3
// Explanation:
// Alice picks the pawn at (2, 4) and captures it in two moves: (0, 0) -> (1, 2) -> (2, 4). Note that the pawn at (1, 2) is not captured.
// Bob picks the pawn at (1, 2) and captures it in one move: (2, 4) -> (1, 2).

// Constraints:
//     0 <= kx, ky <= 49
//     1 <= positions.length <= 15
//     positions[i].length == 2
//     0 <= positions[i][0], positions[i][1] <= 49
//     All positions[i] are unique.
//     The input is generated such that positions[i] != [kx, ky] for all 0 <= i < positions.length.

import "fmt"
import "math/bits"

func maxMoves(kx int, ky int, positions [][]int) int {
    n, m := len(positions), 50
    dx := []int{1, 1, 2, 2, -1, -1, -2, -2}
    dy := []int{2, -2, 1, -1, 2, -2, 1, -1}
    dist := make([][][]int, n + 1)
    for i := range dist {
        dist[i] = make([][]int, m)
        for j := range dist[i] {
            dist[i][j] = make([]int, m)
            for k := range dist[i][j] {
                dist[i][j][k] = -1
            }
        }
    }
    for i := 0; i <= n; i++ {
        x, y := kx, ky
        if i < n {
            x, y = positions[i][0], positions[i][1]
        }
        queue := [][2]int{[2]int{x, y} }
        dist[i][x][y] = 0
        for step := 1; len(queue) > 0; step++ {
            for k := len(queue); k > 0; k-- {
                p := queue[0]
                queue = queue[1:] // pop
                x1, y1 := p[0], p[1]
                for j := 0; j < 8; j++ {
                    x2 := x1 + dx[j]
                    y2 := y1 + dy[j]
                    if x2 >= 0 && x2 < m && y2 >= 0 && y2 < m && dist[i][x2][y2] == -1 { // border & visit check
                        dist[i][x2][y2] = step
                        queue = append(queue, [2]int{x2, y2}) // push
                    }
                }
            }
        }
    }
    dp := make([][][]int, n + 1)
    for i := range dp {
        dp[i] = make([][]int, 1<<n)
        for j := range dp[i] {
            dp[i][j] = make([]int, 2)
            for k := range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(last, state, k int) int
    dfs = func(last, state, k int) int {
        if state == 0 { return 0 }
        if dp[last][state][k] != -1 { return dp[last][state][k] }
        res := 0
        if k == 0 { res = 1 << 31 }
        for i, p := range positions {
            x, y := p[0], p[1]
            if (state >> i) & 1 == 1 {
                t := dfs(i, state^(1<<i), k^1) + dist[last][x][y]
                if k == 1 {
                    res = max(res, t)
                } else {
                    res = min(res, t)
                }
            }
        }
        dp[last][state][k] = res
        return res
    }
    return dfs(n, (1 << n) - 1, 1)
}

func maxMoves1(kx, ky int, positions [][]int) int {
    type Pair struct{ x, y int }
    dirs := []Pair{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
    n := len(positions)
    dis := make([][50][50]int, n) // 计算马到兵的步数，等价于计算兵到其余格子的步数
    for i, pos := range positions {
        d := &dis[i]
        for j := range d {
            for k := range d[j] {
                d[j][k] = -1
            }
        }
        px, py := pos[0], pos[1]
        d[px][py] = 0
        q := []Pair{{px, py}}
        for step := 1; len(q) > 0; step++ {
            tmp := q
            q = nil
            for _, p := range tmp {
                for _, dir := range dirs {
                    x, y := p.x + dir.x, p.y + dir.y
                    if 0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0 {
                        d[x][y] = step
                        q = append(q, Pair{x, y})
                    }
                }
            }
        }
    }
    positions = append(positions, []int{kx, ky})
    u := 1 << n - 1
    f := make([][]int, 1<<n)
    for i := range f {
        f[i] = make([]int, n+1)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    op := func(a, b int, odd bool) int {
        if odd { return min(a, b) }
        return max(a, b)
    }
    for mask := 1; mask < 1<<n; mask++ {
        for i, p := range positions {
            x, y := p[0], p[1]
            odd := bits.OnesCount(uint(u^mask))%2 > 0
            if odd { f[mask][i] = 1 << 31 }
            for s := uint(mask); s > 0; s &= s - 1 {
                j := bits.TrailingZeros(s)
                f[mask][i] = op(f[mask][i], f[mask^1<<j][j]+dis[j][x][y], odd)
            }
        }
    }
    return f[u][n]
}

func main() {
    // Example 1:
    // Input: kx = 1, ky = 1, positions = [[0,0]]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/16/gif3.gif" />
    // The knight takes 4 moves to reach the pawn at (0, 0).
    fmt.Println(maxMoves(1, 1, [][]int{{0,0}})) // 4
    // Example 2:
    // Input: kx = 0, ky = 2, positions = [[1,1],[2,2],[3,3]]
    // Output: 8
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/16/gif4.gif" />
    // Alice picks the pawn at (2, 2) and captures it in two moves: (0, 2) -> (1, 4) -> (2, 2).
    // Bob picks the pawn at (3, 3) and captures it in two moves: (2, 2) -> (4, 1) -> (3, 3).
    // Alice picks the pawn at (1, 1) and captures it in four moves: (3, 3) -> (4, 1) -> (2, 2) -> (0, 3) -> (1, 1).
    fmt.Println(maxMoves(0, 2, [][]int{{1,1},{2,2},{3,3}})) // 8
    // Example 3:
    // Input: kx = 0, ky = 0, positions = [[1,2],[2,4]]
    // Output: 3
    // Explanation:
    // Alice picks the pawn at (2, 4) and captures it in two moves: (0, 0) -> (1, 2) -> (2, 4). Note that the pawn at (1, 2) is not captured.
    // Bob picks the pawn at (1, 2) and captures it in one move: (2, 4) -> (1, 2).
    fmt.Println(maxMoves(0, 0, [][]int{{1,2},{2,4}})) // 3

    fmt.Println(maxMoves1(1, 1, [][]int{{0,0}})) // 4
    fmt.Println(maxMoves1(0, 2, [][]int{{1,1},{2,2},{3,3}})) // 8
    fmt.Println(maxMoves1(0, 0, [][]int{{1,2},{2,4}})) // 3
}