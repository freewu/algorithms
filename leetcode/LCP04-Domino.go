package main

// LCP 04. 覆盖
// 你有一块棋盘，棋盘上有一些格子已经坏掉了。
// 你还有无穷块大小为1 * 2的多米诺骨牌，你想把这些骨牌不重叠地覆盖在完好的格子上，请找出你最多能在棋盘上放多少块骨牌？
// 这些骨牌可以横着或者竖着放。

// 输入：n, m代表棋盘的大小；broken是一个b * 2的二维数组，其中每个元素代表棋盘上每一个坏掉的格子的位置。

// 输出：一个整数，代表最多能在棋盘上放的骨牌数。

// 示例 1：
// 输入：n = 2, m = 3, broken = [[1, 0], [1, 1]]
// 输出：2
// 解释：我们最多可以放两块骨牌：[[0, 0], [0, 1]]以及[[0, 2], [1, 2]]。（见下图）
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/09/domino_example_1.jpg" />

// 示例 2：
// 输入：n = 3, m = 3, broken = []
// 输出：4
// 解释：下图是其中一种可行的摆放方式
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/09/domino_example_2.jpg" />

// 限制：
//     1 <= n <= 8
//     1 <= m <= 8
//     0 <= b <= n * m

import "fmt"

func domino(n int, m int, broken [][]int) int {
    type Node struct { from, to, next int }
    // 构建图变量
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, m)
    }
    direction := [][]int{{1, 0},{-1, 0}, {0, 1}, {0, -1}}
    // 邻接表变量
    res, count := 0, 0
    nodes := make([]*Node, 0)
    head := make([]int, n * m)
    for i := 0; i < len(head); i++ {
        head[i] = -1
    }
    // 二分图变量（匈牙利算法）
    board, boardVisited := make([]int, n * m), make([]bool, n * m)
    for i:=0;i<len(board);i++ {
        board[i] = -1
    }
    for _,v := range broken {
        visited[v[0]][v[1]] = true
    }
    add := func(from, to int) {
        node := &Node{ from: from, to: to, next: head[from] }
        nodes = append(nodes, node)
        head[from] = count
        count++
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if visited[i][j] || (i + j) & 1 == 1 { continue }
            from := i * m + j
            for _, v := range direction {
                dx, dy := i + v[0], j + v[1]
                if dx < 0 || dx >= n || dy < 0 || dy >= m || visited[dx][dy] { continue }
                to := dx * m + dy
                add(from, to)
            }
        }
    }
    var find func(int) bool
    find = func(pos int) bool {
        for i := head[pos]; i != -1; i= nodes[i].next {
            if !boardVisited[nodes[i].to] { continue }
            boardVisited[nodes[i].to] = false
            if board[nodes[i].to] == -1 || find(board[nodes[i].to]) {
                board[nodes[i].to] = pos
                return true
            }
        }
        return false
    }
    for i := 0; i < n * m; i++ {
        for j := 0; j < len(boardVisited); j++ {
            boardVisited[j] = true  // 表示可以用来连接
        }
        if head[i] != -1 {
            if find(i) {
                res++
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：n = 2, m = 3, broken = [[1, 0], [1, 1]]
    // 输出：2
    // 解释：我们最多可以放两块骨牌：[[0, 0], [0, 1]]以及[[0, 2], [1, 2]]。（见下图）
    // <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/09/domino_example_1.jpg" />
    fmt.Println(domino(2, 3, [][]int{{1, 0},{1, 1}})) // 2
    // 示例 2：
    // 输入：n = 3, m = 3, broken = []
    // 输出：4
    // 解释：下图是其中一种可行的摆放方式
    // <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/09/domino_example_2.jpg" />
    fmt.Println(domino(3, 3, [][]int{})) // 4

    fmt.Println(domino(8, 8, [][]int{})) // 32
    fmt.Println(domino(8, 8, [][]int{{1, 0},{1, 1}})) // 31
}