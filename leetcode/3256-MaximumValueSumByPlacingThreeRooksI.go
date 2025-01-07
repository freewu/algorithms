package main

// 3256. Maximum Value Sum by Placing Three Rooks I
// You are given a m x n 2D array board representing a chessboard, where board[i][j] represents the value of the cell (i, j).

// Rooks in the same row or column attack each other. 
// You need to place three rooks on the chessboard such that the rooks do not attack each other.

// Return the maximum sum of the cell values on which the rooks are placed.

// Example 1:
// Input: board = [[-3,1,1,1],[-3,1,-3,1],[-3,2,1,1]]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/08/rooks2.png" />
// We can place the rooks in the cells (0, 2), (1, 3), and (2, 1) for a sum of 1 + 1 + 2 = 4.

// Example 2:
// Input: board = [[1,2,3],[4,5,6],[7,8,9]]
// Output: 15
// Explanation:
// We can place the rooks in the cells (0, 0), (1, 1), and (2, 2) for a sum of 1 + 5 + 9 = 15.

// Example 3:
// Input: board = [[1,1,1],[1,1,1],[1,1,1]]
// Output: 3
// Explanation:
// We can place the rooks in the cells (0, 2), (1, 1), and (2, 0) for a sum of 1 + 1 + 1 = 3.

// Constraints:
//     3 <= m == board.length <= 100
//     3 <= n == board[i].length <= 100
//     -10^9 <= board[i][j] <= 10^9

import "fmt"

func maximumValueSum(board [][]int) int64 {
    inf := 1 << 48
    type Pair struct { num, col int }
    getPair := func() (Pair, Pair, Pair) { return Pair{ num: -inf }, Pair { num: -inf }, Pair{ num: -inf }; }
    get3Max := func(board [][]int) [][]Pair {
        r, c := len(board), len(board[0])
        res := make([][]Pair, r)
        for i := 0; i < r; i++ {
            mo1, mo2, mo3 := getPair()
            for j := 0; j < c; j++ {
                if board[i][j] > mo1.num {
                    mo3.num, mo3.col = mo2.num, mo2.col
                    mo2.num, mo2.col = mo1.num, mo1.col
                    mo1.num, mo1.col = board[i][j], j
                } else if board[i][j] > mo2.num {
                    mo3.num, mo3.col = mo2.num, mo2.col
                    mo2.num, mo2.col = board[i][j], j
                } else if board[i][j] > mo3.num {
                    mo3.num, mo3.col = board[i][j], j
                }
            }
            res[i] = append(res[i], mo1, mo2, mo3)
        }
        return res
    }
    res, n := -inf, len(board)
    boardThreeMax := get3Max(board)
    for i := 0; i < n; i++ {
        for im := 0; im < 3; im++ { // top 3 max
            mx1, mx1c := boardThreeMax[i][im].num, boardThreeMax[i][im].col
            for j := i + 1; j < n; j++ {
                for jm := 0; jm < 3; jm++ { // top 3 max
                    mx2, mx2c := boardThreeMax[j][jm].num, boardThreeMax[j][jm].col
                    for k := j + 1; k < n; k++ {
                        for km := 0; km < 3; km++ { // top 3 max
                            mx3,mx3c := boardThreeMax[k][km].num,  boardThreeMax[k][km].col
                            if mx1c != mx2c && mx2c != mx3c && mx1c != mx3c {
                                res = max(res, mx1 + mx2 + mx3)
                            }
                        }
                    }
                }
            }
        }
    }
    return int64(res)
}

func maximumValueSum1(board [][]int) int64 {
    m, n := len(board), len(board[0])
    type Neighbor struct{ to, rid, cap, cost int }
    g := make([][]Neighbor, m + n + 3)
    addEdge := func(from, to, cap, cost int) {
        g[from] = append(g[from], Neighbor{ to, len(g[to]), cap, cost })
        g[to] = append(g[to], Neighbor{ from, len(g[from]) - 1, 0, -cost })
    }
    R, C, S := m + n,  m + n + 1,  m + n + 2
    for i, row := range board {
        for j, x := range row {
            addEdge(i, m + j, 1, -x)
        }
        addEdge(R, i, 1, 0)
    }
    for j := range board[0] {
        addEdge(m+j, C, 1, 0)
    }
    addEdge(S, R, 3, 0) // 把 3 改成 k 可以支持 k 个车
    // 下面是费用流模板
    dis := make([]int, len(g))
    type vi struct{ v, i int }
    fa := make([]vi, len(g))
    inQ := make([]bool, len(g))
    spfa := func() bool {
        for i := range dis {
            dis[i] = 1 << 32
        }
        dis[S] = 0
        inQ[S] = true
        q := []int{S}
        for len(q) > 0 {
            v := q[0]
            q = q[1:]
            inQ[v] = false
            for i, e := range g[v] {
                if e.cap == 0 {
                    continue
                }
                w := e.to
                newD := dis[v] + e.cost
                if newD < dis[w] {
                    dis[w] = newD
                    fa[w] = vi{v, i}
                    if !inQ[w] {
                        inQ[w] = true
                        q = append(q, w)
                    }
                }
            }
        }
        // 循环结束后所有 inQ[v] 都为 false，无需重置
        return dis[C] < 1 << 32
    }
    minCost := 0
    for spfa() {
        minF := 1 << 32
        for v := C; v != S; {
            p := fa[v]
            minF = min(minF, g[p.v][p.i].cap)
            v = p.v
        }
        for v := C; v != S; {
            p := fa[v]
            e := &g[p.v][p.i]
            e.cap -= minF
            g[v][e.rid].cap += minF
            v = p.v
        }
        minCost += dis[C] * minF
    }
    return int64(-minCost)
}

func main() {
    // Example 1:
    // Input: board = [[-3,1,1,1],[-3,1,-3,1],[-3,2,1,1]]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/08/rooks2.png" />
    // We can place the rooks in the cells (0, 2), (1, 3), and (2, 1) for a sum of 1 + 1 + 2 = 4.
    fmt.Println(maximumValueSum([][]int{{-3,1,1,1},{-3,1,-3,1},{-3,2,1,1}})) // 4
    // Example 2:
    // Input: board = [[1,2,3],[4,5,6],[7,8,9]]
    // Output: 15
    // Explanation:
    // We can place the rooks in the cells (0, 0), (1, 1), and (2, 2) for a sum of 1 + 5 + 9 = 15.
    fmt.Println(maximumValueSum([][]int{{1,2,3},{4,5,6},{7,8,9}})) // 15
    // Example 3:
    // Input: board = [[1,1,1],[1,1,1],[1,1,1]]
    // Output: 3
    // Explanation:
    // We can place the rooks in the cells (0, 2), (1, 1), and (2, 0) for a sum of 1 + 1 + 1 = 3.
    fmt.Println(maximumValueSum([][]int{{1,1,1},{1,1,1},{1,1,1}})) // 3
    fmt.Println(maximumValueSum([][]int{{-1000000000,-1000000000,-1000000000},{-1000000000,-1000000000,-1000000000},{-1000000000,-1000000000,-1000000000}})) // -3000000000

    fmt.Println(maximumValueSum1([][]int{{-3,1,1,1},{-3,1,-3,1},{-3,2,1,1}})) // 4
    fmt.Println(maximumValueSum1([][]int{{1,2,3},{4,5,6},{7,8,9}})) // 15
    fmt.Println(maximumValueSum1([][]int{{1,1,1},{1,1,1},{1,1,1}})) // 3
    fmt.Println(maximumValueSum1([][]int{{-1000000000,-1000000000,-1000000000},{-1000000000,-1000000000,-1000000000},{-1000000000,-1000000000,-1000000000}})) // -3000000000
}