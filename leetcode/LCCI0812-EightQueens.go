package main

// 面试题 08.12. Eight Queens LCCI
// Write an algorithm to print all ways of arranging n queens on an n x n chess board so that none of them share the same row, column, or diagonal. 
// In this case, "diagonal" means all diagonals, not just the two that bisect the board.

// Notes: This problem is a generalization of the original one in the book.

// Example:
// Input: 4
// Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// Explanation: 4 queens has following two solutions
// [
// [".Q..",  // solution 1
// "...Q",
// "Q...",
// "..Q."],
// ["..Q.",  // solution 2
// "Q...",
// "...Q",
// ".Q.."]
// ]

import "fmt"
import "strings"

func solveNQueens(n int) [][]string {
    res, arr := [][]string{}, make([]string, n)  // 有多少个列
    for i := 0; i < n; i ++ {
        arr[i] = strings.Repeat(".", n)
    }
    col, dg, udg  := make([]bool, n), make([]bool, n << 1), make([]bool, n << 1)
    var dfs func(u int)
    dfs = func(u int) {
        if u >= n {
            cp := make([]string, n) // 如何存储res当前的快照，因为切片指向的底层数组都是一个地方
            copy(cp, arr)
            res = append(res, cp)
            return 
        }
        for i := 0; i < n; i ++ {  // 枚举列
            if col[i] || dg[u + i] || udg[u - i + n] { continue }
            col[i], dg[u + i], udg[u - i + n] = true, true, true
            arr[u] = arr[u][:i] + "Q" + arr[u][i + 1:]
            dfs(u + 1)
            arr[u] = arr[u][:i] + "." + arr[u][i + 1:]
            col[i], dg[u + i], udg[u - i + n] = false, false, false
        }
    }
    dfs(0)
    return res
}

func main() {
    // Example:
    // Input: 4
    // Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
    // Explanation: 4 queens has following two solutions
    // [
    // [".Q..",  // solution 1
    // "...Q",
    // "Q...",
    // "..Q."],
    // ["..Q.",  // solution 2
    // "Q...",
    // "...Q",
    // ".Q.."]
    // ]
    fmt.Println(solveNQueens(4)) // [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

    fmt.Println(solveNQueens(1)) // [[Q]]
    fmt.Println(solveNQueens(2)) // []
    fmt.Println(solveNQueens(3)) // []
    fmt.Println(solveNQueens(5)) // [[Q.... ..Q.. ....Q .Q... ...Q.] [Q.... ...Q. .Q... ....Q ..Q..] [.Q... ...Q. Q.... ..Q.. ....Q] [.Q... ....Q ..Q.. Q.... ...Q.] [..Q.. Q.... ...Q. .Q... ....Q] [..Q.. ....Q .Q... ...Q. Q....] [...Q. Q.... ..Q.. ....Q .Q...] [...Q. .Q... ....Q ..Q.. Q....] [....Q .Q... ...Q. Q.... ..Q..] [....Q ..Q.. Q.... ...Q. .Q...]]
    fmt.Println(solveNQueens(6)) // [[.Q.... ...Q.. .....Q Q..... ..Q... ....Q.] [..Q... .....Q .Q.... ....Q. Q..... ...Q..] [...Q.. Q..... ....Q. .Q.... .....Q ..Q...] [....Q. ..Q... Q..... .....Q ...Q.. .Q....]]
    //fmt.Println(solveNQueens(7)) // 
    //fmt.Println(solveNQueens(8)) // 
}