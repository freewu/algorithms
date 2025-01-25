package main

// LCP 13. 寻宝
// 我们得到了一副藏宝图，藏宝图显示，在一个迷宫中存在着未被世人发现的宝藏。

// 迷宫是一个二维矩阵，用一个字符串数组表示。它标识了唯一的入口（用 'S' 表示），和唯一的宝藏地点（用 'T' 表示）。
// 但是，宝藏被一些隐蔽的机关保护了起来。在地图上有若干个机关点（用 'M' 表示），只有所有机关均被触发，才可以拿到宝藏。

// 要保持机关的触发，需要把一个重石放在上面。迷宫中有若干个石堆（用 'O' 表示），每个石堆都有无限个足够触发机关的重石。
// 但是由于石头太重，我们一次只能搬一个石头到指定地点。

// 迷宫中同样有一些墙壁（用 '#' 表示），我们不能走入墙壁。剩余的都是可随意通行的点（用 '.' 表示）。
// 石堆、机关、起点和终点（无论是否能拿到宝藏）也是可以通行的。

// 我们每步可以选择向上/向下/向左/向右移动一格，并且不能移出迷宫。
// 搬起石头和放下石头不算步数。
// 那么，从起点开始，我们最少需要多少步才能最后拿到宝藏呢？如果无法拿到宝藏，返回 -1 。

// 示例 1：
// 输入： ["S#O", "M..", "M.T"]
// 输出：16
// 解释：最优路线为： S->O, cost = 4, 去搬石头 O->第二行的M, cost = 3, M机关触发 第二行的M->O, cost = 3, 我们需要继续回去 O 搬石头。 O->第三行的M, cost = 4, 此时所有机关均触发 第三行的M->T, cost = 2，去T点拿宝藏。 总步数为16。图片.gif
// <img src="https://pic.leetcode-cn.com/6bfff669ad65d494cdc237bcedfec10a2b1ac2f2593c2bf97e9aecb41dc8a08b-%E5%9B%BE%E7%89%87.gif" />

// 示例 2：
// 输入： ["S#O", "M.#", "M.T"]
// 输出：-1
// 解释：我们无法搬到石头触发机关

// 示例 3：
// 输入： ["S#O", "M.T", "M.."]
// 输出：17
// 解释：注意终点也是可以通行的。

// 限制：
//     1 <= maze.length <= 100
//     1 <= maze[i].length <= 100
//     maze[i].length == maze[j].length
//     S 和 T 有且只有一个
//     0 <= M的数量 <= 16
//     0 <= O的数量 <= 40，题目保证当迷宫中存在 M 时，一定存在至少一个 O 。

import "fmt"

func minimalSteps(maze []string) int {
    n, m := len(maze), len(maze[0])
    buttons, stones := [][]int{},  [][]int{} // 机关 & 石头 
    sx, sy, tx, ty := -1, -1, -1, -1 // 起点 & 终点
    dx, dy := []int{ 1, -1, 0, 0 }, []int{ 0, 0, 1, -1 }
    inBound := func(x, y int) bool { return x >= 0 && x < n && y >= 0 && y < m }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            switch maze[i][j] {
            case 'M':
                buttons = append(buttons, []int{i, j})
            case 'O':
                stones = append(stones, []int{i, j})
            case 'S':
                sx, sy = i, j
            case 'T':
                tx, ty = i, j
            }
        }
    }
    nb, ns := len(buttons), len(stones)
    bfs := func (x, y int, maze []string) [][]int {
        res := make([][]int, n)
        for i := 0; i < n; i++ {
            res[i] = make([]int, m)
            for j := 0; j < m; j++ {
                res[i][j] = -1
            }
        }
        res[x][y] = 0
        queue := [][]int{}
        queue = append(queue, []int{x, y})
        for len(queue) > 0 {
            p := queue[0]
            queue = queue[1:]
            curx, cury := p[0], p[1]
            for k := 0; k < 4; k++ {
                nx, ny := curx + dx[k], cury + dy[k]
                if inBound(nx, ny) && maze[nx][ny] != '#' && res[nx][ny] == -1 {
                    res[nx][ny] = res[curx][cury] + 1
                    queue = append(queue, []int{nx, ny})
                }
            }
        }
        return res
    }
    startDist := bfs(sx, sy, maze)
    // 边界情况：没有机关
    if nb == 0 {
        return startDist[tx][ty]
    }
    // 从某个机关到其他机关 / 起点与终点的最短距离。
    dist := make([][]int, nb)
    for i := 0; i < nb; i++ {
        dist[i] = make([]int, nb + 2)
        for j := 0; j < nb + 2; j++ {
            dist[i][j] = -1
        }
    }
    // 中间结果
    dd := make([][][]int, nb)
    for i := 0; i < nb; i++ {
        dd[i] = bfs(buttons[i][0], buttons[i][1], maze)
        // 从某个点到终点不需要拿石头
        dist[i][nb + 1] = dd[i][tx][ty]
    }
    for i := 0; i < nb; i++ {
        tmp := -1
        for k := 0; k < ns; k++ {
            midX, midY := stones[k][0], stones[k][1]
            if dd[i][midX][midY] != -1 && startDist[midX][midY] != -1 {
                if tmp == -1 || tmp > dd[i][midX][midY] + startDist[midX][midY] {
                    tmp = dd[i][midX][midY] + startDist[midX][midY]
                }
            }
        }
        dist[i][nb] = tmp
        for j := i + 1; j < nb; j++ {
            mn := -1
            for k := 0; k < ns; k++ {
                midX, midY := stones[k][0], stones[k][1]
                if dd[i][midX][midY] != -1 && startDist[midX][midY] != -1 {
                    if mn == -1 || mn > dd[i][midX][midY] + dd[j][midX][midY] {
                        mn = dd[i][midX][midY] + dd[j][midX][midY]
                    }
                }
            }
            dist[i][j] = mn
            dist[j][i] = mn
        }
    }
    // 无法达成的情形
    for i := 0; i < nb; i++ {
        if dist[i][nb] == -1 || dist[i][nb + 1] == -1 {
            return -1
        }
    }
    // dp 数组， -1 代表没有遍历到
    dp := make([][]int, 1 << nb)
    for i := 0; i < (1 << nb); i++ {
        dp[i] = make([]int, nb)
        for j := 0; j < nb; j++ {
            dp[i][j] = -1
        }
    }
    for i := 0; i < nb; i++ {
        dp[1 << i][i] = dist[i][nb]
    }
    // 由于更新的状态都比未更新的大，所以直接从小到大遍历即可
    for mask := 1; mask < (1 << nb); mask++ {
        for i := 0; i < nb; i++ {
            // 当前 dp 是合法的
            if mask & (1 << i) != 0 {
                for j := 0; j < nb; j++ {
                    // j 不在 mask 里
                    if mask & (1 << j) == 0 {
                        next := mask | (1 << j)
                        if dp[next][j] == -1 || dp[next][j] > dp[mask][i] + dist[i][j] {
                            dp[next][j] = dp[mask][i] + dist[i][j]
                        }
                    }
                }
            }
        }
    }
    res, finalMask := -1, (1 << nb) - 1
    for i := 0; i < nb; i++ {
        if res == -1 || res > dp[finalMask][i] + dist[i][nb + 1] {
            res = dp[finalMask][i] + dist[i][nb + 1]
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入： ["S#O", "M..", "M.T"]
    // 输出：16
    // 解释：最优路线为： S->O, cost = 4, 去搬石头 O->第二行的M, cost = 3, M机关触发 第二行的M->O, cost = 3, 我们需要继续回去 O 搬石头。 O->第三行的M, cost = 4, 此时所有机关均触发 第三行的M->T, cost = 2，去T点拿宝藏。 总步数为16。图片.gif
    // <img src="https://pic.leetcode-cn.com/6bfff669ad65d494cdc237bcedfec10a2b1ac2f2593c2bf97e9aecb41dc8a08b-%E5%9B%BE%E7%89%87.gif" />
    fmt.Println(minimalSteps([]string{"S#O", "M..", "M.T"})) // 16
    // 示例 2：
    // 输入： ["S#O", "M.#", "M.T"]
    // 输出：-1
    // 解释：我们无法搬到石头触发机关
    fmt.Println(minimalSteps([]string{"S#O", "M.#", "M.T"})) // -1
    // 示例 3：
    // 输入： ["S#O", "M.T", "M.."]
    // 输出：17
    // 解释：注意终点也是可以通行的。
    fmt.Println(minimalSteps([]string{"S#O", "M.T", "M.."})) // 17
}