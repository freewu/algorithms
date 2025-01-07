package main

// 3257. Maximum Value Sum by Placing Three Rooks II
// You are given a m x n 2D array board representing a chessboard, 
// where board[i][j] represents the value of the cell (i, j).

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
//     3 <= m == board.length <= 500
//     3 <= n == board[i].length <= 500
//     -10^9 <= board[i][j] <= 10^9

import "fmt"
import "sort"

func maximumValueSum(board [][]int) int64 {
    m, inf := len(board), 1 << 48
    type Pair struct{ x, j int }
    suf := make([][3]Pair, m)
    p := [3]Pair{} // 最大、次大、第三大
    for i := range p {
        p[i].x = -inf
    }
    update := func(row []int) {
        for j, x := range row {
            if x > p[0].x {
                if p[0].j != j { // 如果相等，仅更新最大
                    if p[1].j != j { // 如果相等，仅更新最大和次大
                        p[2] = p[1]
                    }
                    p[1] = p[0]
                }
                p[0] = Pair{x, j}
            } else if x > p[1].x && j != p[0].j {
                if p[1].j != j { // 如果相等，仅更新次大
                    p[2] = p[1]
                }
                p[1] = Pair{x, j}
            } else if x > p[2].x && j != p[0].j && j != p[1].j {
                p[2] = Pair{x, j}
            }
        }
    }
    for i := m - 1; i > 1; i-- {
        update(board[i])
        suf[i] = p
    }
    res := -inf
    for i := range p {
        p[i].x = -inf // 重置，计算 pre
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, row := range board[:m-2] {
        update(row)
        for j, x := range board[i+1] { // 第二个车
            for _, p := range p { // 第一个车
                for _, q := range suf[i+2] { // 第三个车
                    if p.j != j && q.j != j && p.j != q.j { // 没有同列的车
                        res = max(res, p.x+x+q.x)
                        break
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

func maximumValueSum2(board [][]int) int64 {
    l, w := len(board), len(board[0])
    ww := make([][]int, l)
    for i := range ww {
        tmp := make([]int, 4)
        tmp[0], tmp[1], tmp[2] = -1, -1, -1
        for j := 0; j < w; j++ {
            for k := 2; k >= 0; k-- {
                if tmp[k] != -1 && board[i][j] < board[i][tmp[k]] { break }
                tmp[k+1] = tmp[k]
                tmp[k] = j
            }
        }
        ww[i] = tmp[:3]
    }
    ss := make([][2]int, l*3)
    t := 0
    for i := 0; i < l; i++ {
        for j := 0; j < 3; j++ {
            ss[t] = [2]int{ i, ww[i][j] }
            t++
        }
    }
    sort.Slice(ss, func(i, j int) bool {
        return board[ss[i][0]][ss[i][1]] > board[ss[j][0]][ss[j][1]]
    })
    res := board[0][0] + board[1][1] + board[2][2]
    var tmp0, tmp1, tmp2 [2]int
    for i := 0; i < len(ss); i++ {
        tmp0 = ss[i]
        if t := 3 * board[tmp0[0]][tmp0[1]]; res >= t { break }
        for j := i + 1; j < len(ss); j++ {
            tmp1 = ss[j]
            if tmp0[0] == tmp1[0] || tmp0[1] == tmp1[1] { continue }
            if t := board[tmp0[0]][tmp0[1]] + 2*board[tmp1[0]][tmp1[1]]; res >= t { break }
            for k := j + 1; k < len(ss); k++ {
                tmp2 = ss[k]
                if tmp0[0] == tmp2[0] || tmp0[1] == tmp2[1] || tmp1[0] == tmp2[0] || tmp1[1] == tmp2[1] {
                    continue
                }
                res = max(res, board[tmp0[0]][tmp0[1]] + board[tmp1[0]][tmp1[1]] + board[tmp2[0]][tmp2[1]])
                break
            }
        }
    }
    return int64(res)
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

    fmt.Println(maximumValueSum2([][]int{{-3,1,1,1},{-3,1,-3,1},{-3,2,1,1}})) // 4
    fmt.Println(maximumValueSum2([][]int{{1,2,3},{4,5,6},{7,8,9}})) // 15
    fmt.Println(maximumValueSum2([][]int{{1,1,1},{1,1,1},{1,1,1}})) // 3
    fmt.Println(maximumValueSum2([][]int{{-1000000000,-1000000000,-1000000000},{-1000000000,-1000000000,-1000000000},{-1000000000,-1000000000,-1000000000}})) // -3000000000
}